/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
import { Itx } from '../mibc/itx';
export const protobufPackage = 'liubaninc.m0.mibc';
const baseGenesisState = { itxCount: 0, portId: '' };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.itxList) {
            Itx.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.itxCount !== 0) {
            writer.uint32(24).uint64(message.itxCount);
        }
        if (message.portId !== '') {
            writer.uint32(10).string(message.portId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.itxList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 2:
                    message.itxList.push(Itx.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.itxCount = longToNumber(reader.uint64());
                    break;
                case 1:
                    message.portId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.itxList = [];
        if (object.itxList !== undefined && object.itxList !== null) {
            for (const e of object.itxList) {
                message.itxList.push(Itx.fromJSON(e));
            }
        }
        if (object.itxCount !== undefined && object.itxCount !== null) {
            message.itxCount = Number(object.itxCount);
        }
        else {
            message.itxCount = 0;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = String(object.portId);
        }
        else {
            message.portId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.itxList) {
            obj.itxList = message.itxList.map((e) => (e ? Itx.toJSON(e) : undefined));
        }
        else {
            obj.itxList = [];
        }
        message.itxCount !== undefined && (obj.itxCount = message.itxCount);
        message.portId !== undefined && (obj.portId = message.portId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.itxList = [];
        if (object.itxList !== undefined && object.itxList !== null) {
            for (const e of object.itxList) {
                message.itxList.push(Itx.fromPartial(e));
            }
        }
        if (object.itxCount !== undefined && object.itxCount !== null) {
            message.itxCount = object.itxCount;
        }
        else {
            message.itxCount = 0;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = object.portId;
        }
        else {
            message.portId = '';
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
