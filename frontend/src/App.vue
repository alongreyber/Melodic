<template>
  <div id="app">
    <nav class="navbar is-primary" role="navigation">
      <div class="navbar-brand">
	  <div class="navbar-item">
	      <img src="logo-white.png" height="40" >
	      <span class="logo-text">
	      Melodic
	      </span>
	  </div>
      </div>
      <div class="navbar-start">
	<router-link class="navbar-item" to="/">Home</router-link>
	<router-link class="navbar-item" to="/about">About</router-link>
      </div>
      <div class="navbar-end">
	  <a class="navbar-item" href="#">Our Socials!</a>
      </div>
    </nav>
    <div class="container" id="main-container">
	<template v-for="(message, index) in store.state.messages">
	    <div class="notification" 
		v-bind:class="message.color">
		<button class="delete"
		    @click="store.removeMessageByIndex(index)">
		</button>
		{{message.text}}
	    </div>
	</template>
	<router-view/>
    </div>
  </div>
</template>

<script lang="ts">
import { store } from "./store.ts";
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class App extends Vue {
    // Attach global store so that it becomes reactive
    // This is used to display messages
    store = store
}
</script>

<style>
#main-container {
    margin-top: 15px;
}
.logo-text {
    margin-left: 5px;
    font-family: "Georgia", Serif;
    font-weight: bold;
    font-size: 1.5rem;
}
</style>
