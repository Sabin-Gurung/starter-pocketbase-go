# syntax=docker/dockerfile:1
FROM golang:1.27.0-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal/ internal/
COPY main.go main.go

RUN CGO_ENABLED=0 go build -o /app/server .

# ---- Runtime Stage ----
FROM debian:bookworm-slim

RUN groupadd --system appuser && useradd --system --gid appuser appuser

RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server .
COPY pb_public/ pb_public/

RUN chown -R appuser:appuser /app

EXPOSE 8090

USER appuser

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8090/api/health || exit 1

CMD ["./server", "serve", "--http=0.0.0.0:8090"]
