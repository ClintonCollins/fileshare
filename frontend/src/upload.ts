import UploadFilesHandler from "./utils/uploadfile";
import {removeFullscreenLoader, createFullscreenLoader} from "./components/loader";

const uploadZoneInputElement = document.getElementById('upload-zone-input') as HTMLInputElement
const selectedFilesCardElement = document.getElementById('selected-files-card') as HTMLElement

document.documentElement.addEventListener('drag', (event) => {
    event.preventDefault()
})
document.documentElement.addEventListener('dragenter', (event) => {
    document.documentElement.classList.add('drag-drop-action')
    event.preventDefault()
})
document.documentElement.addEventListener('dragleave', (event) => {
    document.documentElement.classList.remove('drag-drop-action')
    event.preventDefault()
})
document.documentElement.addEventListener('dragover', (event) => {
    document.documentElement.classList.add('drag-drop-action')
    event.preventDefault()
})

document.documentElement.addEventListener('drop', (event) => {
    document.documentElement.classList.remove('drag-drop-action')
    event.preventDefault()
    const dataTransfer = event.dataTransfer
    if (dataTransfer == null) return
    Array.from(dataTransfer.files).forEach(file => {
        void new UploadFilesHandler(selectedFilesCardElement, file)
    })
})

uploadZoneInputElement.addEventListener('change', (event) => {
    event.preventDefault()
    const target = event.target as HTMLInputElement
    const files = target.files as FileList
    Array.from(files).forEach(file => {
        void new UploadFilesHandler(selectedFilesCardElement, file)
    })
})

document.addEventListener('paste', (event) => {
    createFullscreenLoader('Pasting files...')
    const dataTransfer = event.clipboardData
    if (dataTransfer == null) {
        removeFullscreenLoader()
        return
    }
    event.preventDefault()
    Array.from(dataTransfer.files).forEach(file => {
        void new UploadFilesHandler(selectedFilesCardElement, file)
    })
    removeFullscreenLoader()
})

export {}

