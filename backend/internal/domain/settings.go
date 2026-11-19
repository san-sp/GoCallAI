package domain

import "time"

type AssistantSettings struct {
	Mode              string    `json:"mode"`
	PreferredLanguage string    `json:"preferred_language"`
	Greeting          string    `json:"greeting"`
	RecordingEnabled  bool      `json:"recording_enabled"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateAssistantSettings struct {
	Mode              *string `json:"mode"`
	PreferredLanguage *string `json:"preferred_language"`
	Greeting          *string `json:"greeting"`
	RecordingEnabled  *bool   `json:"recording_enabled"`
}
