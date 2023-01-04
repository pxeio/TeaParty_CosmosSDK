import { Client, registry, MissingWalletError } from 'TeaPartyCrypto-partychain-client-ts'

import { OrdersAwaitingFinalizer } from "TeaPartyCrypto-partychain-client-ts/teapartycrypto.partychain.party/types"
import { Params } from "TeaPartyCrypto-partychain-client-ts/teapartycrypto.partychain.party/types"
import { PendingOrders } from "TeaPartyCrypto-partychain-client-ts/teapartycrypto.partychain.party/types"
import { TradeOrders } from "TeaPartyCrypto-partychain-client-ts/teapartycrypto.partychain.party/types"


export { OrdersAwaitingFinalizer, Params, PendingOrders, TradeOrders };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				TradeOrders: {},
				TradeOrdersAll: {},
				PendingOrders: {},
				PendingOrdersAll: {},
				OrdersAwaitingFinalizer: {},
				OrdersAwaitingFinalizerAll: {},
				
				_Structure: {
						OrdersAwaitingFinalizer: getStructure(OrdersAwaitingFinalizer.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PendingOrders: getStructure(PendingOrders.fromPartial({})),
						TradeOrders: getStructure(TradeOrders.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getTradeOrders: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TradeOrders[JSON.stringify(params)] ?? {}
		},
				getTradeOrdersAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TradeOrdersAll[JSON.stringify(params)] ?? {}
		},
				getPendingOrders: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PendingOrders[JSON.stringify(params)] ?? {}
		},
				getPendingOrdersAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PendingOrdersAll[JSON.stringify(params)] ?? {}
		},
				getOrdersAwaitingFinalizer: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OrdersAwaitingFinalizer[JSON.stringify(params)] ?? {}
		},
				getOrdersAwaitingFinalizerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OrdersAwaitingFinalizerAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: teapartycrypto.partychain.party initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTradeOrders({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryTradeOrders( key.index)).data
				
					
				commit('QUERY', { query: 'TradeOrders', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTradeOrders', payload: { options: { all }, params: {...key},query }})
				return getters['getTradeOrders']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTradeOrders API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTradeOrdersAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryTradeOrdersAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.TeapartycryptoPartychainParty.query.queryTradeOrdersAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TradeOrdersAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTradeOrdersAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTradeOrdersAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTradeOrdersAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPendingOrders({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryPendingOrders( key.index)).data
				
					
				commit('QUERY', { query: 'PendingOrders', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPendingOrders', payload: { options: { all }, params: {...key},query }})
				return getters['getPendingOrders']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPendingOrders API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPendingOrdersAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryPendingOrdersAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.TeapartycryptoPartychainParty.query.queryPendingOrdersAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PendingOrdersAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPendingOrdersAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPendingOrdersAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPendingOrdersAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOrdersAwaitingFinalizer({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryOrdersAwaitingFinalizer( key.index)).data
				
					
				commit('QUERY', { query: 'OrdersAwaitingFinalizer', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOrdersAwaitingFinalizer', payload: { options: { all }, params: {...key},query }})
				return getters['getOrdersAwaitingFinalizer']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOrdersAwaitingFinalizer API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOrdersAwaitingFinalizerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.TeapartycryptoPartychainParty.query.queryOrdersAwaitingFinalizerAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.TeapartycryptoPartychainParty.query.queryOrdersAwaitingFinalizerAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OrdersAwaitingFinalizerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOrdersAwaitingFinalizerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOrdersAwaitingFinalizerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOrdersAwaitingFinalizerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSubmitSell({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.TeapartycryptoPartychainParty.tx.sendMsgSubmitSell({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitSell:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitSell:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBuy({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.TeapartycryptoPartychainParty.tx.sendMsgBuy({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuy:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBuy:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.TeapartycryptoPartychainParty.tx.sendMsgCancel({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAccountWatchFailure({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.TeapartycryptoPartychainParty.tx.sendMsgAccountWatchFailure({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAccountWatchFailure:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAccountWatchFailure:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAccountWatchOutcome({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.TeapartycryptoPartychainParty.tx.sendMsgAccountWatchOutcome({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAccountWatchOutcome:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAccountWatchOutcome:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSubmitSell({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.TeapartycryptoPartychainParty.tx.msgSubmitSell({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitSell:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitSell:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBuy({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.TeapartycryptoPartychainParty.tx.msgBuy({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuy:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBuy:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.TeapartycryptoPartychainParty.tx.msgCancel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancel:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAccountWatchFailure({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.TeapartycryptoPartychainParty.tx.msgAccountWatchFailure({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAccountWatchFailure:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAccountWatchFailure:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAccountWatchOutcome({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.TeapartycryptoPartychainParty.tx.msgAccountWatchOutcome({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAccountWatchOutcome:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAccountWatchOutcome:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
