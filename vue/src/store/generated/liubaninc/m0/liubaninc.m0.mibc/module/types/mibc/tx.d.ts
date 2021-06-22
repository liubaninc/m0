import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "liubaninc.m0.mibc";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgSendIbcUTXO {
    sender: string;
    port: string;
    channelID: string;
    timeoutTimestamp: number;
    receiver: string;
    amount: string;
}
export interface MsgSendIbcUTXOResponse {
}
export declare const MsgSendIbcUTXO: {
    encode(message: MsgSendIbcUTXO, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendIbcUTXO;
    fromJSON(object: any): MsgSendIbcUTXO;
    toJSON(message: MsgSendIbcUTXO): unknown;
    fromPartial(object: DeepPartial<MsgSendIbcUTXO>): MsgSendIbcUTXO;
};
export declare const MsgSendIbcUTXOResponse: {
    encode(_: MsgSendIbcUTXOResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendIbcUTXOResponse;
    fromJSON(_: any): MsgSendIbcUTXOResponse;
    toJSON(_: MsgSendIbcUTXOResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendIbcUTXOResponse>): MsgSendIbcUTXOResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SendIbcUTXO(request: MsgSendIbcUTXO): Promise<MsgSendIbcUTXOResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SendIbcUTXO(request: MsgSendIbcUTXO): Promise<MsgSendIbcUTXOResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
