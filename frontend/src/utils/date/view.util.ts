import type { PrimitiveFieldType } from "../../types/primitive-field-type.type"
import { toDate } from "./converters.util";

export const viewDate = (data: unknown, type?: PrimitiveFieldType) => {
    if (type !== 'date') return data;
    return toDate(data as number | Date).toLocaleDateString('ru-RU', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
    })
}
