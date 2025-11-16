# Build Instructions

## Quick Start

Build the server:
```bash
make build
# or
go build -o bin/server cmd/server/main.go
```

Run the server:
```bash
make run
# or
./bin/server
```

## Server Entry Point

The main server entry point is:
- **`cmd/server/main.go`** - Local development server with fx dependency injection

## Build Output

The compiled binary will be in:
- **`bin/server`** - Executable server binary

## Project Structure

```
cmd/server/main.go  # Main entry point
app/options/app.go  # Fx application setup
internal/server/    # HTTP server implementation
graph/              # GraphQL resolvers
app/workflow/       # Workflow layer
app/controllers/    # Controller layer
app/data/           # Data access layer
```

## Dependencies

- Go 1.24+
- Uber Fx for dependency injection
- gqlgen for GraphQL

Install dependencies:
```bash
go mod download
```

## Testing the Build

```bash
# Build
go build -o bin/server cmd/server/main.go

# Run
./bin/server

# Server will start on http://localhost:8080
# GraphQL Playground: http://localhost:8080/
# GraphQL Endpoint: http://localhost:8080/query
```

