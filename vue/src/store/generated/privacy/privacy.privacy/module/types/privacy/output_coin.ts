/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface OutputCoin {
  index: string;
  creator: string;
  is_confidential_asset: boolean;
  pub_key: Uint8Array;
  value: Uint8Array;
}

const baseOutputCoin: object = {
  index: "",
  creator: "",
  is_confidential_asset: false,
};

export const OutputCoin = {
  encode(message: OutputCoin, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.is_confidential_asset === true) {
      writer.uint32(24).bool(message.is_confidential_asset);
    }
    if (message.pub_key.length !== 0) {
      writer.uint32(34).bytes(message.pub_key);
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
          message.creator = reader.string();
          break;
        case 3:
          message.is_confidential_asset = reader.bool();
          break;
        case 4:
          message.pub_key = reader.bytes();
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.is_confidential_asset !== undefined &&
      object.is_confidential_asset !== null
    ) {
      message.is_confidential_asset = Boolean(object.is_confidential_asset);
    } else {
      message.is_confidential_asset = false;
    }
    if (object.pub_key !== undefined && object.pub_key !== null) {
      message.pub_key = bytesFromBase64(object.pub_key);
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = bytesFromBase64(object.value);
    }
    return message;
  },

  toJSON(message: OutputCoin): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.is_confidential_asset !== undefined &&
      (obj.is_confidential_asset = message.is_confidential_asset);
    message.pub_key !== undefined &&
      (obj.pub_key = base64FromBytes(
        message.pub_key !== undefined ? message.pub_key : new Uint8Array()
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
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.is_confidential_asset !== undefined &&
      object.is_confidential_asset !== null
    ) {
      message.is_confidential_asset = object.is_confidential_asset;
    } else {
      message.is_confidential_asset = false;
    }
    if (object.pub_key !== undefined && object.pub_key !== null) {
      message.pub_key = object.pub_key;
    } else {
      message.pub_key = new Uint8Array();
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
