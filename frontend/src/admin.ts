import {ConfirmModal} from './components/ConfirmModal'
import {FormModal} from './components/FormModal'
import {TableSelectActions} from './utils/TableSelectActions'

ConfirmModal.register()
FormModal.register()

const selectAllCheckbox = document.querySelector('.select-all') as HTMLInputElement
const actionForm = document.querySelector('.action-form') as HTMLFormElement
const actionDetails = document.querySelector('.select-action') as HTMLDetailsElement
const actionDelete = document.querySelector('.delete-all-selected-action') as HTMLAnchorElement
const allItemCheckboxes = document.querySelectorAll('.item-selected-checkbox') as NodeListOf<HTMLInputElement>
const searchButton = document.querySelector('.search-button') as HTMLButtonElement
const confirmModalAccounts = document.querySelector('confirm-modal#accounts-confirm') as ConfirmModal
const confirmModalInvitations = document.querySelector('confirm-modal#invitations-confirm') as ConfirmModal
const tableSelectActions = new TableSelectActions(selectAllCheckbox, allItemCheckboxes)
const searchInput = document.querySelector('.search') as HTMLInputElement

const createInvitationButton = document.getElementById('create-invitation-button') as HTMLButtonElement
if (createInvitationButton) {
    createInvitationButton.addEventListener('click', (event) => {
        event.preventDefault()
        const formModal = document.querySelector('form-modal') as FormModal
        const emailInput = formModal.querySelector('#email') as HTMLInputElement
        formModal.showModal()
        setTimeout(() => {
            emailInput.focus()
        }, 50)
    })
}

if (selectAllCheckbox && allItemCheckboxes) {
    const tableSelectActions = new TableSelectActions(selectAllCheckbox, allItemCheckboxes)
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

actionDelete.addEventListener('click', (event) => {
    event.preventDefault()

    if (confirmModalAccounts) {
        const selectedCount = tableSelectActions?.selectedIds.size
        const plural = selectedCount > 1 ? 'accounts' : 'account'
        confirmModalAccounts.setAttribute('title', 'Delete all selected accounts')
        confirmModalAccounts.setAttribute('message',
            `Are you sure you want to delete ${selectedCount} selected ${plural}?`)
        confirmModalAccounts.setAttribute('confirm-button-class', 'button-error')
        confirmModalAccounts.setAttribute('event-data', 'delete-accounts')
        confirmModalAccounts.showModal()
    }
    if (confirmModalInvitations) {
        const selectedCount = tableSelectActions?.selectedIds.size
        const plural = selectedCount > 1 ? 'invitations' : 'invitation'
        confirmModalInvitations.setAttribute('title', 'Delete all selected invitations')
        confirmModalInvitations.setAttribute('message',
            `Are you sure you want to delete ${selectedCount} selected ${plural}?`)
        confirmModalInvitations.setAttribute('confirm-button-class', 'button-error')
        confirmModalInvitations.setAttribute('event-data', 'delete-invitations')
        confirmModalInvitations.showModal()
    }
})

if (confirmModalInvitations) {
    confirmModalInvitations.addEventListener('cancel', () => {
        actionDetails.removeAttribute('open')
    })

    confirmModalInvitations.addEventListener('confirm', (event: any) => {
        actionDetails.removeAttribute('open')
        if (!event.detail) return
        switch (event.detail) {
            case 'delete-invitations':
                actionForm.action = '/admin/invitations/delete'
                actionForm.method = 'POST'
                actionForm.submit()
                break
        }
    })
}

if (confirmModalAccounts) {
    confirmModalAccounts.addEventListener('cancel', () => {
        actionDetails.removeAttribute('open')
    })

    confirmModalAccounts.addEventListener('confirm', (event: any) => {
        actionDetails.removeAttribute('open')
        if (!event.detail) return
        switch (event.detail) {
            case 'delete-accounts':
                actionForm.action = '/admin/accounts/delete'
                actionForm.method = 'POST'
                actionForm.submit()
                break
        }
    })
}

function submitSearch() {
    if (!searchInput || !searchInput.dataset.url) return
    actionForm.action = searchInput.dataset.url
    actionForm.method = 'GET'
    actionForm.submit()
}

searchButton.addEventListener('click', (event) => {
    event.preventDefault()
    submitSearch()
})

document.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
        event.preventDefault()
        if (document.activeElement === searchInput) {
            submitSearch()
            return
        }
    }
})