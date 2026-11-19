CREATE TABLE assistant_settings (
    id SMALLINT PRIMARY KEY CHECK (id = 1),
    mode TEXT NOT NULL DEFAULT 'automatic'
        CHECK (mode IN ('automatic', 'manual_active', 'disabled')),
    preferred_language TEXT NOT NULL DEFAULT 'auto'
        CHECK (preferred_language IN ('auto', 'ml', 'en')),
    greeting TEXT NOT NULL DEFAULT 'Hello, I am the AI assistant. How can I help you?',
    recording_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO assistant_settings (id)
VALUES (1)
ON CONFLICT (id) DO NOTHING;

CREATE TABLE calls (
    id UUID PRIMARY KEY,
    provider TEXT NOT NULL,
    provider_call_id TEXT NOT NULL,
    caller_number TEXT NOT NULL DEFAULT '',
    called_number TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL
        CHECK (status IN ('ringing', 'in_progress', 'completed', 'failed', 'missed')),
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    answered_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    duration_seconds BIGINT NOT NULL DEFAULT 0 CHECK (duration_seconds >= 0),
    recording_object_key TEXT,
    transcript TEXT NOT NULL DEFAULT '',
    summary TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (provider, provider_call_id)
);

CREATE INDEX calls_started_at_idx ON calls (started_at DESC);
CREATE INDEX calls_status_idx ON calls (status);

CREATE TABLE call_events (
    id BIGSERIAL PRIMARY KEY,
    call_id UUID NOT NULL REFERENCES calls(id) ON DELETE CASCADE,
    event_type TEXT NOT NULL,
    payload JSONB NOT NULL DEFAULT '{}'::JSONB,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX call_events_call_id_idx ON call_events (call_id, occurred_at);
