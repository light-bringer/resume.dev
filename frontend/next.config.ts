import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  reactStrictMode: true,
  env: {
    API_URL: process.env.API_URL || 'http://localhost:8080',
  },
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'github.com',
      },
      {
        protocol: 'https',
        hostname: 'raw.githubusercontent.com',
      },
    ],
  },
}

export default nextConfig
