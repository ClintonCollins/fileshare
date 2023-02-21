import {filesize} from "filesize";
import humanizeDuration from "humanize-duration";
import { MediaPreviewModal} from '../components/MediaPreviewModal'

MediaPreviewModal.register()
const previewModal = document.querySelector('media-preview-modal') as MediaPreviewModal

enum UploadStatus {
    Pending = 'pending',
    Uploading = 'uploading',
    Done = 'done',
    Error = 'error',
}

enum FileTypeCategory {
    Image = 'image',
    Video = 'video',
    Audio = 'audio',
    Other = 'other',
}

export default class UploadFilesHandler {
    static nextFileID = 0
    static uploadFilesMap: Map<Number, UploadFilesHandler> = new Map<Number, UploadFilesHandler>()
    static simultaneousUploads = 3
    static uploadURL = '/upload'
    static uploadButtonElement: HTMLButtonElement = document.getElementById('selected-files-upload-button') as HTMLButtonElement
    static uploadButtonVisible = false
    static uploadButtonDisabled = false

    private selectedFilesCardElement!: HTMLElement
    private selectedFilesContainerElement!: HTMLElement
    private selectedFilesRow!: HTMLElement
    private previewBoxElement!: HTMLElement
    private nameBoxElement!: HTMLElement
    private sizeBoxElement!: HTMLElement
    private typeBoxElement!: HTMLElement
    // private passwordBoxElement!: HTMLElement
    private urlBoxElement!: HTMLElement
    private statusBoxElement!: HTMLElement
    private progressBoxElement!: HTMLElement
    private progressBarElement!: HTMLProgressElement
    private progressSpeedBoxElement!: HTMLElement
    private removeFileBoxElement!: HTMLElement
    private fileTypeCategory!: FileTypeCategory
    private progressAlertElement!: HTMLElement
    private uploadStartTime: number = new Date().getTime()
    private readonly file: File
    private readonly id: number
    private _status: UploadStatus
    private _errorMessage: string = ''
    private xhr!: XMLHttpRequest

    public get errorMessage() {
        return this._errorMessage
    }

    public get status() {
        return this._status
    }

    constructor(selectedFilesCardElement: HTMLElement, file: File) {
        this.selectedFilesCardElement = selectedFilesCardElement
        this.file = file
        this.id = UploadFilesHandler.nextFileID
        this._status = UploadStatus.Pending

        this.initializeSelectors()
        this.populateSelectedFileRow()
        this.addEventListeners()

        UploadFilesHandler.uploadFilesMap.set(this.id, this)
        UploadFilesHandler.nextFileID++

        if (!UploadFilesHandler.uploadButtonElement === null) throw new Error('upload button not found')
        if (!UploadFilesHandler.uploadButtonVisible) {
            UploadFilesHandler.uploadButtonElement.classList.remove('opacity-0')
            UploadFilesHandler.uploadButtonVisible = true
        }
        if (UploadFilesHandler.uploadButtonDisabled) {
            UploadFilesHandler.uploadButtonElement.disabled = false
            UploadFilesHandler.uploadButtonElement.classList.remove('disabled')
            UploadFilesHandler.uploadButtonDisabled = false
        }
        this.uploadAllPendingFilesListener()
    }

    private addEventListeners() {
        this.addMediaPreviewClickListener()
        this.addRemoveButtonListener()
    }

