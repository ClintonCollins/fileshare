import {ConfirmModal} from './components/ConfirmModal'
import {MediaPreviewModal} from './components/MediaPreviewModal'
import {TableSelectActions} from './utils/TableSelectActions'

ConfirmModal.register()
MediaPreviewModal.register()

const selectAllCheckbox = document.querySelector('.select-all') as HTMLInputElement
const actionForm = document.querySelector('.action-form') as HTMLFormElement
const actionDetails = document.querySelector('.select-action') as HTMLDetailsElement
const actionDelete = document.querySelector('.delete-all-selected-action') as HTMLAnchorElement
const allItemCheckboxes = document.querySelectorAll('.item-selected-checkbox') as NodeListOf<HTMLInputElement>
const searchButton = document.querySelector('.search-button') as HTMLButtonElement
const confirmModal = document.querySelector('confirm-modal') as ConfirmModal
const previewCells = document.querySelectorAll('.media-preview-cell > *') as NodeListOf<HTMLElement>
const previewModal = document.querySelector('media-preview-modal') as MediaPreviewModal

const tableSelectActions = new TableSelectActions(selectAllCheckbox, allItemCheckboxes)

if (selectAllCheckbox && allItemCheckboxes) {
    document.addEventListener('table-select-change', () => {
        const actionDetails = document.querySelector('.select-action') as HTMLDetailsElement
        if (!actionDetails) return
        if (tableSelectActions.selectedIds.size === 0) {
            actionDetails.classList.add('opacity-0')
        } else {
            actionDetails.classList.remove('opacity-0')
        }
    })
}

function submitSearch() {
    actionForm.action = '/files'
    actionForm.method = 'GET'
    actionForm.submit()
}

searchButton.addEventListener('click', (event) => {
    event.preventDefault()
    submitSearch()
})

document.addEventListener('keyup', (event) => {
    if (event.key === 'Enter') {
        submitSearch()
    }
})

actionForm.addEventListener('submit', (event) => {
    event.preventDefault()
})

actionDelete.addEventListener('click', (event) => {
    event.preventDefault()
    const filesSelectedCount = tableSelectActions?.selectedIds.size
    const plural = filesSelectedCount > 1 ? 'files' : 'file'
    confirmModal.setAttribute('title', 'Delete all selected files')
    confirmModal.setAttribute('message',
        `Are you sure you want to delete ${filesSelectedCount} selected ${plural}?`)
    confirmModal.setAttribute('confirm-button-class', 'button-error')
    confirmModal.setAttribute('event-data', 'delete')
    confirmModal.showModal()
})

const url = new URL(window.location.href)
const urlParams = url.searchParams
const searchQuery = urlParams.get('search')
if (searchQuery != null && searchQuery === '') {
    urlParams.delete('search')
    url.search = urlParams.toString()
    window.history.replaceState({}, '', url.toString())
}

confirmModal.addEventListener('cancel', () => {
    actionDetails.removeAttribute('open')
})

confirmModal.addEventListener('confirm', (event: any) => {
    actionDetails.removeAttribute('open')
    if (event.detail && event.detail === 'delete') {
        actionForm.action = '/files/delete'
        actionForm.method = 'POST'
        actionForm.submit()
    }
})

previewCells.forEach(el => {
    const td = el.parentElement as HTMLTableCellElement
    el.addEventListener('click', () => {
        const mediaType = td.dataset.type
        const mediaRoot = previewModal.shadowRoot?.querySelector('#media-preview-slot') as HTMLDivElement
        const dataURL = td.dataset.url as string
        switch (mediaType) {
            case 'image':
                const imageElement = td.querySelector('img') as HTMLImageElement
                const newImageElement = imageElement.cloneNode(true) as HTMLImageElement
                newImageElement.slot = 'media'
                newImageElement.height = imageElement.naturalHeight
                newImageElement.width = imageElement.naturalWidth
                mediaRoot.replaceChildren(newImageElement)
                previewModal.showModal()
                break
            case 'video':
                const newVideoElement = document.createElement('video') as HTMLVideoElement
                newVideoElement.src = dataURL
                newVideoElement.controls = true
                mediaRoot.replaceChildren(newVideoElement)
                previewModal.showModal()
                break
            case 'audio':
                const newAudioElement = document.createElement('audio') as HTMLAudioElement
                newAudioElement.src = dataURL
                newAudioElement.controls = true
                mediaRoot.replaceChildren(newAudioElement)
                previewModal.showModal()
                break
        }
    })
})

export {}