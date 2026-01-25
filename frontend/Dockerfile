# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Ensure public directory exists (Next.js requires it)
RUN mkdir -p public

# Build the application (standalone output configured in next.config.ts)
RUN npm run build

# Production stage using distroless
FROM gcr.io/distroless/nodejs20-debian12

WORKDIR /app

# Set environment to production
ENV NODE_ENV=production
ENV PORT=3000

# Copy standalone build from builder
# Next.js standalone build creates a minimal server in .next/standalone
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public

# Distroless runs as non-root by default (uid 65532)
# Expose port
EXPOSE 3000

# Start the application using node directly (no npm available in distroless)
CMD ["server.js"]
