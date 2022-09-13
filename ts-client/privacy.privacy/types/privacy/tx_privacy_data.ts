/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface TxPrivacyData {
  index: string;
  creator: string;
  lockTime: number;
  fee: number;
  info: Uint8Array;
  sigPubKey: Uint8Array;
  sig: Uint8Array;
  proof: Uint8Array;
  messages: Uint8Array;
}

const baseTxPrivacyData: object = {
  index: "",
  creator: "",
  lockTime: 0,
  fee: 0,
};

export const TxPrivacyData = {
  encode(message: TxPrivacyData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.lockTime !== 0) {
      writer.uint32(24).uint64(message.lockTime);
    }
    if (message.fee !== 0) {
      writer.uint32(32).uint64(message.fee);
    }
    if (message.info.length !== 0) {
      writer.uint32(42).bytes(message.info);
    }
    if (message.sigPubKey.length !== 0) {
      writer.uint32(50).bytes(message.sigPubKey);
    }
    if (message.sig.length !== 0) {
      writer.uint32(58).bytes(message.sig);
    }
    if (message.proof.length !== 0) {
      writer.uint32(66).bytes(message.proof);
    }
    if (message.messages.length !== 0) {
      writer.uint32(74).bytes(message.messages);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TxPrivacyData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTxPrivacyData } as TxPrivacyData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.lockTime = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.fee = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.info = reader.bytes();
          break;
        case 6:
          message.sigPubKey = reader.bytes();
          break;
        case 7:
          message.sig = reader.bytes();
          break;
        case 8:
          message.proof = reader.bytes();
          break;
        case 9:
          message.messages = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TxPrivacyData {
    const message = { ...baseTxPrivacyData } as TxPrivacyData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.lockTime !== undefined && object.lockTime !== null) {
      message.lockTime = Number(object.lockTime);
    } else {
      message.lockTime = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = Number(object.fee);
    } else {
      message.fee = 0;
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = bytesFromBase64(object.info);
    }
    if (object.sigPubKey !== undefined && object.sigPubKey !== null) {
      message.sigPubKey = bytesFromBase64(object.sigPubKey);
    }
    if (object.sig !== undefined && object.sig !== null) {
      message.sig = bytesFromBase64(object.sig);
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = bytesFromBase64(object.proof);
    }
    if (object.messages !== undefined && object.messages !== null) {
      message.messages = bytesFromBase64(object.messages);
    }
    return message;
  },

  toJSON(message: TxPrivacyData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.lockTime !== undefined && (obj.lockTime = message.lockTime);
    message.fee !== undefined && (obj.fee = message.fee);
    message.info !== undefined &&
      (obj.info = base64FromBytes(
        message.info !== undefined ? message.info : new Uint8Array()
      ));
    message.sigPubKey !== undefined &&
      (obj.sigPubKey = base64FromBytes(
        message.sigPubKey !== undefined ? message.sigPubKey : new Uint8Array()
      ));
    message.sig !== undefined &&
      (obj.sig = base64FromBytes(
        message.sig !== undefined ? message.sig : new Uint8Array()
      ));
    message.proof !== undefined &&
      (obj.proof = base64FromBytes(
        message.proof !== undefined ? message.proof : new Uint8Array()
      ));
    message.messages !== undefined &&
      (obj.messages = base64FromBytes(
        message.messages !== undefined ? message.messages : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<TxPrivacyData>): TxPrivacyData {
    const message = { ...baseTxPrivacyData } as TxPrivacyData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.lockTime !== undefined && object.lockTime !== null) {
      message.lockTime = object.lockTime;
    } else {
      message.lockTime = 0;
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = 0;
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info;
    } else {
      message.info = new Uint8Array();
    }
    if (object.sigPubKey !== undefined && object.sigPubKey !== null) {
      message.sigPubKey = object.sigPubKey;
    } else {
      message.sigPubKey = new Uint8Array();
    }
    if (object.sig !== undefined && object.sig !== null) {
      message.sig = object.sig;
    } else {
      message.sig = new Uint8Array();
    }
    if (object.proof !== undefined && object.proof !== null) {
      message.proof = object.proof;
    } else {
      message.proof = new Uint8Array();
    }
    if (object.messages !== undefined && object.messages !== null) {
      message.messages = object.messages;
    } else {
      message.messages = new Uint8Array();
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
