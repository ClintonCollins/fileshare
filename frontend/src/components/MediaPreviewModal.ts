const template = `
<link rel="stylesheet" href="/static/css/fileshare.css">
<dialog class="media-preview-modal">
    <article class="media-preview-modal-article">
        <a href="#" aria-label="Close" id="form-modal-close" class="close"></a>
        <div id="media-preview-slot">
          <slot name="media"></slot>
        </div>
    </article>
</dialog>
`

export class MediaPreviewModal extends HTMLDialogElement {
    template!: HTMLTemplateElement
    content!: DocumentFragment
    closeButton!: HTMLButtonElement
    dialog!: HTMLDialogElement
    article!: HTMLElement
    visible: boolean = false
    mediaRoot!: HTMLElement
    persist: boolean = false
    animateInOut: boolean = false

    constructor() {
        super()
        this.template = document.createElement('template')
        this.template.innerHTML = template
        this.content = this.template.content

        const shadowRoot = this.attachShadow({mode: 'open'})
        shadowRoot.appendChild(this.content.cloneNode(true))
        this.dialog = shadowRoot.querySelector('dialog') as HTMLDialogElement
        this.article = shadowRoot.querySelector('article') as HTMLElement
        this.closeButton = shadowRoot.querySelector('#form-modal-close') as HTMLButtonElement
        this.mediaRoot = shadowRoot.querySelector('#media-preview-slot') as HTMLElement

    }

    closeModal() {
        if (!this.animateInOut) {
            this.dialog.close()
            this.visible = false
            this.mediaRoot.replaceChildren()
            return
        }
        this.article.style.animation = 'fadeOutUp 0.3s forwards'
        setTimeout(() => {
            this.dialog.close()
            this.visible = false
            this.mediaRoot.replaceChildren()
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
        this.closeModal()
    }

    clickOffArticleListener = (event: MouseEvent) => {
        if (this.persist) return
        if (!this.visible) return
        const articleBounds = this.article.getBoundingClientRect()
        const inArticleBounds = (event.clientX >= articleBounds.left && event.clientX <= articleBounds.right)
            && (event.clientY >= articleBounds.top && event.clientY <= articleBounds.bottom)
        if (!inArticleBounds) {
            this.closeModal()
        }
    }

    connectedCallback() {
        this.persist = this.hasAttribute('persist')
        this.animateInOut = this.hasAttribute('animate')
        this.closeButton.addEventListener('click', this.closeButtonListener)
        document.addEventListener('keydown', (event) => {
            if (event.key === 'Escape' && this.visible) {
                this.closeModal()
            }
        })
        document.addEventListener('click', this.clickOffArticleListener)
    }

    disconnectedCallback() {
        this.closeButton.removeEventListener('click', this.closeButtonListener)
        document.removeEventListener('click', this.clickOffArticleListener)
    }

    static register() {
        customElements.define('media-preview-modal', MediaPreviewModal, {extends: 'dialog'})
    }
}