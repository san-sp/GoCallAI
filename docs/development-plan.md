# Development plan

## Stage 1 — Foundation (current)

- Go API bootstrap and configuration.
- Liveness/readiness endpoints.
- PostgreSQL and Redis local Compose services.
- Architecture and privacy notes.

## Stage 2 — Telephony integration

Before implementing a provider adapter, confirm:
- Indian account eligibility, KYC, number provisioning, and inbound calling.
- AgentStream / bidirectional media-stream entitlement.
- Audio codecs, sample rates, frame sizes, and WebSocket protocol.
- Inbound webhook authentication and callback retry behavior.
- Whether forwarded calls preserve original caller ID.
- Provider-side recording defaults and retention.

Then implement an isolated provider adapter and call lifecycle state machine.

## Stage 3 — Realtime voice

- Streaming audio ingress/egress with bounded buffers and backpressure.
- Realtime model behind an interface so providers can be changed.
- Evaluate Malayalam and English using actual telephony-quality audio.
- Support interruption, silence timeouts, and safe call termination.

## Stage 4 — Persistence and recordings

- Migrations for users, devices, settings, calls, messages, and recording metadata.
- Private object storage for optional recordings.
- Separate caller/assistant tracks where available; otherwise synchronize mixed audio and transcript timestamps.
- Recording disabled by default; disclose and obtain required consent.
- Retention, deletion, encryption, and access authorization.

## Stage 5 — Android

- Kotlin + Jetpack Compose.
- Login, dashboard, automatic/manual mode controls.
- User-guided carrier forwarding workflow and verified status.
- Call history, summaries, transcript, messages.
- FCM notifications and authenticated recording playback.

## Scope discipline

Build a modular monolith first. Avoid microservices until operational needs justify them. Keep provider SDK details behind internal interfaces. Do not treat a successful local build as proof that a carrier call path works.
