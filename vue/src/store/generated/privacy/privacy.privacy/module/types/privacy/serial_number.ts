/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface SerialNumber {
  index: string;
  creator: string;
  token_id: Uint8Array;
  serial_number: Uint8Array;
}

const baseSerialNumber: object = { index: "", creator: "" };

export const SerialNumber = {
  encode(message: SerialNumber, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.token_id.length !== 0) {
      writer.uint32(26).bytes(message.token_id);
    }
    if (message.serial_number.length !== 0) {
      writer.uint32(34).bytes(message.serial_number);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SerialNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSerialNumber } as SerialNumber;
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
          message.serial_number = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SerialNumber {
    const message = { ...baseSerialNumber } as SerialNumber;
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
    if (object.serial_number !== undefined && object.serial_number !== null) {
      message.serial_number = bytesFromBase64(object.serial_number);
    }
    return message;
  },

  toJSON(message: SerialNumber): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.token_id !== undefined &&
      (obj.token_id = base64FromBytes(
        message.token_id !== undefined ? message.token_id : new Uint8Array()
      ));
    message.serial_number !== undefined &&
      (obj.serial_number = base64FromBytes(
        message.serial_number !== undefined
          ? message.serial_number
          : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<SerialNumber>): SerialNumber {
    const message = { ...baseSerialNumber } as SerialNumber;
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
    if (object.serial_number !== undefined && object.serial_number !== null) {
      message.serial_number = object.serial_number;
    } else {
      message.serial_number = new Uint8Array();
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
