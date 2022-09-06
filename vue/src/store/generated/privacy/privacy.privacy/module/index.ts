// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateToken } from "./types/privacy/tx";
import { MsgUpdateOnetimeAddress } from "./types/privacy/tx";
import { MsgCreateOnetimeAddress } from "./types/privacy/tx";
import { MsgTransfer } from "./types/privacy/tx";
import { MsgDeleteOnetimeAddress } from "./types/privacy/tx";
import { MsgUpdateOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgUpdateSerialNumber } from "./types/privacy/tx";
import { MsgDeleteCommitmentIndex } from "./types/privacy/tx";
import { MsgAirdrop } from "./types/privacy/tx";
import { MsgCreateOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgCreateToken } from "./types/privacy/tx";
import { MsgDeleteSerialNumber } from "./types/privacy/tx";
import { MsgUpdateTxPrivacyData } from "./types/privacy/tx";
import { MsgCreateCommitmentIndex } from "./types/privacy/tx";
import { MsgDeleteTxPrivacyData } from "./types/privacy/tx";
import { MsgCreateTx } from "./types/privacy/tx";
import { MsgDeleteCommitment } from "./types/privacy/tx";
import { MsgUpdateCommitment } from "./types/privacy/tx";
import { MsgUpdateOutputCoin } from "./types/privacy/tx";
import { MsgCreateCommitment } from "./types/privacy/tx";
import { MsgCreateOutputCoin } from "./types/privacy/tx";
import { MsgCreateTxPrivacyData } from "./types/privacy/tx";
import { MsgDeleteToken } from "./types/privacy/tx";
import { MsgDeleteOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgCreateSerialNumber } from "./types/privacy/tx";
import { MsgDeleteOutputCoin } from "./types/privacy/tx";
import { MsgUpdateCommitmentIndex } from "./types/privacy/tx";


