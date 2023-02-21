# FileShare
A project for me to get more experience combining traditional server side rendering with a build step to generate static
frontend assets.

## Description
FileShare allows you to deploy a self-contained binary that serves as a file sharing host. It's designed to be invite
only for trusted friends and family. Unique links are generated for each file uploaded and they're assigned to
individual users.

## Installation
Copy `.env.example` to .env or setup matching environment variables to configure the service. You'll need a working
PostgreSQL database. There's a `docker-compose.yml` file included for convenience.


## Building and running

#### Configuration
Follow the installation steps to configure your environment variables.

#### Build
```bash
go build -o fileshare
```

First you'll want to generate a token, since it's invite only.
```bash
./fileshare invite create 
```

Run the service itself.
```bash
./fileshare
```

Login and create a new user with the token you generated.

`http(s)://<host>:<port>/login?invite_token=<token>`

Promote your new user.
```bash
./fileshare user promote
```

## Development
Everything in static_generated is built from the frontend directory. The frontend directory is built using Vite and 
Typescript. The backend is built using Go and Chi.

### Requirements
- [Go 1.20+](https://go.dev)
- [Node 18+](https://nodejs.org/en/)
- [pnpm](https://pnpm.io/)
- [PostgreSQL 14+](https://www.postgresql.org/) or [Docker](https://www.docker.com/)
- [sql-migrate](https://github.com/rubenv/sql-migrate)
- [sqlboiler](https://github.com/volatiletech/sqlboiler)