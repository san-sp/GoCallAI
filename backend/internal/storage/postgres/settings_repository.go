package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/san-sp/GoCallAI/backend/internal/domain"
)

type SettingsRepository struct {
	pool *pgxpool.Pool
}

func NewSettingsRepository(pool *pgxpool.Pool) *SettingsRepository {
	return &SettingsRepository{pool: pool}
}

func (r *SettingsRepository) Get(ctx context.Context) (domain.AssistantSettings, error) {
	var settings domain.AssistantSettings
	err := r.pool.QueryRow(ctx, `
		SELECT mode, preferred_language, greeting, recording_enabled, updated_at
		FROM assistant_settings
		WHERE id = 1
	`).Scan(
		&settings.Mode,
		&settings.PreferredLanguage,
		&settings.Greeting,
		&settings.RecordingEnabled,
		&settings.UpdatedAt,
	)
	if err != nil {
		return domain.AssistantSettings{}, fmt.Errorf("get assistant settings: %w", err)
	}
	return settings, nil
}

func (r *SettingsRepository) Update(ctx context.Context, input domain.UpdateAssistantSettings) (domain.AssistantSettings, error) {
	var settings domain.AssistantSettings
	err := r.pool.QueryRow(ctx, `
		UPDATE assistant_settings
		SET
			mode = COALESCE($1, mode),
			preferred_language = COALESCE($2, preferred_language),
			greeting = COALESCE($3, greeting),
			recording_enabled = COALESCE($4, recording_enabled),
			updated_at = NOW()
		WHERE id = 1
		RETURNING mode, preferred_language, greeting, recording_enabled, updated_at
	`,
		input.Mode,
		input.PreferredLanguage,
		input.Greeting,
		input.RecordingEnabled,
	).Scan(
		&settings.Mode,
		&settings.PreferredLanguage,
		&settings.Greeting,
		&settings.RecordingEnabled,
		&settings.UpdatedAt,
	)
	if err != nil {
		return domain.AssistantSettings{}, fmt.Errorf("update assistant settings: %w", err)
	}
	return settings, nil
}
