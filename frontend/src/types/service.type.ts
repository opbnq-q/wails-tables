export interface IService<T> {
    read(id: number): Promise<T | null>
    readAll(): Promise<T[]>
    create(data: T): Promise<void | never>
    update(data: T): Promise<void | never>
    delete(id: number): Promise<void | never>
}