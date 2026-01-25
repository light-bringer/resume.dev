export interface PersonalInfo {
  name: string
  title: string
  photoUrl?: string // GitHub raw URL for profile photo
  location: string
  phones: string[]
  emails: string[]
  github: string[]
  linkedin: string
}

export interface Experience {
  title: string
  company: string
  location: string
  startDate: string
  endDate?: string | null
  team?: string
  techStack: string[]
  achievements: string[]
}

export interface Education {
  degree: string
  field: string
  institution: string
  startYear: number
  endYear: number
  grade: string
}

export interface TechnicalSkills {
  languages: string[]
  databases: string[]
  cloudInfra: string[]
  specializations: string[]
  versionControl: string[]
}

export interface OpenSource {
  name: string
  description: string
  language?: string
  url?: string
  highlights?: string[]
}

export interface Project {
  name: string
  description: string
  role: string
  techStack: string[]
  startDate: string
  endDate?: string | null
  url?: string
  githubUrl?: string
  highlights: string[]
  status: 'completed' | 'ongoing' | 'archived'
}

export interface Certification {
  name: string
  issuer: string
  issueDate: string
  expiryDate?: string | null
  credentialId?: string
  url?: string
}

export interface Award {
  title: string
  issuer: string
  date: string
  description: string
}

export interface BlogPost {
  id: string
  title: string
  slug: string
  summary: string
  content: string
  author: string
  tags: string[]
  publishedAt: string
  updatedAt: string
  readTimeMinutes: number
  featured: boolean
}

export interface ContactForm {
  name: string
  email: string
  subject?: string
  message: string
}

export interface Resume {
  personalInfo: PersonalInfo
  summary: string
  experience: Experience[]
  education: Education[]
  projects: Project[]
  certifications: Certification[]
  awards?: Award[]
  technicalSkills: TechnicalSkills
  openSource: OpenSource[]
  interests: string[]
}
