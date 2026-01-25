'use client'

import { Card, CardContent, Typography, Chip, Box, Link, IconButton, Stack } from '@mui/material'
import { GitHub, Launch, CalendarToday } from '@mui/icons-material'
import { Project } from '@/types/resume'

interface ProjectCardProps {
    project: Project
}

export default function ProjectCard({ project }: ProjectCardProps) {
    const isOngoing = !project.endDate
    const startYear = new Date(project.startDate).getFullYear()
    const endYear = project.endDate ? new Date(project.endDate).getFullYear() : 'Present'

    return (
        <Card
            sx={{
                height: '100%',
                display: 'flex',
                flexDirection: 'column',
                transition: 'all 0.3s ease',
                '&:hover': {
                    transform: 'translateY(-8px)',
                    boxShadow: 8,
                },
            }}
        >
            <CardContent sx={{ flexGrow: 1, p: 3 }}>
                {/* Header with name and status */}
                <Box display="flex" justifyContent="space-between" alignItems="start" mb={2}>
                    <Typography variant="h5" component="h3" fontWeight={600}>
                        {project.name}
                    </Typography>
                    <Chip
                        label={project.status}
                        size="small"
                        color={isOngoing ? 'primary' : 'success'}
                        sx={{ textTransform: 'capitalize' }}
                    />
                </Box>

                {/* Description */}
                <Typography variant="body2" color="text.secondary" paragraph>
                    {project.description}
                </Typography>

                {/* Role */}
                <Typography variant="subtitle2" color="primary" gutterBottom fontWeight={600}>
                    {project.role}
                </Typography>

                {/* Date range */}
                <Box display="flex" alignItems="center" gap={0.5} mb={2}>
                    <CalendarToday sx={{ fontSize: 16, color: 'text.secondary' }} />
                    <Typography variant="caption" color="text.secondary">
                        {startYear} - {endYear}
                    </Typography>
                </Box>

                {/* Tech stack */}
                <Box display="flex" flexWrap="wrap" gap={1} my={2}>
                    {project.techStack.map((tech) => (
                        <Chip
                            key={tech}
                            label={tech}
                            size="small"
                            variant="outlined"
                            sx={{
                                borderRadius: 2,
                                fontWeight: 500,
                            }}
                        />
                    ))}
                </Box>

                {/* Highlights */}
                <Box component="ul" sx={{ pl: 2, mt: 2, mb: 2 }}>
                    {project.highlights.map((highlight, idx) => (
                        <Typography
                            key={idx}
                            component="li"
                            variant="body2"
                            sx={{
                                mb: 1,
                                '&::marker': {
                                    color: 'primary.main',
                                },
                            }}
                        >
                            {highlight}
                        </Typography>
                    ))}
                </Box>

                {/* Links */}
                <Stack direction="row" spacing={1} mt="auto">
                    {project.githubUrl && (
                        <IconButton
                            component={Link}
                            href={project.githubUrl}
                            target="_blank"
                            rel="noopener noreferrer"
                            size="small"
                            sx={{
                                border: 1,
                                borderColor: 'divider',
                                '&:hover': {
                                    borderColor: 'primary.main',
                                    color: 'primary.main',
                                },
                            }}
                        >
                            <GitHub fontSize="small" />
                        </IconButton>
                    )}
                    {project.url && (
                        <IconButton
                            component={Link}
                            href={project.url}
                            target="_blank"
                            rel="noopener noreferrer"
                            size="small"
                            sx={{
                                border: 1,
                                borderColor: 'divider',
                                '&:hover': {
                                    borderColor: 'primary.main',
                                    color: 'primary.main',
                                },
                            }}
                        >
                            <Launch fontSize="small" />
                        </IconButton>
                    )}
                </Stack>
            </CardContent>
        </Card>
    )
}
