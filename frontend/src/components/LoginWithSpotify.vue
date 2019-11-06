<template>
    <button class="button is-primary" @click="login()">
	Log in with Spotify
    </button>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import axios from "axios";

@Component
export default class LoginWithSpotify extends Vue {
    public redirectURI = "";
    public state = "";

    public mounted() {
	const component = this;
	axios.get("http://localhost:5000/api/getCallbackURL").then((response: any) => {
	    component.redirectURI = response.data.data.url;
	    component.state = response.data.data.state;
	}).catch((error: any) => {
	    // Dunno what to do here
	});
    }
    public login() {
	// Save state in a cookie
	document.cookie = "spotifyState=" + this.state;
	window.location.href = this.redirectURI;
    }
}
</script>
