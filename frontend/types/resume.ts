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

export interface Resume {
  personalInfo: PersonalInfo
  summary: string
  experience: Experience[]
  education: Education[]
  technicalSkills: TechnicalSkills
  openSource: OpenSource[]
  interests: string[]
}
