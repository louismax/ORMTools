import {
	createStore
} from 'vuex';



export const store = createStore({
	state: {
		userConfig: {}
	},
	mutations: {
		setUserConfig(state, cfg) {
			state.userConfig = cfg
		}
	},
	actions: {},
	modules: {}
});
