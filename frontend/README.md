# Resume Frontend

A modern, responsive resume website built with Next.js, TypeScript, and Tailwind CSS.

## Getting Started

1. Install dependencies:
```bash
npm install
```

2. Set up environment variables:
```bash
cp .env.example .env.local
```

3. Run the development server:
```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) to see the resume.

## Environment Variables

- `API_URL`: The URL of the Go API backend (default: `http://localhost:8080`)

## Deployment

This project is configured for deployment on Vercel:

1. Push your code to GitHub
2. Import the project in Vercel
3. Set the `API_URL` environment variable to your deployed Go API URL
4. Deploy

## Tech Stack

- Next.js 15
- React 19
- TypeScript
- Tailwind CSS
