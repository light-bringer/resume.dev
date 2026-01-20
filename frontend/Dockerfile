# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:20-alpine AS runner

WORKDIR /app

# Set environment to production
ENV NODE_ENV=production

# Copy package files and install production dependencies only
COPY package*.json ./
RUN npm ci --only=production

# Copy built application from builder
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

# Create a non-root user
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
RUN chown -R nextjs:nodejs /app

# Switch to non-root user
USER nextjs

# Expose port
EXPOSE 3000

# Set environment variable for Next.js
ENV PORT=3000

# Start the application
CMD ["npm", "start"]
