/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface OutputCoin {
  index: string;
  serialNumber: Uint8Array;
  isConfidentialAsset: boolean;
  pubKey: Uint8Array;
  value: Uint8Array;
}

const baseOutputCoin: object = { index: "", isConfidentialAsset: false };

export const OutputCoin = {
  encode(message: OutputCoin, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.serialNumber.length !== 0) {
      writer.uint32(18).bytes(message.serialNumber);
    }
    if (message.isConfidentialAsset === true) {
      writer.uint32(24).bool(message.isConfidentialAsset);
    }
    if (message.pubKey.length !== 0) {
      writer.uint32(34).bytes(message.pubKey);
    }
    if (message.value.length !== 0) {
      writer.uint32(42).bytes(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OutputCoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOutputCoin } as OutputCoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.serialNumber = reader.bytes();
          break;
        case 3:
          message.isConfidentialAsset = reader.bool();
          break;
        case 4:
          message.pubKey = reader.bytes();
          break;
        case 5:
          message.value = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OutputCoin {
    const message = { ...baseOutputCoin } as OutputCoin;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = bytesFromBase64(object.serialNumber);
    }
    if (
      object.isConfidentialAsset !== undefined &&
      object.isConfidentialAsset !== null
    ) {
      message.isConfidentialAsset = Boolean(object.isConfidentialAsset);
    } else {
      message.isConfidentialAsset = false;
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = bytesFromBase64(object.pubKey);
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = bytesFromBase64(object.value);
    }
    return message;
  },

  toJSON(message: OutputCoin): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.serialNumber !== undefined &&
      (obj.serialNumber = base64FromBytes(
        message.serialNumber !== undefined
          ? message.serialNumber
          : new Uint8Array()
      ));
    message.isConfidentialAsset !== undefined &&
      (obj.isConfidentialAsset = message.isConfidentialAsset);
    message.pubKey !== undefined &&
      (obj.pubKey = base64FromBytes(
        message.pubKey !== undefined ? message.pubKey : new Uint8Array()
      ));
    message.value !== undefined &&
      (obj.value = base64FromBytes(
        message.value !== undefined ? message.value : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<OutputCoin>): OutputCoin {
    const message = { ...baseOutputCoin } as OutputCoin;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber;
    } else {
      message.serialNumber = new Uint8Array();
    }
    if (
      object.isConfidentialAsset !== undefined &&
      object.isConfidentialAsset !== null
    ) {
      message.isConfidentialAsset = object.isConfidentialAsset;
    } else {
      message.isConfidentialAsset = false;
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = new Uint8Array();
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = new Uint8Array();
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
