# resume.dev

A modern developer resume website with a Go API backend and Next.js frontend.

## Project Structure

```
.
├── backend/          # Go API server
│   ├── main.go      # HTTP server and routes
│   ├── models/      # Data models
│   ├── data/        # Resume data
│   └── Makefile     # Backend build commands
├── frontend/         # Next.js frontend
│   ├── app/         # Next.js app router pages
│   ├── components/  # React components
│   ├── types/       # TypeScript types
│   └── package.json # Frontend dependencies
└── Makefile         # Root-level commands
```

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn

### Installation

Install all dependencies:

```bash
make install
```

### Development

Run both backend and frontend:

```bash
make dev
```

Or run them separately:

```bash
# Terminal 1: Backend API (http://localhost:8080)
make backend

# Terminal 2: Frontend (http://localhost:3000)
make frontend
```

### Building

Build both projects:

```bash
make build
```

## Docker Setup (Recommended)

Run the entire system with Docker Compose - **only the frontend is exposed to the host**.

### Quick Docker Start

```bash
# Start all services
make docker-up

# View logs
make docker-logs

# Stop all services
make docker-down
```

**Access**: Open `http://localhost:3000`

The backend runs internally and is only accessible to the frontend container.

### Docker Commands

```bash
make docker-up       # Start all services (detached)
make docker-down     # Stop all services
make docker-logs     # View live logs
make docker-rebuild  # Rebuild and restart
make docker-clean    # Remove all containers and images
```

**For detailed Docker documentation**, see [DOCKER.md](DOCKER.md).

### Why Use Docker?

- ✅ **No local dependencies** - Only Docker required
- ✅ **Production-like environment** - Same as deployment
- ✅ **Isolated services** - No port conflicts
- ✅ **One command setup** - `make docker-up`
- ✅ **Backend not exposed** - Secure internal networking

## Backend API

The Go backend serves resume data via a REST API.

**Endpoints:**
- `GET /api/resume` - Returns complete resume data
- `GET /health` - Health check endpoint

**Development:**

```bash
cd backend
make run          # Run the API
make test         # Run tests
make build        # Build binary
```

## Frontend

The Next.js frontend displays the resume with a modern, responsive design.

**Development:**

```bash
cd frontend
npm install       # Install dependencies
npm run dev       # Start dev server
npm run build     # Build for production
npm start         # Start production server
```

### Environment Variables

Create `frontend/.env.local`:

```env
API_URL=http://localhost:8080
```

## Deployment

This project includes automated CI/CD pipelines for deploying to Google Cloud Run (backend) and Vercel (frontend).

### Quick Deploy

See [QUICKSTART.md](QUICKSTART.md) for the fastest way to get deployed.

### Automated CI/CD

The project uses GitHub Actions for continuous deployment:

**On every push to main:**
- Backend changes automatically deploy to Google Cloud Run
- Frontend changes automatically deploy to Vercel
- All changes are tested before deployment

### Backend (Google Cloud Run)

Free tier: 2 million requests/month

**Features:**
- Containerized with Docker
- Auto-scales to zero when idle
- Global availability
- Automatic HTTPS

**Manual deployment:**
```bash
cd backend
gcloud builds submit --tag gcr.io/PROJECT_ID/resume-api
gcloud run deploy resume-api --image gcr.io/PROJECT_ID/resume-api --region us-central1
```

### Frontend (Vercel)

Free tier: Unlimited deployments, 100GB bandwidth/month

**Features:**
- Global CDN
- Automatic HTTPS
- Preview deployments for PRs
- Zero-config deployment

**Manual deployment:**
```bash
cd frontend
vercel --prod
```

### Setup Instructions

For detailed CI/CD setup instructions, see [DEPLOYMENT.md](DEPLOYMENT.md).

Required GitHub secrets:
- `GCP_PROJECT_ID`, `GCP_SA_KEY` (for Cloud Run)
- `VERCEL_TOKEN`, `VERCEL_ORG_ID`, `VERCEL_PROJECT_ID`, `API_URL` (for Vercel)

## Tech Stack

**Backend:**
- Go 1.21+
- Standard library HTTP server
- CORS enabled for frontend requests

**Frontend:**
- Next.js 15 (App Router)
- React 19
- TypeScript
- Tailwind CSS
- Server-side rendering

## License

MIT
