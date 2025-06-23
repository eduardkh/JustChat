# JustChat

JustChat is a lightweight real-time chat server written in Go. It serves a web UI using server-rendered HTML and HTMX and exposes a REST + WebSocket API for a mobile client.

## Running

```bash
go run ./cmd/server
```

The server listens on port `8080` by default and stores data in a local `justchat.db` SQLite file.

## Web UI

After starting the server, open `http://localhost:8080/` in your browser to access the web interface. You can register a new user using the **Register** link on the home page.

The HTML templates are stored in the `templates/` directory and rendered by Echo.
To build a standalone binary:
```bash
go build -o justchat ./cmd/server
```
Run the resulting binary and open `http://localhost:8080/` to access the UI.
Successful login redirects you to `/chat`, a placeholder chat screen.
