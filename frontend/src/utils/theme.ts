const darkModeToggle = document.querySelector('.dark-mode-toggle') as HTMLElement

/**
 * Enum for a website theme.
 * @readonly
 * @enum {string}
 */
export const Theme = Object.freeze({
    Light: 'light',
    Dark: 'dark',
})

/**
 * Gets the current theme.
 * @returns {string}
 */
function getTheme() {
    const currentTheme = document.documentElement.dataset.theme
    switch (currentTheme) {
        case 'light':
            return Theme.Light
        default:
            return Theme.Dark
    }
}

function persistThemeToStorage(theme: string) {
    window.localStorage.setItem('theme', theme)
}

/** Gets the theme from local storage.
 * @returns {Theme|string|null}
 */
function getThemeFromStorage() {
    const theme = window.localStorage.getItem('theme')
    if (theme === 'dark' || theme === 'light') {
        return theme
    }
    return null
}

/**
 * Sets the tooltip for the dark mode toggle.
 */
function setThemeTooltip() {
    darkModeToggle.dataset.tooltip = getTheme() === 'light' ? 'Dark mode' : 'Light mode'
}

/**
 * Sets the theme.
 * @param {Theme|string} theme The theme to set.
 */
export function SetTheme(theme: string) {
    document.documentElement.dataset.theme = theme
    persistThemeToStorage(theme)
    setThemeTooltip()
}

/**
 * Loads the theme preference from local storage or the user's OS settings.
 */
export function LoadThemePreference() {
    switch (getThemeFromStorage()) {
        case 'light':
            document.documentElement.dataset.theme = 'light'
            break
        case 'dark':
            document.documentElement.dataset.theme = 'dark'
            break
        default:
            document.documentElement.dataset.theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
    }
    setThemeTooltip()
}

/**
 * Watches for theme changes based on the dark mode toggle element and updates the theme accordingly.
 */
export function WatchForThemeChanges() {
    darkModeToggle.addEventListener('click', () => {
        SetTheme(getTheme() === 'light' ? 'dark' : 'light')
    })
}