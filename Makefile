.PHONY: backend-run backend-build infra-up infra-down stack-up stack-down

backend-run:
	cd backend && go run ./cmd/api

backend-build:
	cd backend && go build ./...

infra-up:
	docker compose -f deployments/docker-compose.yml up -d postgres redis

infra-down:
	docker compose -f deployments/docker-compose.yml down

stack-up:
	docker compose -f deployments/docker-compose.yml up -d --build

stack-down:
	docker compose -f deployments/docker-compose.yml down
