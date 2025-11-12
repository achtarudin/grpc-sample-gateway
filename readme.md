# grpc-sample-gateway

HTTP/JSON gateway and Swagger UI for the gRPC Sample services. This service translates RESTful HTTP requests into gRPC calls using grpc-gateway v2 and exposes interactive API docs powered by the OpenAPI spec embedded in the shared proto module.

## Overview

This gateway sits in front of the gRPC server and provides:

- HTTP to gRPC translation via grpc-gateway v2 runtime
- A minimal HTTP router using Gorilla Mux
- Swagger UI at `/doc/` and raw OpenAPI at `/swagger.json`
- Graceful shutdown and simple runtime configuration via environment variables

It relies on the shared module `github.com/achtarudin/grpc-sample` for generated gateway bindings and an embedded OpenAPI document.

Related repositories in this workspace:

- grpc-sample: Protobuf definitions, generated code, and embedded OpenAPI assets
- grpc-sample-server: The gRPC backend serving the actual business logic
- grpc-sample-client: Example CLI client for calling the gRPC services directly

## Features

- JSON over HTTP to gRPC bridging using `github.com/grpc-ecosystem/grpc-gateway/v2`
- Built-in Swagger UI backed by the embedded OpenAPI spec from `grpc-sample`
- Hot-reload dev mode with `gow`
- Single binary deploy with simple Docker/Compose setup

## Runtime configuration

Environment variables (with defaults from `cmd/server/main.go`):

- `GRPC_REMOTE_SERVER` (default: `localhost:7000`) — gRPC backend address the gateway connects to
- `GRPC_TLS` (default: `false`) — enable TLS when dialing the gRPC backend
- `GATEWAY_PORT` (default: `8081`) — HTTP port exposed by the gateway

## Endpoints

- `GET /doc/` — Swagger UI
- `GET /swagger.json` — OpenAPI document
- Other paths are routed to the grpc-gateway mux and depend on the service definitions. Refer to the Swagger UI for the full list of REST endpoints.

## Quick start (local)

Prerequisites:

- Go 1.21+ (tested with go.mod requirements)
- Running gRPC backend (see `grpc-sample-server`) or adjust `GRPC_REMOTE_SERVER`

Common tasks via Makefile:

- Install tools: `make install-tools`
- Fetch deps: `make install-deps`
- Dev server with hot reload: `make dev-server`
- Build binary: `make build-server` (outputs `./bin/grpc-sample-gateway`)
- Build and run binary: `make prod-server`

The dev server reads environment variables (from your shell or a `.env` file if present). Example `.env`:

```
GATEWAY_PORT=8081
GRPC_REMOTE_SERVER=localhost:7000
GRPC_TLS=false
```

## Run with Docker Compose (dev)

The provided `docker-compose.yml` supports development with live code mounts and a shared Go module cache. It expects an external Docker network named `grpc_sample_network` (so the gateway can reach the server container by name if you run the server with the same network).

Steps:

1) Create the external network once (if you don’t already have it):
	- `docker network create grpc_sample_network`
2) Provide a `.env` file (or environment variables) containing at least `GATEWAY_PORT`.
3) Start the dev container:
	- `docker compose up --build dev`

The gateway will be available on `http://localhost:${GATEWAY_PORT}`.

## Run with Docker Compose (prod-like)

`docker-compose.prod.yml` builds a production image and exposes the service on port 6000 by default. It includes Traefik labels for routing under `grpc-gateway.cutbray.tech` when connected to a Traefik-managed network.

- Image name: `${IMAGE_NAME:-grpc-sample-gateway}:${IMAGE_VERSION:-latest}`
- Container port: `6000` (env `GATEWAY_PORT=6000`)

Example:

1) Ensure your Traefik network exists and is named `traefik-network`.
2) Build and run:
	- `docker compose -f docker-compose.prod.yml up --build -d`

## Project layout

- `cmd/server/main.go` — bootstraps the HTTP server, grpc-gateway mux, Swagger routes, and graceful shutdown
- `internal/adapter/gateway` — gateway configuration and registration with the grpc-gateway runtime
- `internal/adapter/http/handler` — HTTP handlers for Swagger UI and grpc-gateway routing
- `internal/adapter/logging` — simple log formatter
- `internal/helper` — environment helpers
- `Makefile` — dev/prod targets and dependency management

## Notes

- The gateway depends on the gRPC server being reachable at `GRPC_REMOTE_SERVER`. If running everything locally via Docker, prefer using service names on the same Docker network (e.g., `grpc-sample-server:7000`).
- The OpenAPI document is embedded in the shared `grpc-sample` module and served at `/swagger.json`; Swagger UI at `/doc/` points to it automatically.

## License

This is a sample project for educational and demonstration purposes.

