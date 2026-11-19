# Backend

Go HTTP API for GoCallAI.

## Requirements

- Go 1.27.4 or newer in the supported release line.

## Run locally

From this directory:

```bash
go mod tidy
go run ./cmd/api
```

The API listens on `:8080` by default.

- `GET /healthz`
- `GET /readyz`
- `GET /api/v1/version`

The current readiness response intentionally reports database, Redis, telephony, and AI as not configured. The next backend milestone is database persistence and provider-neutral call lifecycle management.
