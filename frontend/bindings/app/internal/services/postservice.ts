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
import * as $models from "./models.js";

export function Count(): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3109924027) as any;
    return $resultPromise;
}

export function Create(item: $models.Post): Promise<$models.Post> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1443399856, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Delete(id: number): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2924549135, id) as any;
    return $resultPromise;
}

export function ExportToExcel(): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(75322242) as any;
    return $resultPromise;
}

export function GetAll(): Promise<($models.Post | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(65691059) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetById(id: number): Promise<$models.Post | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4074736792, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SortedByOrder(fieldsSortOrder: { [_: string]: string }): Promise<($models.Post | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(471862744, fieldsSortOrder) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Update(item: $models.Post): Promise<$models.Post> & { cancel(): void } {
    let $resultPromise = $Call.ByID(137798821, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = models$0.Post.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
