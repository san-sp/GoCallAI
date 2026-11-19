# Build from the backend directory:
# docker build -t gocallai-api -f Dockerfile .
FROM golang:1.27.4-alpine AS builder

WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/gocallai-api ./cmd/api

FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /
COPY --from=builder /out/gocallai-api /gocallai-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/gocallai-api"]
