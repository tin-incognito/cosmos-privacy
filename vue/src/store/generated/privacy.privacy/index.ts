import { Client, registry, MissingWalletError } from 'privacy-client-ts'

import { Commitment } from "privacy-client-ts/privacy.privacy/types"
import { CommitmentIndex } from "privacy-client-ts/privacy.privacy/types"
import { OnetimeAddress } from "privacy-client-ts/privacy.privacy/types"
import { OTACoin } from "privacy-client-ts/privacy.privacy/types"
import { OutputCoin } from "privacy-client-ts/privacy.privacy/types"
import { OutputCoinLength } from "privacy-client-ts/privacy.privacy/types"
import { OutputCoinSerialNumber } from "privacy-client-ts/privacy.privacy/types"
import { Params } from "privacy-client-ts/privacy.privacy/types"
import { SerialNumber } from "privacy-client-ts/privacy.privacy/types"
import { Token } from "privacy-client-ts/privacy.privacy/types"
import { MsgTransfer_PaymentInfo } from "privacy-client-ts/privacy.privacy/types"
import { TxPrivacyData } from "privacy-client-ts/privacy.privacy/types"


export { Commitment, CommitmentIndex, OnetimeAddress, OTACoin, OutputCoin, OutputCoinLength, OutputCoinSerialNumber, Params, SerialNumber, Token, MsgTransfer_PaymentInfo, TxPrivacyData };

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
				SerialNumber: {},
				SerialNumberAll: {},
				OutputCoin: {},
				OutputCoinAll: {},
				Commitment: {},
				CommitmentAll: {},
				CommitmentIndex: {},
				CommitmentIndexAll: {},
				Token: {},
				TokenAll: {},
				OnetimeAddress: {},
				OnetimeAddressAll: {},
				TxPrivacyData: {},
				TxPrivacyDataAll: {},
				Balance: {},
				OTACoin: {},
				OTACoinAll: {},
				OutputCoinSerialNumber: {},
				
				_Structure: {
						Commitment: getStructure(Commitment.fromPartial({})),
						CommitmentIndex: getStructure(CommitmentIndex.fromPartial({})),
						OnetimeAddress: getStructure(OnetimeAddress.fromPartial({})),
						OTACoin: getStructure(OTACoin.fromPartial({})),
						OutputCoin: getStructure(OutputCoin.fromPartial({})),
						OutputCoinLength: getStructure(OutputCoinLength.fromPartial({})),
						OutputCoinSerialNumber: getStructure(OutputCoinSerialNumber.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						SerialNumber: getStructure(SerialNumber.fromPartial({})),
						Token: getStructure(Token.fromPartial({})),
						MsgTransfer_PaymentInfo: getStructure(MsgTransfer_PaymentInfo.fromPartial({})),
						TxPrivacyData: getStructure(TxPrivacyData.fromPartial({})),
						
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
				getSerialNumber: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SerialNumber[JSON.stringify(params)] ?? {}
		},
				getSerialNumberAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SerialNumberAll[JSON.stringify(params)] ?? {}
		},
				getOutputCoin: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutputCoin[JSON.stringify(params)] ?? {}
		},
				getOutputCoinAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutputCoinAll[JSON.stringify(params)] ?? {}
		},
				getCommitment: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Commitment[JSON.stringify(params)] ?? {}
		},
				getCommitmentAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CommitmentAll[JSON.stringify(params)] ?? {}
		},
				getCommitmentIndex: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CommitmentIndex[JSON.stringify(params)] ?? {}
		},
				getCommitmentIndexAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CommitmentIndexAll[JSON.stringify(params)] ?? {}
		},
				getToken: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Token[JSON.stringify(params)] ?? {}
		},
				getTokenAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenAll[JSON.stringify(params)] ?? {}
		},
				getOnetimeAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OnetimeAddress[JSON.stringify(params)] ?? {}
		},
				getOnetimeAddressAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OnetimeAddressAll[JSON.stringify(params)] ?? {}
		},
				getTxPrivacyData: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxPrivacyData[JSON.stringify(params)] ?? {}
		},
				getTxPrivacyDataAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxPrivacyDataAll[JSON.stringify(params)] ?? {}
		},
				getBalance: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Balance[JSON.stringify(params)] ?? {}
		},
				getOTACoin: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OTACoin[JSON.stringify(params)] ?? {}
		},
				getOTACoinAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OTACoinAll[JSON.stringify(params)] ?? {}
		},
				getOutputCoinSerialNumber: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutputCoinSerialNumber[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: privacy.privacy initialized!')
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
				let value= (await client.PrivacyPrivacy.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySerialNumber({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.querySerialNumber( key.index)).data
				
					
				commit('QUERY', { query: 'SerialNumber', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySerialNumber', payload: { options: { all }, params: {...key},query }})
				return getters['getSerialNumber']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySerialNumber API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySerialNumberAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.querySerialNumberAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.querySerialNumberAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SerialNumberAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySerialNumberAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSerialNumberAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySerialNumberAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutputCoin({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOutputCoin( key.index)).data
				
					
				commit('QUERY', { query: 'OutputCoin', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutputCoin', payload: { options: { all }, params: {...key},query }})
				return getters['getOutputCoin']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutputCoin API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutputCoinAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOutputCoinAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryOutputCoinAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OutputCoinAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutputCoinAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOutputCoinAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutputCoinAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCommitment({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryCommitment( key.index)).data
				
					
				commit('QUERY', { query: 'Commitment', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCommitment', payload: { options: { all }, params: {...key},query }})
				return getters['getCommitment']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCommitment API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCommitmentAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryCommitmentAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryCommitmentAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CommitmentAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCommitmentAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCommitmentAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCommitmentAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCommitmentIndex({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryCommitmentIndex( key.index)).data
				
					
				commit('QUERY', { query: 'CommitmentIndex', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCommitmentIndex', payload: { options: { all }, params: {...key},query }})
				return getters['getCommitmentIndex']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCommitmentIndex API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCommitmentIndexAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryCommitmentIndexAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryCommitmentIndexAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CommitmentIndexAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCommitmentIndexAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCommitmentIndexAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCommitmentIndexAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryToken({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryToken( key.index)).data
				
					
				commit('QUERY', { query: 'Token', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryToken', payload: { options: { all }, params: {...key},query }})
				return getters['getToken']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryToken API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryTokenAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryTokenAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokenAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOnetimeAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOnetimeAddress( key.index)).data
				
					
				commit('QUERY', { query: 'OnetimeAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOnetimeAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getOnetimeAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOnetimeAddress API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOnetimeAddressAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOnetimeAddressAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryOnetimeAddressAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OnetimeAddressAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOnetimeAddressAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOnetimeAddressAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOnetimeAddressAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTxPrivacyData({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryTxPrivacyData( key.index)).data
				
					
				commit('QUERY', { query: 'TxPrivacyData', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTxPrivacyData', payload: { options: { all }, params: {...key},query }})
				return getters['getTxPrivacyData']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTxPrivacyData API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTxPrivacyDataAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryTxPrivacyDataAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryTxPrivacyDataAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TxPrivacyDataAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTxPrivacyDataAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTxPrivacyDataAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTxPrivacyDataAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBalance({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryBalance( key.privateKey)).data
				
					
				commit('QUERY', { query: 'Balance', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBalance', payload: { options: { all }, params: {...key},query }})
				return getters['getBalance']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBalance API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOTACoin({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOTACoin( key.index)).data
				
					
				commit('QUERY', { query: 'OTACoin', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOTACoin', payload: { options: { all }, params: {...key},query }})
				return getters['getOTACoin']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOTACoin API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOTACoinAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOTACoinAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.PrivacyPrivacy.query.queryOTACoinAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OTACoinAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOTACoinAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOTACoinAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOTACoinAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutputCoinSerialNumber({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.PrivacyPrivacy.query.queryOutputCoinSerialNumber()).data
				
					
				commit('QUERY', { query: 'OutputCoinSerialNumber', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutputCoinSerialNumber', payload: { options: { all }, params: {...key},query }})
				return getters['getOutputCoinSerialNumber']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutputCoinSerialNumber API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDeleteCommitment({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteCommitment({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteCommitment:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteCommitment:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPrivacyData({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgPrivacyData({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPrivacyData:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPrivacyData:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateOutputCoin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateOutputCoin({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutputCoin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateOutputCoin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateOutputCoin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateOutputCoin({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOutputCoin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOutputCoin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateOutputCoinSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateOutputCoinSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateOutputCoinSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateCommitmentIndex({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateCommitmentIndex({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCommitmentIndex:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteOutputCoinSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteOutputCoinSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteOutputCoinSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateCommitmentIndex({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateCommitmentIndex({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateCommitmentIndex:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateOutputCoinSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateOutputCoinSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOutputCoinSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateCommitment({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateCommitment({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateCommitment:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateCommitment:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateCommitment({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateCommitment({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCommitment:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCommitment:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgTransfer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgTransfer({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTransfer:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgTransfer:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteOutputCoin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteOutputCoin({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOutputCoin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteOutputCoin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteOnetimeAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteOnetimeAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteOnetimeAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteCommitmentIndex({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteCommitmentIndex({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteCommitmentIndex:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateOnetimeAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateOnetimeAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOnetimeAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteTxPrivacyData({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteTxPrivacyData({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteTxPrivacyData:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateTx({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateTx({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTx:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateTx:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgDeleteToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateOnetimeAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateOnetimeAddress({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateOnetimeAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateSerialNumber({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateSerialNumber({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSerialNumber:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateSerialNumber:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAirdrop({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgAirdrop({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAirdrop:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAirdrop:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateTxPrivacyData({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgCreateTxPrivacyData({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateTxPrivacyData:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateTxPrivacyData({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.PrivacyPrivacy.tx.sendMsgUpdateTxPrivacyData({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateTxPrivacyData:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDeleteCommitment({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteCommitment({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteCommitment:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteCommitment:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPrivacyData({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgPrivacyData({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPrivacyData:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPrivacyData:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateOutputCoin({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateOutputCoin({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutputCoin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateOutputCoin:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateOutputCoin({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateOutputCoin({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOutputCoin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOutputCoin:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateOutputCoinSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateOutputCoinSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateOutputCoinSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateCommitmentIndex({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateCommitmentIndex({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCommitmentIndex:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteOutputCoinSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteOutputCoinSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteOutputCoinSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateCommitmentIndex({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateCommitmentIndex({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateCommitmentIndex:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateOutputCoinSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateOutputCoinSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOutputCoinSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOutputCoinSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateCommitment({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateCommitment({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateCommitment:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateCommitment:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateCommitment({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateCommitment({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCommitment:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCommitment:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgTransfer({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgTransfer({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTransfer:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgTransfer:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteOutputCoin({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteOutputCoin({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOutputCoin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteOutputCoin:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteOnetimeAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteOnetimeAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteOnetimeAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteCommitmentIndex({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteCommitmentIndex({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteCommitmentIndex:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteCommitmentIndex:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateOnetimeAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateOnetimeAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOnetimeAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteTxPrivacyData({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteTxPrivacyData({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteTxPrivacyData:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateTx({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateTx({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTx:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateTx:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgDeleteToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateOnetimeAddress({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateOnetimeAddress({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOnetimeAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateOnetimeAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateSerialNumber({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateSerialNumber({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSerialNumber:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateSerialNumber:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAirdrop({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgAirdrop({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAirdrop:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAirdrop:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateTxPrivacyData({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgCreateTxPrivacyData({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateTxPrivacyData:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateTxPrivacyData({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.PrivacyPrivacy.tx.msgUpdateTxPrivacyData({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateTxPrivacyData:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateTxPrivacyData:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
