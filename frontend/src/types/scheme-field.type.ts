import type { IEntity } from "./entity.type"
import type { PrimitiveFieldType } from "./primitive-field-type.type"

export interface ISchemeField<T extends IEntity> {
    type?: {
        primitive?: PrimitiveFieldType
        nested?: {
            field: string[]
            values: T[]
        }
    }
    many?: boolean
    russian?: string
    hidden?: boolean
    readonly?: boolean
    customWindow?: {
        common?: boolean
        create?: boolean
        update?: boolean
    }
}