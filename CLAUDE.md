# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

resume.dev is a developer resume application with a Go API backend and Next.js frontend. The application displays a professional resume with data served from a RESTful API.

## Tech Stack

**Backend:**
- Go 1.21+ with standard library HTTP server
- CORS enabled for frontend requests
- JSON API responses

**Frontend:**
- Next.js 15 with App Router
- React 19 with Server Components
- TypeScript for type safety
- Tailwind CSS for styling
- Configured for Vercel deployment

## Project Structure

```
.
├── backend/              # Go API server
│   ├── main.go          # HTTP server, routes, CORS middleware
│   ├── models/          # Data structures (Resume, Experience, Education, etc.)
│   │   └── resume.go
│   ├── data/            # Resume data
│   │   └── resume_data.go
│   ├── go.mod           # Go module definition
│   └── Makefile         # Backend build commands
│
├── frontend/            # Next.js application
│   ├── app/            # Next.js App Router
│   │   ├── layout.tsx  # Root layout
│   │   ├── page.tsx    # Home page (fetches and displays resume)
│   │   └── globals.css # Global styles
│   ├── components/     # React components
│   │   ├── Header.tsx
│   │   ├── Summary.tsx
│   │   ├── Experience.tsx
│   │   ├── Skills.tsx
│   │   ├── Education.tsx
│   │   ├── OpenSource.tsx
│   │   └── Interests.tsx
│   ├── types/          # TypeScript type definitions
│   │   └── resume.ts
│   ├── package.json
│   ├── tsconfig.json
│   ├── tailwind.config.ts
│   └── next.config.ts
│
├── Makefile            # Root-level commands for running both services
├── README.md           # Project documentation
└── CLAUDE.md          # This file
```

## Common Development Commands

### Full Stack Development

Run both backend and frontend together:
```bash
make install    # Install all dependencies
make dev        # Run both services (backend on :8080, frontend on :3000)
make build      # Build both projects
make clean      # Clean build artifacts
```

### Backend Only

```bash
cd backend
make run        # Run the Go API server (port 8080)
make build      # Build the binary (creates resume-api)
make test       # Run tests
make fmt        # Format Go code
make deps       # Download and tidy dependencies
```

Backend runs on `http://localhost:8080`

### Frontend Only

```bash
cd frontend
npm install     # Install dependencies
npm run dev     # Start development server (port 3000)
npm run build   # Build for production
npm start       # Start production server
npm run lint    # Run ESLint
```

Frontend runs on `http://localhost:3000`

### Docker Compose

Run the entire system with Docker (recommended for testing):
```bash
make docker-up       # Start all services (only frontend exposed on :3000)
make docker-down     # Stop all services
make docker-logs     # View logs
make docker-rebuild  # Rebuild and restart after code changes
make docker-clean    # Remove all containers and images
```

**Architecture:**
- Backend runs internally (not exposed to host)
- Frontend exposed on port 3000
- Internal Docker network for service communication
- Frontend connects to backend at `http://backend:8080`

See [DOCKER.md](DOCKER.md) for comprehensive Docker documentation.

## API Architecture

### Endpoints

- `GET /api/resume` - Returns complete resume data as JSON
- `GET /health` - Health check endpoint

### Data Model

The resume data is structured with the following models (see backend/models/resume.go):
- `Resume` - Top-level structure containing all resume sections
- `PersonalInfo` - Contact and personal information
- `Experience` - Work experience entries with achievements
- `Education` - Educational qualifications
- `TechnicalSkills` - Grouped technical competencies
- `OpenSource` - Open source projects and contributions

All data is currently static and defined in `backend/data/resume_data.go`. To update resume content, modify this file.

## Frontend Architecture

- Uses Next.js App Router with Server Components for optimal performance
- Data fetching happens server-side in `app/page.tsx`
- Resume sections are split into focused, reusable components
- Tailwind CSS for responsive, mobile-first design
- TypeScript types mirror the Go backend models

