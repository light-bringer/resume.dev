package data

import (
	"time"

	"github.com/light-bringer/resume.dev/models"
)

// GetResumeData returns the resume data
func GetResumeData() *models.Resume {
	// Helper function to create time pointers
	mustParseDate := func(layout, value string) time.Time {
		t, _ := time.Parse(layout, value)
		return t
	}

	endDate := func(layout, value string) *time.Time {
		t := mustParseDate(layout, value)
		return &t
	}

	return &models.Resume{
		PersonalInfo: models.PersonalInfo{
			Name:     "Debapriya Das",
			Title:    "Senior Backend Engineer",
			PhotoURL: "https://github.com/yodebu.png", // GitHub profile picture
			Location: "Bangalore, India",
			Phones:   []string{"+91-8553450562", "+91-7001673957"},
			Emails:   []string{"yodebu@gmail.com", "debapriya.nitd@gmail.com"},
			GitHub:   []string{"@yodebu", "@light-bringer"},
			LinkedIn: "linkedin.com/in/yodebu",
		},
		Summary: "Senior Backend Engineer with 9+ years of experience designing, building and maintaining large-scale distributed systems. Proficient in Golang, Python, TypeScript, Java, Scala, Postgres, Kafka, AWS, and GCP. Active open-source contributor and a Linux Fanboy. Loves Road Trips.",
		Experience: []models.Experience{
			{
				Title:     "Senior Software Engineer",
				Company:   "Twilio Segment",
				Location:  "Bangalore, India (Remote)",
				StartDate: mustParseDate("2006-01", "2022-02"),
				EndDate:   nil, // Current position
				Team:      "Core Platform - Data Management & Compliance/Governance",
				TechStack: []string{"Go", "Postgres", "Kafka", "AWS", "GCP", "Terraform"},
				Achievements: []string{
					"Designed, developed and acted as a lead on the latest data retention project, built an Orkes based orchestrator and subsequent DAG workflows for automatically identifying churned accounts and implementing billing-based retention flows, reducing Twilio's storage usage by ~40%",
					"Led re-architecture of deletion orchestration service, written in Go, scaling throughput 30x from legacy system while eliminating major SEVs and maintaining 100% uptime. Enabled public-apis to make sure that the deletions system can be programmatically invoked across the customers.",
					"Orchestrated phased migration of all high-volume customers to new service with zero incidents, successfully onboarding 100% of largest customers.",
					"Served as service owner of various deletion systems and deletion manager, and primary/secondary on-call, coordinating and mitigating SEVs through rapid incident response and root cause resolution.",
					"Built end-to-end testing framework based on Orkes to internally test the load bearing capacity, integrations and deletion correctness across the different deletion systems, GDPR, BCR, Retention.",
					"Implemented rate-limiting improvements across the overall system for deletion orchestration, smoothing traffic bursts and protecting external facing and internal systems; exposed rate-limiting controls via public APIs for customer self-service.",
					"Served as infrastructure lead across team projects, managing service and cloud infrastructure using Terraform.",
					"Built an internal reporting service for the deletion orchestration system to detect anomalies on the overall deletion systems and SLO breaches.",
					"Built a new custom identifier-based deletion system on archived object storage, enabling cost-effective targeted purges and exposing self-service customer API, scaled up this system to use EMR based hadoop jobs, reducing reliance on AWS S3 APIs, leading to efficiency in terms of cost.",
					"Designed and implemented a P0 encryption-service for PII data plane, ensuring HIPAA compliance and restricting data access to authorized personnel only.",
					"Optimized consent management filtering service 5x through improved caching strategies, achieving significant infrastructure cost savings with existing resources.",
					"Led 3-person team to design and develop Segment's internal data-plane auditing service from scratch, vital for ensuring Segment data stores adhere to PII compliance standards.",
					"Developed self-service data migration tooling and internal dashboards, enabling seamless migration of 500+ workspaces across regions for GDPR compliance.",
				},
			},
			{
				Title:     "Senior Backend Developer",
				Company:   "Healthifyme",
				Location:  "Bangalore, India",
				StartDate: mustParseDate("2006-01", "2021-09"),
				EndDate:   endDate("2006-01", "2022-01"),
				TechStack: []string{"Python", "Django", "Flask", "Golang", "AWS"},
				Achievements: []string{
					"Led 3-person team optimizing Django monolith, scaling services 10x for peak traffic while maintaining performance SLAs.",
					"Built Webhook Service in Golang to offload SMS/call status updates from monolith, reducing database load.",
					"Developed User Configuration service with caching layer, significantly improving app load times.",
					"Optimized AWS infrastructure for cost efficiency under heavy traffic, integrating with Twilio SMS, Twilio Video, and Zoom platforms.",
				},
			},
			{
				Title:     "Software Engineer II",
				Company:   "Ajio.com, JIO Platforms",
				Location:  "Bangalore, India",
				StartDate: mustParseDate("2006-01", "2020-06"),
				EndDate:   endDate("2006-01", "2021-08"),
				Team:      "Internal Tools & DevOps Team",
				TechStack: []string{"Python", "C++", "Golang", "Java"},
				Achievements: []string{
					"Built large-scale ELK clusters from scratch for data aggregation, ingesting 30M events/hour. Developed elastic strategies for demand based cluster creation.",
					"Built the Ajio and later Jio E-commerce observability stack based on Prometheus.",
					"Designed monitoring dashboards using Grafana, and ELK for real-time system observability across the services onboarded.",
					"Built automated CI/CD pipelines with unit testing, containerization, monitoring, deployment, and alerting for seamless application delivery.",
				},
			},
			{
				Title:     "Software Engineer",
				Company:   "Electronics For Imaging (EFI)",
				Location:  "Bangalore, India",
				StartDate: mustParseDate("2006-01", "2017-03"),
				EndDate:   endDate("2006-01", "2020-05"),
				Team:      "Fiery Core Engineering - Internal Tools & DevOps",
				TechStack: []string{"Python", "C++", "Ruby", "Golang", "JavaScript", "Bash", "Make", "gcc"},
				Achievements: []string{
					"Acted as a service owner for all the internal tools and applications.",
					"Designed and developed end-to-end patching framework from creation to customer delivery and deployment.",
					"Built automation framework for detecting performance degradation and system regressions.",
					"Maintained release management framework, CI/CD systems, and testing infrastructure enabling seamless developer contributions.",
				},
			},
			{
				Title:     "Associate Software Engineer",
				Company:   "Incture",
				Location:  "Bangalore, India",
				StartDate: mustParseDate("2006-01", "2016-08"),
				EndDate:   endDate("2006-01", "2017-02"),
				Team:      "CherryWorks - Backend Engineer",
				TechStack: []string{"Node.js", "Express.js", "JavaScript", "Bash", "AWS"},
				Achievements: []string{
					"Worked on requirements and developed HR Apps for CherryWorks, coded the recruitment and the onboarding flow.",
					"Maintained releases, databases for on-prem infrastructure and helped migrate to AWS later.",
				},
			},
		},
		Education: []models.Education{
			{
				Degree:      "Bachelor of Technology",
				Field:       "Computer Science and Engineering",
				Institution: "National Institute of Technology, Durgapur",
				StartYear:   2012,
				EndYear:     2016,
				Grade:       "First Class Degree",
			},
		},
		Projects: []models.Project{
			{
				Name:        "Convoy - Ride Sharing Platform",
				Description: "Full-stack ride-sharing platform with real-time location tracking, trip management, and user authentication. Built with microservices architecture featuring Auth and Trip services.",
				Role:        "Lead Developer & Architect",
				TechStack:   []string{"Go", "gRPC", "PostgreSQL", "Redis", "Next.js", "TypeScript", "Docker", "Kubernetes", "Terraform"},
				StartDate:   mustParseDate("2006-01", "2025-12"),
				EndDate:     nil,
				GitHubURL:   "https://github.com/light-bringer/convoy",
				Highlights: []string{
					"Designed microservices architecture with Auth and Trip services communicating via gRPC",
					"Implemented real-time location tracking using Redis geospatial indexing for proximity queries",
					"Built containerized integration testing framework with Docker Compose",
					"Created comprehensive CI/CD pipeline with GitHub Actions for automated testing and deployment",
					"Deployed to Google Cloud Run with infrastructure-as-code using Terraform",
				},
				Status: "ongoing",
			},
			{
				Name:        "Stats - OpenTelemetry Metrics Library",
				Description: "Production-ready, high-performance metrics library with full OpenTelemetry SDK compliance and multi-backend support. Achieves >200k events/sec with sub-nanosecond push latency.",
				Role:        "Creator & Maintainer",
				TechStack:   []string{"Go", "OpenTelemetry", "Prometheus", "Datadog", "CloudWatch"},
				StartDate:   mustParseDate("2006-01", "2024-01"),
				EndDate:     endDate("2006-01", "2025-06"),
				GitHubURL:   "https://github.com/yodebu/stats",
				Highlights: []string{
					"Achieves >200k events/sec with sub-nanosecond latency using lock-free ring buffers",
					"Implements adaptive batching for optimal throughput across varying load patterns",
					"Multi-backend support (Datadog, Prometheus, CloudWatch) with parallel exporter architecture",
					"Circuit breakers, panic recovery, and graceful degradation for production reliability",
					"Memory-bounded operation with <0.1% drop rate under sustained load",
					"Dual-mode API (legacy + OTel) sharing high-performance pipeline",
				},
				Status: "completed",
			},
			{
				Name:        "StatLib - Lock-Free Statistics Aggregator",
				Description: "High-performance, thread-safe statistics library using atomic operations and cache-line padding to prevent false sharing. Achieves sub-nanosecond read latency and 11M events/sec write throughput.",
				Role:        "Creator",
				TechStack:   []string{"Go", "Atomic Operations", "Concurrent Programming"},
				StartDate:   mustParseDate("2006-01", "2024-03"),
				EndDate:     endDate("2006-01", "2024-08"),
				GitHubURL:   "https://github.com/yodebu/statlib",
				Highlights: []string{
					"Sub-nanosecond read latency with 11M events/sec write throughput on 8 cores",
					"Lock-free CAS loops for min/max updates with overflow protection",
					"Cache-line padding prevents false sharing in multi-core environments",
					"Extensively tested with race detector and benchmarked across 1-8 CPUs",
					"Demonstrates near-linear read scaling with increasing CPU count",
				},
				Status: "completed",
			},
			{
				Name:        "Trading System Performance Enhancement",
				Description: "Advanced algorithmic trading system with machine learning-based entry timing, volatility regime detection, and Greeks-based position management.",
				Role:        "Backend Developer",
				TechStack:   []string{"Python", "Backtrader", "Pandas", "NumPy", "Machine Learning"},
				StartDate:   mustParseDate("2006-01", "2025-11"),
				EndDate:     endDate("2006-01", "2025-11"),
				Highlights: []string{
					"Implemented volatility regime detection for adaptive trading strategies",
					"Dynamic profit targets and stop losses based on market conditions",
					"Transaction cost filtering to improve risk-adjusted returns",
					"Greeks-based position management for options trading",
					"Machine learning models for entry timing optimization",
				},
				Status: "completed",
			},
			{
				Name:        "Electoral Search Tool",
				Description: "Command-line tool for searching and analyzing electoral data with optimized search algorithms and comprehensive test coverage.",
				Role:        "Developer",
				TechStack:   []string{"Python", "CLI"},
				StartDate:   mustParseDate("2006-01", "2025-12"),
				EndDate:     endDate("2006-01", "2025-12"),
				Highlights: []string{
					"Refactored codebase reducing cyclomatic complexity from 41 to <10",
					"Extracted helper functions for improved maintainability and readability",
					"Comprehensive test coverage ensuring reliability",
					"Efficient search algorithms for large datasets",
				},
				Status: "completed",
			},
		},
		Certifications: []models.Certification{
			// Add certifications here as they are obtained
		},
		Awards: []models.Award{
			// Add awards here if any
		},
		TechnicalSkills: models.TechnicalSkills{
			Languages:       []string{"Golang", "Python", "C++", "JavaScript/TypeScript", "Java/Scala", "Bash", "Make"},
			Databases:       []string{"PostgreSQL", "MySQL", "MongoDB", "DynamoDB", "Redis", "Kafka", "SQS", "Snowflake", "Redshift", "BigQuery"},
			CloudInfra:      []string{"AWS (EKS, ECS, Lambda, ELB, EMR, MemoryDB)", "GCP (Kubernetes, Firebase, VPC)", "Docker", "Kubernetes", "Terraform", "Orkes"},
			Specializations: []string{"Distributed Systems", "Microservices Architecture", "Compilers", "Linux", "Data Compliance (GDPR/HIPAA/BCR)", "CI/CD", "Observability", "Incident Response"},
			VersionControl:  []string{"Git", "GitHub", "SVN", "Clearcase"},
		},
		OpenSource: []models.OpenSource{
			{
				Name:        "Stats - OpenTelemetry-Compliant Metrics Library",
				Language:    "Golang",
				Description: "Production-ready, non-blocking UDP metrics library with full OpenTelemetry SDK compliance and multi-backend support (Datadog, Prometheus, CloudWatch). Achieves >200k events/sec with sub-nanosecond push latency using lock-free ring buffers and adaptive batching. Features circuit breakers, panic recovery, graceful degradation, and memory-bounded operation with <0.1% drop rate under load. Dual-mode API (legacy + OTel) shares high-performance pipeline with parallel exporter architecture and comprehensive observability.",
			},
			{
				Name:        "StatLib - Lock-Free Statistics Aggregator",
				Language:    "Golang",
				Description: "High-performance, thread-safe statistics library using atomic operations and cache-line padding to prevent false sharing. Achieves sub-nanosecond read latency and 11M events/sec write throughput on 8 cores. Implements lock-free CAS loops for min/max updates with overflow protection. Extensively tested with race detector and benchmarked across 1-8 CPUs demonstrating near-linear read scaling.",
			},
			{
				Name:        "Node-Db-File-Logger",
				Description: "Simplified JavaScript logging module supporting stdout/file/database outputs. Available on GitHub (@light-bringer) and NPM registry.",
				URL:         "https://github.com/light-bringer",
			},
			{
				Name:        "Active Contributor",
				Description: "Regular contributions to open-source repositories across backend tooling, infrastructure automation, and developer productivity tools.",
				URL:         "https://github.com/yodebu",
				Highlights:  []string{"github.com/yodebu", "github.com/light-bringer"},
			},
		},
		Interests: []string{
			"Football: Hardcore Barcelona fan, follows Champions League, La Liga, and Premier League",
			"Photography: Capturing nature's beauty - Instagram: @Yodebu",
			"Music: Trained in North Indian Classical Music, plays keyboard, loves a good jam!",
			"Adventure: Biking, Roadtrips, backpacking, camping and trekking in the Himalayas.",
		},
	}
}
