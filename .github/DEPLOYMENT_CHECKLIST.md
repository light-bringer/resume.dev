# Deployment Checklist

Use this checklist to track your CI/CD setup progress.

## Pre-Deployment

- [ ] Code pushed to GitHub repository
- [ ] Local development working (`make dev` runs successfully)
- [ ] Backend tests passing (`cd backend && make test`)
- [ ] Frontend builds successfully (`cd frontend && npm run build`)

## Google Cloud Run Setup (Backend)

- [ ] GCP account created
- [ ] GCP project created
- [ ] Billing enabled (required even for free tier)
- [ ] Cloud Run API enabled
- [ ] Container Registry API enabled
- [ ] Service account created
- [ ] Service account key downloaded (key.json)
- [ ] Service account has required permissions:
  - [ ] `roles/run.admin`
  - [ ] `roles/storage.admin`
  - [ ] `roles/iam.serviceAccountUser`

## GitHub Secrets - Backend

- [ ] `GCP_PROJECT_ID` added to GitHub secrets
- [ ] `GCP_SA_KEY` added to GitHub secrets (entire key.json content)

## Vercel Setup (Frontend)

- [ ] Vercel account created
- [ ] Vercel CLI installed (`npm install -g vercel`)
- [ ] Logged in to Vercel (`vercel login`)
- [ ] Project linked (`cd frontend && vercel link`)
- [ ] Vercel token generated
- [ ] Organization ID obtained
- [ ] Project ID obtained

## GitHub Secrets - Frontend

- [ ] `VERCEL_TOKEN` added to GitHub secrets
- [ ] `VERCEL_ORG_ID` added to GitHub secrets
- [ ] `VERCEL_PROJECT_ID` added to GitHub secrets
- [ ] `VERCEL_DOMAIN` added to GitHub secrets
- [ ] `API_URL` added to GitHub secrets (will be updated after backend deployment)

## Initial Deployments

### Backend
- [ ] First manual deployment completed:
  ```bash
  cd backend
  gcloud builds submit --tag gcr.io/PROJECT_ID/resume-api
  gcloud run deploy resume-api --image gcr.io/PROJECT_ID/resume-api --region us-central1
  ```
- [ ] Backend URL obtained
- [ ] Health check endpoint working (`curl https://YOUR_URL/health`)

### Frontend
- [ ] Updated `API_URL` in Vercel with backend URL
- [ ] Updated `API_URL` GitHub secret with backend URL
- [ ] First manual deployment completed:
  ```bash
  cd frontend
  vercel --prod
  ```
- [ ] Frontend accessible at Vercel URL
- [ ] Resume data loading correctly

## CI/CD Verification

- [ ] Push test commit to main branch
- [ ] CI workflow runs successfully
- [ ] Backend deployment workflow triggers (if backend changed)
- [ ] Frontend deployment workflow triggers (if frontend changed)
- [ ] Deployments complete successfully
- [ ] Services health checks pass

## Post-Deployment

- [ ] Backend Cloud Run URL accessible
- [ ] Frontend Vercel URL accessible
- [ ] Resume displays correctly
- [ ] No CORS errors in browser console
- [ ] Custom domain configured (optional)
- [ ] SSL certificate active (automatic)

## Monitoring Setup

- [ ] Cloud Run monitoring dashboard bookmarked
- [ ] Vercel dashboard bookmarked
- [ ] GitHub Actions tab bookmarked
- [ ] GCP billing alerts configured
- [ ] Vercel usage alerts configured (if needed)

## Documentation

- [ ] `API_URL` documented for team
- [ ] Deployment URLs shared
- [ ] GitHub secrets documented (securely)
- [ ] Team members have necessary access

## Optional Enhancements

- [ ] Custom domain for backend
- [ ] Custom domain for frontend
- [ ] Environment-specific deployments (staging/production)
- [ ] Slack/Discord notifications for deployments
- [ ] Monitoring and alerting setup
- [ ] Performance monitoring (Lighthouse CI)
- [ ] Error tracking (Sentry, etc.)

## Troubleshooting Checklist

If deployment fails, check:

- [ ] GitHub secrets are correct (no extra spaces)
- [ ] GCP project ID matches in all places
- [ ] Service account has all required permissions
- [ ] APIs are enabled in GCP
- [ ] Billing is enabled in GCP
- [ ] Vercel token is valid
- [ ] `API_URL` is correct in frontend environment
- [ ] CORS headers are configured in backend
- [ ] Docker image builds successfully locally
- [ ] Frontend builds successfully locally

## Success Criteria

Your deployment is successful when:

- ✅ Backend API responds to requests
- ✅ Frontend displays resume correctly
- ✅ Push to main triggers automatic deployments
- ✅ Both services stay within free tier limits
- ✅ No errors in production logs
- ✅ HTTPS works on both services
- ✅ Mobile responsive design works

## Next Steps After Deployment

1. Share your resume URL with others
2. Monitor usage in first week
3. Update resume content as needed
4. Consider adding analytics (optional)
5. Set up uptime monitoring (optional)

---

**Need help?** See [DEPLOYMENT.md](../DEPLOYMENT.md) for detailed instructions.
