import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "liubaninc.m0.mibc";
export interface Itx {
    creator: string;
    id: number;
    hash: string;
    source: boolean;
    chain: string;
    log: string;
}
export declare const Itx: {
    encode(message: Itx, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Itx;
    fromJSON(object: any): Itx;
    toJSON(message: Itx): unknown;
    fromPartial(object: DeepPartial<Itx>): Itx;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
