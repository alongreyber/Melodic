<template>
    <button class="button is-primary" @click="login()">
	Log in with Spotify
    </button>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class LoginWithSpotify extends Vue {
    login() {
	const redirect_uri = "http://localhost:8080/spotify_callback"
	const client_id = "e9eb61d2a082412caa493f4c9ef86774"

	var state = Math.random().toString(36).substring(2, 15);
	// Save state in a cookie
	document.cookie = "spotifyState=" + state;
	var scopes = 'user-read-private user-read-email';
	// Redirect
	window.location.href = 'https://accounts.spotify.com/authorize' +
	    '?response_type=code' +
	    '&client_id=' + client_id +
	    (scopes ? '&scope=' + encodeURIComponent(scopes) : '') +
	    '&state=' + state +
	    '&redirect_uri=' + encodeURIComponent(redirect_uri);
    }
}
</script>
