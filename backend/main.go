package main

import (
	"log"
	"net/http"
	"os"

	"github.com/justinas/alice"
	"github.com/light-bringer/resume.dev/handlers"
	"github.com/light-bringer/resume.dev/middleware"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/api/resume", handlers.HandleResume)
	mux.HandleFunc("/api/projects", handlers.HandleProjects)
	mux.HandleFunc("/api/certifications", handlers.HandleCertifications)
	mux.HandleFunc("/api/skills", handlers.HandleSkills)
	mux.HandleFunc("/api/contact", handlers.HandleContact)
	mux.HandleFunc("/health", handlers.HandleHealth)
	mux.HandleFunc("/metrics", handlers.HandleMetrics)

	// Create middleware chain using Alice for cleaner composition
	chain := alice.New(
		middleware.LoggingMiddleware,
		middleware.CORSMiddleware,
		middleware.RateLimitMiddleware(100), // 100 requests per second
	).Then(mux)

	// Log available endpoints
	log.Printf("Starting Resume API server on port %s", port)
	log.Printf("Endpoints available:")
	log.Printf("  - GET  /api/resume          Complete resume data")
	log.Printf("  - GET  /api/projects        Projects only")
	log.Printf("  - GET  /api/certifications  Certifications only")
	log.Printf("  - GET  /api/skills          Technical skills only")
	log.Printf("  - POST /api/contact         Contact form submission")
	log.Printf("  - GET  /health              Health check")
	log.Printf("  - GET  /metrics             Server metrics")
	log.Printf("")
	log.Printf("Middleware enabled:")
	log.Printf("  - Request logging")
	log.Printf("  - CORS (Access-Control-Allow-Origin: *)")
	log.Printf("  - Rate limiting (100 requests/second using Uber's token bucket)")

	// Start server
	if err := http.ListenAndServe(":"+port, chain); err != nil {
		log.Fatal(err)
	}
}
