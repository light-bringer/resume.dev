# Quick Start Guide

Get your resume website running in minutes!

## Option 1: Docker (Recommended - Fastest!)

**Prerequisites**: Docker installed

### 1. Start Everything

```bash
make docker-up
```

### 2. View Your Resume

Open `http://localhost:3000` in your browser.

That's it! 🎉

**Only the frontend is exposed** to your host machine. The backend runs internally and is not accessible from outside.

### Docker Commands

```bash
make docker-up       # Start all services
make docker-logs     # View logs
make docker-down     # Stop all services
make docker-rebuild  # Rebuild after code changes
```

**See [DOCKER.md](DOCKER.md) for detailed Docker documentation.**

---

## Option 2: Local Development

### 1. Install Dependencies

```bash
make install
```

This installs:
- Go modules for backend
- npm packages for frontend

### 2. Run Development Servers

```bash
make dev
```

This starts:
- Backend API at `http://localhost:8080`
- Frontend at `http://localhost:3000`

### 3. View Your Resume

Open `http://localhost:3000` in your browser.

## Deployment

### Quick Deploy (Recommended)

1. **Push to GitHub**
   ```bash
   git add .
   git commit -m "Initial deployment"
   git push origin main
   ```

2. **Set up Google Cloud Run** (5 minutes)
   - Create GCP project
   - Enable Cloud Run API
   - Create service account
   - Add GitHub secrets (see DEPLOYMENT.md)

3. **Set up Vercel** (2 minutes)
   - Connect GitHub repo to Vercel
   - Add environment variable `API_URL`
   - Deploy

### Automatic Deployments

Once set up, every push to `main` automatically deploys:
- Backend changes → Google Cloud Run
- Frontend changes → Vercel

## Updating Resume Content

1. Edit `backend/data/resume_data.go`
2. Commit and push:
   ```bash
   git add backend/data/resume_data.go
   git commit -m "Update resume"
   git push origin main
   ```
3. Wait 2-3 minutes for automatic deployment

## Common Commands

```bash
# Development
make dev              # Run both services
make backend          # Run backend only
make frontend         # Run frontend only

# Building
make build            # Build both projects
make clean            # Clean build artifacts

# Testing
cd backend && make test           # Test backend
cd frontend && npm run lint       # Lint frontend
cd frontend && npm run build      # Build frontend
```

## Troubleshooting

### Backend won't start
- Check if port 8080 is available
- Run `cd backend && go mod tidy`

### Frontend won't start
- Check if port 3000 is available
- Run `cd frontend && rm -rf node_modules && npm install`

### Can't connect to API
- Ensure backend is running on port 8080
- Check `frontend/.env.local` has `API_URL=http://localhost:8080`

## Next Steps

- 📖 Read [DEPLOYMENT.md](DEPLOYMENT.md) for detailed deployment instructions
- 🏗️ Read [CLAUDE.md](CLAUDE.md) for architecture and development guidelines
- 📝 Read [README.md](README.md) for project documentation

## Free Tier Limits

Both platforms are free for personal projects:

**Google Cloud Run:**
- 2 million requests/month
- Always free if traffic is low

**Vercel:**
- Unlimited deployments
- 100 GB bandwidth/month

Your resume website will easily stay within these limits!
