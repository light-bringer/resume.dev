'use client'

import { Container, Typography, Box } from '@mui/material'
import { Project } from '@/types/resume'
import ProjectCard from './ProjectCard'

interface ProjectsSectionProps {
    projects: Project[]
}

export default function ProjectsSection({ projects }: ProjectsSectionProps) {
    if (!projects || projects.length === 0) {
        return null
    }

    return (
        <Box component="section" id="projects" sx={{ py: 8 }}>
            <Container maxWidth="lg">
                <Typography
                    variant="h3"
                    component="h2"
                    gutterBottom
                    fontWeight={700}
                    sx={{
                        mb: 4,
                        background: (theme) =>
                            theme.palette.mode === 'dark'
                                ? 'linear-gradient(45deg, #818cf8 30%, #f472b6 90%)'
                                : 'linear-gradient(45deg, #6366f1 30%, #ec4899 90%)',
                        WebkitBackgroundClip: 'text',
                        WebkitTextFillColor: 'transparent',
                    }}
                >
                    Projects
                </Typography>

                <Box
                    sx={{
                        display: 'grid',
                        gridTemplateColumns: {
                            xs: '1fr',
                            md: 'repeat(2, 1fr)',
                        },
                        gap: 3,
                    }}
                >
                    {projects.map((project, index) => (
                        <ProjectCard key={index} project={project} />
                    ))}
                </Box>
            </Container>
        </Box>
    )
}
