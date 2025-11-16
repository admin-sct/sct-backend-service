# SCT Backend Service

A GraphQL server built with Go using the builder pattern for flexible configuration.

## Features

- 🚀 GraphQL server powered by [gqlgen](https://github.com/99designs/gqlgen)
- 🏗️ Builder pattern for server configuration
- 🎮 GraphQL Playground for interactive testing
- ⚡ Graceful shutdown handling

## Prerequisites

- Go 1.24 or higher
- Make (optional, for convenience commands)

## Getting Started

### 1. Install Dependencies

```bash
go mod download
```

### 2. Generate GraphQL Code

```bash
make generate
# or
go run github.com/99designs/gqlgen generate
```

### 3. Run the Server

For local development and debugging:

```bash
make run
# or
go run cmd/server/main.go
```

Or with custom options:

```bash
go run cmd/server/main.go -port 3000 -debug=true -playground=true
```

Available flags:
- `-port`: Server port (default: 8080)
- `-host`: Server host (default: 0.0.0.0)
- `-debug`: Enable debug mode (default: false)
- `-playground`: Enable GraphQL playground (default: true)

The server will start on `http://localhost:8080` with:
- GraphQL endpoint: `http://localhost:8080/query`
- GraphQL Playground: `http://localhost:8080/`

## Vercel Deployment

This project is configured for Vercel serverless deployment. See [VERCEL.md](./VERCEL.md) for detailed deployment instructions.

**Quick Deploy:**
```bash
npm install -g vercel
vercel login
vercel
```

**Endpoints on Vercel:**
- GraphQL: `https://your-project.vercel.app/api/graphql`
- Playground: `https://your-project.vercel.app/api/playground`

## Architecture

This project follows a layered architecture with the builder pattern:

1. **Options Layer** (`options/`) - Configuration builders for server, controller, and workflow options
2. **Controller Layer** (`controller/`) - Handles GraphQL requests and delegates to workflows
3. **Workflow Layer** (`workflow/`) - Contains business logic for user operations
4. **Utils Layer** (`utils/`) - Utility functions for ID generation, validation, etc.
5. **Graph Layer** (`graph/`) - GraphQL resolvers that connect controllers to the schema
6. **Server Layer** (`internal/server/`) - HTTP server setup and configuration

## Usage

### Builder Pattern Examples

#### Server Options Builder

```go
opts, err := options.NewBuilder().
    WithPort(8080).
    WithHost("0.0.0.0").
    WithDebugMode(true).
    WithPlayground(true).
    WithLogLevel("debug").
    Build()
```

#### Controller Builder (via Options)

```go
// Build workflow
userWorkflow, err := options.BuildWorkflow()

// Create controller options
controllerOpts := &controller.ControllerOptions{
    UserWorkflow: userWorkflow,
}

// Create controller
ctrl := controller.NewController(controllerOpts)
```

#### Server Builder

```go
srv, err := server.NewServerBuilder().
    WithPort(8080).
    WithHost("0.0.0.0").
    WithReadTimeout(15 * time.Second).
    WithWriteTimeout(15 * time.Second).
    WithPlayground(true).
    WithGraphQLPath("/query").
    WithResolvers(ctrl.GetResolver()).
    Build()
```

### Available Builder Methods

**Options Builder:**
- `WithPort(int)` - Set server port (default: 8080)
- `WithHost(string)` - Set server host (default: "0.0.0.0")
- `WithReadTimeout(time.Duration)` - Set read timeout (default: 15s)
- `WithWriteTimeout(time.Duration)` - Set write timeout (default: 15s)
- `WithIdleTimeout(time.Duration)` - Set idle timeout (default: 60s)
- `WithPlayground(bool)` - Enable/disable playground (default: true)
- `WithPlaygroundPath(string)` - Set playground path (default: "/")
- `WithGraphQLPath(string)` - Set GraphQL endpoint path (default: "/query")
- `WithDebugMode(bool)` - Enable debug mode (default: false)
- `WithLogLevel(string)` - Set log level (default: "info")

## GraphQL Schema

### Queries

```graphql
query {
  hello
  user(id: "1") {
    id
    name
    email
    createdAt
  }
  users {
    id
    name
    email
    createdAt
  }
}
```

### Mutations

```graphql
mutation {
  createUser(input: { name: "John Doe", email: "john@example.com" }) {
    id
    name
    email
    createdAt
  }
  
  updateUser(id: "1", input: { name: "Jane Doe" }) {
    id
    name
    email
  }
  
  deleteUser(id: "1")
}
```

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go              # Server entry point for local development/debugging
├── options/
│   ├── builder.go               # Options builder pattern (ServerOptions, ControllerOptions, WorkflowOptions)
│   └── server.go                # Local server setup helpers
├── controller/
│   └── controller.go            # Controller layer (handles GraphQL requests, delegates to workflows)
├── workflow/
│   └── user_workflow.go         # Business logic layer (handles user operations)
├── utils/
│   ├── id.go                    # ID generation utilities
│   └── validation.go            # Validation utilities
├── graph/
│   ├── schema.graphqls          # GraphQL schema definition
│   ├── resolver.go              # Resolver struct and dependencies
│   ├── schema.resolvers.go      # Resolver implementations
│   ├── generated/               # Generated code (auto-generated)
│   └── model/                   # Generated models (auto-generated)
├── internal/
│   └── server/
│       └── builder.go           # HTTP server builder pattern implementation
├── go.mod                       # Go module definition
├── gqlgen.yml                   # gqlgen configuration
└── README.md                    # This file
```

## Development

### Regenerate GraphQL Code

After modifying `graph/schema.graphqls`, regenerate the code:

```bash
go run github.com/99designs/gqlgen generate
```

### Running Tests

```bash
go test ./...
```

## License

MIT

