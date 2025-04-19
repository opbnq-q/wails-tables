import type { IEntity } from "./entity.type";
import type { ISchemeField } from "./scheme-field.type";

export type Scheme<T extends IEntity, S extends IEntity=any> = { entityId: string } & Record<keyof T, ISchemeField<S>>