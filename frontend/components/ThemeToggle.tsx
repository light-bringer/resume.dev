'use client'

import { IconButton, Tooltip } from '@mui/material'
import LightModeIcon from '@mui/icons-material/LightMode'
import DarkModeIcon from '@mui/icons-material/DarkMode'
import { useTheme } from './ThemeProvider'

export default function ThemeToggle() {
  const { mode, toggleTheme } = useTheme()

  return (
    <Tooltip title={mode === 'light' ? 'Switch to dark mode' : 'Switch to light mode'}>
      <IconButton
        onClick={toggleTheme}
        aria-label="Toggle theme"
        sx={{
          position: 'fixed',
          top: 24,
          right: 24,
          zIndex: 50,
          bgcolor: 'background.paper',
          backdropFilter: 'blur(8px)',
          border: 1,
          borderColor: 'divider',
          boxShadow: 3,
          '&:hover': {
            boxShadow: 6,
            transform: 'scale(1.05)',
          },
          transition: 'all 0.2s ease-in-out',
        }}
      >
        {mode === 'light' ? (
          <DarkModeIcon sx={{ color: 'text.primary' }} />
        ) : (
          <LightModeIcon sx={{ color: 'text.primary' }} />
        )}
      </IconButton>
    </Tooltip>
  )
}
