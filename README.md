# Go PocketBase Starter

A lightweight, modular, production-ready Go + PocketBase backend starter template. Designed for rapid bootstrapping of custom backend microservices powered by embedded PocketBase.

---

## Features

* **Embedded PocketBase Framework**: Uses PocketBase (`v0.39+`) as a Go library with full access to hooks, events, and router custom endpoints.
* **Auto Admin Bootstrapping**: Automatically creates a initial superuser in `DEV_MODE` if one does not exist.
* **Modular Clean Architecture**:
  * `main.go`: Application initialization and server binding.
  * `internal/config`: Centralized environment configuration loader.
  * `internal/bootstrap`: Admin seeding and startup bootstrapping routines.
  * `internal/api`: Custom HTTP endpoints via PocketBase Chi router (`core.RequestEvent`).
  * `internal/db`: PocketBase event hooks and database utilities.
* **Docker Ready**: Production-grade multi-stage `Dockerfile` and `docker-compose.yml`.

---

## Project Structure

```
.
├── Dockerfile             # Multi-stage production build
├── docker-compose.yml     # Local compose setup
├── main.go                # Server entrypoint
├── go.mod / go.sum        # Go dependencies
├── internal/
│   ├── api/               # Custom HTTP route handlers (/api/health, /api/hello)
│   ├── bootstrap/         # Initial admin bootstrapping
│   ├── config/            # Environment variable configuration
│   └── db/                # PocketBase event hooks
└── pb_public/             # Static file directory for frontend web assets
```

---

## Quick Start

### 1. Run Locally

```sh
go run .
# or
go run . serve --http=0.0.0.0:8090
```

* **PocketBase Admin UI**: [http://localhost:8090/_/](http://localhost:8090/_/)
* **Default Admin**: `admin@example.com` / `supersecret` (when `DEV_MODE=true`)
* **Health Check**: [http://localhost:8090/api/health](http://localhost:8090/api/health)

---

### 2. Run via Docker Compose

```sh
docker compose up --build -d
```

---

## How to use this template for a new backend service

1. **Copy repository directory**:
   ```sh
   cp -r go-pocketbase-starter my-new-backend
   cd my-new-backend
   ```

2. **Update Go module name**:
   ```sh
   go mod edit -module github.com/your-org/my-new-backend
   ```
   *Update package import paths in `main.go` and `internal/` files accordingly.*

3. **Add Custom Routes**:
   Define new endpoints in `internal/api/` and register them inside `RegisterRoutes(se *core.ServeEvent)`.

4. **Add Event Hooks**:
   Register custom record hooks (e.g. `OnRecordCreateExecute`) in `internal/db/hooks.go`.

---

## Environment Variables

| Parameter | Default | Description |
|-----------|---------|-------------|
| `PORT` | `8090` | HTTP listening port |
| `DEV_MODE` | `false` | Enables auto admin bootstrapping on start |
| `ADMIN_EMAIL` | `admin@example.com` | Initial superuser email |
| `ADMIN_PASSWORD` | `supersecret` | Initial superuser password |
