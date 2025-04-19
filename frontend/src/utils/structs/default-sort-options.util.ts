import type { IEntity } from "../../types/entity.type";
import type { Scheme } from "../../types/scheme.type";
import type { SortOptions } from "../../types/sort-options.type";

export const getDefaultSortOptions = <T extends IEntity>(scheme: Scheme<T>): SortOptions<T> => {
    const keys = Object.keys(scheme) as (keyof T)[]
    const result: Partial<SortOptions<T>> = {}

    keys.forEach(key => {
        if (!scheme[key].hidden && key !== 'entityId' && !scheme[key].many) {
            result[key] = 'NONE'
        }
    })

    return result as SortOptions<T>
}