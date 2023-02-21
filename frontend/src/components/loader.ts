export function createFullscreenLoader(text: string) {
    const documentFragment = document.createDocumentFragment()
    const divElement = document.createElement('div')
    divElement.classList.add('full-screen-loader')
    const iconSpinnerElement = document.createElement('i')
    iconSpinnerElement.className = 'fa-solid fa-circle-notch fa-spin'
    const spanElement = document.createElement('span')
    spanElement.textContent = text
    divElement.appendChild(iconSpinnerElement)
    divElement.appendChild(spanElement)
    documentFragment.appendChild(divElement)
    document.documentElement.appendChild(documentFragment)
    console.log(documentFragment)
}

export function removeFullscreenLoader() {
    const divElement = document.querySelector('.full-screen-loader')
    if (divElement) {
        divElement.remove()
    }
}