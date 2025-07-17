# Greenlight

Greenlight is a minimal web RESTful application built to practice backend development, system architecture, and deployment workflow. While the business logic is simple, the project focuses on clean code structure, Docker-based environment management, and production-ready deployment practices.

## Stack and Tools

- **Language**: Go (Golang)
- **Routing**: [`httprouter`](https://github.com/julienschmidt/httprouter)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Containerization**: Docker, Docker Compose
- **Environment Configuration**: `.env`, `envconfig`
- **Deployment**: EC2 (Ubuntu), managed via systemd and Makefile
- **Build Tools**: Makefile, Bash scripts
- **Observability**: Structured logging, trace ID injection

## Features

- Modular project structure (`cmd`, `internal`, `migrations`)
- Graceful shutdown, timeouts, and context propagation
- Channel-based background workers
- Docker-based local development with isolated DB
- Production-ready configuration with `.env` and secrets management
- Zero-downtime deploys via `make reload`

## Architecture Overview

This is the production and staging architecture used to deploy Greenlight. The infrastructure includes a reverse proxy, separated environments, and cloud services for storage, email, and monitoring.

```mermaid
flowchart LR
    subgraph PublicInternet
        User
    end

    subgraph EntryPoint["Edge Layer"]
        Cloudflare["Cloudflare (TLS, DNS, CDN)"]
    end

    subgraph AWS
        subgraph EC2 Instance
            Caddy[[Reverse Proxy]]
            stage["Stage Server (port 4000)"]
            production["Production Server (port 4001) "]
        end
        subgraph Monitoring["Monitoring"]
            cloudwatch[CloudWatch]
        end

        subgraph Email["Email Service"]
            ses[AWS SES]
        end
    end

    subgraph DB
        db[(RDS PostgreSQL)]
    end

    subgraph R2
        r2[(Cloudflare R2)]
    end

    PublicInternet -- "TCP 443 (HTTPS)" --> Cloudflare
    Cloudflare -- "stage.domain" --> Caddy
    Cloudflare --"production.domain" --> Caddy
    Caddy -- "stage.domain → :4000" --> stage
    Caddy -- "prod.domain → :4001" --> production
    stage -- "5432 (Stage DB Conn)" --> db
    production -- "5432 (Prod DB Conn)" --> db
    stage -."R2 API".-> r2
    production -."R2 API".-> r2
    stage -.-> ses
    production -.-> ses
