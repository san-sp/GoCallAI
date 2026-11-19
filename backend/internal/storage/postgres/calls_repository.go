package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/san-sp/GoCallAI/backend/internal/domain"
)

type CallsRepository struct {
	pool *pgxpool.Pool
}

func NewCallsRepository(pool *pgxpool.Pool) *CallsRepository {
	return &CallsRepository{pool: pool}
}

func (r *CallsRepository) List(ctx context.Context, limit, offset int) ([]domain.CallSummary, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, caller_number, status, started_at, ended_at,
		       duration_seconds, recording_object_key IS NOT NULL, summary
		FROM calls
		ORDER BY started_at DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list calls: %w", err)
	}
	defer rows.Close()

	calls := make([]domain.CallSummary, 0)
	for rows.Next() {
		var call domain.CallSummary
		if err := rows.Scan(
			&call.ID,
			&call.CallerNumber,
			&call.Status,
			&call.StartedAt,
			&call.EndedAt,
			&call.DurationSeconds,
			&call.HasRecording,
			&call.Summary,
		); err != nil {
			return nil, fmt.Errorf("scan call: %w", err)
		}
		calls = append(calls, call)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate calls: %w", err)
	}
	return calls, nil
}

func (r *CallsRepository) Get(ctx context.Context, id uuid.UUID) (domain.CallDetail, error) {
	var call domain.CallDetail
	err := r.pool.QueryRow(ctx, `
		SELECT id, caller_number, status, started_at, ended_at,
		       duration_seconds, recording_object_key IS NOT NULL, summary, transcript
		FROM calls
		WHERE id = $1
	`, id).Scan(
		&call.ID,
		&call.CallerNumber,
		&call.Status,
		&call.StartedAt,
		&call.EndedAt,
		&call.DurationSeconds,
		&call.HasRecording,
		&call.Summary,
		&call.Transcript,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.CallDetail{}, pgx.ErrNoRows
		}
		return domain.CallDetail{}, fmt.Errorf("get call: %w", err)
	}
	return call, nil
}
