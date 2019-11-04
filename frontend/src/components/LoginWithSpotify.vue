<template>
    <button class="button is-primary" @click="login()">
	Log in with Spotify
    </button>
</template>

<script lang="ts">
const axios = require('axios');
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class LoginWithSpotify extends Vue {
    redirectURI = ""
    state = ""

    mounted() {
	let component = this;
	axios.get('http://localhost:5000/api/getCallbackURL').then(function(response: any) {
	    component.redirectURI = response.data.data.url;
	    component.state = response.data.data.state;
	}).catch(function(error: any) {
	    // Dunno what to do here
	});
    }
    login() {
	// Save state in a cookie
	document.cookie = "spotifyState=" + this.state;
	window.location.href = this.redirectURI;
    }
}
</script>
