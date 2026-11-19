.PHONY: backend-run backend-build infra-up infra-down

backend-run:
	cd backend && go run ./cmd/api

backend-build:
	cd backend && go build ./...

infra-up:
	docker compose -f deployments/docker-compose.yml up -d

infra-down:
	docker compose -f deployments/docker-compose.yml down
