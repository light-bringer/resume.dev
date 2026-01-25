package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/light-bringer/resume.dev/data"
	"github.com/light-bringer/resume.dev/models"
)

var startTime = time.Now()

// HandleResume returns the complete resume data
func HandleResume(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resume := data.GetResumeData()
	respondJSON(w, resume)
}

// HandleProjects returns only the projects section
func HandleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resume := data.GetResumeData()
	setCacheHeaders(w, 3600) // Cache for 1 hour
	respondJSON(w, resume.Projects)
}

// HandleCertifications returns only certifications
func HandleCertifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resume := data.GetResumeData()
	setCacheHeaders(w, 3600)
	respondJSON(w, resume.Certifications)
}

// HandleSkills returns technical skills
func HandleSkills(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resume := data.GetResumeData()
	setCacheHeaders(w, 3600)
	respondJSON(w, resume.TechnicalSkills)
}

// HandleContact processes contact form submissions
func HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var contact models.ContactForm
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate contact form
	if contact.Email == "" || contact.Message == "" {
		http.Error(w, "Email and message are required", http.StatusBadRequest)
		return
	}

	// TODO: Send email notification via SendGrid, AWS SES, etc.
	log.Printf("Contact form submission from: %s (%s)", contact.Name, contact.Email)
	log.Printf("Subject: %s", contact.Subject)
	log.Printf("Message: %s", contact.Message)

	respondJSON(w, map[string]string{
		"status":  "success",
		"message": "Thank you for your message! I'll get back to you soon.",
	})
}

// HandleHealth returns service health status
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"uptime":    time.Since(startTime).Seconds(),
		"version":   "1.0.0",
	}
	respondJSON(w, health)
}

// HandleMetrics returns basic metrics
func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := map[string]interface{}{
		"uptime_seconds": time.Since(startTime).Seconds(),
		"go_version":     "1.21+",
	}
	respondJSON(w, metrics)
}

// Helper functions

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func setCacheHeaders(w http.ResponseWriter, maxAge int) {
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
	// Simple ETag based on current time (in production, use content hash)
	w.Header().Set("ETag", fmt.Sprintf("\"%d\"", time.Now().Unix()/int64(maxAge)))
}