    private initializeSelectors() {
        const selectedFilesTemplate = document.getElementById('selected-files-template') as HTMLTemplateElement
        if (selectedFilesTemplate === null || selectedFilesTemplate.content.firstElementChild === null) {
            throw new Error('selected-files-template not found')
        }
        this.selectedFilesRow = selectedFilesTemplate.content.firstElementChild.cloneNode(true) as HTMLElement
        this.selectedFilesRow.dataset.fileId = this.id.toString()
        this.previewBoxElement = this.selectedFilesRow.querySelector('.selected-file-preview-box') as HTMLElement
        this.nameBoxElement = this.selectedFilesRow.querySelector('.selected-file-name-box') as HTMLElement
        this.sizeBoxElement = this.selectedFilesRow.querySelector('.selected-file-size-box') as HTMLElement
        this.urlBoxElement = this.selectedFilesRow.querySelector('.selected-file-url-box') as HTMLElement
        this.statusBoxElement = this.selectedFilesRow.querySelector('.selected-file-status-box') as HTMLElement
        this.progressBoxElement = this.selectedFilesRow.querySelector('.selected-file-progress-box') as HTMLElement
        this.progressSpeedBoxElement = this.selectedFilesRow.querySelector('.selected-file-progress-speed-box') as HTMLElement
        this.removeFileBoxElement = this.selectedFilesRow.querySelector('.selected-file-remove-box') as HTMLElement
        this.typeBoxElement = this.selectedFilesRow.querySelector('.selected-file-type-box') as HTMLElement
        this.progressBarElement = this.progressBoxElement.querySelector('progress') as HTMLProgressElement
        this.progressAlertElement = this.progressBoxElement.querySelector('fileshare-alert') as HTMLElement
        this.selectedFilesContainerElement = document.getElementById('selected-files-container') as HTMLElement

        const missingElements: string[] = []
        if (this.previewBoxElement === null) {
            missingElements.push('selected-file-preview-box')
        }
        if (this.nameBoxElement === null) {
            missingElements.push('selected-file-name-box')
        }
        if (this.sizeBoxElement === null) {
            missingElements.push('selected-file-size-box')
        }
        if (this.urlBoxElement === null) {
            missingElements.push('selected-file-url-box')
        }
        if (this.statusBoxElement === null) {
            missingElements.push('selected-file-status-box')
        }
        if (this.progressBoxElement === null) {
            missingElements.push('selected-file-progress-box')
        }
        if (this.progressSpeedBoxElement === null) {
            missingElements.push('selected-file-progress-speed-box')
        }
        if (this.removeFileBoxElement === null) {
            missingElements.push('selected-file-remove-box')
        }
        if (this.progressBarElement === null) {
            missingElements.push('selected-file-progress-box progress')
        }

        if (this.selectedFilesContainerElement === null) {
            missingElements.push('selected-files-container')
        }

        if (missingElements.length > 0) {
            throw new Error('template elements not found: ' + missingElements.join(', '))
        }
    }

    private populateSelectedFileRow() {
        this.nameBoxElement.textContent = this.file.name
        this.statusBoxElement.textContent = this._status
        this.typeBoxElement.textContent = this.file.type
        this.sizeBoxElement.textContent = filesize(this.file.size, {output: 'string'}) as string
        if (this.file.type.startsWith('image/')) {
            this.fileTypeCategory = FileTypeCategory.Image
        } else if (this.file.type.startsWith('video/')) {
            this.fileTypeCategory = FileTypeCategory.Video
        } else if (this.file.type.startsWith('audio/')) {
            this.fileTypeCategory = FileTypeCategory.Audio
        } else {
            this.fileTypeCategory = FileTypeCategory.Other
        }
        switch (this.fileTypeCategory) {
            case FileTypeCategory.Image:
                const reader = new FileReader()
                reader.onload = (event) => {
                    const target = event.target
                    if (target == null) throw new Error('file target reader is null')
                    const dataURL = target.result as string
                    if (dataURL == null) throw new Error('file dataURL is null')
                    const imageElement = document.createElement('img')
                    imageElement.src = dataURL
                    this.previewBoxElement.appendChild(imageElement)
                }
                reader.readAsDataURL(this.file)
                break
            case FileTypeCategory.Video:
                const videoElement = document.createElement('video')
                videoElement.src = URL.createObjectURL(this.file)
                videoElement.controls = false
                this.previewBoxElement.appendChild(videoElement)
                break
            case FileTypeCategory.Audio:
                const audioElement = document.createElement('audio')
                audioElement.src = URL.createObjectURL(this.file)
                audioElement.controls = true
                this.previewBoxElement.appendChild(audioElement)
                break
            default:
                const imageElement = document.createElement('img')
                imageElement.src = '/static/images/no_preview.webp'
                this.previewBoxElement.appendChild(imageElement)
        }
        this.selectedFilesContainerElement.prepend(this.selectedFilesRow)
        this.selectedFilesCardElement.classList.remove('hidden')
        void this.selectedFilesRow.offsetWidth
    }

