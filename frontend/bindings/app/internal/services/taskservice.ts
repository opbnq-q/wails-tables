// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as models$0 from "../models/models.js";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as utils$0 from "../utils/models.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function Count(): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2626684778) as any;
    return $resultPromise;
}

export function Create(item: $models.Task): Promise<$models.Task> & { cancel(): void } {
    let $resultPromise = $Call.ByID(893779179, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Delete(id: number): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3033756968, id) as any;
    return $resultPromise;
}

export function GetAll(): Promise<($models.Task | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1660059028) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetById(id: number): Promise<$models.Task | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3152185033, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SearchByAllTextFields(phrase: string): Promise<($models.Task | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3905257587, phrase) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SortedByOrder(fieldsSortingOrder: utils$0.SortField[]): Promise<($models.Task | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2942045693, fieldsSortingOrder) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Update(item: $models.Task): Promise<$models.Task> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3997472442, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = models$0.Task.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
