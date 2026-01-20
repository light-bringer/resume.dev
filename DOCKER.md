# Docker Setup Guide

Run the entire resume system using Docker Compose with a single command!

## Overview

The Docker Compose setup includes:
- **Backend API** (Go) - Internal network only, not exposed to host
- **Frontend** (Next.js) - Exposed on port 3000
- Automatic health checks
- Internal Docker networking
- Auto-restart on failure

## Prerequisites

- Docker installed (version 20.10+)
- Docker Compose installed (version 2.0+)

### Install Docker

**macOS:**
```bash
brew install --cask docker
```

**Linux:**
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```

**Windows:**
Download Docker Desktop from https://www.docker.com/products/docker-desktop

## Quick Start

### 1. Build and Start All Services

```bash
docker-compose up --build
```

This will:
- Build the backend Docker image
- Build the frontend Docker image
- Start both services
- Only expose frontend on port 3000

### 2. Access the Application

Open your browser to:
```
http://localhost:3000
```

The frontend will automatically connect to the backend via the internal Docker network.

### 3. Stop All Services

Press `Ctrl+C` in the terminal, then:

```bash
docker-compose down
```

## Docker Compose Commands

### Start Services (Background)

```bash
docker-compose up -d
```

### View Logs

```bash
# All services
docker-compose logs -f

# Backend only
docker-compose logs -f backend

# Frontend only
docker-compose logs -f frontend
```

### Stop Services

```bash
docker-compose stop
```

### Restart Services

```bash
docker-compose restart
```

### Remove Containers and Networks

```bash
docker-compose down
```

### Remove Everything (Including Images)

```bash
docker-compose down --rmi all --volumes
```

### Rebuild Services

```bash
# Rebuild all
docker-compose build --no-cache

# Rebuild specific service
docker-compose build --no-cache backend
docker-compose build --no-cache frontend
```

### Check Service Status

```bash
docker-compose ps
```

### Run Commands in Containers

```bash
# Execute command in backend
docker-compose exec backend sh

# Execute command in frontend
docker-compose exec frontend sh
```

## Architecture

```
┌─────────────────────────────────────┐
│           Host Machine               │
│                                      │
│   Port 3000 ←→ Frontend Container   │
│                     ↓                │
│                     ↓                │
│              Internal Network        │
│                     ↓                │
│              Backend Container       │
│              (Not Exposed)           │
└─────────────────────────────────────┘
```

### Network Configuration

- **Network Name**: `resume-network` (bridge driver)
- **Backend**: Accessible at `http://backend:8080` from frontend
- **Frontend**: Accessible at `http://localhost:3000` from host
- **Backend**: NOT accessible from host (internal only)

## Service Details

### Backend Service

- **Container Name**: `resume-api`
- **Internal Port**: 8080
- **Host Access**: None (internal only)
- **Health Check**: `http://localhost:8080/health`
- **Restart Policy**: unless-stopped

### Frontend Service

- **Container Name**: `resume-frontend`
- **Internal Port**: 3000
- **Host Port**: 3000
- **Environment**: `API_URL=http://backend:8080`
- **Health Check**: `http://localhost:3000`
- **Restart Policy**: unless-stopped

## Environment Variables

Default environment variables are set in `docker-compose.yml`:

```yaml
Backend:
  - PORT=8080

Frontend:
  - API_URL=http://backend:8080
  - NODE_ENV=production
```

### Custom Environment Variables

Create a `.env` file in the root directory:

```env
# Override default ports if needed
FRONTEND_PORT=3001
```

Then reference in docker-compose.yml:
```yaml
ports:
  - "${FRONTEND_PORT:-3000}:3000"
```

## Troubleshooting

### Port Already in Use

If port 3000 is already in use:

```bash
# Find process using port 3000
lsof -i :3000

# Kill the process
kill -9 <PID>

# Or change the port in docker-compose.yml
ports:
  - "3001:3000"
```

### Services Won't Start

Check logs:
```bash
docker-compose logs backend
docker-compose logs frontend
```

Common issues:
1. **Build failed**: Check Dockerfile syntax
2. **Health check failed**: Service might be slow to start
3. **Network error**: Backend not accessible from frontend

### Backend Not Accessible from Frontend

Verify network configuration:
```bash
docker network inspect resume-dev_resume-network
```

Check frontend logs:
```bash
docker-compose logs frontend | grep "API_URL"
```

### Rebuild After Code Changes

```bash
# Stop services
docker-compose down

# Rebuild and start
docker-compose up --build
```

### Clear Docker Cache

If builds are stuck or using old cache:

```bash
# Remove all stopped containers
docker container prune

# Remove all unused images
docker image prune -a

# Remove all unused volumes
docker volume prune

# Nuclear option - remove everything
docker system prune -a --volumes
```

### Check Health Status

