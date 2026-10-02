FROM golang:1.27.0-bookworm AS builder

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", ".", "serve", "--http=0.0.0.0:8090"]
