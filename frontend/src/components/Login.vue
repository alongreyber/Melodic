<template>
    <button class="button is-primary is-large" :class="{'is-loading' : is_loading}" @click="login">
    Log in with Spotify
</button>
</template>

<script>

import { postApiJson, getApi } from '../api';

export default {
    name: 'Login',
    props: ['is_loading'],
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
