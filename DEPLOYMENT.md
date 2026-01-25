# Deployment Guide

This guide explains how to set up CI/CD for deploying the backend to Google Cloud Run and the frontend to Vercel.

## Overview

- **Backend**: Deployed to Google Cloud Run (free tier available)
- **Frontend**: Deployed to Vercel (free tier available)
- **CI/CD**: GitHub Actions automates deployments on push to main

## Prerequisites

1. GitHub repository with this code
2. Google Cloud Platform account
3. Vercel account
4. `gcloud` CLI installed locally (for initial setup)

## Part 1: Google Cloud Run Setup (Backend)

### 1.1 Create a GCP Project

```bash
# Create a new project
gcloud projects create resume-api-PROJECT_ID --name="Resume API"

# Set the project as default
gcloud config set project resume-api-PROJECT_ID

# Enable required APIs
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com
```

### 1.2 Create a Service Account

```bash
# Create service account
gcloud iam service-accounts create github-actions \
    --display-name="GitHub Actions"

# Grant necessary permissions
gcloud projects add-iam-policy-binding resume-api-PROJECT_ID \
    --member="serviceAccount:github-actions@resume-api-PROJECT_ID.iam.gserviceaccount.com" \
    --role="roles/run.admin"

gcloud projects add-iam-policy-binding resume-api-PROJECT_ID \
    --member="serviceAccount:github-actions@resume-api-PROJECT_ID.iam.gserviceaccount.com" \
    --role="roles/storage.admin"

gcloud projects add-iam-policy-binding resume-api-PROJECT_ID \
    --member="serviceAccount:github-actions@resume-api-PROJECT_ID.iam.gserviceaccount.com" \
    --role="roles/iam.serviceAccountUser"

# Create and download key
gcloud iam service-accounts keys create key.json \
    --iam-account=github-actions@resume-api-PROJECT_ID.iam.gserviceaccount.com
```

### 1.3 Add GitHub Secrets

Go to your GitHub repository → Settings → Secrets and variables → Actions

Add these secrets:

| Secret Name | Value | Description |
|-------------|-------|-------------|
| `GCP_PROJECT_ID` | `resume-api-PROJECT_ID` | Your GCP project ID |
| `GCP_SA_KEY` | Contents of `key.json` | Service account key JSON |

### 1.4 Initial Deployment (Manual)

```bash
# Build and push the first image
cd backend
gcloud builds submit --tag gcr.io/resume-api-PROJECT_ID/resume-api

# Deploy to Cloud Run
gcloud run deploy resume-api \
    --image gcr.io/resume-api-PROJECT_ID/resume-api \
    --platform managed \
    --region us-central1 \
    --allow-unauthenticated \
    --port 8080 \
    --memory 256Mi \
    --cpu 1 \
    --min-instances 0 \
    --max-instances 10

# Get the service URL
gcloud run services describe resume-api \
    --platform managed \
    --region us-central1 \
    --format 'value(status.url)'
```

Save this URL - you'll need it for the frontend configuration.

## Part 2: Vercel Setup (Frontend)

### 2.1 Install Vercel CLI

```bash
npm install -g vercel
```

### 2.2 Login to Vercel

```bash
vercel login
```

### 2.3 Link Project

```bash
cd frontend
vercel link
```

Follow the prompts to create a new project or link to an existing one.

### 2.4 Get Vercel Credentials

```bash
# Get your Vercel org ID and project ID from .vercel/project.json
cat .vercel/project.json
```

Generate a Vercel token:
1. Go to https://vercel.com/account/tokens
2. Create a new token
3. Copy the token

### 2.5 Add GitHub Secrets

Add these secrets to your GitHub repository:

| Secret Name | Value | Description |
|-------------|-------|-------------|
| `VERCEL_TOKEN` | Your Vercel token | Authentication token |
| `VERCEL_ORG_ID` | From `.vercel/project.json` | Organization ID |
| `VERCEL_PROJECT_ID` | From `.vercel/project.json` | Project ID |
| `API_URL` | Cloud Run service URL | Backend API URL |
| `VERCEL_DOMAIN` | Your Vercel domain | For testing (e.g., `your-project.vercel.app`) |

### 2.6 Configure Environment Variables in Vercel

