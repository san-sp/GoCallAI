# Backend

Go HTTP API for GoCallAI.

## Requirements

- Go 1.27.4+ for local development.
- Docker Compose for PostgreSQL and Redis.

## Start the complete local stack

From the repository root:

```bash
cp .env.example .env
# Replace API_ACCESS_TOKEN with: openssl rand -hex 32
docker compose -f deployments/docker-compose.yml up -d --build
```

The API listens on `:8080`. PostgreSQL and Redis are exposed locally on ports 5438 and 6381.

## Endpoints

Public:
- `GET /healthz`
- `GET /readyz`

Bearer-token protected:
- `GET /api/v1/version`
- `GET /api/v1/settings`
- `PATCH /api/v1/settings`
- `GET /api/v1/calls?limit=20&offset=0`
- `GET /api/v1/calls/{id}`

Pass `Authorization: Bearer $API_ACCESS_TOKEN` for protected endpoints.

Database migrations are embedded in the binary and applied at startup. The current migration creates the single-owner assistant settings record, calls, and call events tables.

The call-history endpoints return an empty list until a telephony provider adapter is implemented. Recording is intentionally unavailable until recording, disclosure, and consent handling are built.