    private addRemoveButtonListener() {
        this.removeFileBoxElement.addEventListener('click', () => {
            this.selectedFilesRow.style.animationName = 'fadeOutRightBig'
            this.selectedFilesRow.style.animationDuration = '0.4s'
            setTimeout(() => {
                UploadFilesHandler.uploadFilesMap.delete(this.id)
                this.selectedFilesRow.remove()
                if (UploadFilesHandler.uploadFilesMap.size === 0) {
                    UploadFilesHandler.uploadButtonElement.classList.remove('disabled')
                    UploadFilesHandler.uploadButtonElement.classList.add('opacity-0')
                    UploadFilesHandler.uploadButtonVisible = false
                    this.selectedFilesCardElement.classList.add('hidden')
                }
            }, 350)
        })
    }

    private uploadAllPendingFilesListener() {
        UploadFilesHandler.uploadButtonElement.addEventListener('click', () => {
            if (!UploadFilesHandler.uploadButtonVisible) return
            if (UploadFilesHandler.uploadButtonDisabled) return
            UploadFilesHandler.uploadAllSelectedFiles()
        })
    }

    private addMediaPreviewClickListener() {
        this.previewBoxElement.addEventListener('click', () => {
            const mediaRoot = previewModal.shadowRoot?.querySelector('#media-preview-slot') as HTMLDivElement
            switch (this.fileTypeCategory) {
                case FileTypeCategory.Video:
                    const videoElement = this.previewBoxElement.querySelector('video') as HTMLVideoElement
                    if (videoElement === null) return
                    const newVideoElement = videoElement.cloneNode(true) as HTMLVideoElement
                    newVideoElement.controls = true
                    mediaRoot.replaceChildren(newVideoElement)
                    previewModal.showModal()
                    break
                case FileTypeCategory.Audio:
                    const audioElement = this.previewBoxElement.querySelector('audio') as HTMLAudioElement
                    if (audioElement === null) return
                    const newAudioElement = audioElement.cloneNode(true) as HTMLAudioElement
                    newAudioElement.controls = true
                    mediaRoot.replaceChildren(newAudioElement)
                    previewModal.showModal()
                    break
                default:
                    const imageElement = this.previewBoxElement.querySelector('img') as HTMLImageElement
                    if (imageElement === null) return
                    const newImageElement = imageElement.cloneNode(true) as HTMLImageElement
                    mediaRoot.replaceChildren(newImageElement)
                    previewModal.showModal()
                    break
            }
        })
    }

    private calculateSpeedBetweenProgressUpdates(lastBytes: number, currentBytes: number,
                                                 lastUpdateInSeconds: number, currentUpdateInSeconds: number): number {
        const bytesDifference = currentBytes - lastBytes
        const timeDifference = currentUpdateInSeconds - lastUpdateInSeconds
        return this.calculateSpeedPerSecond(bytesDifference, timeDifference)
    }

    private calculateSpeedPerSecond(bytes: number, timeInSeconds: number): number {
        return bytes / timeInSeconds
    }

