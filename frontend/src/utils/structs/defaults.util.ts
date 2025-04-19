import type { IEntity } from "../../types/entity.type";
import type { Scheme } from "../../types/scheme.type";
import { toTimestamp } from "../date/converters.util";

export const getDefaultValues = <T extends IEntity>(scheme: Scheme<T>) => {
    const keys = Object.keys(scheme) as (keyof typeof scheme)[]
    let obj: any = {}

    for (let key of keys) {
        const primitive = scheme[key]?.type?.primitive
        if (scheme[key].hidden) continue;
        if ((primitive == 'string' || primitive == 'multiple')) {
            obj[key] = ''
        } else if (primitive == 'date') {
            obj[key] = toTimestamp(new Date)
        } else if (primitive == 'boolean') {
            obj[key] = false
        } else if (primitive == 'number') {
            obj[key] = 0
        } else if (scheme[key]?.many) {
            obj[key] = []
        } else if (scheme[key]?.type?.nested?.values) {
            obj[key] = scheme[key].type.nested.values[0]
        }
    }

    return obj as T
}