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
    let $resultPromise = $Call.ByID(2983034394) as any;
    return $resultPromise;
}

export function Create(item: $models.Workshop): Promise<$models.Workshop> & { cancel(): void } {
    let $resultPromise = $Call.ByID(982947003, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Delete(id: number): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3145051352, id) as any;
    return $resultPromise;
}

export function GetAll(): Promise<($models.Workshop | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(806282052) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetById(id: number): Promise<$models.Workshop | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1829452217, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SearchByAllTextFields(phrase: string): Promise<($models.Workshop | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(283835139, phrase) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SortedByOrder(fieldsSortingOrder: utils$0.SortField[]): Promise<($models.Workshop | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3581052397, fieldsSortingOrder) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Update(item: $models.Workshop): Promise<$models.Workshop> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1795717546, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = models$0.Workshop.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
