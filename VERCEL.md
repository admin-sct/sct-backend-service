# Vercel Deployment Guide

This project is configured to deploy on Vercel as a serverless GraphQL API.

## Structure

- **`api/index.go`** - Vercel serverless function handler
- **`vercel.json`** - Vercel configuration
- **`cmd/server/main.go`** - Local development server (not used in Vercel)

## Deployment

### Prerequisites

1. Install Vercel CLI:
```bash
npm install -g vercel
```

2. Login to Vercel:
```bash
vercel login
```

### Deploy

```bash
vercel
```

Or for production:
```bash
vercel --prod
```

## Endpoints

Once deployed, your GraphQL API will be available at:

- **GraphQL Endpoint**: `https://your-project.vercel.app/api/graphql`
- **GraphQL Playground**: `https://your-project.vercel.app/api/playground`
- **Query Endpoint**: `https://your-project.vercel.app/api/query`

## Environment Variables

Set these in the Vercel dashboard under Project Settings → Environment Variables:

- `SLACK_WEBHOOK_URL` (if you want to override the default in `app/keys/slack.go`)

## Local Development

For local development, use:

```bash
make run
# or
go run cmd/server/main.go
```

This uses the full server setup with fx dependency injection.

## Notes

- The `api/index.go` handler initializes dependencies on first request (singleton pattern)
- Vercel will automatically detect Go from `go.mod`
- The handler routes requests based on path:
  - `/api/playground` → GraphQL Playground
  - `/api/graphql` or `/api/query` → GraphQL endpoint
  - Default → GraphQL endpoint

