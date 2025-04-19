import type { IEntity } from "./entity.type";

export type ValidateResult = { status: "error" | "success", message?: string }

export type Validate<T extends IEntity> = (data: T, mode: "update" | "create") => ValidateResult | Promise<ValidateResult>