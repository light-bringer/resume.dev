package models

import "time"

// Resume represents the complete resume data structure
type Resume struct {
	PersonalInfo   PersonalInfo     `json:"personalInfo"`
	Summary        string           `json:"summary"`
	Experience     []Experience     `json:"experience"`
	Education      []Education      `json:"education"`
	TechnicalSkills TechnicalSkills `json:"technicalSkills"`
	OpenSource     []OpenSource     `json:"openSource"`
	Interests      []string         `json:"interests"`
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
