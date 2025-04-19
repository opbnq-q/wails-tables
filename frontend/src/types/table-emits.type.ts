export type TableEmits = {
    (e: 'onCreateOpen'): Promise<void> | void
    (e: 'onCreateClose', data: any): Promise<void> | void
    (e: 'onUpdateOpen'): Promise<void> | void
    (e: 'onUpdateClose', data: any): Promise<void> | void
    (e: 'onOpen'): Promise<void> | void
    (e: 'onClose', data: any): Promise<void> | void
    (e: 'onDelete', data: any): Promise<void> | void
    (e: 'onSaveUpdate', data: any): Promise<void> | void
    (e: 'onSaveCreate', data: any): Promise<void> | void
    (e: 'onSave', data: any): Promise<void> | void
    (e: 'onSearch', input: string): Promise<void> | void
}