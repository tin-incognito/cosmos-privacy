/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

export interface MsgCreateSerialNumber {
  creator: string;
  index: string;
}

export interface MsgCreateSerialNumberResponse {}

export interface MsgUpdateSerialNumber {
  creator: string;
  index: string;
}

export interface MsgUpdateSerialNumberResponse {}

export interface MsgDeleteSerialNumber {
  creator: string;
  index: string;
}

export interface MsgDeleteSerialNumberResponse {}

export interface MsgCreateOutputCoin {
  creator: string;
  index: string;
}

export interface MsgCreateOutputCoinResponse {}

export interface MsgUpdateOutputCoin {
  creator: string;
  index: string;
}

export interface MsgUpdateOutputCoinResponse {}

export interface MsgDeleteOutputCoin {
  creator: string;
  index: string;
}

export interface MsgDeleteOutputCoinResponse {}

export interface MsgCreateCommitment {
  creator: string;
  index: string;
}

export interface MsgCreateCommitmentResponse {}

export interface MsgUpdateCommitment {
  creator: string;
  index: string;
}

export interface MsgUpdateCommitmentResponse {}

export interface MsgDeleteCommitment {
  creator: string;
  index: string;
}

export interface MsgDeleteCommitmentResponse {}

export interface MsgCreateCommitmentIndex {
  creator: string;
  index: string;
}

export interface MsgCreateCommitmentIndexResponse {}

export interface MsgUpdateCommitmentIndex {
  creator: string;
  index: string;
}

export interface MsgUpdateCommitmentIndexResponse {}

export interface MsgDeleteCommitmentIndex {
  creator: string;
  index: string;
}

export interface MsgDeleteCommitmentIndexResponse {}

export interface MsgCreateToken {
  creator: string;
  index: string;
}

export interface MsgCreateTokenResponse {}

export interface MsgUpdateToken {
  creator: string;
  index: string;
}

export interface MsgUpdateTokenResponse {}

export interface MsgDeleteToken {
  creator: string;
  index: string;
}

export interface MsgDeleteTokenResponse {}

export interface MsgCreateOnetimeAddress {
  creator: string;
  index: string;
}

export interface MsgCreateOnetimeAddressResponse {}

export interface MsgUpdateOnetimeAddress {
  creator: string;
  index: string;
}

export interface MsgUpdateOnetimeAddressResponse {}

export interface MsgDeleteOnetimeAddress {
  creator: string;
  index: string;
}

export interface MsgDeleteOnetimeAddressResponse {}

export interface MsgCreateTx {
  creator: string;
  value: Uint8Array;
}

export interface MsgCreateTxResponse {}

export interface MsgCreateTxPrivacyData {
  creator: string;
  index: string;
}

export interface MsgCreateTxPrivacyDataResponse {}

export interface MsgUpdateTxPrivacyData {
  creator: string;
  index: string;
}

export interface MsgUpdateTxPrivacyDataResponse {}

export interface MsgDeleteTxPrivacyData {
  creator: string;
  index: string;
}

export interface MsgDeleteTxPrivacyDataResponse {}

const baseMsgCreateSerialNumber: object = { creator: "", index: "" };

