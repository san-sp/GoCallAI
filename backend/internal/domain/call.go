package domain

import "time"

type CallSummary struct {
	ID              string     `json:"id"`
	CallerNumber    string     `json:"caller_number"`
	Status          string     `json:"status"`
	StartedAt       time.Time  `json:"started_at"`
	EndedAt         *time.Time `json:"ended_at,omitempty"`
	DurationSeconds int64      `json:"duration_seconds"`
	HasRecording    bool       `json:"has_recording"`
	Summary         string     `json:"summary,omitempty"`
}

type CallDetail struct {
	CallSummary
	Transcript string `json:"transcript,omitempty"`
}
