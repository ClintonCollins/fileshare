export class TableSelectActions {
    private selectAllCheckboxElement: HTMLInputElement
    private itemCheckboxElements: NodeListOf<HTMLInputElement>
    private bulkSelect: boolean = false
    private selectedItems: Set<string> = new Set()

    constructor(selectAllCheckboxElement: HTMLInputElement, itemCheckboxElements: NodeListOf<HTMLInputElement>) {
        this.selectAllCheckboxElement = selectAllCheckboxElement
        this.itemCheckboxElements = itemCheckboxElements

        this.addSelectAllCheckboxEventListener()
        this.addCheckboxEventListeners()
    }

    public get selectedIds(): Set<string> {
        return this.selectedItems
    }

    private emitChange = () => {
        const changeEvent = new CustomEvent('table-select-change', {detail: this.selectedIds})
        document.dispatchEvent(changeEvent)
    }

    private selectAllCheckboxChanged = (event: Event) => {
        const target = event.target as HTMLInputElement
        this.bulkSelect = true
        this.itemCheckboxElements.forEach((checkbox) => {
            if (!checkbox.dataset.id) return
            if (target.checked) {
                this.selectedItems.add(checkbox.dataset.id)
            } else {
                this.selectedItems.delete(checkbox.dataset.id)
            }
            checkbox.checked = target.checked
        })
        this.calculateChanges()
        this.bulkSelect = false
    }

    private calculateChanges = () => {
        if (this.selectedItems.size === 0) {
            this.selectAllCheckboxElement.checked = false
        }
        if (this.selectedItems.size === this.itemCheckboxElements.length) {
            this.selectAllCheckboxElement.checked = true
        }
        this.emitChange()
    }

    private itemCheckboxChanged = (event: Event) => {
        if (this.bulkSelect) return
        const target = event.target as HTMLInputElement
        if (!target.dataset.id) return
        if (target.checked) {
            this.selectedItems.add(target.dataset.id)
        } else {
            this.selectedItems.delete(target.dataset.id)
        }
        this.calculateChanges()
    }

    private addSelectAllCheckboxEventListener = () => {
        this.selectAllCheckboxElement.addEventListener('change', this.selectAllCheckboxChanged)
    }

    private addCheckboxEventListeners = () => {
        this.itemCheckboxElements.forEach(checkbox => {
            checkbox.addEventListener('change', this.itemCheckboxChanged)
        })
    }

    public removeSelectAllCheckboxEventListener = () => {
        this.selectAllCheckboxElement.removeEventListener('change', this.selectAllCheckboxChanged)
    }

    public removeCheckboxEventListeners = () => {
        this.itemCheckboxElements.forEach(checkbox => {
            checkbox.removeEventListener('change', this.itemCheckboxChanged)
        })
    }
}