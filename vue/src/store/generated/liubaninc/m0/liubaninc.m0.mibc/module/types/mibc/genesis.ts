/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { Itx } from '../mibc/itx'

export const protobufPackage = 'liubaninc.m0.mibc'

/** GenesisState defines the mibc module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  itxList: Itx[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  itxCount: number
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  portId: string
}

const baseGenesisState: object = { itxCount: 0, portId: '' }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.itxList) {
      Itx.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    if (message.itxCount !== 0) {
      writer.uint32(24).uint64(message.itxCount)
    }
    if (message.portId !== '') {
      writer.uint32(10).string(message.portId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.itxList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 2:
          message.itxList.push(Itx.decode(reader, reader.uint32()))
          break
        case 3:
          message.itxCount = longToNumber(reader.uint64() as Long)
          break
        case 1:
          message.portId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.itxList = []
    if (object.itxList !== undefined && object.itxList !== null) {
      for (const e of object.itxList) {
        message.itxList.push(Itx.fromJSON(e))
      }
    }
    if (object.itxCount !== undefined && object.itxCount !== null) {
      message.itxCount = Number(object.itxCount)
    } else {
      message.itxCount = 0
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = String(object.portId)
    } else {
      message.portId = ''
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.itxList) {
      obj.itxList = message.itxList.map((e) => (e ? Itx.toJSON(e) : undefined))
    } else {
      obj.itxList = []
    }
    message.itxCount !== undefined && (obj.itxCount = message.itxCount)
    message.portId !== undefined && (obj.portId = message.portId)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.itxList = []
    if (object.itxList !== undefined && object.itxList !== null) {
      for (const e of object.itxList) {
        message.itxList.push(Itx.fromPartial(e))
      }
    }
    if (object.itxCount !== undefined && object.itxCount !== null) {
      message.itxCount = object.itxCount
    } else {
      message.itxCount = 0
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = object.portId
    } else {
      message.portId = ''
    }
    return message
  }
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
