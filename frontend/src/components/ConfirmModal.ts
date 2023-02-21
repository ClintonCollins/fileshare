const template = `
	<link rel="stylesheet" href="/static/css/fileshare.css">
	<dialog id="confirm-modal">
		<article id="confirm-article">
			<a href="#" aria-label="Close" id="confirm-modal-close" class="close"></a>
			<h3>
				<slot id="confirm-modal-title" name="title">Title</slot>
			</h3>
			<p>
				<slot id="confirm-modal-message" name="message">Are you sure you want to do this?</slot>
			</p>
			<footer id="confirm-modal-action-footer">
				<button id="confirm-modal-cancel-button">Cancel</button>
				<button id="confirm-modal-confirm-button">Confirm</button>
			</footer>
		</article>
	</dialog>`

export class ConfirmModal extends HTMLDialogElement {
    template: HTMLTemplateElement
    content: DocumentFragment
    closeButton: HTMLButtonElement
    confirmButton: HTMLButtonElement
    cancelButton: HTMLButtonElement
    dialog: HTMLDialogElement
    article: HTMLElement
    animateInOut: boolean = false
    eventData: any
    visible: boolean = false
    persistEventHandler: Function | null = null
    titleSlotElement: HTMLElement
    messageSlotElement: HTMLElement


    static get observedAttributes() {
        return ['animate', 'persist', 'confirm-button-class', 'cancel-button-class',
            'event-data', 'title', 'message']
    }

    constructor() {
        super()
        this.template = document.createElement('template')
        this.template.innerHTML = template
        this.content = this.template.content

        const shadowRoot = this.attachShadow({mode: 'open'})
        shadowRoot.appendChild(this.content.cloneNode(true))
        this.closeButton = shadowRoot.querySelector('#confirm-modal-close') as HTMLButtonElement
        this.confirmButton = shadowRoot.querySelector('#confirm-modal-confirm-button') as HTMLButtonElement
        this.cancelButton = shadowRoot.querySelector('#confirm-modal-cancel-button') as HTMLButtonElement
        this.dialog = shadowRoot.querySelector('#confirm-modal') as HTMLDialogElement
        this.article = shadowRoot.querySelector('#confirm-article') as HTMLElement
        this.titleSlotElement = shadowRoot.querySelector('#confirm-modal-title') as HTMLElement
        this.messageSlotElement = shadowRoot.querySelector('#confirm-modal-message') as HTMLElement
    }

    closeModal = () => {
        if (!this.animateInOut) {
            this.dialog.close()
            this.visible = false
            return
        }
        this.article.style.animation = 'fadeOutUp 0.3s forwards'
        setTimeout(() => {
            this.dialog.close()
            this.visible = false
        }, 300)
    }

    showModal = () => {
        setTimeout(() => {
            this.visible = true
        }, 100)
        if (!this.animateInOut) {
            this.dialog.showModal()
            return
        }
        this.article.style.animation = 'fadeInDown 0.3s forwards'
        this.dialog.showModal()
    }

    closeButtonListener = () => {
        this.dispatchEvent(new CustomEvent('cancel', {detail: this.eventData}))
        this.closeModal()
    }

    cancelButtonListener = () => {
        this.dispatchEvent(new CustomEvent('cancel', {detail: this.eventData}))
        this.closeModal()
    }

    confirmButtonListener = () => {
        this.dispatchEvent(new CustomEvent('confirm', {detail: this.eventData}))
        this.closeModal()
    }

    connectedCallback() {
        this.closeButton.addEventListener('click', this.closeButtonListener)
        this.confirmButton.addEventListener('click', this.confirmButtonListener)
        this.cancelButton.addEventListener('click', this.cancelButtonListener)

        const persist = this.getAttribute('persist')?.toLowerCase() !== 'true'
        if (persist) {
            document.documentElement.addEventListener('click', this.noPersistHandler)
        }
        this.animateInOut = this.getAttribute('animate')?.toLowerCase() === 'true'
        this.eventData = this.getAttribute('event-data') ?? null

        this.confirmButton.className = this.getAttribute('confirm-button-class') ?? 'primary'
        this.cancelButton.className = this.getAttribute('cancel-button-class') ?? 'secondary'
    }

    disconnectedCallback() {
        this.closeButton.removeEventListener('click', this.closeButtonListener)
        this.confirmButton.removeEventListener('click', this.confirmButtonListener)
        this.cancelButton.removeEventListener('click', this.cancelButtonListener)
    }

    noPersistHandler = (event: any) => {
        if (!this.visible) return
        if (!this.article.contains(event.target as HTMLElement)) {
            this.closeModal()
        }
    }

    attributeChangedCallback(name: string, _: string, newValue: string) {
        switch (name) {
            case 'animate':
                this.animateInOut = newValue.toLowerCase() === 'true'
                break
            case 'persist':
                const persist = newValue.toLowerCase() === 'true'
                document.documentElement.removeEventListener('click', this.noPersistHandler)
                if (persist) {
                    return
                }
                document.documentElement.addEventListener('click', this.noPersistHandler)
                break
            case 'confirm-button-class':
                this.confirmButton.className = newValue
                break
            case 'cancel-button-class':
                this.cancelButton.className = newValue
                break
            case 'event-data':
                this.eventData = newValue
                break
            case 'title':
                this.titleSlotElement.textContent = newValue
                break
            case 'message':
                this.messageSlotElement.textContent = newValue
                break
        }
    }

    static register() {
        customElements.define('confirm-modal', ConfirmModal, {extends: 'dialog'})
    }
}