# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Personal website for Ilia Zalesskii (ilia.fi) - a minimal Go HTTP server that renders Markdown pages with goldmark. The server embeds Markdown files at compile time using `//go:embed` and converts them to HTML on-the-fly.

## Architecture

- **server/server.go**: Main HTTP server with markdown rendering
  - Embeds pages from `server/pages/` directory at compile time using `//go:embed`
  - Uses goldmark for Markdown to HTML conversion
  - Wraps rendered content in minimal HTML template with charset and title
  - Serves on port 6969

- **server/pages/**: Markdown content files embedded into the binary

## Development Commands

### Build and Run
```bash
# Build the server
go build -o app ./server/server.go

# Run locally
go run server/server.go
```

### Docker
```bash
# Build Docker image (multi-stage: golang:1.23-alpine -> alpine:latest)
docker build -t ilia-fi .

# Run with Docker Compose (includes Caddy reverse proxy)
docker-compose up
```

The Docker setup uses:
- Multi-stage build to minimize image size
- Caddy as reverse proxy on ports 80/443
- External network named "web"

### Testing
Access the server at http://localhost:6969

## Deployment

Automatic deployment via GitHub Actions (.github/workflows/deploy_image.yml):
- Triggers on push to main branch or manual workflow dispatch
- Builds Docker image and pushes to GitHub Container Registry (ghcr.io)
- Image tagged as `ghcr.io/hedgeho/ilia.fi:latest`