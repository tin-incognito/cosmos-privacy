import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgTransfer } from "./types/privacy/tx";
import { MsgDeleteSerialNumber } from "./types/privacy/tx";
import { MsgDeleteCommitmentIndex } from "./types/privacy/tx";
import { MsgCreateOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgDeleteOutputCoin } from "./types/privacy/tx";
import { MsgCreateOnetimeAddress } from "./types/privacy/tx";
import { MsgCreateOutputCoin } from "./types/privacy/tx";
import { MsgUpdateOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgDeleteTxPrivacyData } from "./types/privacy/tx";
import { MsgPrivacyData } from "./types/privacy/tx";
import { MsgUpdateCommitmentIndex } from "./types/privacy/tx";
import { MsgUpdateSerialNumber } from "./types/privacy/tx";
import { MsgCreateCommitment } from "./types/privacy/tx";
import { MsgUpdateCommitment } from "./types/privacy/tx";
import { MsgUpdateOnetimeAddress } from "./types/privacy/tx";
import { MsgUpdateToken } from "./types/privacy/tx";
import { MsgCreateTx } from "./types/privacy/tx";
import { MsgAirdrop } from "./types/privacy/tx";
import { MsgCreateToken } from "./types/privacy/tx";
import { MsgCreateSerialNumber } from "./types/privacy/tx";
import { MsgUpdateOutputCoin } from "./types/privacy/tx";
import { MsgCreateTxPrivacyData } from "./types/privacy/tx";
import { MsgCreateCommitmentIndex } from "./types/privacy/tx";
import { MsgUpdateTxPrivacyData } from "./types/privacy/tx";
import { MsgDeleteToken } from "./types/privacy/tx";
import { MsgDeleteOnetimeAddress } from "./types/privacy/tx";
import { MsgDeleteOutputCoinSerialNumber } from "./types/privacy/tx";
import { MsgDeleteCommitment } from "./types/privacy/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/privacy.privacy.MsgTransfer", MsgTransfer],
    ["/privacy.privacy.MsgDeleteSerialNumber", MsgDeleteSerialNumber],
    ["/privacy.privacy.MsgDeleteCommitmentIndex", MsgDeleteCommitmentIndex],
    ["/privacy.privacy.MsgCreateOutputCoinSerialNumber", MsgCreateOutputCoinSerialNumber],
    ["/privacy.privacy.MsgDeleteOutputCoin", MsgDeleteOutputCoin],
    ["/privacy.privacy.MsgCreateOnetimeAddress", MsgCreateOnetimeAddress],
    ["/privacy.privacy.MsgCreateOutputCoin", MsgCreateOutputCoin],
    ["/privacy.privacy.MsgUpdateOutputCoinSerialNumber", MsgUpdateOutputCoinSerialNumber],
    ["/privacy.privacy.MsgDeleteTxPrivacyData", MsgDeleteTxPrivacyData],
    ["/privacy.privacy.MsgPrivacyData", MsgPrivacyData],
    ["/privacy.privacy.MsgUpdateCommitmentIndex", MsgUpdateCommitmentIndex],
    ["/privacy.privacy.MsgUpdateSerialNumber", MsgUpdateSerialNumber],
    ["/privacy.privacy.MsgCreateCommitment", MsgCreateCommitment],
    ["/privacy.privacy.MsgUpdateCommitment", MsgUpdateCommitment],
    ["/privacy.privacy.MsgUpdateOnetimeAddress", MsgUpdateOnetimeAddress],
    ["/privacy.privacy.MsgUpdateToken", MsgUpdateToken],
    ["/privacy.privacy.MsgCreateTx", MsgCreateTx],
    ["/privacy.privacy.MsgAirdrop", MsgAirdrop],
    ["/privacy.privacy.MsgCreateToken", MsgCreateToken],
    ["/privacy.privacy.MsgCreateSerialNumber", MsgCreateSerialNumber],
    ["/privacy.privacy.MsgUpdateOutputCoin", MsgUpdateOutputCoin],
    ["/privacy.privacy.MsgCreateTxPrivacyData", MsgCreateTxPrivacyData],
    ["/privacy.privacy.MsgCreateCommitmentIndex", MsgCreateCommitmentIndex],
    ["/privacy.privacy.MsgUpdateTxPrivacyData", MsgUpdateTxPrivacyData],
    ["/privacy.privacy.MsgDeleteToken", MsgDeleteToken],
    ["/privacy.privacy.MsgDeleteOnetimeAddress", MsgDeleteOnetimeAddress],
    ["/privacy.privacy.MsgDeleteOutputCoinSerialNumber", MsgDeleteOutputCoinSerialNumber],
    ["/privacy.privacy.MsgDeleteCommitment", MsgDeleteCommitment],
    
];

export { msgTypes }