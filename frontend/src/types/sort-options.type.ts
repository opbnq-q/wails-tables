import type { IEntity } from "./entity.type";
import type { Sorting } from "./sorting.type";

export type SortOptions<T extends IEntity> = {
    [key in keyof Partial<T>]: Sorting
}