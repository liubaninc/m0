/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Itx } from '../mibc/itx'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'liubaninc.m0.mibc'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetItxRequest {
  id: number
}

export interface QueryGetItxResponse {
  Itx: Itx | undefined
}

export interface QueryAllItxRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllItxResponse {
  Itx: Itx[]
  pagination: PageResponse | undefined
}

const baseQueryGetItxRequest: object = { id: 0 }

export const QueryGetItxRequest = {
  encode(message: QueryGetItxRequest, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetItxRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetItxRequest } as QueryGetItxRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetItxRequest {
    const message = { ...baseQueryGetItxRequest } as QueryGetItxRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    return message
  },

  toJSON(message: QueryGetItxRequest): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetItxRequest>): QueryGetItxRequest {
    const message = { ...baseQueryGetItxRequest } as QueryGetItxRequest
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    return message
  }
}

const baseQueryGetItxResponse: object = {}

export const QueryGetItxResponse = {
  encode(message: QueryGetItxResponse, writer: Writer = Writer.create()): Writer {
    if (message.Itx !== undefined) {
      Itx.encode(message.Itx, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetItxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetItxResponse } as QueryGetItxResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Itx = Itx.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetItxResponse {
    const message = { ...baseQueryGetItxResponse } as QueryGetItxResponse
    if (object.Itx !== undefined && object.Itx !== null) {
      message.Itx = Itx.fromJSON(object.Itx)
    } else {
      message.Itx = undefined
    }
    return message
  },

  toJSON(message: QueryGetItxResponse): unknown {
    const obj: any = {}
    message.Itx !== undefined && (obj.Itx = message.Itx ? Itx.toJSON(message.Itx) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetItxResponse>): QueryGetItxResponse {
    const message = { ...baseQueryGetItxResponse } as QueryGetItxResponse
    if (object.Itx !== undefined && object.Itx !== null) {
      message.Itx = Itx.fromPartial(object.Itx)
    } else {
      message.Itx = undefined
    }
    return message
  }
}

const baseQueryAllItxRequest: object = {}

export const QueryAllItxRequest = {
  encode(message: QueryAllItxRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllItxRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllItxRequest } as QueryAllItxRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllItxRequest {
    const message = { ...baseQueryAllItxRequest } as QueryAllItxRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllItxRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllItxRequest>): QueryAllItxRequest {
    const message = { ...baseQueryAllItxRequest } as QueryAllItxRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllItxResponse: object = {}

export const QueryAllItxResponse = {
  encode(message: QueryAllItxResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.Itx) {
      Itx.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllItxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllItxResponse } as QueryAllItxResponse
    message.Itx = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Itx.push(Itx.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllItxResponse {
    const message = { ...baseQueryAllItxResponse } as QueryAllItxResponse
    message.Itx = []
    if (object.Itx !== undefined && object.Itx !== null) {
      for (const e of object.Itx) {
        message.Itx.push(Itx.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllItxResponse): unknown {
    const obj: any = {}
    if (message.Itx) {
      obj.Itx = message.Itx.map((e) => (e ? Itx.toJSON(e) : undefined))
    } else {
      obj.Itx = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllItxResponse>): QueryAllItxResponse {
    const message = { ...baseQueryAllItxResponse } as QueryAllItxResponse
    message.Itx = []
    if (object.Itx !== undefined && object.Itx !== null) {
      for (const e of object.Itx) {
        message.Itx.push(Itx.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a itx by id. */
  Itx(request: QueryGetItxRequest): Promise<QueryGetItxResponse>
  /** Queries a list of itx items. */
  ItxAll(request: QueryAllItxRequest): Promise<QueryAllItxResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Itx(request: QueryGetItxRequest): Promise<QueryGetItxResponse> {
    const data = QueryGetItxRequest.encode(request).finish()
    const promise = this.rpc.request('liubaninc.m0.mibc.Query', 'Itx', data)
    return promise.then((data) => QueryGetItxResponse.decode(new Reader(data)))
  }

  ItxAll(request: QueryAllItxRequest): Promise<QueryAllItxResponse> {
    const data = QueryAllItxRequest.encode(request).finish()
    const promise = this.rpc.request('liubaninc.m0.mibc.Query', 'ItxAll', data)
    return promise.then((data) => QueryAllItxResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