const types = [
  ["/privacy.privacy.MsgUpdateToken", MsgUpdateToken],
  ["/privacy.privacy.MsgUpdateOnetimeAddress", MsgUpdateOnetimeAddress],
  ["/privacy.privacy.MsgCreateOnetimeAddress", MsgCreateOnetimeAddress],
  ["/privacy.privacy.MsgTransfer", MsgTransfer],
  ["/privacy.privacy.MsgDeleteOnetimeAddress", MsgDeleteOnetimeAddress],
  ["/privacy.privacy.MsgUpdateOutputCoinSerialNumber", MsgUpdateOutputCoinSerialNumber],
  ["/privacy.privacy.MsgUpdateSerialNumber", MsgUpdateSerialNumber],
  ["/privacy.privacy.MsgDeleteCommitmentIndex", MsgDeleteCommitmentIndex],
  ["/privacy.privacy.MsgAirdrop", MsgAirdrop],
  ["/privacy.privacy.MsgCreateOutputCoinSerialNumber", MsgCreateOutputCoinSerialNumber],
  ["/privacy.privacy.MsgCreateToken", MsgCreateToken],
  ["/privacy.privacy.MsgDeleteSerialNumber", MsgDeleteSerialNumber],
  ["/privacy.privacy.MsgUpdateTxPrivacyData", MsgUpdateTxPrivacyData],
  ["/privacy.privacy.MsgCreateCommitmentIndex", MsgCreateCommitmentIndex],
  ["/privacy.privacy.MsgDeleteTxPrivacyData", MsgDeleteTxPrivacyData],
  ["/privacy.privacy.MsgCreateTx", MsgCreateTx],
  ["/privacy.privacy.MsgDeleteCommitment", MsgDeleteCommitment],
  ["/privacy.privacy.MsgUpdateCommitment", MsgUpdateCommitment],
  ["/privacy.privacy.MsgUpdateOutputCoin", MsgUpdateOutputCoin],
  ["/privacy.privacy.MsgCreateCommitment", MsgCreateCommitment],
  ["/privacy.privacy.MsgCreateOutputCoin", MsgCreateOutputCoin],
  ["/privacy.privacy.MsgCreateTxPrivacyData", MsgCreateTxPrivacyData],
  ["/privacy.privacy.MsgDeleteToken", MsgDeleteToken],
  ["/privacy.privacy.MsgDeleteOutputCoinSerialNumber", MsgDeleteOutputCoinSerialNumber],
  ["/privacy.privacy.MsgCreateSerialNumber", MsgCreateSerialNumber],
  ["/privacy.privacy.MsgDeleteOutputCoin", MsgDeleteOutputCoin],
  ["/privacy.privacy.MsgUpdateCommitmentIndex", MsgUpdateCommitmentIndex],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgUpdateToken: (data: MsgUpdateToken): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateToken", value: MsgUpdateToken.fromPartial( data ) }),
    msgUpdateOnetimeAddress: (data: MsgUpdateOnetimeAddress): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateOnetimeAddress", value: MsgUpdateOnetimeAddress.fromPartial( data ) }),
    msgCreateOnetimeAddress: (data: MsgCreateOnetimeAddress): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateOnetimeAddress", value: MsgCreateOnetimeAddress.fromPartial( data ) }),
    msgTransfer: (data: MsgTransfer): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgTransfer", value: MsgTransfer.fromPartial( data ) }),
    msgDeleteOnetimeAddress: (data: MsgDeleteOnetimeAddress): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteOnetimeAddress", value: MsgDeleteOnetimeAddress.fromPartial( data ) }),
    msgUpdateOutputCoinSerialNumber: (data: MsgUpdateOutputCoinSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateOutputCoinSerialNumber", value: MsgUpdateOutputCoinSerialNumber.fromPartial( data ) }),
    msgUpdateSerialNumber: (data: MsgUpdateSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateSerialNumber", value: MsgUpdateSerialNumber.fromPartial( data ) }),
    msgDeleteCommitmentIndex: (data: MsgDeleteCommitmentIndex): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteCommitmentIndex", value: MsgDeleteCommitmentIndex.fromPartial( data ) }),
    msgAirdrop: (data: MsgAirdrop): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgAirdrop", value: MsgAirdrop.fromPartial( data ) }),
    msgCreateOutputCoinSerialNumber: (data: MsgCreateOutputCoinSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateOutputCoinSerialNumber", value: MsgCreateOutputCoinSerialNumber.fromPartial( data ) }),
    msgCreateToken: (data: MsgCreateToken): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateToken", value: MsgCreateToken.fromPartial( data ) }),
    msgDeleteSerialNumber: (data: MsgDeleteSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteSerialNumber", value: MsgDeleteSerialNumber.fromPartial( data ) }),
    msgUpdateTxPrivacyData: (data: MsgUpdateTxPrivacyData): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateTxPrivacyData", value: MsgUpdateTxPrivacyData.fromPartial( data ) }),
    msgCreateCommitmentIndex: (data: MsgCreateCommitmentIndex): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateCommitmentIndex", value: MsgCreateCommitmentIndex.fromPartial( data ) }),
    msgDeleteTxPrivacyData: (data: MsgDeleteTxPrivacyData): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteTxPrivacyData", value: MsgDeleteTxPrivacyData.fromPartial( data ) }),
    msgCreateTx: (data: MsgCreateTx): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateTx", value: MsgCreateTx.fromPartial( data ) }),
    msgDeleteCommitment: (data: MsgDeleteCommitment): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteCommitment", value: MsgDeleteCommitment.fromPartial( data ) }),
    msgUpdateCommitment: (data: MsgUpdateCommitment): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateCommitment", value: MsgUpdateCommitment.fromPartial( data ) }),
    msgUpdateOutputCoin: (data: MsgUpdateOutputCoin): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateOutputCoin", value: MsgUpdateOutputCoin.fromPartial( data ) }),
    msgCreateCommitment: (data: MsgCreateCommitment): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateCommitment", value: MsgCreateCommitment.fromPartial( data ) }),
    msgCreateOutputCoin: (data: MsgCreateOutputCoin): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateOutputCoin", value: MsgCreateOutputCoin.fromPartial( data ) }),
    msgCreateTxPrivacyData: (data: MsgCreateTxPrivacyData): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateTxPrivacyData", value: MsgCreateTxPrivacyData.fromPartial( data ) }),
    msgDeleteToken: (data: MsgDeleteToken): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteToken", value: MsgDeleteToken.fromPartial( data ) }),
    msgDeleteOutputCoinSerialNumber: (data: MsgDeleteOutputCoinSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteOutputCoinSerialNumber", value: MsgDeleteOutputCoinSerialNumber.fromPartial( data ) }),
    msgCreateSerialNumber: (data: MsgCreateSerialNumber): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgCreateSerialNumber", value: MsgCreateSerialNumber.fromPartial( data ) }),
    msgDeleteOutputCoin: (data: MsgDeleteOutputCoin): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgDeleteOutputCoin", value: MsgDeleteOutputCoin.fromPartial( data ) }),
    msgUpdateCommitmentIndex: (data: MsgUpdateCommitmentIndex): EncodeObject => ({ typeUrl: "/privacy.privacy.MsgUpdateCommitmentIndex", value: MsgUpdateCommitmentIndex.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
