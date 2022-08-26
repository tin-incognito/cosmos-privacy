/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface OTACoin {
  index: string;
  output_coin_index: string;
  height: Uint8Array;
}

const baseOTACoin: object = { index: "", output_coin_index: "" };

export const OTACoin = {
  encode(message: OTACoin, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.output_coin_index !== "") {
      writer.uint32(18).string(message.output_coin_index);
    }
    if (message.height.length !== 0) {
      writer.uint32(26).bytes(message.height);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OTACoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOTACoin } as OTACoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.output_coin_index = reader.string();
          break;
        case 3:
          message.height = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OTACoin {
    const message = { ...baseOTACoin } as OTACoin;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (
      object.output_coin_index !== undefined &&
      object.output_coin_index !== null
    ) {
      message.output_coin_index = String(object.output_coin_index);
    } else {
      message.output_coin_index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = bytesFromBase64(object.height);
    }
    return message;
  },

  toJSON(message: OTACoin): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.output_coin_index !== undefined &&
      (obj.output_coin_index = message.output_coin_index);
    message.height !== undefined &&
      (obj.height = base64FromBytes(
        message.height !== undefined ? message.height : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<OTACoin>): OTACoin {
    const message = { ...baseOTACoin } as OTACoin;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (
      object.output_coin_index !== undefined &&
      object.output_coin_index !== null
    ) {
      message.output_coin_index = object.output_coin_index;
    } else {
      message.output_coin_index = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = new Uint8Array();
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
