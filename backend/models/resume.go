package models

import "time"

// Resume represents the complete resume data structure
type Resume struct {
	PersonalInfo    PersonalInfo     `json:"personalInfo"`
	Summary         string           `json:"summary"`
	Experience      []Experience     `json:"experience"`
	Education       []Education      `json:"education"`
	Projects        []Project        `json:"projects"`
	Certifications  []Certification  `json:"certifications"`
	Awards          []Award          `json:"awards,omitempty"`
	TechnicalSkills TechnicalSkills  `json:"technicalSkills"`
	OpenSource      []OpenSource     `json:"openSource"`
	Interests       []string         `json:"interests"`
}

// PersonalInfo contains basic personal and contact information
type PersonalInfo struct {
	Name     string   `json:"name"`
	Title    string   `json:"title"`
	PhotoURL string   `json:"photoUrl,omitempty"` // GitHub raw URL for profile photo
	Location string   `json:"location"`
	Phones   []string `json:"phones"`
	Emails   []string `json:"emails"`
	GitHub   []string `json:"github"`
	LinkedIn string   `json:"linkedin"`
}

// Experience represents a work experience entry
type Experience struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"` // nil for current position
	Team        string    `json:"team,omitempty"`
	TechStack   []string  `json:"techStack"`
	Achievements []string `json:"achievements"`
}

// Education represents an educational qualification
type Education struct {
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	Institution string `json:"institution"`
	StartYear   int    `json:"startYear"`
	EndYear     int    `json:"endYear"`
	Grade       string `json:"grade"`
}

// TechnicalSkills groups technical competencies
type TechnicalSkills struct {
	Languages       []string `json:"languages"`
	Databases       []string `json:"databases"`
	CloudInfra      []string `json:"cloudInfra"`
	Specializations []string `json:"specializations"`
	VersionControl  []string `json:"versionControl"`
}

// OpenSource represents an open source project
type OpenSource struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Language    string   `json:"language,omitempty"`
	URL         string   `json:"url,omitempty"`
	Highlights  []string `json:"highlights,omitempty"`
}

// Project represents a personal or professional project
type Project struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Role        string     `json:"role"`
	TechStack   []string   `json:"techStack"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"` // nil for ongoing projects
	URL         string     `json:"url,omitempty"`
	GitHubURL   string     `json:"githubUrl,omitempty"`
	Highlights  []string   `json:"highlights"`
	Status      string     `json:"status"` // "completed", "ongoing", "archived"
}

// Certification represents a professional certification
type Certification struct {
	Name         string     `json:"name"`
	Issuer       string     `json:"issuer"`
	IssueDate    time.Time  `json:"issueDate"`
	ExpiryDate   *time.Time `json:"expiryDate,omitempty"`
	CredentialID string     `json:"credentialId,omitempty"`
	URL          string     `json:"url,omitempty"`
}

// Award represents an award or recognition
type Award struct {
	Title       string    `json:"title"`
	Issuer      string    `json:"issuer"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// BlogPost represents a blog post
type BlogPost struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"publishedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ReadTime    int       `json:"readTimeMinutes"`
	Featured    bool      `json:"featured"`
}

// ContactForm represents a contact form submission
type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject,omitempty"`
	Message string `json:"message"`
}
