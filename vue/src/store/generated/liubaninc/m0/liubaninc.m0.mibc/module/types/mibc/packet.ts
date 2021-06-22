/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'liubaninc.m0.mibc'

export interface MibcPacketData {
  noData: NoData | undefined
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  ibcUTXOPacket: IbcUTXOPacketData | undefined
}

export interface NoData {}

/**
 * this line is used by starport scaffolding # ibc/packet/proto/message
 * IbcUTXOPacketData defines a struct for the packet payload
 */
export interface IbcUTXOPacketData {
  creator: string
  receiver: string
  amount: string
}

/** IbcUTXOPacketAck defines a struct for the packet acknowledgment */
export interface IbcUTXOPacketAck {
  hash: string
}

const baseMibcPacketData: object = {}

export const MibcPacketData = {
  encode(message: MibcPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim()
    }
    if (message.ibcUTXOPacket !== undefined) {
      IbcUTXOPacketData.encode(message.ibcUTXOPacket, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MibcPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMibcPacketData } as MibcPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32())
          break
        case 2:
          message.ibcUTXOPacket = IbcUTXOPacketData.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MibcPacketData {
    const message = { ...baseMibcPacketData } as MibcPacketData
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData)
    } else {
      message.noData = undefined
    }
    if (object.ibcUTXOPacket !== undefined && object.ibcUTXOPacket !== null) {
      message.ibcUTXOPacket = IbcUTXOPacketData.fromJSON(object.ibcUTXOPacket)
    } else {
      message.ibcUTXOPacket = undefined
    }
    return message
  },

  toJSON(message: MibcPacketData): unknown {
    const obj: any = {}
    message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined)
    message.ibcUTXOPacket !== undefined && (obj.ibcUTXOPacket = message.ibcUTXOPacket ? IbcUTXOPacketData.toJSON(message.ibcUTXOPacket) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MibcPacketData>): MibcPacketData {
    const message = { ...baseMibcPacketData } as MibcPacketData
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData)
    } else {
      message.noData = undefined
    }
    if (object.ibcUTXOPacket !== undefined && object.ibcUTXOPacket !== null) {
      message.ibcUTXOPacket = IbcUTXOPacketData.fromPartial(object.ibcUTXOPacket)
    } else {
      message.ibcUTXOPacket = undefined
    }
    return message
  }
}

const baseNoData: object = {}

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNoData } as NoData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData
    return message
  },

  toJSON(_: NoData): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData
    return message
  }
}

const baseIbcUTXOPacketData: object = { creator: '', receiver: '', amount: '' }

export const IbcUTXOPacketData = {
  encode(message: IbcUTXOPacketData, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.receiver !== '') {
      writer.uint32(18).string(message.receiver)
    }
    if (message.amount !== '') {
      writer.uint32(26).string(message.amount)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): IbcUTXOPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseIbcUTXOPacketData } as IbcUTXOPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.receiver = reader.string()
          break
        case 3:
          message.amount = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): IbcUTXOPacketData {
    const message = { ...baseIbcUTXOPacketData } as IbcUTXOPacketData
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = String(object.receiver)
    } else {
      message.receiver = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount)
    } else {
      message.amount = ''
    }
    return message
  },

  toJSON(message: IbcUTXOPacketData): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.receiver !== undefined && (obj.receiver = message.receiver)
    message.amount !== undefined && (obj.amount = message.amount)
    return obj
  },

  fromPartial(object: DeepPartial<IbcUTXOPacketData>): IbcUTXOPacketData {
    const message = { ...baseIbcUTXOPacketData } as IbcUTXOPacketData
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.receiver !== undefined && object.receiver !== null) {
      message.receiver = object.receiver
    } else {
      message.receiver = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = ''
    }
    return message
  }
}

const baseIbcUTXOPacketAck: object = { hash: '' }

export const IbcUTXOPacketAck = {
  encode(message: IbcUTXOPacketAck, writer: Writer = Writer.create()): Writer {
    if (message.hash !== '') {
      writer.uint32(10).string(message.hash)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): IbcUTXOPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseIbcUTXOPacketAck } as IbcUTXOPacketAck
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.hash = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): IbcUTXOPacketAck {
    const message = { ...baseIbcUTXOPacketAck } as IbcUTXOPacketAck
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = String(object.hash)
    } else {
      message.hash = ''
    }
    return message
  },

  toJSON(message: IbcUTXOPacketAck): unknown {
    const obj: any = {}
    message.hash !== undefined && (obj.hash = message.hash)
    return obj
  },

  fromPartial(object: DeepPartial<IbcUTXOPacketAck>): IbcUTXOPacketAck {
    const message = { ...baseIbcUTXOPacketAck } as IbcUTXOPacketAck
    if (object.hash !== undefined && object.hash !== null) {
      message.hash = object.hash
    } else {
      message.hash = ''
    }
    return message
  }
}

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