    private uploadFilePromise() {
        this.setStatus(UploadStatus.Uploading)
        this.progressBarElement.classList.remove('hidden')
        this.progressSpeedBoxElement.classList.remove('hidden')
        this.progressAlertElement.classList.add('hidden')
        return new Promise((resolve, reject) => {
            let lastProgressUpdateBytes = 0
            let lastProgressUpdateAt = new Date().getTime()
            this.uploadStartTime = new Date().getTime()
            this.xhr = new XMLHttpRequest()
            this.xhr.open('POST', UploadFilesHandler.uploadURL, true)
            this.xhr.responseType = 'json'
            this.xhr.onload = () => {
                if (this.xhr.status === 200) {
                    resolve(null)
                } else {
                    reject()
                }
            }
            this.xhr.onerror = () => {
                reject()
            }
            this.xhr.upload.onprogress = (event) => {
                if (event.lengthComputable) {
                    this.progressBarElement.value = event.loaded
                    this.progressBarElement.max = event.total
                    const lastUpdateInSeconds = Math.round(lastProgressUpdateAt / 1000)
                    const currentUpdateInSeconds = Math.round(new Date().getTime() / 1000)
                    const bytesPerSecond = this.calculateSpeedBetweenProgressUpdates(
                        lastProgressUpdateBytes, event.loaded,
                        lastUpdateInSeconds, currentUpdateInSeconds,
                    )
                    this.progressSpeedBoxElement.textContent = `${filesize(bytesPerSecond, {output: 'string'})}/s`
                }
            }
            this.xhr.upload.onloadend = (event) => {
                const timeTaken = new Date().getTime() - this.uploadStartTime
                const humanizedTimeTaken = humanizeDuration(timeTaken, {round: true})
                const bytesPerSecond = this.calculateSpeedPerSecond(event.loaded, timeTaken / 1000)
                const bytesPerSecondString = `${filesize(bytesPerSecond, {output: 'string'})}/s`
                this.progressSpeedBoxElement.textContent = `Took ${humanizedTimeTaken} @ ${bytesPerSecondString}`

            }
            const formData = new FormData()
            formData.append('file', this.file)
            this.xhr.send(formData)
        })
    }

    private handleUploadSuccess() {
        this.setStatus(UploadStatus.Done)
        const response = this.xhr.response as { error: boolean, message: string, url: string }
        if (response.error) {
            this.setStatus(UploadStatus.Error)
            this.handleUploadFailure()
            return
        }
        const uploadFileURLAnchorElement = document.createElement('a')
        uploadFileURLAnchorElement.href = response.url
        uploadFileURLAnchorElement.textContent = response.url
        uploadFileURLAnchorElement.target = '_blank'
        this.urlBoxElement.replaceChildren(uploadFileURLAnchorElement)

        this.progressBarElement.classList.add('hidden')
        this.progressAlertElement.classList.remove('hidden')
        this.progressAlertElement.classList.remove('alert-error')
        this.progressAlertElement.classList.add('alert-success')
        this.progressAlertElement.textContent = 'Upload successful'
    }


    private handleUploadFailure() {
        this.setStatus(UploadStatus.Error)
        this.progressBarElement.classList.add('hidden')
        this.progressAlertElement.classList.remove('hidden')
        this.progressAlertElement.classList.remove('alert-success')
        this.progressAlertElement.classList.add('alert-error')
        this.progressAlertElement.textContent = 'Failed to upload file! ' + (this.xhr.responseText !== '' ? this.xhr.responseText : this.xhr.statusText)
        this._errorMessage = this.xhr.responseText !== '' ? this.xhr.responseText : this.xhr.statusText
    }

    private setStatus(status: UploadStatus) {
        this._status = status
        switch (this._status) {
            case UploadStatus.Pending:
                this.statusBoxElement.textContent = 'Pending'
                break
            case UploadStatus.Uploading:
                this.statusBoxElement.textContent = 'Uploading'
                break
            case UploadStatus.Done:
                this.statusBoxElement.textContent = 'Completed'
                break
            case UploadStatus.Error:
                this.statusBoxElement.textContent = 'Error'
        }
    }


    public static uploadAllSelectedFiles() {
        UploadFilesHandler.uploadButtonElement.disabled = true
        UploadFilesHandler.uploadButtonElement.classList.add('disabled')
        UploadFilesHandler.uploadButtonDisabled = true
        let pendingFileUploads = Array<UploadFilesHandler>()
        UploadFilesHandler.uploadFilesMap.forEach((uploadFilesHandler) => {
            if (uploadFilesHandler._status === UploadStatus.Pending) {
                pendingFileUploads.push(uploadFilesHandler)
            }
        })

        if (pendingFileUploads.length === 0) {
            return
        }

        for (let i = 0; i < UploadFilesHandler.simultaneousUploads; i++) {
            const runNextTask = async () => {
                const uploadFile = pendingFileUploads.shift()
                if (uploadFile) {
                    try {
                        await uploadFile.uploadFilePromise()
                        uploadFile.handleUploadSuccess()
                    } catch (xhr: any) {
                        uploadFile.handleUploadFailure()
                    } finally {
                        await runNextTask()
                    }
                }
            }
            void runNextTask()
        }
    }

}