'use client'

import createCache from '@emotion/cache'

// This implementation is from emotion-js.
// https://github.com/emotion-js/emotion/issues/2928
export default function createEmotionCache() {
    return createCache({ key: 'css', prepend: true })
}
