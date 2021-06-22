import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "liubaninc.m0.mibc";
export interface MibcPacketData {
    noData: NoData | undefined;
    /** this line is used by starport scaffolding # ibc/packet/proto/field */
    ibcUTXOPacket: IbcUTXOPacketData | undefined;
}
export interface NoData {
}
/**
 * this line is used by starport scaffolding # ibc/packet/proto/message
 * IbcUTXOPacketData defines a struct for the packet payload
 */
export interface IbcUTXOPacketData {
    creator: string;
    receiver: string;
    amount: string;
}
/** IbcUTXOPacketAck defines a struct for the packet acknowledgment */
export interface IbcUTXOPacketAck {
    hash: string;
}
export declare const MibcPacketData: {
    encode(message: MibcPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MibcPacketData;
    fromJSON(object: any): MibcPacketData;
    toJSON(message: MibcPacketData): unknown;
    fromPartial(object: DeepPartial<MibcPacketData>): MibcPacketData;
};
export declare const NoData: {
    encode(_: NoData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NoData;
    fromJSON(_: any): NoData;
    toJSON(_: NoData): unknown;
    fromPartial(_: DeepPartial<NoData>): NoData;
};
export declare const IbcUTXOPacketData: {
    encode(message: IbcUTXOPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IbcUTXOPacketData;
    fromJSON(object: any): IbcUTXOPacketData;
    toJSON(message: IbcUTXOPacketData): unknown;
    fromPartial(object: DeepPartial<IbcUTXOPacketData>): IbcUTXOPacketData;
};
export declare const IbcUTXOPacketAck: {
    encode(message: IbcUTXOPacketAck, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): IbcUTXOPacketAck;
    fromJSON(object: any): IbcUTXOPacketAck;
    toJSON(message: IbcUTXOPacketAck): unknown;
    fromPartial(object: DeepPartial<IbcUTXOPacketAck>): IbcUTXOPacketAck;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
