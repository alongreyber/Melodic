<template>
  <div id="app">
      <nav v-if="loggedIn" class="navbar is-primary">
	  <div class="navbar-brand">
	      <div class="navbar-item">
		  <img src="./assets/logo-white.png" alt="Melodic Logo" height="40" >
		  <span class="logo-text">
		  Melodic
		  </span>
	      </div>
	      <a class="navbar-burger" @click="show_menu = !show_menu">
		  <span></span><span></span><span></span>
              </a>
	  </div>

	  <div class="navbar-menu" :class="{'is-active' : show_menu}">
	      <div class="navbar-start">
		<a class="navbar-item">Home</a>
	      </div>

	      <div class="navbar-end">
		  <div class="navbar-item">
		      Welcome {{ $store.state.user.display_name }}
		  </div>
		  <div class="navbar-item">
		      <div class="buttons">
			  <button class="button is-light" v-on:click="logout">
			      Log Out
			  </button>
		      </div>
		  </div>
	      </div>
	  </div>
      </nav>
      <br />
      <div class="container" id="main-container">
	  <div class="modal" :class="{'is-active' : $store.state.modal }">
	      <div class="modal-background"></div>
	      <div class="modal-card">
		  <header class="modal-card-head">
		      <p class="modal-card-title">{{ $store.state.modal ? $store.state.modal.title : '' }}</p>
		      <button class="delete" @click="$store.commit('clearModal')"></button>
		  </header>
		  <section class="modal-card-body">
		      {{ $store.state.modal ? $store.state.modal.text : '' }}
		  </section>
	      </div>
	  </div>
	  <div v-for="n in $store.state.messages" class="notification space-above" :class="n.color">
	      <button class="delete" @click="$store.commit('clearMessage', n)"></button>
	      {{ n.text }}
	  </div>
	
	<router-view/>
      </div>
  </div>
</template>

<script>
import { getApi } from './api';

export default {
    name: 'App',
    data: function() {
	return {
	    // Only used on mobile
	    show_menu : false
	}
    },
    computed: {
	loggedIn: function() {
	    return this.$store.state.loggedIn;
	}
    },
    methods: {
	logout: async function() {
	    let result = await getApi('/logout');
	    if(result) {
		this.$store.commit('logOut');
		this.$router.push('/');
	    }
	}
    }
}
</script>

<style lang="scss">
.logo-text {
    margin-left: 5px;
    font-family: "Georgia", Serif;
    font-weight: bold;
    font-size: 1.5rem;
}

</style>
