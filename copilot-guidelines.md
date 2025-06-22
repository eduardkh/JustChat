# JustChat Copilot Coding Guidelines

## Project Overview

JustChat is a Go-based real-time chat server with both a lightweight web UI and an Android client. It uses WebSockets for real-time messaging, REST for mobile client interactions, and server-rendered HTML + HTMX for the web interface.

## Stack

- Go (Echo framework)
- WebSockets for realtime communication
- SQLite (start) or Postgres (optional)
- JWT for stateless authentication
- Android Client: Kotlin + WebSocket + Retrofit
- Web UI: Echo templates (Go's `html/template`) + HTMX

## Code Style

- Idiomatic Go (`go fmt`, `go vet`, `golangci-lint`)
- Clean, layered architecture:
  - `handlers/` → `services/` → `storage/`
- Use interfaces for database and external systems
- Make all request-scoped logic context-aware
- Avoid global state

## Security

- Passwords must be hashed with bcrypt
- JWT tokens with expiration and refresh tokens
- CORS configuration for API endpoints
- Rate limiting middleware per IP/user
- CSRF protection for HTMX routes (token-based)

## Web UI (HTMX + Echo)

- Use Go `html/template` engine
- Minimal JS: use HTMX for AJAX/partial updates
- Pages:
  - Login/Register
  - Chat List
  - Single Chat View (HTMX swap)
  - Settings/Profile

## Android Client

- Kotlin preferred
- Retrofit for REST
- OkHttp/WebSocket for live chat
- FCM optional for push notifications

## Prompting Copilot
