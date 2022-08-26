/* eslint-disable */
import { Params } from "../privacy/params";
import { SerialNumber } from "../privacy/serial_number";
import { OutputCoin } from "../privacy/output_coin";
import { Commitment } from "../privacy/commitment";
import { CommitmentIndex } from "../privacy/commitment_index";
import { Token } from "../privacy/token";
import { OnetimeAddress } from "../privacy/onetime_address";
import { TxPrivacyData } from "../privacy/tx_privacy_data";
import { OTACoin } from "../privacy/ota_coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "privacy.privacy";

/** GenesisState defines the privacy module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  serialNumberList: SerialNumber[];
  outputCoinList: OutputCoin[];
  commitmentList: Commitment[];
  commitmentIndexList: CommitmentIndex[];
  tokenList: Token[];
  onetimeAddressList: OnetimeAddress[];
  txPrivacyDataList: TxPrivacyData[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  oTACoinList: OTACoin[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.serialNumberList) {
      SerialNumber.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.outputCoinList) {
      OutputCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.commitmentList) {
      Commitment.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.commitmentIndexList) {
      CommitmentIndex.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.tokenList) {
      Token.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.onetimeAddressList) {
      OnetimeAddress.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.txPrivacyDataList) {
      TxPrivacyData.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.oTACoinList) {
      OTACoin.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.serialNumberList = [];
    message.outputCoinList = [];
    message.commitmentList = [];
    message.commitmentIndexList = [];
    message.tokenList = [];
    message.onetimeAddressList = [];
    message.txPrivacyDataList = [];
    message.oTACoinList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.serialNumberList.push(
            SerialNumber.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.outputCoinList.push(
            OutputCoin.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.commitmentList.push(
            Commitment.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.commitmentIndexList.push(
            CommitmentIndex.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.tokenList.push(Token.decode(reader, reader.uint32()));
          break;
        case 7:
          message.onetimeAddressList.push(
            OnetimeAddress.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.txPrivacyDataList.push(
            TxPrivacyData.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.oTACoinList.push(OTACoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.serialNumberList = [];
    message.outputCoinList = [];
    message.commitmentList = [];
    message.commitmentIndexList = [];
    message.tokenList = [];
    message.onetimeAddressList = [];
    message.txPrivacyDataList = [];
    message.oTACoinList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.serialNumberList !== undefined &&
      object.serialNumberList !== null
    ) {
      for (const e of object.serialNumberList) {
        message.serialNumberList.push(SerialNumber.fromJSON(e));
      }
    }
    if (object.outputCoinList !== undefined && object.outputCoinList !== null) {
      for (const e of object.outputCoinList) {
        message.outputCoinList.push(OutputCoin.fromJSON(e));
      }
    }
    if (object.commitmentList !== undefined && object.commitmentList !== null) {
      for (const e of object.commitmentList) {
        message.commitmentList.push(Commitment.fromJSON(e));
      }
    }
    if (
      object.commitmentIndexList !== undefined &&
      object.commitmentIndexList !== null
    ) {
      for (const e of object.commitmentIndexList) {
        message.commitmentIndexList.push(CommitmentIndex.fromJSON(e));
      }
    }
    if (object.tokenList !== undefined && object.tokenList !== null) {
      for (const e of object.tokenList) {
        message.tokenList.push(Token.fromJSON(e));
      }
    }
    if (
      object.onetimeAddressList !== undefined &&
      object.onetimeAddressList !== null
    ) {
      for (const e of object.onetimeAddressList) {
        message.onetimeAddressList.push(OnetimeAddress.fromJSON(e));
      }
    }
    if (
      object.txPrivacyDataList !== undefined &&
      object.txPrivacyDataList !== null
    ) {
      for (const e of object.txPrivacyDataList) {
        message.txPrivacyDataList.push(TxPrivacyData.fromJSON(e));
      }
    }
    if (object.oTACoinList !== undefined && object.oTACoinList !== null) {
      for (const e of object.oTACoinList) {
        message.oTACoinList.push(OTACoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.serialNumberList) {
      obj.serialNumberList = message.serialNumberList.map((e) =>
        e ? SerialNumber.toJSON(e) : undefined
      );
    } else {
      obj.serialNumberList = [];
    }
    if (message.outputCoinList) {
      obj.outputCoinList = message.outputCoinList.map((e) =>
        e ? OutputCoin.toJSON(e) : undefined
      );
    } else {
      obj.outputCoinList = [];
    }
    if (message.commitmentList) {
      obj.commitmentList = message.commitmentList.map((e) =>
        e ? Commitment.toJSON(e) : undefined
      );
    } else {
      obj.commitmentList = [];
    }
    if (message.commitmentIndexList) {
      obj.commitmentIndexList = message.commitmentIndexList.map((e) =>
        e ? CommitmentIndex.toJSON(e) : undefined
      );
    } else {
      obj.commitmentIndexList = [];
    }
    if (message.tokenList) {
      obj.tokenList = message.tokenList.map((e) =>
        e ? Token.toJSON(e) : undefined
      );
    } else {
      obj.tokenList = [];
    }
    if (message.onetimeAddressList) {
      obj.onetimeAddressList = message.onetimeAddressList.map((e) =>
        e ? OnetimeAddress.toJSON(e) : undefined
      );
    } else {
      obj.onetimeAddressList = [];
    }
    if (message.txPrivacyDataList) {
      obj.txPrivacyDataList = message.txPrivacyDataList.map((e) =>
        e ? TxPrivacyData.toJSON(e) : undefined
      );
    } else {
      obj.txPrivacyDataList = [];
    }
    if (message.oTACoinList) {
      obj.oTACoinList = message.oTACoinList.map((e) =>
        e ? OTACoin.toJSON(e) : undefined
      );
    } else {
      obj.oTACoinList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.serialNumberList = [];
    message.outputCoinList = [];
    message.commitmentList = [];
    message.commitmentIndexList = [];
    message.tokenList = [];
    message.onetimeAddressList = [];
    message.txPrivacyDataList = [];
    message.oTACoinList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.serialNumberList !== undefined &&
      object.serialNumberList !== null
    ) {
      for (const e of object.serialNumberList) {
        message.serialNumberList.push(SerialNumber.fromPartial(e));
      }
    }
    if (object.outputCoinList !== undefined && object.outputCoinList !== null) {
      for (const e of object.outputCoinList) {
        message.outputCoinList.push(OutputCoin.fromPartial(e));
      }
    }
    if (object.commitmentList !== undefined && object.commitmentList !== null) {
      for (const e of object.commitmentList) {
        message.commitmentList.push(Commitment.fromPartial(e));
      }
    }
    if (
      object.commitmentIndexList !== undefined &&
      object.commitmentIndexList !== null
    ) {
      for (const e of object.commitmentIndexList) {
        message.commitmentIndexList.push(CommitmentIndex.fromPartial(e));
      }
    }
    if (object.tokenList !== undefined && object.tokenList !== null) {
      for (const e of object.tokenList) {
        message.tokenList.push(Token.fromPartial(e));
      }
    }
    if (
      object.onetimeAddressList !== undefined &&
      object.onetimeAddressList !== null
    ) {
      for (const e of object.onetimeAddressList) {
        message.onetimeAddressList.push(OnetimeAddress.fromPartial(e));
      }
    }
    if (
      object.txPrivacyDataList !== undefined &&
      object.txPrivacyDataList !== null
    ) {
      for (const e of object.txPrivacyDataList) {
        message.txPrivacyDataList.push(TxPrivacyData.fromPartial(e));
      }
    }
    if (object.oTACoinList !== undefined && object.oTACoinList !== null) {
      for (const e of object.oTACoinList) {
        message.oTACoinList.push(OTACoin.fromPartial(e));
      }
    }
    return message;
  },
};

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
