import {WatchForThemeChanges, LoadThemePreference} from './utils/theme'

/**
 * Handles loading the theme data and watching for theme changes.
 * */
function handleThemeData() {
    LoadThemePreference()
    WatchForThemeChanges()
}

const flashNotificationElement = document.querySelector('.flash-notification') as HTMLDivElement
if (flashNotificationElement) {
    flashNotificationElement.style.animation = 'bounceInDown 0.75s forwards'
    setTimeout(() => {
        flashNotificationElement.style.animation = 'bounceOutUp 0.75s forwards'
    }, 5000)
}

// Manage theme changes
handleThemeData()