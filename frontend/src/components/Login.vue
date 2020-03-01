<template>
<button class="button is-primary" @click="login">
    Log in with Spotify
</button>
</template>

<script>

import { postApiJson, getApi } from '../api';

export default {
    name: 'Login',
    data: function() {
	return {
	    redirectURI: "",
	    state: ""
	}
    },
    beforeCreate: async function() {
	let data = await getApi('/getCallbackURL');
	this.state = data.state;
	this.redirectURI = data.url;
    },
    methods: {
	login: function() {
	    // Save state in a cookie
	    document.cookie = "spotifyState=" + this.state;
	    // Redirect to URI
	    window.location.href = this.redirectURI;
	}
    }

}

</script>