export const MsgCreateSerialNumber = {
  encode(
    message: MsgCreateSerialNumber,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSerialNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSerialNumber } as MsgCreateSerialNumber;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSerialNumber {
    const message = { ...baseMsgCreateSerialNumber } as MsgCreateSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateSerialNumber): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSerialNumber>
  ): MsgCreateSerialNumber {
    const message = { ...baseMsgCreateSerialNumber } as MsgCreateSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateSerialNumberResponse: object = {};

export const MsgCreateSerialNumberResponse = {
  encode(
    _: MsgCreateSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSerialNumberResponse,
    } as MsgCreateSerialNumberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateSerialNumberResponse {
    const message = {
      ...baseMsgCreateSerialNumberResponse,
    } as MsgCreateSerialNumberResponse;
    return message;
  },

  toJSON(_: MsgCreateSerialNumberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateSerialNumberResponse>
  ): MsgCreateSerialNumberResponse {
    const message = {
      ...baseMsgCreateSerialNumberResponse,
    } as MsgCreateSerialNumberResponse;
    return message;
  },
};

const baseMsgUpdateSerialNumber: object = { creator: "", index: "" };

export const MsgUpdateSerialNumber = {
  encode(
    message: MsgUpdateSerialNumber,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateSerialNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateSerialNumber } as MsgUpdateSerialNumber;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSerialNumber {
    const message = { ...baseMsgUpdateSerialNumber } as MsgUpdateSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateSerialNumber): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateSerialNumber>
  ): MsgUpdateSerialNumber {
    const message = { ...baseMsgUpdateSerialNumber } as MsgUpdateSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateSerialNumberResponse: object = {};

export const MsgUpdateSerialNumberResponse = {
  encode(
    _: MsgUpdateSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSerialNumberResponse,
    } as MsgUpdateSerialNumberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateSerialNumberResponse {
    const message = {
      ...baseMsgUpdateSerialNumberResponse,
    } as MsgUpdateSerialNumberResponse;
    return message;
  },

  toJSON(_: MsgUpdateSerialNumberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSerialNumberResponse>
  ): MsgUpdateSerialNumberResponse {
    const message = {
      ...baseMsgUpdateSerialNumberResponse,
    } as MsgUpdateSerialNumberResponse;
    return message;
  },
};

const baseMsgDeleteSerialNumber: object = { creator: "", index: "" };

export const MsgDeleteSerialNumber = {
  encode(
    message: MsgDeleteSerialNumber,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteSerialNumber {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteSerialNumber } as MsgDeleteSerialNumber;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteSerialNumber {
    const message = { ...baseMsgDeleteSerialNumber } as MsgDeleteSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteSerialNumber): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSerialNumber>
  ): MsgDeleteSerialNumber {
    const message = { ...baseMsgDeleteSerialNumber } as MsgDeleteSerialNumber;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteSerialNumberResponse: object = {};

export const MsgDeleteSerialNumberResponse = {
  encode(
    _: MsgDeleteSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSerialNumberResponse,
    } as MsgDeleteSerialNumberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteSerialNumberResponse {
    const message = {
      ...baseMsgDeleteSerialNumberResponse,
    } as MsgDeleteSerialNumberResponse;
    return message;
  },

  toJSON(_: MsgDeleteSerialNumberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSerialNumberResponse>
  ): MsgDeleteSerialNumberResponse {
    const message = {
      ...baseMsgDeleteSerialNumberResponse,
    } as MsgDeleteSerialNumberResponse;
    return message;
  },
};

const baseMsgCreateOutputCoin: object = { creator: "", index: "" };

export const MsgCreateOutputCoin = {
  encode(
    message: MsgCreateOutputCoin,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOutputCoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOutputCoin } as MsgCreateOutputCoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOutputCoin {
    const message = { ...baseMsgCreateOutputCoin } as MsgCreateOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateOutputCoin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateOutputCoin>): MsgCreateOutputCoin {
    const message = { ...baseMsgCreateOutputCoin } as MsgCreateOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateOutputCoinResponse: object = {};

export const MsgCreateOutputCoinResponse = {
  encode(
    _: MsgCreateOutputCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateOutputCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateOutputCoinResponse,
    } as MsgCreateOutputCoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateOutputCoinResponse {
    const message = {
      ...baseMsgCreateOutputCoinResponse,
    } as MsgCreateOutputCoinResponse;
    return message;
  },

  toJSON(_: MsgCreateOutputCoinResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateOutputCoinResponse>
  ): MsgCreateOutputCoinResponse {
    const message = {
      ...baseMsgCreateOutputCoinResponse,
    } as MsgCreateOutputCoinResponse;
    return message;
  },
};

const baseMsgUpdateOutputCoin: object = { creator: "", index: "" };

export const MsgUpdateOutputCoin = {
  encode(
    message: MsgUpdateOutputCoin,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateOutputCoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateOutputCoin } as MsgUpdateOutputCoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateOutputCoin {
    const message = { ...baseMsgUpdateOutputCoin } as MsgUpdateOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateOutputCoin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateOutputCoin>): MsgUpdateOutputCoin {
    const message = { ...baseMsgUpdateOutputCoin } as MsgUpdateOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateOutputCoinResponse: object = {};

export const MsgUpdateOutputCoinResponse = {
  encode(
    _: MsgUpdateOutputCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateOutputCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateOutputCoinResponse,
    } as MsgUpdateOutputCoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateOutputCoinResponse {
    const message = {
      ...baseMsgUpdateOutputCoinResponse,
    } as MsgUpdateOutputCoinResponse;
    return message;
  },

  toJSON(_: MsgUpdateOutputCoinResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateOutputCoinResponse>
  ): MsgUpdateOutputCoinResponse {
    const message = {
      ...baseMsgUpdateOutputCoinResponse,
    } as MsgUpdateOutputCoinResponse;
    return message;
  },
};

const baseMsgDeleteOutputCoin: object = { creator: "", index: "" };

export const MsgDeleteOutputCoin = {
  encode(
    message: MsgDeleteOutputCoin,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteOutputCoin {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteOutputCoin } as MsgDeleteOutputCoin;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteOutputCoin {
    const message = { ...baseMsgDeleteOutputCoin } as MsgDeleteOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteOutputCoin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteOutputCoin>): MsgDeleteOutputCoin {
    const message = { ...baseMsgDeleteOutputCoin } as MsgDeleteOutputCoin;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteOutputCoinResponse: object = {};

export const MsgDeleteOutputCoinResponse = {
  encode(
    _: MsgDeleteOutputCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteOutputCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteOutputCoinResponse,
    } as MsgDeleteOutputCoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteOutputCoinResponse {
    const message = {
      ...baseMsgDeleteOutputCoinResponse,
    } as MsgDeleteOutputCoinResponse;
    return message;
  },

  toJSON(_: MsgDeleteOutputCoinResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteOutputCoinResponse>
  ): MsgDeleteOutputCoinResponse {
    const message = {
      ...baseMsgDeleteOutputCoinResponse,
    } as MsgDeleteOutputCoinResponse;
    return message;
  },
};

const baseMsgCreateCommitment: object = { creator: "", index: "" };

export const MsgCreateCommitment = {
  encode(
    message: MsgCreateCommitment,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCommitment {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCommitment } as MsgCreateCommitment;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCommitment {
    const message = { ...baseMsgCreateCommitment } as MsgCreateCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateCommitment): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCommitment>): MsgCreateCommitment {
    const message = { ...baseMsgCreateCommitment } as MsgCreateCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateCommitmentResponse: object = {};

export const MsgCreateCommitmentResponse = {
  encode(
    _: MsgCreateCommitmentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateCommitmentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateCommitmentResponse,
    } as MsgCreateCommitmentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateCommitmentResponse {
    const message = {
      ...baseMsgCreateCommitmentResponse,
    } as MsgCreateCommitmentResponse;
    return message;
  },

  toJSON(_: MsgCreateCommitmentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateCommitmentResponse>
  ): MsgCreateCommitmentResponse {
    const message = {
      ...baseMsgCreateCommitmentResponse,
    } as MsgCreateCommitmentResponse;
    return message;
  },
};

const baseMsgUpdateCommitment: object = { creator: "", index: "" };

export const MsgUpdateCommitment = {
  encode(
    message: MsgUpdateCommitment,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateCommitment {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateCommitment } as MsgUpdateCommitment;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateCommitment {
    const message = { ...baseMsgUpdateCommitment } as MsgUpdateCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateCommitment): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateCommitment>): MsgUpdateCommitment {
    const message = { ...baseMsgUpdateCommitment } as MsgUpdateCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateCommitmentResponse: object = {};

export const MsgUpdateCommitmentResponse = {
  encode(
    _: MsgUpdateCommitmentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateCommitmentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateCommitmentResponse,
    } as MsgUpdateCommitmentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateCommitmentResponse {
    const message = {
      ...baseMsgUpdateCommitmentResponse,
    } as MsgUpdateCommitmentResponse;
    return message;
  },

  toJSON(_: MsgUpdateCommitmentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateCommitmentResponse>
  ): MsgUpdateCommitmentResponse {
    const message = {
      ...baseMsgUpdateCommitmentResponse,
    } as MsgUpdateCommitmentResponse;
    return message;
  },
};

const baseMsgDeleteCommitment: object = { creator: "", index: "" };

export const MsgDeleteCommitment = {
  encode(
    message: MsgDeleteCommitment,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteCommitment {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteCommitment } as MsgDeleteCommitment;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteCommitment {
    const message = { ...baseMsgDeleteCommitment } as MsgDeleteCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteCommitment): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteCommitment>): MsgDeleteCommitment {
    const message = { ...baseMsgDeleteCommitment } as MsgDeleteCommitment;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteCommitmentResponse: object = {};

export const MsgDeleteCommitmentResponse = {
  encode(
    _: MsgDeleteCommitmentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteCommitmentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteCommitmentResponse,
    } as MsgDeleteCommitmentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteCommitmentResponse {
    const message = {
      ...baseMsgDeleteCommitmentResponse,
    } as MsgDeleteCommitmentResponse;
    return message;
  },

  toJSON(_: MsgDeleteCommitmentResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteCommitmentResponse>
  ): MsgDeleteCommitmentResponse {
    const message = {
      ...baseMsgDeleteCommitmentResponse,
    } as MsgDeleteCommitmentResponse;
    return message;
  },
};

const baseMsgCreateCommitmentIndex: object = { creator: "", index: "" };

export const MsgCreateCommitmentIndex = {
  encode(
    message: MsgCreateCommitmentIndex,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateCommitmentIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateCommitmentIndex,
    } as MsgCreateCommitmentIndex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCommitmentIndex {
    const message = {
      ...baseMsgCreateCommitmentIndex,
    } as MsgCreateCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateCommitmentIndex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateCommitmentIndex>
  ): MsgCreateCommitmentIndex {
    const message = {
      ...baseMsgCreateCommitmentIndex,
    } as MsgCreateCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateCommitmentIndexResponse: object = {};

export const MsgCreateCommitmentIndexResponse = {
  encode(
    _: MsgCreateCommitmentIndexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateCommitmentIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateCommitmentIndexResponse,
    } as MsgCreateCommitmentIndexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateCommitmentIndexResponse {
    const message = {
      ...baseMsgCreateCommitmentIndexResponse,
    } as MsgCreateCommitmentIndexResponse;
    return message;
  },

  toJSON(_: MsgCreateCommitmentIndexResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateCommitmentIndexResponse>
  ): MsgCreateCommitmentIndexResponse {
    const message = {
      ...baseMsgCreateCommitmentIndexResponse,
    } as MsgCreateCommitmentIndexResponse;
    return message;
  },
};

const baseMsgUpdateCommitmentIndex: object = { creator: "", index: "" };

export const MsgUpdateCommitmentIndex = {
  encode(
    message: MsgUpdateCommitmentIndex,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateCommitmentIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateCommitmentIndex,
    } as MsgUpdateCommitmentIndex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateCommitmentIndex {
    const message = {
      ...baseMsgUpdateCommitmentIndex,
    } as MsgUpdateCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateCommitmentIndex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateCommitmentIndex>
  ): MsgUpdateCommitmentIndex {
    const message = {
      ...baseMsgUpdateCommitmentIndex,
    } as MsgUpdateCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateCommitmentIndexResponse: object = {};

export const MsgUpdateCommitmentIndexResponse = {
  encode(
    _: MsgUpdateCommitmentIndexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateCommitmentIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateCommitmentIndexResponse,
    } as MsgUpdateCommitmentIndexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateCommitmentIndexResponse {
    const message = {
      ...baseMsgUpdateCommitmentIndexResponse,
    } as MsgUpdateCommitmentIndexResponse;
    return message;
  },

  toJSON(_: MsgUpdateCommitmentIndexResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateCommitmentIndexResponse>
  ): MsgUpdateCommitmentIndexResponse {
    const message = {
      ...baseMsgUpdateCommitmentIndexResponse,
    } as MsgUpdateCommitmentIndexResponse;
    return message;
  },
};

const baseMsgDeleteCommitmentIndex: object = { creator: "", index: "" };

export const MsgDeleteCommitmentIndex = {
  encode(
    message: MsgDeleteCommitmentIndex,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteCommitmentIndex {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteCommitmentIndex,
    } as MsgDeleteCommitmentIndex;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteCommitmentIndex {
    const message = {
      ...baseMsgDeleteCommitmentIndex,
    } as MsgDeleteCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteCommitmentIndex): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteCommitmentIndex>
  ): MsgDeleteCommitmentIndex {
    const message = {
      ...baseMsgDeleteCommitmentIndex,
    } as MsgDeleteCommitmentIndex;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteCommitmentIndexResponse: object = {};

export const MsgDeleteCommitmentIndexResponse = {
  encode(
    _: MsgDeleteCommitmentIndexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteCommitmentIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteCommitmentIndexResponse,
    } as MsgDeleteCommitmentIndexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteCommitmentIndexResponse {
    const message = {
      ...baseMsgDeleteCommitmentIndexResponse,
    } as MsgDeleteCommitmentIndexResponse;
    return message;
  },

  toJSON(_: MsgDeleteCommitmentIndexResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteCommitmentIndexResponse>
  ): MsgDeleteCommitmentIndexResponse {
    const message = {
      ...baseMsgDeleteCommitmentIndexResponse,
    } as MsgDeleteCommitmentIndexResponse;
    return message;
  },
};

const baseMsgCreateToken: object = { creator: "", index: "" };

export const MsgCreateToken = {
  encode(message: MsgCreateToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateToken } as MsgCreateToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateToken {
    const message = { ...baseMsgCreateToken } as MsgCreateToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateToken>): MsgCreateToken {
    const message = { ...baseMsgCreateToken } as MsgCreateToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateTokenResponse: object = {};

export const MsgCreateTokenResponse = {
  encode(_: MsgCreateTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTokenResponse } as MsgCreateTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTokenResponse {
    const message = { ...baseMsgCreateTokenResponse } as MsgCreateTokenResponse;
    return message;
  },

  toJSON(_: MsgCreateTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateTokenResponse>): MsgCreateTokenResponse {
    const message = { ...baseMsgCreateTokenResponse } as MsgCreateTokenResponse;
    return message;
  },
};

const baseMsgUpdateToken: object = { creator: "", index: "" };

export const MsgUpdateToken = {
  encode(message: MsgUpdateToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateToken } as MsgUpdateToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateToken {
    const message = { ...baseMsgUpdateToken } as MsgUpdateToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateToken>): MsgUpdateToken {
    const message = { ...baseMsgUpdateToken } as MsgUpdateToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateTokenResponse: object = {};

export const MsgUpdateTokenResponse = {
  encode(_: MsgUpdateTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateTokenResponse } as MsgUpdateTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateTokenResponse {
    const message = { ...baseMsgUpdateTokenResponse } as MsgUpdateTokenResponse;
    return message;
  },

  toJSON(_: MsgUpdateTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdateTokenResponse>): MsgUpdateTokenResponse {
    const message = { ...baseMsgUpdateTokenResponse } as MsgUpdateTokenResponse;
    return message;
  },
};

const baseMsgDeleteToken: object = { creator: "", index: "" };

export const MsgDeleteToken = {
  encode(message: MsgDeleteToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteToken } as MsgDeleteToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteToken {
    const message = { ...baseMsgDeleteToken } as MsgDeleteToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteToken>): MsgDeleteToken {
    const message = { ...baseMsgDeleteToken } as MsgDeleteToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteTokenResponse: object = {};

export const MsgDeleteTokenResponse = {
  encode(_: MsgDeleteTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteTokenResponse } as MsgDeleteTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteTokenResponse {
    const message = { ...baseMsgDeleteTokenResponse } as MsgDeleteTokenResponse;
    return message;
  },

  toJSON(_: MsgDeleteTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteTokenResponse>): MsgDeleteTokenResponse {
    const message = { ...baseMsgDeleteTokenResponse } as MsgDeleteTokenResponse;
    return message;
  },
};

const baseMsgCreateOnetimeAddress: object = { creator: "", index: "" };

export const MsgCreateOnetimeAddress = {
  encode(
    message: MsgCreateOnetimeAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOnetimeAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateOnetimeAddress,
    } as MsgCreateOnetimeAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOnetimeAddress {
    const message = {
      ...baseMsgCreateOnetimeAddress,
    } as MsgCreateOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateOnetimeAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateOnetimeAddress>
  ): MsgCreateOnetimeAddress {
    const message = {
      ...baseMsgCreateOnetimeAddress,
    } as MsgCreateOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateOnetimeAddressResponse: object = {};

export const MsgCreateOnetimeAddressResponse = {
  encode(
    _: MsgCreateOnetimeAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateOnetimeAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateOnetimeAddressResponse,
    } as MsgCreateOnetimeAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateOnetimeAddressResponse {
    const message = {
      ...baseMsgCreateOnetimeAddressResponse,
    } as MsgCreateOnetimeAddressResponse;
    return message;
  },

  toJSON(_: MsgCreateOnetimeAddressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateOnetimeAddressResponse>
  ): MsgCreateOnetimeAddressResponse {
    const message = {
      ...baseMsgCreateOnetimeAddressResponse,
    } as MsgCreateOnetimeAddressResponse;
    return message;
  },
};

const baseMsgUpdateOnetimeAddress: object = { creator: "", index: "" };

export const MsgUpdateOnetimeAddress = {
  encode(
    message: MsgUpdateOnetimeAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateOnetimeAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateOnetimeAddress,
    } as MsgUpdateOnetimeAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateOnetimeAddress {
    const message = {
      ...baseMsgUpdateOnetimeAddress,
    } as MsgUpdateOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateOnetimeAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateOnetimeAddress>
  ): MsgUpdateOnetimeAddress {
    const message = {
      ...baseMsgUpdateOnetimeAddress,
    } as MsgUpdateOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateOnetimeAddressResponse: object = {};

export const MsgUpdateOnetimeAddressResponse = {
  encode(
    _: MsgUpdateOnetimeAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateOnetimeAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateOnetimeAddressResponse,
    } as MsgUpdateOnetimeAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateOnetimeAddressResponse {
    const message = {
      ...baseMsgUpdateOnetimeAddressResponse,
    } as MsgUpdateOnetimeAddressResponse;
    return message;
  },

  toJSON(_: MsgUpdateOnetimeAddressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateOnetimeAddressResponse>
  ): MsgUpdateOnetimeAddressResponse {
    const message = {
      ...baseMsgUpdateOnetimeAddressResponse,
    } as MsgUpdateOnetimeAddressResponse;
    return message;
  },
};

const baseMsgDeleteOnetimeAddress: object = { creator: "", index: "" };

export const MsgDeleteOnetimeAddress = {
  encode(
    message: MsgDeleteOnetimeAddress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteOnetimeAddress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteOnetimeAddress,
    } as MsgDeleteOnetimeAddress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteOnetimeAddress {
    const message = {
      ...baseMsgDeleteOnetimeAddress,
    } as MsgDeleteOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteOnetimeAddress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteOnetimeAddress>
  ): MsgDeleteOnetimeAddress {
    const message = {
      ...baseMsgDeleteOnetimeAddress,
    } as MsgDeleteOnetimeAddress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteOnetimeAddressResponse: object = {};

export const MsgDeleteOnetimeAddressResponse = {
  encode(
    _: MsgDeleteOnetimeAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteOnetimeAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteOnetimeAddressResponse,
    } as MsgDeleteOnetimeAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteOnetimeAddressResponse {
    const message = {
      ...baseMsgDeleteOnetimeAddressResponse,
    } as MsgDeleteOnetimeAddressResponse;
    return message;
  },

  toJSON(_: MsgDeleteOnetimeAddressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteOnetimeAddressResponse>
  ): MsgDeleteOnetimeAddressResponse {
    const message = {
      ...baseMsgDeleteOnetimeAddressResponse,
    } as MsgDeleteOnetimeAddressResponse;
    return message;
  },
};

const baseMsgCreateTx: object = { creator: "" };

export const MsgCreateTx = {
  encode(message: MsgCreateTx, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.value.length !== 0) {
      writer.uint32(18).bytes(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTx {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTx } as MsgCreateTx;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.value = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTx {
    const message = { ...baseMsgCreateTx } as MsgCreateTx;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = bytesFromBase64(object.value);
    }
    return message;
  },

  toJSON(message: MsgCreateTx): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.value !== undefined &&
      (obj.value = base64FromBytes(
        message.value !== undefined ? message.value : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTx>): MsgCreateTx {
    const message = { ...baseMsgCreateTx } as MsgCreateTx;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = new Uint8Array();
    }
    return message;
  },
};

const baseMsgCreateTxResponse: object = {};

export const MsgCreateTxResponse = {
  encode(_: MsgCreateTxResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTxResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTxResponse } as MsgCreateTxResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTxResponse {
    const message = { ...baseMsgCreateTxResponse } as MsgCreateTxResponse;
    return message;
  },

  toJSON(_: MsgCreateTxResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateTxResponse>): MsgCreateTxResponse {
    const message = { ...baseMsgCreateTxResponse } as MsgCreateTxResponse;
    return message;
  },
};

const baseMsgCreateTxPrivacyData: object = { creator: "", index: "" };

export const MsgCreateTxPrivacyData = {
  encode(
    message: MsgCreateTxPrivacyData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTxPrivacyData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTxPrivacyData } as MsgCreateTxPrivacyData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTxPrivacyData {
    const message = { ...baseMsgCreateTxPrivacyData } as MsgCreateTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTxPrivacyData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTxPrivacyData>
  ): MsgCreateTxPrivacyData {
    const message = { ...baseMsgCreateTxPrivacyData } as MsgCreateTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgCreateTxPrivacyDataResponse: object = {};

export const MsgCreateTxPrivacyDataResponse = {
  encode(
    _: MsgCreateTxPrivacyDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateTxPrivacyDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTxPrivacyDataResponse,
    } as MsgCreateTxPrivacyDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateTxPrivacyDataResponse {
    const message = {
      ...baseMsgCreateTxPrivacyDataResponse,
    } as MsgCreateTxPrivacyDataResponse;
    return message;
  },

  toJSON(_: MsgCreateTxPrivacyDataResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateTxPrivacyDataResponse>
  ): MsgCreateTxPrivacyDataResponse {
    const message = {
      ...baseMsgCreateTxPrivacyDataResponse,
    } as MsgCreateTxPrivacyDataResponse;
    return message;
  },
};

const baseMsgUpdateTxPrivacyData: object = { creator: "", index: "" };

export const MsgUpdateTxPrivacyData = {
  encode(
    message: MsgUpdateTxPrivacyData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateTxPrivacyData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateTxPrivacyData } as MsgUpdateTxPrivacyData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateTxPrivacyData {
    const message = { ...baseMsgUpdateTxPrivacyData } as MsgUpdateTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateTxPrivacyData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateTxPrivacyData>
  ): MsgUpdateTxPrivacyData {
    const message = { ...baseMsgUpdateTxPrivacyData } as MsgUpdateTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgUpdateTxPrivacyDataResponse: object = {};

export const MsgUpdateTxPrivacyDataResponse = {
  encode(
    _: MsgUpdateTxPrivacyDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateTxPrivacyDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTxPrivacyDataResponse,
    } as MsgUpdateTxPrivacyDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateTxPrivacyDataResponse {
    const message = {
      ...baseMsgUpdateTxPrivacyDataResponse,
    } as MsgUpdateTxPrivacyDataResponse;
    return message;
  },

  toJSON(_: MsgUpdateTxPrivacyDataResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateTxPrivacyDataResponse>
  ): MsgUpdateTxPrivacyDataResponse {
    const message = {
      ...baseMsgUpdateTxPrivacyDataResponse,
    } as MsgUpdateTxPrivacyDataResponse;
    return message;
  },
};

const baseMsgDeleteTxPrivacyData: object = { creator: "", index: "" };

export const MsgDeleteTxPrivacyData = {
  encode(
    message: MsgDeleteTxPrivacyData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteTxPrivacyData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteTxPrivacyData } as MsgDeleteTxPrivacyData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteTxPrivacyData {
    const message = { ...baseMsgDeleteTxPrivacyData } as MsgDeleteTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteTxPrivacyData): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteTxPrivacyData>
  ): MsgDeleteTxPrivacyData {
    const message = { ...baseMsgDeleteTxPrivacyData } as MsgDeleteTxPrivacyData;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseMsgDeleteTxPrivacyDataResponse: object = {};

export const MsgDeleteTxPrivacyDataResponse = {
  encode(
    _: MsgDeleteTxPrivacyDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteTxPrivacyDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTxPrivacyDataResponse,
    } as MsgDeleteTxPrivacyDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteTxPrivacyDataResponse {
    const message = {
      ...baseMsgDeleteTxPrivacyDataResponse,
    } as MsgDeleteTxPrivacyDataResponse;
    return message;
  },

  toJSON(_: MsgDeleteTxPrivacyDataResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteTxPrivacyDataResponse>
  ): MsgDeleteTxPrivacyDataResponse {
    const message = {
      ...baseMsgDeleteTxPrivacyDataResponse,
    } as MsgDeleteTxPrivacyDataResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateSerialNumber(
    request: MsgCreateSerialNumber
  ): Promise<MsgCreateSerialNumberResponse>;
  UpdateSerialNumber(
    request: MsgUpdateSerialNumber
  ): Promise<MsgUpdateSerialNumberResponse>;
  DeleteSerialNumber(
    request: MsgDeleteSerialNumber
  ): Promise<MsgDeleteSerialNumberResponse>;
  CreateOutputCoin(
    request: MsgCreateOutputCoin
  ): Promise<MsgCreateOutputCoinResponse>;
  UpdateOutputCoin(
    request: MsgUpdateOutputCoin
  ): Promise<MsgUpdateOutputCoinResponse>;
  DeleteOutputCoin(
    request: MsgDeleteOutputCoin
  ): Promise<MsgDeleteOutputCoinResponse>;
  CreateCommitment(
    request: MsgCreateCommitment
  ): Promise<MsgCreateCommitmentResponse>;
  UpdateCommitment(
    request: MsgUpdateCommitment
  ): Promise<MsgUpdateCommitmentResponse>;
  DeleteCommitment(
    request: MsgDeleteCommitment
  ): Promise<MsgDeleteCommitmentResponse>;
  CreateCommitmentIndex(
    request: MsgCreateCommitmentIndex
  ): Promise<MsgCreateCommitmentIndexResponse>;
  UpdateCommitmentIndex(
    request: MsgUpdateCommitmentIndex
  ): Promise<MsgUpdateCommitmentIndexResponse>;
  DeleteCommitmentIndex(
    request: MsgDeleteCommitmentIndex
  ): Promise<MsgDeleteCommitmentIndexResponse>;
  CreateToken(request: MsgCreateToken): Promise<MsgCreateTokenResponse>;
  UpdateToken(request: MsgUpdateToken): Promise<MsgUpdateTokenResponse>;
  DeleteToken(request: MsgDeleteToken): Promise<MsgDeleteTokenResponse>;
  CreateOnetimeAddress(
    request: MsgCreateOnetimeAddress
  ): Promise<MsgCreateOnetimeAddressResponse>;
  UpdateOnetimeAddress(
    request: MsgUpdateOnetimeAddress
  ): Promise<MsgUpdateOnetimeAddressResponse>;
  DeleteOnetimeAddress(
    request: MsgDeleteOnetimeAddress
  ): Promise<MsgDeleteOnetimeAddressResponse>;
  CreateTx(request: MsgCreateTx): Promise<MsgCreateTxResponse>;
  CreateTxPrivacyData(
    request: MsgCreateTxPrivacyData
  ): Promise<MsgCreateTxPrivacyDataResponse>;
  UpdateTxPrivacyData(
    request: MsgUpdateTxPrivacyData
  ): Promise<MsgUpdateTxPrivacyDataResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteTxPrivacyData(
    request: MsgDeleteTxPrivacyData
  ): Promise<MsgDeleteTxPrivacyDataResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateSerialNumber(
    request: MsgCreateSerialNumber
  ): Promise<MsgCreateSerialNumberResponse> {
    const data = MsgCreateSerialNumber.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateSerialNumber",
      data
    );
    return promise.then((data) =>
      MsgCreateSerialNumberResponse.decode(new Reader(data))
    );
  }

  UpdateSerialNumber(
    request: MsgUpdateSerialNumber
  ): Promise<MsgUpdateSerialNumberResponse> {
    const data = MsgUpdateSerialNumber.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateSerialNumber",
      data
    );
    return promise.then((data) =>
      MsgUpdateSerialNumberResponse.decode(new Reader(data))
    );
  }

  DeleteSerialNumber(
    request: MsgDeleteSerialNumber
  ): Promise<MsgDeleteSerialNumberResponse> {
    const data = MsgDeleteSerialNumber.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteSerialNumber",
      data
    );
    return promise.then((data) =>
      MsgDeleteSerialNumberResponse.decode(new Reader(data))
    );
  }

  CreateOutputCoin(
    request: MsgCreateOutputCoin
  ): Promise<MsgCreateOutputCoinResponse> {
    const data = MsgCreateOutputCoin.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateOutputCoin",
      data
    );
    return promise.then((data) =>
      MsgCreateOutputCoinResponse.decode(new Reader(data))
    );
  }

  UpdateOutputCoin(
    request: MsgUpdateOutputCoin
  ): Promise<MsgUpdateOutputCoinResponse> {
    const data = MsgUpdateOutputCoin.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateOutputCoin",
      data
    );
    return promise.then((data) =>
      MsgUpdateOutputCoinResponse.decode(new Reader(data))
    );
  }

  DeleteOutputCoin(
    request: MsgDeleteOutputCoin
  ): Promise<MsgDeleteOutputCoinResponse> {
    const data = MsgDeleteOutputCoin.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteOutputCoin",
      data
    );
    return promise.then((data) =>
      MsgDeleteOutputCoinResponse.decode(new Reader(data))
    );
  }

  CreateCommitment(
    request: MsgCreateCommitment
  ): Promise<MsgCreateCommitmentResponse> {
    const data = MsgCreateCommitment.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateCommitment",
      data
    );
    return promise.then((data) =>
      MsgCreateCommitmentResponse.decode(new Reader(data))
    );
  }

  UpdateCommitment(
    request: MsgUpdateCommitment
  ): Promise<MsgUpdateCommitmentResponse> {
    const data = MsgUpdateCommitment.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateCommitment",
      data
    );
    return promise.then((data) =>
      MsgUpdateCommitmentResponse.decode(new Reader(data))
    );
  }

  DeleteCommitment(
    request: MsgDeleteCommitment
  ): Promise<MsgDeleteCommitmentResponse> {
    const data = MsgDeleteCommitment.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteCommitment",
      data
    );
    return promise.then((data) =>
      MsgDeleteCommitmentResponse.decode(new Reader(data))
    );
  }

  CreateCommitmentIndex(
    request: MsgCreateCommitmentIndex
  ): Promise<MsgCreateCommitmentIndexResponse> {
    const data = MsgCreateCommitmentIndex.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateCommitmentIndex",
      data
    );
    return promise.then((data) =>
      MsgCreateCommitmentIndexResponse.decode(new Reader(data))
    );
  }

  UpdateCommitmentIndex(
    request: MsgUpdateCommitmentIndex
  ): Promise<MsgUpdateCommitmentIndexResponse> {
    const data = MsgUpdateCommitmentIndex.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateCommitmentIndex",
      data
    );
    return promise.then((data) =>
      MsgUpdateCommitmentIndexResponse.decode(new Reader(data))
    );
  }

  DeleteCommitmentIndex(
    request: MsgDeleteCommitmentIndex
  ): Promise<MsgDeleteCommitmentIndexResponse> {
    const data = MsgDeleteCommitmentIndex.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteCommitmentIndex",
      data
    );
    return promise.then((data) =>
      MsgDeleteCommitmentIndexResponse.decode(new Reader(data))
    );
  }

  CreateToken(request: MsgCreateToken): Promise<MsgCreateTokenResponse> {
    const data = MsgCreateToken.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateToken",
      data
    );
    return promise.then((data) =>
      MsgCreateTokenResponse.decode(new Reader(data))
    );
  }

  UpdateToken(request: MsgUpdateToken): Promise<MsgUpdateTokenResponse> {
    const data = MsgUpdateToken.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateToken",
      data
    );
    return promise.then((data) =>
      MsgUpdateTokenResponse.decode(new Reader(data))
    );
  }

  DeleteToken(request: MsgDeleteToken): Promise<MsgDeleteTokenResponse> {
    const data = MsgDeleteToken.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteToken",
      data
    );
    return promise.then((data) =>
      MsgDeleteTokenResponse.decode(new Reader(data))
    );
  }

  CreateOnetimeAddress(
    request: MsgCreateOnetimeAddress
  ): Promise<MsgCreateOnetimeAddressResponse> {
    const data = MsgCreateOnetimeAddress.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateOnetimeAddress",
      data
    );
    return promise.then((data) =>
      MsgCreateOnetimeAddressResponse.decode(new Reader(data))
    );
  }

  UpdateOnetimeAddress(
    request: MsgUpdateOnetimeAddress
  ): Promise<MsgUpdateOnetimeAddressResponse> {
    const data = MsgUpdateOnetimeAddress.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateOnetimeAddress",
      data
    );
    return promise.then((data) =>
      MsgUpdateOnetimeAddressResponse.decode(new Reader(data))
    );
  }

  DeleteOnetimeAddress(
    request: MsgDeleteOnetimeAddress
  ): Promise<MsgDeleteOnetimeAddressResponse> {
    const data = MsgDeleteOnetimeAddress.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteOnetimeAddress",
      data
    );
    return promise.then((data) =>
      MsgDeleteOnetimeAddressResponse.decode(new Reader(data))
    );
  }

  CreateTx(request: MsgCreateTx): Promise<MsgCreateTxResponse> {
    const data = MsgCreateTx.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Msg", "CreateTx", data);
    return promise.then((data) => MsgCreateTxResponse.decode(new Reader(data)));
  }

  CreateTxPrivacyData(
    request: MsgCreateTxPrivacyData
  ): Promise<MsgCreateTxPrivacyDataResponse> {
    const data = MsgCreateTxPrivacyData.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "CreateTxPrivacyData",
      data
    );
    return promise.then((data) =>
      MsgCreateTxPrivacyDataResponse.decode(new Reader(data))
    );
  }

  UpdateTxPrivacyData(
    request: MsgUpdateTxPrivacyData
  ): Promise<MsgUpdateTxPrivacyDataResponse> {
    const data = MsgUpdateTxPrivacyData.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "UpdateTxPrivacyData",
      data
    );
    return promise.then((data) =>
      MsgUpdateTxPrivacyDataResponse.decode(new Reader(data))
    );
  }

  DeleteTxPrivacyData(
    request: MsgDeleteTxPrivacyData
  ): Promise<MsgDeleteTxPrivacyDataResponse> {
    const data = MsgDeleteTxPrivacyData.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Msg",
      "DeleteTxPrivacyData",
      data
    );
    return promise.then((data) =>
      MsgDeleteTxPrivacyDataResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
