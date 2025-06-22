# JustChat

JustChat is a lightweight real-time chat server written in Go. It serves a web UI using server-rendered HTML and HTMX and exposes a REST + WebSocket API for a mobile client.

## Running

```bash
go run ./cmd/server
```

The server listens on port 8080 by default and stores data in a local `justchat.db` SQLite file.

