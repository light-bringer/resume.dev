'use client'

import * as React from 'react'
import { CacheProvider } from '@emotion/react'
import createEmotionCache from './createEmotionCache'

// Client-side cache shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache()

export interface EmotionCacheProviderProps {
    children: React.ReactNode
}

export default function EmotionCacheProvider({ children }: EmotionCacheProviderProps) {
    return (
        <CacheProvider value={clientSideEmotionCache}>
            {children}
        </CacheProvider>
    )
}
