# GoCallAI

A personal AI phone assistant for calls forwarded from an existing mobile number.

> Current status: backend persistence and authenticated settings/call-history APIs. Carrier forwarding, live telephony, AI voice, Android UI, and recording pipeline are not integrated yet.

## Product goals

- Answer calls forwarded by the user's mobile operator.
- Support Malayalam and English conversations.
- Provide automatic forwarding and user-activated unavailable mode.
- Capture caller messages and generate summaries.
- Optionally record calls with disclosure, required consent, retention controls, and deletion.
- Provide an Android app for settings, notifications, call history, and playback.

## Repository layout

- `backend/`: Go API and provider/voice adapters.
- `android/`: Kotlin + Jetpack Compose application (planned).
- `deployments/`: local infrastructure.
- `docs/`: architecture and development plan.

The telephony provider and cloud backend must answer calls independently of the Android app.

## Run locally

Requirements: Go 1.27.4+ and Docker Compose.

1. Copy `.env.example` to `.env`.
2. Generate a private API token and replace the placeholder:

```bash
openssl rand -hex 32
```

3. Start the complete local stack:

```bash
docker compose -f deployments/docker-compose.yml up -d --build
```

4. Check readiness:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

The API reads the root `.env` file. The Compose stack exposes the API on 8080, PostgreSQL on 5438, and Redis on 6381.

## API

All `/api/v1` endpoints require:

```http
Authorization: Bearer <API_ACCESS_TOKEN>
```

- `GET /api/v1/version`
- `GET /api/v1/settings`
- `PATCH /api/v1/settings`
- `GET /api/v1/calls?limit=20&offset=0`
- `GET /api/v1/calls/{id}`

Example settings update:

```bash
curl -X PATCH http://localhost:8080/api/v1/settings \
  -H "Authorization: Bearer $API_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"mode":"automatic","preferred_language":"ml"}'
```

Call history will remain empty until a telephony provider is connected.

## Privacy and security defaults

- No provider credentials are committed.
- Call recording is disabled by default and cannot yet be enabled through the API; the recording pipeline and consent/disclosure flow are not implemented.
- Transcripts, recordings, and summaries are sensitive; define access controls and retention/deletion before production.
- Provider webhooks must be authenticated and call events idempotent.
- Requested assistant mode and verified carrier forwarding state are separate.
- The current bearer token is a single-owner development gate, not a substitute for production user authentication and token rotation.

See [architecture](docs/architecture.md) and [development plan](docs/development-plan.md).
