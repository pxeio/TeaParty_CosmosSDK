import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBuy } from "./types/partychain/party/tx";
import { MsgCancel } from "./types/partychain/party/tx";
import { MsgAccountWatchFailure } from "./types/partychain/party/tx";
import { MsgAccountWatchOutcome } from "./types/partychain/party/tx";
import { MsgSubmitSell } from "./types/partychain/party/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/teapartycrypto.partychain.party.MsgBuy", MsgBuy],
    ["/teapartycrypto.partychain.party.MsgCancel", MsgCancel],
    ["/teapartycrypto.partychain.party.MsgAccountWatchFailure", MsgAccountWatchFailure],
    ["/teapartycrypto.partychain.party.MsgAccountWatchOutcome", MsgAccountWatchOutcome],
    ["/teapartycrypto.partychain.party.MsgSubmitSell", MsgSubmitSell],
    
];

export { msgTypes }