In Vercel dashboard:
1. Go to Project Settings → Environment Variables
2. Add `API_URL` with your Cloud Run URL
3. Set it for Production, Preview, and Development

### 2.7 Initial Deployment (Manual)

```bash
cd frontend
vercel --prod
```

## Part 3: CI/CD Workflows

The repository includes three GitHub Actions workflows:

### 3.1 CI Workflow (`.github/workflows/ci.yml`)

Runs on every push and PR:
- Tests and lints backend Go code
- Tests and lints frontend TypeScript code
- Builds Docker image
- Validates the build

### 3.2 Backend Deployment (`.github/workflows/deploy-backend.yml`)

Triggers on:
- Push to main (when backend files change)
- Manual workflow dispatch

Deploys:
1. Builds Docker image
2. Pushes to Google Container Registry
3. Deploys to Cloud Run
4. Runs health check

### 3.3 Frontend Deployment (`.github/workflows/deploy-frontend.yml`)

Triggers on:
- Push to main (when frontend files change)
- Manual workflow dispatch

Deploys:
1. Builds Next.js application
2. Deploys to Vercel
3. Runs accessibility check

## Part 4: Testing the Setup

### 4.1 Test Backend Deployment

Make a small change to the backend:

```bash
# Edit backend/main.go or backend/data/resume_data.go
git add backend/
git commit -m "Update backend"
git push origin main
```

Check the Actions tab in GitHub to see the deployment progress.

### 4.2 Test Frontend Deployment

Make a small change to the frontend:

```bash
# Edit frontend/app/page.tsx or any component
git add frontend/
git commit -m "Update frontend"
git push origin main
```

Check the Actions tab in GitHub to see the deployment progress.

## Part 5: Manual Deployments

### Deploy Backend Manually

```bash
# From repository root
cd backend
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/resume-api
gcloud run deploy resume-api --image gcr.io/YOUR_PROJECT_ID/resume-api --region us-central1
```

### Deploy Frontend Manually

```bash
# From repository root
cd frontend
vercel --prod
```

### Trigger GitHub Actions Manually

Go to Actions → Choose workflow → Run workflow

## Monitoring

### Cloud Run Logs

```bash
gcloud run services logs read resume-api --region us-central1
```

Or visit: https://console.cloud.google.com/run

### Vercel Logs

Visit: https://vercel.com/dashboard

## Cost Optimization

### Cloud Run (Free Tier Limits)

- 2M requests/month
- 360,000 GB-seconds memory
- 180,000 vCPU-seconds
- 1 GB network egress

Stay within limits:
- Use `--min-instances 0` (scales to zero)
- Use `--memory 256Mi` (minimum needed)
- Monitor usage in GCP Console

### Vercel (Free Tier)

- Unlimited deployments
- 100 GB bandwidth
- Automatic HTTPS
- Global CDN

## Troubleshooting

### Backend Fails to Deploy

Check:
1. GCP project ID is correct
2. Service account has proper permissions
3. APIs are enabled (Run, Container Registry)
4. Docker image builds successfully locally

### Frontend Fails to Deploy

Check:
1. Vercel token is valid
2. Org ID and Project ID match
3. `API_URL` environment variable is set
4. Build succeeds locally with `npm run build`

### CORS Errors

Ensure backend `main.go` has proper CORS headers (already configured):
```go
w.Header().Set("Access-Control-Allow-Origin", "*")
```

For production, update to specific domain:
```go
w.Header().Set("Access-Control-Allow-Origin", "https://your-domain.vercel.app")
```

## Security Best Practices

1. **Rotate service account keys** regularly
2. **Use specific CORS origins** in production (not `*`)
3. **Enable Cloud Run authentication** if API should be private
4. **Monitor GCP billing** to avoid unexpected charges
5. **Review Vercel deployment logs** for security issues

## Updating Resume Content

Simply edit `backend/data/resume_data.go` and push to main:

```bash
# Edit resume data
vim backend/data/resume_data.go

# Commit and push
git add backend/data/resume_data.go
git commit -m "Update resume content"
git push origin main
```

The CI/CD pipeline will automatically:
1. Build and test the changes
2. Deploy to Cloud Run
3. Update the live API

The frontend will fetch the new data on next visit.

## Resources

- [Google Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Vercel Documentation](https://vercel.com/docs)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
