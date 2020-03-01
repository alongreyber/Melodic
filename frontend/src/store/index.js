import Vue from 'vue'
import Vuex from "vuex"

import { getApi } from '../api';

// Add Vuex
Vue.use(Vuex);

const store = new Vuex.Store(
    {
        state: {
            loggedIn: false,
	    user: {},
	    messages: [],
	    modal: null,
        },
        mutations: {
            logIn(state, user) {
                state.loggedIn = true;
		state.user = user;
            },
            logOut(state) {
                state.loggedIn = false;
		state.user = {};
            },
	    displayMessage(state, msg) {
		state.messages.push(msg);
	    },
	    clearMessage(state, msg) {
		var index = state.messages.indexOf(msg);
		if (index !== -1) state.messages.splice(index, 1);
	    },
	    clearAllMessages(state) {
		state.messages.splice(0, state.messages.length);
	    },
	    displayModal(state, msg) {
		state.modal = msg;
	    },
	    clearModal(state) {
		state.modal = null; 
	    },
        },
	getters: {
	},
	actions: {
	},
    }
);

export default store

