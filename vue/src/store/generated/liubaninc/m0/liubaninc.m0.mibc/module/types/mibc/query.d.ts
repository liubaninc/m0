import { Reader, Writer } from 'protobufjs/minimal';
import { Itx } from '../mibc/itx';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "liubaninc.m0.mibc";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetItxRequest {
    id: number;
}
export interface QueryGetItxResponse {
    Itx: Itx | undefined;
}
export interface QueryAllItxRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllItxResponse {
    Itx: Itx[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetItxRequest: {
    encode(message: QueryGetItxRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetItxRequest;
    fromJSON(object: any): QueryGetItxRequest;
    toJSON(message: QueryGetItxRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetItxRequest>): QueryGetItxRequest;
};
export declare const QueryGetItxResponse: {
    encode(message: QueryGetItxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetItxResponse;
    fromJSON(object: any): QueryGetItxResponse;
    toJSON(message: QueryGetItxResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetItxResponse>): QueryGetItxResponse;
};
export declare const QueryAllItxRequest: {
    encode(message: QueryAllItxRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllItxRequest;
    fromJSON(object: any): QueryAllItxRequest;
    toJSON(message: QueryAllItxRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllItxRequest>): QueryAllItxRequest;
};
export declare const QueryAllItxResponse: {
    encode(message: QueryAllItxResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllItxResponse;
    fromJSON(object: any): QueryAllItxResponse;
    toJSON(message: QueryAllItxResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllItxResponse>): QueryAllItxResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a itx by id. */
    Itx(request: QueryGetItxRequest): Promise<QueryGetItxResponse>;
    /** Queries a list of itx items. */
    ItxAll(request: QueryAllItxRequest): Promise<QueryAllItxResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Itx(request: QueryGetItxRequest): Promise<QueryGetItxResponse>;
    ItxAll(request: QueryAllItxRequest): Promise<QueryAllItxResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