## Environment Variables

### Frontend

Create `frontend/.env.local` for local development:
```env
API_URL=http://localhost:8080
```

For production (Vercel), set:
- `API_URL` - URL of the deployed Go API backend

### Backend

The Go API reads:
- `PORT` - HTTP server port (defaults to 8080)

## CI/CD Pipeline

The project uses GitHub Actions for automated deployments:

### Workflows

1. **CI Workflow** (`.github/workflows/ci.yml`)
   - Runs on all pushes and PRs
   - Tests backend Go code
   - Tests frontend TypeScript code
   - Validates Docker build

2. **Backend Deployment** (`.github/workflows/deploy-backend.yml`)
   - Triggers on push to main (backend changes only)
   - Builds Docker image
   - Deploys to Google Cloud Run
   - Runs health checks

3. **Frontend Deployment** (`.github/workflows/deploy-frontend.yml`)
   - Triggers on push to main (frontend changes only)
   - Builds Next.js application
   - Deploys to Vercel
   - Validates deployment

### Required GitHub Secrets

For Cloud Run backend:
- `GCP_PROJECT_ID` - Google Cloud project ID
- `GCP_SA_KEY` - Service account JSON key

For Vercel frontend:
- `VERCEL_TOKEN` - Vercel authentication token
- `VERCEL_ORG_ID` - Vercel organization ID
- `VERCEL_PROJECT_ID` - Vercel project ID
- `API_URL` - Backend API URL (Cloud Run URL)
- `VERCEL_DOMAIN` - Your Vercel domain for testing

See `DEPLOYMENT.md` for complete setup instructions.

## Deployment

### Backend (Google Cloud Run)

Deploy to Google Cloud Run (free tier: 2M requests/month):
- Containerized with Docker
- Auto-scales to zero when idle
- 256Mi memory, 1 CPU configured
- Deployed to `us-central1` region

Manual deployment:
```bash
cd backend
gcloud builds submit --tag gcr.io/PROJECT_ID/resume-api
gcloud run deploy resume-api --image gcr.io/PROJECT_ID/resume-api --region us-central1
```

Automatic deployment via GitHub Actions on push to main.

### Frontend (Vercel)

Deploy to Vercel (free tier with unlimited deployments):
- Automatic deployments on push to main
- Global CDN distribution
- Automatic HTTPS

Manual deployment:
```bash
cd frontend
vercel --prod
```

Automatic deployment via GitHub Actions on push to main.

## Development Workflow

### Branch Strategy
- Main branch: `main`
- Feature branches follow the pattern: `copilot/create-developer-resume-app`

### Adding New Resume Sections

1. Update Go models in `backend/models/resume.go`
2. Add data in `backend/data/resume_data.go`
3. Update TypeScript types in `frontend/types/resume.ts`
4. Create new component in `frontend/components/`
5. Import and use component in `frontend/app/page.tsx`

### Making Changes to Resume Content

Resume content is in `backend/data/resume_data.go`. Edit this file and restart the backend to see changes.

When deployed, changes pushed to main trigger automatic redeployment of the backend.

## CI/CD Best Practices

### Testing Locally Before Push

Always test changes locally before pushing:
```bash
# Test backend
cd backend && make test && make build

# Test frontend
cd frontend && npm run lint && npm run build

# Test Docker build
cd backend && docker build -t test .
```

### Manual Workflow Triggers

All workflows can be triggered manually from GitHub Actions UI:
1. Go to Actions tab
2. Select the workflow
3. Click "Run workflow"
4. Choose branch and run

### Monitoring Deployments

- **Backend**: Check Cloud Run logs in GCP Console
- **Frontend**: Check deployment logs in Vercel Dashboard
- **CI/CD**: Monitor GitHub Actions for build status

### Cost Management

Both platforms offer generous free tiers:
- **Cloud Run**: Stay within free tier by using min-instances=0
- **Vercel**: Free tier includes unlimited deployments and 100GB bandwidth
