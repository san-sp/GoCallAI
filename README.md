# GoCallAI

A personal AI phone assistant for calls forwarded from an existing mobile number.

> Status: foundation. Telephony and AI integrations are not configured until provider access, credentials, and the supported audio-streaming contract are confirmed.

## Product goals

- Answer calls forwarded by the user's mobile operator.
- Support Malayalam and English conversations.
- Provide automatic forwarding and user-activated unavailable mode.
- Capture caller messages and generate summaries.
- Optionally record calls with disclosure, required consent, retention controls, and deletion.
- Provide an Android app for settings, notifications, call history, and playback.

## Repository layout

- `backend/`: Go API and provider/voice adapters.
- `android/`: Kotlin + Jetpack Compose application (next implementation stage).
- `deployments/`: local infrastructure.
- `docs/`: architecture and development plan.

The telephony provider and cloud backend must answer calls independently of the Android app.

## Run backend

Requirements: Go 1.23+.

```bash
cd backend
go mod tidy
go run ./cmd/api
```

- `GET /healthz`: liveness.
- `GET /readyz`: current configuration readiness; external dependencies are not integrated yet.
- `GET /api/v1/version`: service version.

## Local infrastructure

```bash
docker compose -f deployments/docker-compose.yml up -d
```

## Privacy and security defaults

- No credentials are committed.
- Call recording is disabled by default.
- Do not persist raw audio until caller disclosure, consent, applicable legal requirements, provider retention, and access controls have been reviewed.
- Transcripts, recordings, and summaries are sensitive; define access controls and retention/deletion before production.
- Authenticate provider webhooks and make call events idempotent.
- Requested assistant mode and verified carrier forwarding state are separate.

## Milestones

1. Backend foundation and local infrastructure.
2. Provider account, inbound number, webhook validation, and verified forwarding.
3. Bidirectional streaming and call lifecycle.
4. Realtime voice integration and Malayalam/English evaluation.
5. Recording, object storage, transcripts, and summaries.
6. Android app, FCM notifications, call history, and playback.
7. Production hardening and deployment.

See [architecture](docs/architecture.md) and [development plan](docs/development-plan.md).
