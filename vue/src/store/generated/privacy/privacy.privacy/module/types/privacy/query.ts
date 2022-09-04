/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../privacy/params";
import { SerialNumber } from "../privacy/serial_number";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { OutputCoin } from "../privacy/output_coin";
import { Commitment } from "../privacy/commitment";
import { CommitmentIndex } from "../privacy/commitment_index";
import { Token } from "../privacy/token";
import { OnetimeAddress } from "../privacy/onetime_address";
import { TxPrivacyData } from "../privacy/tx_privacy_data";
import { OTACoin } from "../privacy/ota_coin";
import { OutputCoinSerialNumber } from "../privacy/output_coin_serial_number";

export const protobufPackage = "privacy.privacy";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSerialNumberRequest {
  index: string;
}

export interface QueryGetSerialNumberResponse {
  serialNumber: SerialNumber | undefined;
}

export interface QueryAllSerialNumberRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSerialNumberResponse {
  serialNumber: SerialNumber[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOutputCoinRequest {
  index: string;
}

export interface QueryGetOutputCoinResponse {
  outputCoin: OutputCoin | undefined;
}

export interface QueryAllOutputCoinRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOutputCoinResponse {
  outputCoin: OutputCoin[];
  pagination: PageResponse | undefined;
}

export interface QueryGetCommitmentRequest {
  index: string;
}

export interface QueryGetCommitmentResponse {
  commitment: Commitment | undefined;
}

export interface QueryAllCommitmentRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCommitmentResponse {
  commitment: Commitment[];
  pagination: PageResponse | undefined;
}

export interface QueryGetCommitmentIndexRequest {
  index: string;
}

export interface QueryGetCommitmentIndexResponse {
  commitmentIndex: CommitmentIndex | undefined;
}

export interface QueryAllCommitmentIndexRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCommitmentIndexResponse {
  commitmentIndex: CommitmentIndex[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTokenRequest {
  index: string;
}

export interface QueryGetTokenResponse {
  token: Token | undefined;
}

export interface QueryAllTokenRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokenResponse {
  token: Token[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOnetimeAddressRequest {
  index: string;
}

export interface QueryGetOnetimeAddressResponse {
  onetimeAddress: OnetimeAddress | undefined;
}

export interface QueryAllOnetimeAddressRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOnetimeAddressResponse {
  onetimeAddress: OnetimeAddress[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTxPrivacyDataRequest {
  index: string;
}

export interface QueryGetTxPrivacyDataResponse {
  txPrivacyData: TxPrivacyData | undefined;
}

export interface QueryAllTxPrivacyDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTxPrivacyDataResponse {
  txPrivacyData: TxPrivacyData[];
  pagination: PageResponse | undefined;
}

export interface QueryBalanceRequest {
  privateKey: string;
}

export interface QueryBalanceResponse {
  value: number;
}

export interface QueryGetOTACoinRequest {
  index: string;
}

export interface QueryGetOTACoinResponse {
  oTACoin: OTACoin | undefined;
}

export interface QueryAllOTACoinRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOTACoinResponse {
  oTACoin: OTACoin[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOutputCoinSerialNumberRequest {}

export interface QueryGetOutputCoinSerialNumberResponse {
  OutputCoinSerialNumber: OutputCoinSerialNumber | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetSerialNumberRequest: object = { index: "" };

export const QueryGetSerialNumberRequest = {
  encode(
    message: QueryGetSerialNumberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSerialNumberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSerialNumberRequest,
    } as QueryGetSerialNumberRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSerialNumberRequest {
    const message = {
      ...baseQueryGetSerialNumberRequest,
    } as QueryGetSerialNumberRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetSerialNumberRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSerialNumberRequest>
  ): QueryGetSerialNumberRequest {
    const message = {
      ...baseQueryGetSerialNumberRequest,
    } as QueryGetSerialNumberRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetSerialNumberResponse: object = {};

export const QueryGetSerialNumberResponse = {
  encode(
    message: QueryGetSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.serialNumber !== undefined) {
      SerialNumber.encode(
        message.serialNumber,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSerialNumberResponse,
    } as QueryGetSerialNumberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serialNumber = SerialNumber.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSerialNumberResponse {
    const message = {
      ...baseQueryGetSerialNumberResponse,
    } as QueryGetSerialNumberResponse;
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = SerialNumber.fromJSON(object.serialNumber);
    } else {
      message.serialNumber = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSerialNumberResponse): unknown {
    const obj: any = {};
    message.serialNumber !== undefined &&
      (obj.serialNumber = message.serialNumber
        ? SerialNumber.toJSON(message.serialNumber)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSerialNumberResponse>
  ): QueryGetSerialNumberResponse {
    const message = {
      ...baseQueryGetSerialNumberResponse,
    } as QueryGetSerialNumberResponse;
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = SerialNumber.fromPartial(object.serialNumber);
    } else {
      message.serialNumber = undefined;
    }
    return message;
  },
};

const baseQueryAllSerialNumberRequest: object = {};

export const QueryAllSerialNumberRequest = {
  encode(
    message: QueryAllSerialNumberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSerialNumberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSerialNumberRequest,
    } as QueryAllSerialNumberRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSerialNumberRequest {
    const message = {
      ...baseQueryAllSerialNumberRequest,
    } as QueryAllSerialNumberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSerialNumberRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSerialNumberRequest>
  ): QueryAllSerialNumberRequest {
    const message = {
      ...baseQueryAllSerialNumberRequest,
    } as QueryAllSerialNumberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSerialNumberResponse: object = {};

export const QueryAllSerialNumberResponse = {
  encode(
    message: QueryAllSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.serialNumber) {
      SerialNumber.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSerialNumberResponse,
    } as QueryAllSerialNumberResponse;
    message.serialNumber = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serialNumber.push(
            SerialNumber.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSerialNumberResponse {
    const message = {
      ...baseQueryAllSerialNumberResponse,
    } as QueryAllSerialNumberResponse;
    message.serialNumber = [];
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      for (const e of object.serialNumber) {
        message.serialNumber.push(SerialNumber.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSerialNumberResponse): unknown {
    const obj: any = {};
    if (message.serialNumber) {
      obj.serialNumber = message.serialNumber.map((e) =>
        e ? SerialNumber.toJSON(e) : undefined
      );
    } else {
      obj.serialNumber = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSerialNumberResponse>
  ): QueryAllSerialNumberResponse {
    const message = {
      ...baseQueryAllSerialNumberResponse,
    } as QueryAllSerialNumberResponse;
    message.serialNumber = [];
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      for (const e of object.serialNumber) {
        message.serialNumber.push(SerialNumber.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetOutputCoinRequest: object = { index: "" };

export const QueryGetOutputCoinRequest = {
  encode(
    message: QueryGetOutputCoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOutputCoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOutputCoinRequest,
    } as QueryGetOutputCoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOutputCoinRequest {
    const message = {
      ...baseQueryGetOutputCoinRequest,
    } as QueryGetOutputCoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOutputCoinRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOutputCoinRequest>
  ): QueryGetOutputCoinRequest {
    const message = {
      ...baseQueryGetOutputCoinRequest,
    } as QueryGetOutputCoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOutputCoinResponse: object = {};

export const QueryGetOutputCoinResponse = {
  encode(
    message: QueryGetOutputCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.outputCoin !== undefined) {
      OutputCoin.encode(message.outputCoin, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOutputCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOutputCoinResponse,
    } as QueryGetOutputCoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outputCoin = OutputCoin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOutputCoinResponse {
    const message = {
      ...baseQueryGetOutputCoinResponse,
    } as QueryGetOutputCoinResponse;
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      message.outputCoin = OutputCoin.fromJSON(object.outputCoin);
    } else {
      message.outputCoin = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOutputCoinResponse): unknown {
    const obj: any = {};
    message.outputCoin !== undefined &&
      (obj.outputCoin = message.outputCoin
        ? OutputCoin.toJSON(message.outputCoin)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOutputCoinResponse>
  ): QueryGetOutputCoinResponse {
    const message = {
      ...baseQueryGetOutputCoinResponse,
    } as QueryGetOutputCoinResponse;
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      message.outputCoin = OutputCoin.fromPartial(object.outputCoin);
    } else {
      message.outputCoin = undefined;
    }
    return message;
  },
};

const baseQueryAllOutputCoinRequest: object = {};

export const QueryAllOutputCoinRequest = {
  encode(
    message: QueryAllOutputCoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOutputCoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOutputCoinRequest,
    } as QueryAllOutputCoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOutputCoinRequest {
    const message = {
      ...baseQueryAllOutputCoinRequest,
    } as QueryAllOutputCoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOutputCoinRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOutputCoinRequest>
  ): QueryAllOutputCoinRequest {
    const message = {
      ...baseQueryAllOutputCoinRequest,
    } as QueryAllOutputCoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOutputCoinResponse: object = {};

export const QueryAllOutputCoinResponse = {
  encode(
    message: QueryAllOutputCoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.outputCoin) {
      OutputCoin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOutputCoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOutputCoinResponse,
    } as QueryAllOutputCoinResponse;
    message.outputCoin = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.outputCoin.push(OutputCoin.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOutputCoinResponse {
    const message = {
      ...baseQueryAllOutputCoinResponse,
    } as QueryAllOutputCoinResponse;
    message.outputCoin = [];
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      for (const e of object.outputCoin) {
        message.outputCoin.push(OutputCoin.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOutputCoinResponse): unknown {
    const obj: any = {};
    if (message.outputCoin) {
      obj.outputCoin = message.outputCoin.map((e) =>
        e ? OutputCoin.toJSON(e) : undefined
      );
    } else {
      obj.outputCoin = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOutputCoinResponse>
  ): QueryAllOutputCoinResponse {
    const message = {
      ...baseQueryAllOutputCoinResponse,
    } as QueryAllOutputCoinResponse;
    message.outputCoin = [];
    if (object.outputCoin !== undefined && object.outputCoin !== null) {
      for (const e of object.outputCoin) {
        message.outputCoin.push(OutputCoin.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetCommitmentRequest: object = { index: "" };

export const QueryGetCommitmentRequest = {
  encode(
    message: QueryGetCommitmentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCommitmentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCommitmentRequest,
    } as QueryGetCommitmentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCommitmentRequest {
    const message = {
      ...baseQueryGetCommitmentRequest,
    } as QueryGetCommitmentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetCommitmentRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCommitmentRequest>
  ): QueryGetCommitmentRequest {
    const message = {
      ...baseQueryGetCommitmentRequest,
    } as QueryGetCommitmentRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetCommitmentResponse: object = {};

export const QueryGetCommitmentResponse = {
  encode(
    message: QueryGetCommitmentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.commitment !== undefined) {
      Commitment.encode(message.commitment, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCommitmentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCommitmentResponse,
    } as QueryGetCommitmentResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commitment = Commitment.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCommitmentResponse {
    const message = {
      ...baseQueryGetCommitmentResponse,
    } as QueryGetCommitmentResponse;
    if (object.commitment !== undefined && object.commitment !== null) {
      message.commitment = Commitment.fromJSON(object.commitment);
    } else {
      message.commitment = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetCommitmentResponse): unknown {
    const obj: any = {};
    message.commitment !== undefined &&
      (obj.commitment = message.commitment
        ? Commitment.toJSON(message.commitment)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCommitmentResponse>
  ): QueryGetCommitmentResponse {
    const message = {
      ...baseQueryGetCommitmentResponse,
    } as QueryGetCommitmentResponse;
    if (object.commitment !== undefined && object.commitment !== null) {
      message.commitment = Commitment.fromPartial(object.commitment);
    } else {
      message.commitment = undefined;
    }
    return message;
  },
};

const baseQueryAllCommitmentRequest: object = {};

export const QueryAllCommitmentRequest = {
  encode(
    message: QueryAllCommitmentRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCommitmentRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCommitmentRequest,
    } as QueryAllCommitmentRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCommitmentRequest {
    const message = {
      ...baseQueryAllCommitmentRequest,
    } as QueryAllCommitmentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCommitmentRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCommitmentRequest>
  ): QueryAllCommitmentRequest {
    const message = {
      ...baseQueryAllCommitmentRequest,
    } as QueryAllCommitmentRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllCommitmentResponse: object = {};

export const QueryAllCommitmentResponse = {
  encode(
    message: QueryAllCommitmentResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.commitment) {
      Commitment.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCommitmentResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCommitmentResponse,
    } as QueryAllCommitmentResponse;
    message.commitment = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commitment.push(Commitment.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCommitmentResponse {
    const message = {
      ...baseQueryAllCommitmentResponse,
    } as QueryAllCommitmentResponse;
    message.commitment = [];
    if (object.commitment !== undefined && object.commitment !== null) {
      for (const e of object.commitment) {
        message.commitment.push(Commitment.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCommitmentResponse): unknown {
    const obj: any = {};
    if (message.commitment) {
      obj.commitment = message.commitment.map((e) =>
        e ? Commitment.toJSON(e) : undefined
      );
    } else {
      obj.commitment = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCommitmentResponse>
  ): QueryAllCommitmentResponse {
    const message = {
      ...baseQueryAllCommitmentResponse,
    } as QueryAllCommitmentResponse;
    message.commitment = [];
    if (object.commitment !== undefined && object.commitment !== null) {
      for (const e of object.commitment) {
        message.commitment.push(Commitment.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetCommitmentIndexRequest: object = { index: "" };

export const QueryGetCommitmentIndexRequest = {
  encode(
    message: QueryGetCommitmentIndexRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCommitmentIndexRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCommitmentIndexRequest,
    } as QueryGetCommitmentIndexRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCommitmentIndexRequest {
    const message = {
      ...baseQueryGetCommitmentIndexRequest,
    } as QueryGetCommitmentIndexRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetCommitmentIndexRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCommitmentIndexRequest>
  ): QueryGetCommitmentIndexRequest {
    const message = {
      ...baseQueryGetCommitmentIndexRequest,
    } as QueryGetCommitmentIndexRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetCommitmentIndexResponse: object = {};

export const QueryGetCommitmentIndexResponse = {
  encode(
    message: QueryGetCommitmentIndexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.commitmentIndex !== undefined) {
      CommitmentIndex.encode(
        message.commitmentIndex,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCommitmentIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCommitmentIndexResponse,
    } as QueryGetCommitmentIndexResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commitmentIndex = CommitmentIndex.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCommitmentIndexResponse {
    const message = {
      ...baseQueryGetCommitmentIndexResponse,
    } as QueryGetCommitmentIndexResponse;
    if (
      object.commitmentIndex !== undefined &&
      object.commitmentIndex !== null
    ) {
      message.commitmentIndex = CommitmentIndex.fromJSON(
        object.commitmentIndex
      );
    } else {
      message.commitmentIndex = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetCommitmentIndexResponse): unknown {
    const obj: any = {};
    message.commitmentIndex !== undefined &&
      (obj.commitmentIndex = message.commitmentIndex
        ? CommitmentIndex.toJSON(message.commitmentIndex)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCommitmentIndexResponse>
  ): QueryGetCommitmentIndexResponse {
    const message = {
      ...baseQueryGetCommitmentIndexResponse,
    } as QueryGetCommitmentIndexResponse;
    if (
      object.commitmentIndex !== undefined &&
      object.commitmentIndex !== null
    ) {
      message.commitmentIndex = CommitmentIndex.fromPartial(
        object.commitmentIndex
      );
    } else {
      message.commitmentIndex = undefined;
    }
    return message;
  },
};

const baseQueryAllCommitmentIndexRequest: object = {};

export const QueryAllCommitmentIndexRequest = {
  encode(
    message: QueryAllCommitmentIndexRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCommitmentIndexRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCommitmentIndexRequest,
    } as QueryAllCommitmentIndexRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCommitmentIndexRequest {
    const message = {
      ...baseQueryAllCommitmentIndexRequest,
    } as QueryAllCommitmentIndexRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCommitmentIndexRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCommitmentIndexRequest>
  ): QueryAllCommitmentIndexRequest {
    const message = {
      ...baseQueryAllCommitmentIndexRequest,
    } as QueryAllCommitmentIndexRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllCommitmentIndexResponse: object = {};

export const QueryAllCommitmentIndexResponse = {
  encode(
    message: QueryAllCommitmentIndexResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.commitmentIndex) {
      CommitmentIndex.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCommitmentIndexResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCommitmentIndexResponse,
    } as QueryAllCommitmentIndexResponse;
    message.commitmentIndex = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.commitmentIndex.push(
            CommitmentIndex.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCommitmentIndexResponse {
    const message = {
      ...baseQueryAllCommitmentIndexResponse,
    } as QueryAllCommitmentIndexResponse;
    message.commitmentIndex = [];
    if (
      object.commitmentIndex !== undefined &&
      object.commitmentIndex !== null
    ) {
      for (const e of object.commitmentIndex) {
        message.commitmentIndex.push(CommitmentIndex.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCommitmentIndexResponse): unknown {
    const obj: any = {};
    if (message.commitmentIndex) {
      obj.commitmentIndex = message.commitmentIndex.map((e) =>
        e ? CommitmentIndex.toJSON(e) : undefined
      );
    } else {
      obj.commitmentIndex = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCommitmentIndexResponse>
  ): QueryAllCommitmentIndexResponse {
    const message = {
      ...baseQueryAllCommitmentIndexResponse,
    } as QueryAllCommitmentIndexResponse;
    message.commitmentIndex = [];
    if (
      object.commitmentIndex !== undefined &&
      object.commitmentIndex !== null
    ) {
      for (const e of object.commitmentIndex) {
        message.commitmentIndex.push(CommitmentIndex.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTokenRequest: object = { index: "" };

export const QueryGetTokenRequest = {
  encode(
    message: QueryGetTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenRequest {
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetTokenRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetTokenRequest>): QueryGetTokenRequest {
    const message = { ...baseQueryGetTokenRequest } as QueryGetTokenRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetTokenResponse: object = {};

export const QueryGetTokenResponse = {
  encode(
    message: QueryGetTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.token !== undefined) {
      Token.encode(message.token, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = Token.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenResponse {
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokenResponse): unknown {
    const obj: any = {};
    message.token !== undefined &&
      (obj.token = message.token ? Token.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenResponse>
  ): QueryGetTokenResponse {
    const message = { ...baseQueryGetTokenResponse } as QueryGetTokenResponse;
    if (object.token !== undefined && object.token !== null) {
      message.token = Token.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenRequest: object = {};

export const QueryAllTokenRequest = {
  encode(
    message: QueryAllTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokenRequest {
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllTokenRequest>): QueryAllTokenRequest {
    const message = { ...baseQueryAllTokenRequest } as QueryAllTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenResponse: object = {};

export const QueryAllTokenResponse = {
  encode(
    message: QueryAllTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.token) {
      Token.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token.push(Token.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTokenResponse {
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    if (object.token !== undefined && object.token !== null) {
      for (const e of object.token) {
        message.token.push(Token.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenResponse): unknown {
    const obj: any = {};
    if (message.token) {
      obj.token = message.token.map((e) => (e ? Token.toJSON(e) : undefined));
    } else {
      obj.token = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenResponse>
  ): QueryAllTokenResponse {
    const message = { ...baseQueryAllTokenResponse } as QueryAllTokenResponse;
    message.token = [];
    if (object.token !== undefined && object.token !== null) {
      for (const e of object.token) {
        message.token.push(Token.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetOnetimeAddressRequest: object = { index: "" };

export const QueryGetOnetimeAddressRequest = {
  encode(
    message: QueryGetOnetimeAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOnetimeAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOnetimeAddressRequest,
    } as QueryGetOnetimeAddressRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOnetimeAddressRequest {
    const message = {
      ...baseQueryGetOnetimeAddressRequest,
    } as QueryGetOnetimeAddressRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOnetimeAddressRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOnetimeAddressRequest>
  ): QueryGetOnetimeAddressRequest {
    const message = {
      ...baseQueryGetOnetimeAddressRequest,
    } as QueryGetOnetimeAddressRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOnetimeAddressResponse: object = {};

export const QueryGetOnetimeAddressResponse = {
  encode(
    message: QueryGetOnetimeAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.onetimeAddress !== undefined) {
      OnetimeAddress.encode(
        message.onetimeAddress,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOnetimeAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOnetimeAddressResponse,
    } as QueryGetOnetimeAddressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.onetimeAddress = OnetimeAddress.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOnetimeAddressResponse {
    const message = {
      ...baseQueryGetOnetimeAddressResponse,
    } as QueryGetOnetimeAddressResponse;
    if (object.onetimeAddress !== undefined && object.onetimeAddress !== null) {
      message.onetimeAddress = OnetimeAddress.fromJSON(object.onetimeAddress);
    } else {
      message.onetimeAddress = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOnetimeAddressResponse): unknown {
    const obj: any = {};
    message.onetimeAddress !== undefined &&
      (obj.onetimeAddress = message.onetimeAddress
        ? OnetimeAddress.toJSON(message.onetimeAddress)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOnetimeAddressResponse>
  ): QueryGetOnetimeAddressResponse {
    const message = {
      ...baseQueryGetOnetimeAddressResponse,
    } as QueryGetOnetimeAddressResponse;
    if (object.onetimeAddress !== undefined && object.onetimeAddress !== null) {
      message.onetimeAddress = OnetimeAddress.fromPartial(
        object.onetimeAddress
      );
    } else {
      message.onetimeAddress = undefined;
    }
    return message;
  },
};

const baseQueryAllOnetimeAddressRequest: object = {};

export const QueryAllOnetimeAddressRequest = {
  encode(
    message: QueryAllOnetimeAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOnetimeAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOnetimeAddressRequest,
    } as QueryAllOnetimeAddressRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOnetimeAddressRequest {
    const message = {
      ...baseQueryAllOnetimeAddressRequest,
    } as QueryAllOnetimeAddressRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOnetimeAddressRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOnetimeAddressRequest>
  ): QueryAllOnetimeAddressRequest {
    const message = {
      ...baseQueryAllOnetimeAddressRequest,
    } as QueryAllOnetimeAddressRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOnetimeAddressResponse: object = {};

export const QueryAllOnetimeAddressResponse = {
  encode(
    message: QueryAllOnetimeAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.onetimeAddress) {
      OnetimeAddress.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllOnetimeAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOnetimeAddressResponse,
    } as QueryAllOnetimeAddressResponse;
    message.onetimeAddress = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.onetimeAddress.push(
            OnetimeAddress.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOnetimeAddressResponse {
    const message = {
      ...baseQueryAllOnetimeAddressResponse,
    } as QueryAllOnetimeAddressResponse;
    message.onetimeAddress = [];
    if (object.onetimeAddress !== undefined && object.onetimeAddress !== null) {
      for (const e of object.onetimeAddress) {
        message.onetimeAddress.push(OnetimeAddress.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOnetimeAddressResponse): unknown {
    const obj: any = {};
    if (message.onetimeAddress) {
      obj.onetimeAddress = message.onetimeAddress.map((e) =>
        e ? OnetimeAddress.toJSON(e) : undefined
      );
    } else {
      obj.onetimeAddress = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOnetimeAddressResponse>
  ): QueryAllOnetimeAddressResponse {
    const message = {
      ...baseQueryAllOnetimeAddressResponse,
    } as QueryAllOnetimeAddressResponse;
    message.onetimeAddress = [];
    if (object.onetimeAddress !== undefined && object.onetimeAddress !== null) {
      for (const e of object.onetimeAddress) {
        message.onetimeAddress.push(OnetimeAddress.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTxPrivacyDataRequest: object = { index: "" };

export const QueryGetTxPrivacyDataRequest = {
  encode(
    message: QueryGetTxPrivacyDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxPrivacyDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxPrivacyDataRequest,
    } as QueryGetTxPrivacyDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTxPrivacyDataRequest {
    const message = {
      ...baseQueryGetTxPrivacyDataRequest,
    } as QueryGetTxPrivacyDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetTxPrivacyDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTxPrivacyDataRequest>
  ): QueryGetTxPrivacyDataRequest {
    const message = {
      ...baseQueryGetTxPrivacyDataRequest,
    } as QueryGetTxPrivacyDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetTxPrivacyDataResponse: object = {};

export const QueryGetTxPrivacyDataResponse = {
  encode(
    message: QueryGetTxPrivacyDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.txPrivacyData !== undefined) {
      TxPrivacyData.encode(
        message.txPrivacyData,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTxPrivacyDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTxPrivacyDataResponse,
    } as QueryGetTxPrivacyDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.txPrivacyData = TxPrivacyData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTxPrivacyDataResponse {
    const message = {
      ...baseQueryGetTxPrivacyDataResponse,
    } as QueryGetTxPrivacyDataResponse;
    if (object.txPrivacyData !== undefined && object.txPrivacyData !== null) {
      message.txPrivacyData = TxPrivacyData.fromJSON(object.txPrivacyData);
    } else {
      message.txPrivacyData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTxPrivacyDataResponse): unknown {
    const obj: any = {};
    message.txPrivacyData !== undefined &&
      (obj.txPrivacyData = message.txPrivacyData
        ? TxPrivacyData.toJSON(message.txPrivacyData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTxPrivacyDataResponse>
  ): QueryGetTxPrivacyDataResponse {
    const message = {
      ...baseQueryGetTxPrivacyDataResponse,
    } as QueryGetTxPrivacyDataResponse;
    if (object.txPrivacyData !== undefined && object.txPrivacyData !== null) {
      message.txPrivacyData = TxPrivacyData.fromPartial(object.txPrivacyData);
    } else {
      message.txPrivacyData = undefined;
    }
    return message;
  },
};

const baseQueryAllTxPrivacyDataRequest: object = {};

export const QueryAllTxPrivacyDataRequest = {
  encode(
    message: QueryAllTxPrivacyDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTxPrivacyDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTxPrivacyDataRequest,
    } as QueryAllTxPrivacyDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTxPrivacyDataRequest {
    const message = {
      ...baseQueryAllTxPrivacyDataRequest,
    } as QueryAllTxPrivacyDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTxPrivacyDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTxPrivacyDataRequest>
  ): QueryAllTxPrivacyDataRequest {
    const message = {
      ...baseQueryAllTxPrivacyDataRequest,
    } as QueryAllTxPrivacyDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTxPrivacyDataResponse: object = {};

export const QueryAllTxPrivacyDataResponse = {
  encode(
    message: QueryAllTxPrivacyDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.txPrivacyData) {
      TxPrivacyData.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTxPrivacyDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTxPrivacyDataResponse,
    } as QueryAllTxPrivacyDataResponse;
    message.txPrivacyData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.txPrivacyData.push(
            TxPrivacyData.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTxPrivacyDataResponse {
    const message = {
      ...baseQueryAllTxPrivacyDataResponse,
    } as QueryAllTxPrivacyDataResponse;
    message.txPrivacyData = [];
    if (object.txPrivacyData !== undefined && object.txPrivacyData !== null) {
      for (const e of object.txPrivacyData) {
        message.txPrivacyData.push(TxPrivacyData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTxPrivacyDataResponse): unknown {
    const obj: any = {};
    if (message.txPrivacyData) {
      obj.txPrivacyData = message.txPrivacyData.map((e) =>
        e ? TxPrivacyData.toJSON(e) : undefined
      );
    } else {
      obj.txPrivacyData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTxPrivacyDataResponse>
  ): QueryAllTxPrivacyDataResponse {
    const message = {
      ...baseQueryAllTxPrivacyDataResponse,
    } as QueryAllTxPrivacyDataResponse;
    message.txPrivacyData = [];
    if (object.txPrivacyData !== undefined && object.txPrivacyData !== null) {
      for (const e of object.txPrivacyData) {
        message.txPrivacyData.push(TxPrivacyData.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryBalanceRequest: object = { privateKey: "" };

export const QueryBalanceRequest = {
  encode(
    message: QueryBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.privateKey !== "") {
      writer.uint32(10).string(message.privateKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBalanceRequest } as QueryBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.privateKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBalanceRequest {
    const message = { ...baseQueryBalanceRequest } as QueryBalanceRequest;
    if (object.privateKey !== undefined && object.privateKey !== null) {
      message.privateKey = String(object.privateKey);
    } else {
      message.privateKey = "";
    }
    return message;
  },

  toJSON(message: QueryBalanceRequest): unknown {
    const obj: any = {};
    message.privateKey !== undefined && (obj.privateKey = message.privateKey);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBalanceRequest>): QueryBalanceRequest {
    const message = { ...baseQueryBalanceRequest } as QueryBalanceRequest;
    if (object.privateKey !== undefined && object.privateKey !== null) {
      message.privateKey = object.privateKey;
    } else {
      message.privateKey = "";
    }
    return message;
  },
};

const baseQueryBalanceResponse: object = { value: 0 };

export const QueryBalanceResponse = {
  encode(
    message: QueryBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.value !== 0) {
      writer.uint32(8).uint64(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBalanceResponse } as QueryBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBalanceResponse {
    const message = { ...baseQueryBalanceResponse } as QueryBalanceResponse;
    if (object.value !== undefined && object.value !== null) {
      message.value = Number(object.value);
    } else {
      message.value = 0;
    }
    return message;
  },

  toJSON(message: QueryBalanceResponse): unknown {
    const obj: any = {};
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBalanceResponse>): QueryBalanceResponse {
    const message = { ...baseQueryBalanceResponse } as QueryBalanceResponse;
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = 0;
    }
    return message;
  },
};

const baseQueryGetOTACoinRequest: object = { index: "" };

export const QueryGetOTACoinRequest = {
  encode(
    message: QueryGetOTACoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOTACoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOTACoinRequest } as QueryGetOTACoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOTACoinRequest {
    const message = { ...baseQueryGetOTACoinRequest } as QueryGetOTACoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOTACoinRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOTACoinRequest>
  ): QueryGetOTACoinRequest {
    const message = { ...baseQueryGetOTACoinRequest } as QueryGetOTACoinRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOTACoinResponse: object = {};

export const QueryGetOTACoinResponse = {
  encode(
    message: QueryGetOTACoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.oTACoin !== undefined) {
      OTACoin.encode(message.oTACoin, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOTACoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOTACoinResponse,
    } as QueryGetOTACoinResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.oTACoin = OTACoin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOTACoinResponse {
    const message = {
      ...baseQueryGetOTACoinResponse,
    } as QueryGetOTACoinResponse;
    if (object.oTACoin !== undefined && object.oTACoin !== null) {
      message.oTACoin = OTACoin.fromJSON(object.oTACoin);
    } else {
      message.oTACoin = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOTACoinResponse): unknown {
    const obj: any = {};
    message.oTACoin !== undefined &&
      (obj.oTACoin = message.oTACoin
        ? OTACoin.toJSON(message.oTACoin)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOTACoinResponse>
  ): QueryGetOTACoinResponse {
    const message = {
      ...baseQueryGetOTACoinResponse,
    } as QueryGetOTACoinResponse;
    if (object.oTACoin !== undefined && object.oTACoin !== null) {
      message.oTACoin = OTACoin.fromPartial(object.oTACoin);
    } else {
      message.oTACoin = undefined;
    }
    return message;
  },
};

const baseQueryAllOTACoinRequest: object = {};

export const QueryAllOTACoinRequest = {
  encode(
    message: QueryAllOTACoinRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOTACoinRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOTACoinRequest } as QueryAllOTACoinRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOTACoinRequest {
    const message = { ...baseQueryAllOTACoinRequest } as QueryAllOTACoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOTACoinRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOTACoinRequest>
  ): QueryAllOTACoinRequest {
    const message = { ...baseQueryAllOTACoinRequest } as QueryAllOTACoinRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOTACoinResponse: object = {};

export const QueryAllOTACoinResponse = {
  encode(
    message: QueryAllOTACoinResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.oTACoin) {
      OTACoin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOTACoinResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllOTACoinResponse,
    } as QueryAllOTACoinResponse;
    message.oTACoin = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.oTACoin.push(OTACoin.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOTACoinResponse {
    const message = {
      ...baseQueryAllOTACoinResponse,
    } as QueryAllOTACoinResponse;
    message.oTACoin = [];
    if (object.oTACoin !== undefined && object.oTACoin !== null) {
      for (const e of object.oTACoin) {
        message.oTACoin.push(OTACoin.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOTACoinResponse): unknown {
    const obj: any = {};
    if (message.oTACoin) {
      obj.oTACoin = message.oTACoin.map((e) =>
        e ? OTACoin.toJSON(e) : undefined
      );
    } else {
      obj.oTACoin = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOTACoinResponse>
  ): QueryAllOTACoinResponse {
    const message = {
      ...baseQueryAllOTACoinResponse,
    } as QueryAllOTACoinResponse;
    message.oTACoin = [];
    if (object.oTACoin !== undefined && object.oTACoin !== null) {
      for (const e of object.oTACoin) {
        message.oTACoin.push(OTACoin.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetOutputCoinSerialNumberRequest: object = {};

export const QueryGetOutputCoinSerialNumberRequest = {
  encode(
    _: QueryGetOutputCoinSerialNumberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOutputCoinSerialNumberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOutputCoinSerialNumberRequest,
    } as QueryGetOutputCoinSerialNumberRequest;
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

  fromJSON(_: any): QueryGetOutputCoinSerialNumberRequest {
    const message = {
      ...baseQueryGetOutputCoinSerialNumberRequest,
    } as QueryGetOutputCoinSerialNumberRequest;
    return message;
  },

  toJSON(_: QueryGetOutputCoinSerialNumberRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetOutputCoinSerialNumberRequest>
  ): QueryGetOutputCoinSerialNumberRequest {
    const message = {
      ...baseQueryGetOutputCoinSerialNumberRequest,
    } as QueryGetOutputCoinSerialNumberRequest;
    return message;
  },
};

const baseQueryGetOutputCoinSerialNumberResponse: object = {};

export const QueryGetOutputCoinSerialNumberResponse = {
  encode(
    message: QueryGetOutputCoinSerialNumberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.OutputCoinSerialNumber !== undefined) {
      OutputCoinSerialNumber.encode(
        message.OutputCoinSerialNumber,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetOutputCoinSerialNumberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetOutputCoinSerialNumberResponse,
    } as QueryGetOutputCoinSerialNumberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.OutputCoinSerialNumber = OutputCoinSerialNumber.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOutputCoinSerialNumberResponse {
    const message = {
      ...baseQueryGetOutputCoinSerialNumberResponse,
    } as QueryGetOutputCoinSerialNumberResponse;
    if (
      object.OutputCoinSerialNumber !== undefined &&
      object.OutputCoinSerialNumber !== null
    ) {
      message.OutputCoinSerialNumber = OutputCoinSerialNumber.fromJSON(
        object.OutputCoinSerialNumber
      );
    } else {
      message.OutputCoinSerialNumber = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOutputCoinSerialNumberResponse): unknown {
    const obj: any = {};
    message.OutputCoinSerialNumber !== undefined &&
      (obj.OutputCoinSerialNumber = message.OutputCoinSerialNumber
        ? OutputCoinSerialNumber.toJSON(message.OutputCoinSerialNumber)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOutputCoinSerialNumberResponse>
  ): QueryGetOutputCoinSerialNumberResponse {
    const message = {
      ...baseQueryGetOutputCoinSerialNumberResponse,
    } as QueryGetOutputCoinSerialNumberResponse;
    if (
      object.OutputCoinSerialNumber !== undefined &&
      object.OutputCoinSerialNumber !== null
    ) {
      message.OutputCoinSerialNumber = OutputCoinSerialNumber.fromPartial(
        object.OutputCoinSerialNumber
      );
    } else {
      message.OutputCoinSerialNumber = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SerialNumber by index. */
  SerialNumber(
    request: QueryGetSerialNumberRequest
  ): Promise<QueryGetSerialNumberResponse>;
  /** Queries a list of SerialNumber items. */
  SerialNumberAll(
    request: QueryAllSerialNumberRequest
  ): Promise<QueryAllSerialNumberResponse>;
  /** Queries a OutputCoin by index. */
  OutputCoin(
    request: QueryGetOutputCoinRequest
  ): Promise<QueryGetOutputCoinResponse>;
  /** Queries a list of OutputCoin items. */
  OutputCoinAll(
    request: QueryAllOutputCoinRequest
  ): Promise<QueryAllOutputCoinResponse>;
  /** Queries a Commitment by index. */
  Commitment(
    request: QueryGetCommitmentRequest
  ): Promise<QueryGetCommitmentResponse>;
  /** Queries a list of Commitment items. */
  CommitmentAll(
    request: QueryAllCommitmentRequest
  ): Promise<QueryAllCommitmentResponse>;
  /** Queries a CommitmentIndex by index. */
  CommitmentIndex(
    request: QueryGetCommitmentIndexRequest
  ): Promise<QueryGetCommitmentIndexResponse>;
  /** Queries a list of CommitmentIndex items. */
  CommitmentIndexAll(
    request: QueryAllCommitmentIndexRequest
  ): Promise<QueryAllCommitmentIndexResponse>;
  /** Queries a Token by index. */
  Token(request: QueryGetTokenRequest): Promise<QueryGetTokenResponse>;
  /** Queries a list of Token items. */
  TokenAll(request: QueryAllTokenRequest): Promise<QueryAllTokenResponse>;
  /** Queries a OnetimeAddress by index. */
  OnetimeAddress(
    request: QueryGetOnetimeAddressRequest
  ): Promise<QueryGetOnetimeAddressResponse>;
  /** Queries a list of OnetimeAddress items. */
  OnetimeAddressAll(
    request: QueryAllOnetimeAddressRequest
  ): Promise<QueryAllOnetimeAddressResponse>;
  /** Queries a TxPrivacyData by index. */
  TxPrivacyData(
    request: QueryGetTxPrivacyDataRequest
  ): Promise<QueryGetTxPrivacyDataResponse>;
  /** Queries a list of TxPrivacyData items. */
  TxPrivacyDataAll(
    request: QueryAllTxPrivacyDataRequest
  ): Promise<QueryAllTxPrivacyDataResponse>;
  /** Queries a list of Balance items. */
  Balance(request: QueryBalanceRequest): Promise<QueryBalanceResponse>;
  /** Queries a OTACoin by index. */
  OTACoin(request: QueryGetOTACoinRequest): Promise<QueryGetOTACoinResponse>;
  /** Queries a list of OTACoin items. */
  OTACoinAll(request: QueryAllOTACoinRequest): Promise<QueryAllOTACoinResponse>;
  /** Queries a OutputCoinSerialNumber by index. */
  OutputCoinSerialNumber(
    request: QueryGetOutputCoinSerialNumberRequest
  ): Promise<QueryGetOutputCoinSerialNumberResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SerialNumber(
    request: QueryGetSerialNumberRequest
  ): Promise<QueryGetSerialNumberResponse> {
    const data = QueryGetSerialNumberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "SerialNumber",
      data
    );
    return promise.then((data) =>
      QueryGetSerialNumberResponse.decode(new Reader(data))
    );
  }

  SerialNumberAll(
    request: QueryAllSerialNumberRequest
  ): Promise<QueryAllSerialNumberResponse> {
    const data = QueryAllSerialNumberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "SerialNumberAll",
      data
    );
    return promise.then((data) =>
      QueryAllSerialNumberResponse.decode(new Reader(data))
    );
  }

  OutputCoin(
    request: QueryGetOutputCoinRequest
  ): Promise<QueryGetOutputCoinResponse> {
    const data = QueryGetOutputCoinRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OutputCoin",
      data
    );
    return promise.then((data) =>
      QueryGetOutputCoinResponse.decode(new Reader(data))
    );
  }

  OutputCoinAll(
    request: QueryAllOutputCoinRequest
  ): Promise<QueryAllOutputCoinResponse> {
    const data = QueryAllOutputCoinRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OutputCoinAll",
      data
    );
    return promise.then((data) =>
      QueryAllOutputCoinResponse.decode(new Reader(data))
    );
  }

  Commitment(
    request: QueryGetCommitmentRequest
  ): Promise<QueryGetCommitmentResponse> {
    const data = QueryGetCommitmentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "Commitment",
      data
    );
    return promise.then((data) =>
      QueryGetCommitmentResponse.decode(new Reader(data))
    );
  }

  CommitmentAll(
    request: QueryAllCommitmentRequest
  ): Promise<QueryAllCommitmentResponse> {
    const data = QueryAllCommitmentRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "CommitmentAll",
      data
    );
    return promise.then((data) =>
      QueryAllCommitmentResponse.decode(new Reader(data))
    );
  }

  CommitmentIndex(
    request: QueryGetCommitmentIndexRequest
  ): Promise<QueryGetCommitmentIndexResponse> {
    const data = QueryGetCommitmentIndexRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "CommitmentIndex",
      data
    );
    return promise.then((data) =>
      QueryGetCommitmentIndexResponse.decode(new Reader(data))
    );
  }

  CommitmentIndexAll(
    request: QueryAllCommitmentIndexRequest
  ): Promise<QueryAllCommitmentIndexResponse> {
    const data = QueryAllCommitmentIndexRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "CommitmentIndexAll",
      data
    );
    return promise.then((data) =>
      QueryAllCommitmentIndexResponse.decode(new Reader(data))
    );
  }

  Token(request: QueryGetTokenRequest): Promise<QueryGetTokenResponse> {
    const data = QueryGetTokenRequest.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Query", "Token", data);
    return promise.then((data) =>
      QueryGetTokenResponse.decode(new Reader(data))
    );
  }

  TokenAll(request: QueryAllTokenRequest): Promise<QueryAllTokenResponse> {
    const data = QueryAllTokenRequest.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Query", "TokenAll", data);
    return promise.then((data) =>
      QueryAllTokenResponse.decode(new Reader(data))
    );
  }

  OnetimeAddress(
    request: QueryGetOnetimeAddressRequest
  ): Promise<QueryGetOnetimeAddressResponse> {
    const data = QueryGetOnetimeAddressRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OnetimeAddress",
      data
    );
    return promise.then((data) =>
      QueryGetOnetimeAddressResponse.decode(new Reader(data))
    );
  }

  OnetimeAddressAll(
    request: QueryAllOnetimeAddressRequest
  ): Promise<QueryAllOnetimeAddressResponse> {
    const data = QueryAllOnetimeAddressRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OnetimeAddressAll",
      data
    );
    return promise.then((data) =>
      QueryAllOnetimeAddressResponse.decode(new Reader(data))
    );
  }

  TxPrivacyData(
    request: QueryGetTxPrivacyDataRequest
  ): Promise<QueryGetTxPrivacyDataResponse> {
    const data = QueryGetTxPrivacyDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "TxPrivacyData",
      data
    );
    return promise.then((data) =>
      QueryGetTxPrivacyDataResponse.decode(new Reader(data))
    );
  }

  TxPrivacyDataAll(
    request: QueryAllTxPrivacyDataRequest
  ): Promise<QueryAllTxPrivacyDataResponse> {
    const data = QueryAllTxPrivacyDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "TxPrivacyDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllTxPrivacyDataResponse.decode(new Reader(data))
    );
  }

  Balance(request: QueryBalanceRequest): Promise<QueryBalanceResponse> {
    const data = QueryBalanceRequest.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Query", "Balance", data);
    return promise.then((data) =>
      QueryBalanceResponse.decode(new Reader(data))
    );
  }

  OTACoin(request: QueryGetOTACoinRequest): Promise<QueryGetOTACoinResponse> {
    const data = QueryGetOTACoinRequest.encode(request).finish();
    const promise = this.rpc.request("privacy.privacy.Query", "OTACoin", data);
    return promise.then((data) =>
      QueryGetOTACoinResponse.decode(new Reader(data))
    );
  }

  OTACoinAll(
    request: QueryAllOTACoinRequest
  ): Promise<QueryAllOTACoinResponse> {
    const data = QueryAllOTACoinRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OTACoinAll",
      data
    );
    return promise.then((data) =>
      QueryAllOTACoinResponse.decode(new Reader(data))
    );
  }

  OutputCoinSerialNumber(
    request: QueryGetOutputCoinSerialNumberRequest
  ): Promise<QueryGetOutputCoinSerialNumberResponse> {
    const data = QueryGetOutputCoinSerialNumberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "privacy.privacy.Query",
      "OutputCoinSerialNumber",
      data
    );
    return promise.then((data) =>
      QueryGetOutputCoinSerialNumberResponse.decode(new Reader(data))
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
