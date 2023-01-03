// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgAccountWatchOutcome } from "./types/partychain/party/tx";
import { MsgSubmitSell } from "./types/partychain/party/tx";
import { MsgAccountWatchFailure } from "./types/partychain/party/tx";
import { MsgBuy } from "./types/partychain/party/tx";
import { MsgCancel } from "./types/partychain/party/tx";
import { MsgTransactionResult } from "./types/partychain/party/tx";


export { MsgAccountWatchOutcome, MsgSubmitSell, MsgAccountWatchFailure, MsgBuy, MsgCancel, MsgTransactionResult };

type sendMsgAccountWatchOutcomeParams = {
  value: MsgAccountWatchOutcome,
  fee?: StdFee,
  memo?: string
};

type sendMsgSubmitSellParams = {
  value: MsgSubmitSell,
  fee?: StdFee,
  memo?: string
};

type sendMsgAccountWatchFailureParams = {
  value: MsgAccountWatchFailure,
  fee?: StdFee,
  memo?: string
};

type sendMsgBuyParams = {
  value: MsgBuy,
  fee?: StdFee,
  memo?: string
};

type sendMsgCancelParams = {
  value: MsgCancel,
  fee?: StdFee,
  memo?: string
};

type sendMsgTransactionResultParams = {
  value: MsgTransactionResult,
  fee?: StdFee,
  memo?: string
};


type msgAccountWatchOutcomeParams = {
  value: MsgAccountWatchOutcome,
};

type msgSubmitSellParams = {
  value: MsgSubmitSell,
};

type msgAccountWatchFailureParams = {
  value: MsgAccountWatchFailure,
};

type msgBuyParams = {
  value: MsgBuy,
};

type msgCancelParams = {
  value: MsgCancel,
};

type msgTransactionResultParams = {
  value: MsgTransactionResult,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgAccountWatchOutcome({ value, fee, memo }: sendMsgAccountWatchOutcomeParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAccountWatchOutcome: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAccountWatchOutcome({ value: MsgAccountWatchOutcome.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAccountWatchOutcome: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSubmitSell({ value, fee, memo }: sendMsgSubmitSellParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSubmitSell: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSubmitSell({ value: MsgSubmitSell.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSubmitSell: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAccountWatchFailure({ value, fee, memo }: sendMsgAccountWatchFailureParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAccountWatchFailure: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAccountWatchFailure({ value: MsgAccountWatchFailure.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAccountWatchFailure: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBuy({ value, fee, memo }: sendMsgBuyParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBuy: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBuy({ value: MsgBuy.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBuy: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCancel({ value, fee, memo }: sendMsgCancelParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCancel: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCancel({ value: MsgCancel.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCancel: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgTransactionResult({ value, fee, memo }: sendMsgTransactionResultParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgTransactionResult: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgTransactionResult({ value: MsgTransactionResult.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgTransactionResult: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgAccountWatchOutcome({ value }: msgAccountWatchOutcomeParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgAccountWatchOutcome", value: MsgAccountWatchOutcome.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAccountWatchOutcome: Could not create message: ' + e.message)
			}
		},
		
		msgSubmitSell({ value }: msgSubmitSellParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgSubmitSell", value: MsgSubmitSell.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSubmitSell: Could not create message: ' + e.message)
			}
		},
		
		msgAccountWatchFailure({ value }: msgAccountWatchFailureParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgAccountWatchFailure", value: MsgAccountWatchFailure.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAccountWatchFailure: Could not create message: ' + e.message)
			}
		},
		
		msgBuy({ value }: msgBuyParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgBuy", value: MsgBuy.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBuy: Could not create message: ' + e.message)
			}
		},
		
		msgCancel({ value }: msgCancelParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgCancel", value: MsgCancel.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCancel: Could not create message: ' + e.message)
			}
		},
		
		msgTransactionResult({ value }: msgTransactionResultParams): EncodeObject {
			try {
				return { typeUrl: "/teapartycrypto.partychain.party.MsgTransactionResult", value: MsgTransactionResult.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgTransactionResult: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			TeapartycryptoPartychainParty: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;