/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface OnetimeAddress {
  index: string;
  creator: string;
  token_id: Uint8Array;
  public_key: Uint8Array;
  i: Uint8Array;
  status: number;
}

const baseOnetimeAddress: object = { index: "", creator: "", status: 0 };

export const OnetimeAddress = {
  encode(message: OnetimeAddress, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.token_id.length !== 0) {
      writer.uint32(26).bytes(message.token_id);
    }
    if (message.public_key.length !== 0) {
      writer.uint32(34).bytes(message.public_key);
    }
    if (message.i.length !== 0) {
      writer.uint32(42).bytes(message.i);
    }
    if (message.status !== 0) {
      writer.uint32(48).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OnetimeAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOnetimeAddress } as OnetimeAddress;
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
          message.token_id = reader.bytes();
          break;
        case 4:
          message.public_key = reader.bytes();
          break;
        case 5:
          message.i = reader.bytes();
          break;
        case 6:
          message.status = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OnetimeAddress {
    const message = { ...baseOnetimeAddress } as OnetimeAddress;
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
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = bytesFromBase64(object.token_id);
    }
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = bytesFromBase64(object.public_key);
    }
    if (object.i !== undefined && object.i !== null) {
      message.i = bytesFromBase64(object.i);
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = Number(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: OnetimeAddress): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.token_id !== undefined &&
      (obj.token_id = base64FromBytes(
        message.token_id !== undefined ? message.token_id : new Uint8Array()
      ));
    message.public_key !== undefined &&
      (obj.public_key = base64FromBytes(
        message.public_key !== undefined ? message.public_key : new Uint8Array()
      ));
    message.i !== undefined &&
      (obj.i = base64FromBytes(
        message.i !== undefined ? message.i : new Uint8Array()
      ));
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<OnetimeAddress>): OnetimeAddress {
    const message = { ...baseOnetimeAddress } as OnetimeAddress;
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
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = object.token_id;
    } else {
      message.token_id = new Uint8Array();
    }
    if (object.public_key !== undefined && object.public_key !== null) {
      message.public_key = object.public_key;
    } else {
      message.public_key = new Uint8Array();
    }
    if (object.i !== undefined && object.i !== null) {
      message.i = object.i;
    } else {
      message.i = new Uint8Array();
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
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