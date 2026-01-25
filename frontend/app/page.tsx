import { Resume } from '@/types/resume'
import Header from '@/components/Header'
import Summary from '@/components/Summary'
import Experience from '@/components/Experience'
import Education from '@/components/Education'
import ProjectsSection from '@/components/ProjectsSection'
import Skills from '@/components/Skills'
import OpenSource from '@/components/OpenSource'
import Interests from '@/components/Interests'
import AnimatedSection from '@/components/AnimatedSection'
import ThemeToggle from '@/components/ThemeToggle'

// Force dynamic rendering - don't try to fetch during build
export const dynamic = 'force-dynamic'

async function getResumeData(): Promise<Resume> {
  const apiUrl = process.env.API_URL || 'http://localhost:8080'

  try {
    const res = await fetch(`${apiUrl}/api/resume`, {
      cache: 'no-store',
    })

    if (!res.ok) {
      throw new Error('Failed to fetch resume data')
    }

    return res.json()
  } catch (error) {
    console.error('Error fetching resume:', error)
    throw error
  }
}

export default async function Home() {
  const resume = await getResumeData()

  return (
    <>
      <ThemeToggle />
      <main className="min-h-screen py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-4xl mx-auto space-y-6">
          <AnimatedSection delay={0}>
            <Header personalInfo={resume.personalInfo} />
          </AnimatedSection>

          <AnimatedSection delay={100}>
            <Summary summary={resume.summary} />
          </AnimatedSection>

          <AnimatedSection delay={150}>
            <Experience experiences={resume.experience} />
          </AnimatedSection>

          <AnimatedSection delay={200}>
            <Skills skills={resume.technicalSkills} />
          </AnimatedSection>

          <AnimatedSection delay={250}>
            <Education education={resume.education} />
          </AnimatedSection>

          <AnimatedSection delay={300}>
            <ProjectsSection projects={resume.projects} />
          </AnimatedSection>

          <AnimatedSection delay={350}>
            <OpenSource projects={resume.openSource} />
          </AnimatedSection>

          <AnimatedSection delay={400}>
            <Interests interests={resume.interests} />
          </AnimatedSection>
        </div>
      </main>
    </>
  )
}
