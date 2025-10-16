# Me (profile) HTTP Service

This small Go HTTP service exposes a single endpoint `/me` which returns a JSON profile including a cat fact fetched from an external API.

## Requirements

- Go 1.25 (go 1.25.1 is used in this project)
- Network access to `https://catfact.ninja/fact` to fetch cat facts

## Setup (install dependencies)

This project uses the standard Go toolchain and modules. From the repository root (where `go.mod` lives) run:

```bash
# ensure you're in the project root
cd /path/to/project/root

# download modules
go mod download
```

No other external package managers are required.

## Run locally

To build and run the server locally:

```bash
# build
go build -o me ./src

# run (or simply `go run ./src`)
./me

# or
go run ./src
```

The server listens on port 8080 by default. Open `http://localhost:8080/me` in your browser or use curl:

```bash
curl http://localhost:8080/me
```

You should receive a JSON response containing `status`, `user`, `timestamp` and `fact`.

## Project structure

- `go.mod` - module and Go version
- `src/main.go` - application entrypoint (starts HTTP server)
- `api/` - handler that builds the response
- `data/` - code that fetches the cat fact from the external API

