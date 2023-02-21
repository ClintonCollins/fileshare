const template = `
<link rel="stylesheet" href="/static/css/fileshare.css">
<dialog id="form-modal">
	<article id="form-article">
		<a href="#" aria-label="Close" id="form-modal-close" class="close"></a>
		<h3>
			<slot name="title">Title</slot>
		</h3>
		<slot name="form">Form</slot>
	</article>
</dialog>
`

export class FormModal extends HTMLDialogElement {
    template!: HTMLTemplateElement
    content!: DocumentFragment
    closeButton!: HTMLButtonElement
    dialog!: HTMLDialogElement
    article!: HTMLElement
    animateInOut: boolean = false
    eventData: any
    visible: boolean = false
    persistEventHandler: Function | null = null
    titleSlotElement!: HTMLElement


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
        this.closeButton = shadowRoot.querySelector('#form-modal-close') as HTMLButtonElement
        this.dialog = shadowRoot.querySelector('#form-modal') as HTMLDialogElement
        this.article = shadowRoot.querySelector('#form-article') as HTMLElement
        this.titleSlotElement = shadowRoot.querySelector('#form-modal-title') as HTMLElement
    }

    closeModal() {
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

    showModal() {
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

    connectedCallback() {
        this.closeButton.addEventListener('click', this.closeButtonListener)

        const persist: boolean = this.getAttribute('persist')?.toLowerCase() !== 'true'
        if (persist) {
            document.documentElement.addEventListener('click', this.noPersistHandler)
        }
        this.animateInOut = this.getAttribute('animate')?.toLowerCase() === 'true'
        this.eventData = this.getAttribute('event-data') ?? null
    }

    disconnectedCallback() {
        this.closeButton.removeEventListener('click', this.closeButtonListener)
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
            case 'event-data':
                this.eventData = newValue
                break
            case 'title':
                this.titleSlotElement.textContent = newValue
                break
            case 'form':
                break
        }
    }

    static register() {
        customElements.define('form-modal', FormModal, {extends: 'dialog'})
    }
}