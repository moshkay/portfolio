# Portfolio Website

A fast, single-binary personal portfolio website built with Go. Templates and
static assets are compiled into the executable via `go:embed`, so deployment is
just one file — no external dependencies, no CDN required.

## Features

- **Zero-dependency runtime** — pure Go standard library (`net/http`, `html/template`, `embed`).
- **Self-contained binary** — HTML, CSS, and JS are embedded at build time.
- **Modern, responsive UI** — light/dark theme toggle, scroll-reveal animations, mobile friendly.
- **Production-minded server** — structured JSON logging, panic recovery, request logging, read/write timeouts, and graceful shutdown.
- **`/healthz` endpoint** — ready for Kubernetes liveness/readiness probes.
- **Docker-ready** — multi-stage build, non-root user, healthcheck.

## Project structure

```
portfolio-website/
├── cmd/portfolio/        # main entrypoint (config + graceful shutdown)
├── internal/server/      # HTTP server, routes, middleware, content model
│   ├── server.go
│   ├── middleware.go
│   ├── data.go           # <-- edit your portfolio content here
│   └── icons.go
├── web/
│   ├── embed.go          # embeds templates + static assets
│   ├── templates/        # index.html
│   └── static/           # css/, js/
├── Dockerfile
└── go.mod
```

## Getting started

Requires Go 1.21+.

```bash
# Run locally
go run ./cmd/portfolio

# Then open http://localhost:8080
```

Configure the port with the `PORT` environment variable (defaults to `8080`):

```bash
PORT=3000 go run ./cmd/portfolio
```

## Build a binary

```bash
go build -trimpath -ldflags="-s -w" -o portfolio ./cmd/portfolio
./portfolio
```

## Run with Docker

```bash
docker build -t portfolio .
docker run --rm -p 8080:8080 portfolio
```

## Customizing content

All content lives in `internal/server/data.go` in the `defaultPortfolio()`
function — update your name, role, about text, skills, projects, experience, and
social links there. Adjust colors and layout in `web/static/css/style.css`.

## Endpoints

| Method | Path        | Description              |
| ------ | ----------- | ------------------------ |
| GET    | `/`         | The portfolio page       |
| GET    | `/healthz`  | Health check (JSON)      |
| GET    | `/static/*` | Embedded static assets   |