```bash
# View detailed status
docker-compose ps

# Check health of specific service
docker inspect --format='{{.State.Health.Status}}' resume-api
docker inspect --format='{{.State.Health.Status}}' resume-frontend
```

## Development Workflow

### 1. Make Code Changes

Edit files in `backend/` or `frontend/`

### 2. Rebuild Affected Service

```bash
# Backend changes
docker-compose up -d --build backend

# Frontend changes
docker-compose up -d --build frontend
```

### 3. View Logs

```bash
docker-compose logs -f
```

### 4. Test Changes

Open `http://localhost:3000`

## Production Considerations

This Docker Compose setup is designed for **local development and testing**.

For production deployment, use:
- **Backend**: Google Cloud Run (see [DEPLOYMENT.md](DEPLOYMENT.md))
- **Frontend**: Vercel (see [DEPLOYMENT.md](DEPLOYMENT.md))

### Why Not Use Docker Compose in Production?

1. **No auto-scaling**: Fixed resources
2. **Single point of failure**: One server
3. **No load balancing**: Cannot distribute traffic
4. **Manual updates**: No CI/CD integration
5. **Limited monitoring**: Basic health checks only

Cloud Run and Vercel provide:
- Auto-scaling
- Global distribution
- Built-in monitoring
- Automatic deployments
- Free tiers

## Resource Usage

Monitor resource usage:

```bash
# Real-time stats
docker stats

# Specific containers
docker stats resume-api resume-frontend
```

### Typical Resource Usage

**Backend:**
- CPU: ~5% idle, ~20% under load
- Memory: ~20 MB

**Frontend:**
- CPU: ~10% idle, ~30% under load
- Memory: ~80 MB

## Advanced Configuration

### Custom Dockerfile Build Args

Add build arguments to docker-compose.yml:

```yaml
services:
  backend:
    build:
      context: ./backend
      args:
        GO_VERSION: 1.21
```

### Volume Mounts for Development

Mount code directories for live reload:

```yaml
services:
  backend:
    volumes:
      - ./backend:/app
  frontend:
    volumes:
      - ./frontend:/app
      - /app/node_modules
```

**Note**: This is for development only. Remove for production.

### Multiple Compose Files

Create separate configs:

```bash
# Base config
docker-compose.yml

# Development overrides
docker-compose.dev.yml

# Use both
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up
```

## Testing the Setup

### 1. Build Images

```bash
docker-compose build
```

### 2. Start Services

```bash
docker-compose up -d
```

### 3. Wait for Health Checks

```bash
# Wait up to 30 seconds for healthy status
for i in {1..30}; do
  STATUS=$(docker inspect --format='{{.State.Health.Status}}' resume-api)
  if [ "$STATUS" = "healthy" ]; then
    echo "Backend is healthy!"
    break
  fi
  sleep 1
done
```

### 4. Test Backend Internally

```bash
docker-compose exec frontend wget -qO- http://backend:8080/health
```

### 5. Test Frontend

```bash
curl http://localhost:3000
```

### 6. Verify Backend is Not Exposed

This should fail:
```bash
curl http://localhost:8080/health
# Connection refused (expected)
```

## Cleanup

### Remove Everything

```bash
# Stop and remove containers, networks
docker-compose down

# Also remove images
docker-compose down --rmi all

# Nuclear option - remove all Docker resources
docker system prune -a --volumes
```

## Integration with Makefile

Add to root `Makefile`:

```makefile
.PHONY: docker-up docker-down docker-logs docker-rebuild

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-rebuild:
	docker-compose up -d --build
```

Usage:
```bash
make docker-up
make docker-logs
make docker-down
```

## Comparison: Docker vs Local Development

| Aspect | Docker Compose | Local Development |
|--------|---------------|-------------------|
| Setup Time | 5 minutes | 2 minutes |
| Isolation | ✅ Complete | ❌ Shared system |
| Port Conflicts | ✅ Rare | ⚠️ Common |
| Consistency | ✅ Same everywhere | ⚠️ System dependent |
| Hot Reload | ❌ Need rebuild | ✅ Automatic |
| Resource Usage | ⚠️ Higher | ✅ Lower |
| Production Parity | ✅ High | ⚠️ Lower |

**Recommendation**: Use Docker Compose for testing production-like environments, use local development for active coding.

## Quick Reference

```bash
# Start everything
docker-compose up -d

# View logs
docker-compose logs -f

# Stop everything
docker-compose down

# Rebuild after changes
docker-compose up -d --build

# Check status
docker-compose ps

# Access application
open http://localhost:3000
```

## Next Steps

- ✅ System running locally with Docker
- 📖 Read [DEPLOYMENT.md](DEPLOYMENT.md) for production deployment
- 🏗️ Read [CLAUDE.md](CLAUDE.md) for architecture details
- 🚀 Deploy to Cloud Run and Vercel for production
