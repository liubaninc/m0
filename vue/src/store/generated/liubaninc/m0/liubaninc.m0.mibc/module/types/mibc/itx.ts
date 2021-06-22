/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'liubaninc.m0.mibc'

export interface Itx {
  creator: string
  id: number
  hash: string
  source: boolean
  chain: string
  log: string
}

const baseItx: object = { creator: '', id: 0, hash: '', source: false, chain: '', log: '' }

export const Itx = {
  encode(message: Itx, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id)
    }
    if (message.hash !== '') {
      writer.uint32(26).string(message.hash)
    }
    if (message.source === true) {
      writer.uint32(32).bool(message.source)
    }
    if (message.chain !== '') {
      writer.uint32(42).string(message.chain)
    }
    if (message.log !== '') {
      writer.uint32(50).string(message.log)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Itx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseItx } as Itx
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.hash = reader.string()
          break
        case 4:
          message.source = reader.bool()
          break
        case 5:
          message.chain = reader.string()
          break
        case 6:
          message.log = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Itx {
    const message = { ...baseItx } as Itx
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = String(object.hash)
    } else {
      message.hash = ''
    }
    if (object.source !== undefined && object.source !== null) {
      message.source = Boolean(object.source)
    } else {
      message.source = false
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain)
    } else {
      message.chain = ''
    }
    if (object.log !== undefined && object.log !== null) {
      message.log = String(object.log)
    } else {
      message.log = ''
    }
    return message
  },

  toJSON(message: Itx): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.hash !== undefined && (obj.hash = message.hash)
    message.source !== undefined && (obj.source = message.source)
    message.chain !== undefined && (obj.chain = message.chain)
    message.log !== undefined && (obj.log = message.log)
    return obj
  },

  fromPartial(object: DeepPartial<Itx>): Itx {
    const message = { ...baseItx } as Itx
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = object.hash
    } else {
      message.hash = ''
    }
    if (object.source !== undefined && object.source !== null) {
      message.source = object.source
    } else {
      message.source = false
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain
    } else {
      message.chain = ''
    }
    if (object.log !== undefined && object.log !== null) {
      message.log = object.log
    } else {
      message.log = ''
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
