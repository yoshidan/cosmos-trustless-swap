import { Client, registry, MissingWalletError } from 'swap-client-ts'

import { Params } from "swap-client-ts/swap.swap/types"
import { Swap } from "swap-client-ts/swap.swap/types"
import { NFTSwap } from "swap-client-ts/swap.swap/types"


export { Params, Swap, NFTSwap };

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
				Show: {},
				ShowNFT: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						Swap: getStructure(Swap.fromPartial({})),
						NFTSwap: getStructure(NFTSwap.fromPartial({})),
						
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
				getShow: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Show[JSON.stringify(params)] ?? {}
		},
				getShowNFT: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ShowNFT[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: swap.swap initialized!')
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
				let value= (await client.SwapSwap.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryShow({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SwapSwap.query.queryShow( key.id)).data
				
					
				commit('QUERY', { query: 'Show', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryShow', payload: { options: { all }, params: {...key},query }})
				return getters['getShow']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryShow API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryShowNFT({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SwapSwap.query.queryShowNFT( key.id)).data
				
					
				commit('QUERY', { query: 'ShowNFT', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryShowNFT', payload: { options: { all }, params: {...key},query }})
				return getters['getShowNFT']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryShowNFT API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCancelNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgCancelNFT({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgCancel({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSend({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgSend({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSend:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSend:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgReceive({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgReceive({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReceive:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgReceive:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgReceiveNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgReceiveNFT({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReceiveNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgReceiveNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendNFT({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SwapSwap.tx.sendMsgSendNFT({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendNFT:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendNFT:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCancelNFT({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgCancelNFT({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgCancel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancel:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSend({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgSend({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSend:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSend:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgReceive({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgReceive({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReceive:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgReceive:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgReceiveNFT({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgReceiveNFT({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgReceiveNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgReceiveNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendNFT({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SwapSwap.tx.msgSendNFT({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendNFT:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendNFT:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
