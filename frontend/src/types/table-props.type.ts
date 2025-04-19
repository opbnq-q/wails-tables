import type { IEntity } from "./entity.type";
import type { Scheme } from "./scheme.type";
import type { IService } from "./service.type";
import type { Validate } from "./validate.type";

export interface TableProps<T extends IEntity> {
    service: IService<T>
    scheme: Scheme<T>
    name?: string
    getDefaults: () => Partial<T>
    validate: Validate<T>
    load(): Promise<T[]>
    items: T[]
    colorize?: (data: T) => string
}