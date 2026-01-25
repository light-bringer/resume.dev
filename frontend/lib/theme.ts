'use client'

import { createTheme, ThemeOptions } from '@mui/material/styles'

// Light theme configuration
const lightThemeOptions: ThemeOptions = {
    palette: {
        mode: 'light',
        primary: {
            main: '#6366f1', // Indigo
            light: '#818cf8',
            dark: '#4f46e5',
            contrastText: '#ffffff',
        },
        secondary: {
            main: '#ec4899', // Pink
            light: '#f472b6',
            dark: '#db2777',
            contrastText: '#ffffff',
        },
        background: {
            default: '#f8fafc',
            paper: '#ffffff',
        },
        text: {
            primary: '#0f172a',
            secondary: '#475569',
        },
        divider: 'rgba(0, 0, 0, 0.08)',
    },
    typography: {
        fontFamily: '"Inter", "Roboto", "Helvetica", "Arial", sans-serif',
        h1: {
            fontSize: '3rem',
            fontWeight: 700,
            lineHeight: 1.2,
            letterSpacing: '-0.02em',
        },
        h2: {
            fontSize: '2.25rem',
            fontWeight: 600,
            lineHeight: 1.3,
            letterSpacing: '-0.01em',
        },
        h3: {
            fontSize: '1.5rem',
            fontWeight: 600,
            lineHeight: 1.4,
        },
        h4: {
            fontSize: '1.25rem',
            fontWeight: 600,
            lineHeight: 1.5,
        },
        body1: {
            fontSize: '1rem',
            lineHeight: 1.7,
        },
        body2: {
            fontSize: '0.875rem',
            lineHeight: 1.6,
        },
    },
    shape: {
        borderRadius: 12,
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    textTransform: 'none',
                    fontWeight: 600,
                    borderRadius: 8,
                    padding: '10px 24px',
                },
                contained: {
                    boxShadow: '0 4px 14px 0 rgba(99, 102, 241, 0.25)',
                    '&:hover': {
                        boxShadow: '0 6px 20px 0 rgba(99, 102, 241, 0.35)',
                    },
                },
            },
        },
        MuiCard: {
            styleOverrides: {
                root: {
                    borderRadius: 16,
                    boxShadow: '0 4px 24px rgba(0, 0, 0, 0.06)',
                    border: '1px solid rgba(0, 0, 0, 0.05)',
                },
            },
        },
        MuiChip: {
            styleOverrides: {
                root: {
                    borderRadius: 8,
                    fontWeight: 500,
                },
            },
        },
        MuiPaper: {
            styleOverrides: {
                root: {
                    backgroundImage: 'none',
                },
            },
        },
    },
}

// Dark theme configuration
const darkThemeOptions: ThemeOptions = {
    palette: {
        mode: 'dark',
        primary: {
            main: '#818cf8', // Lighter indigo for dark mode
            light: '#a5b4fc',
            dark: '#6366f1',
            contrastText: '#0f172a',
        },
        secondary: {
            main: '#f472b6', // Lighter pink for dark mode
            light: '#f9a8d4',
            dark: '#ec4899',
            contrastText: '#0f172a',
        },
        background: {
            default: '#0f172a',
            paper: '#1e293b',
        },
        text: {
            primary: '#f1f5f9',
            secondary: '#94a3b8',
        },
        divider: 'rgba(255, 255, 255, 0.08)',
    },
    typography: lightThemeOptions.typography,
    shape: lightThemeOptions.shape,
    components: {
        ...lightThemeOptions.components,
        MuiCard: {
            styleOverrides: {
                root: {
                    borderRadius: 16,
                    boxShadow: '0 4px 24px rgba(0, 0, 0, 0.3)',
                    border: '1px solid rgba(255, 255, 255, 0.08)',
                    background: 'linear-gradient(145deg, rgba(30, 41, 59, 0.9), rgba(30, 41, 59, 0.7))',
                    backdropFilter: 'blur(10px)',
                },
            },
        },
        MuiButton: {
            styleOverrides: {
                root: {
                    textTransform: 'none',
                    fontWeight: 600,
                    borderRadius: 8,
                    padding: '10px 24px',
                },
                contained: {
                    boxShadow: '0 4px 14px 0 rgba(129, 140, 248, 0.3)',
                    '&:hover': {
                        boxShadow: '0 6px 20px 0 rgba(129, 140, 248, 0.4)',
                    },
                },
            },
        },
    },
}

export const lightTheme = createTheme(lightThemeOptions)
export const darkTheme = createTheme(darkThemeOptions)
