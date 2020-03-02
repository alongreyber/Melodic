<template>

<div class="has-text-centered">
    <!-- this is in the html because we don't want it to apply to every page -->
    <br/>
    <br/>
    <br/>
    <br/>
    <div class="title is-1 has-text-white">
	Welcome to Melodic
    </div>
    <Login v-bind:is_loading="is_loading" ></Login>
    <div class="title is-5 has-text-white" v-if="is_loading">
	Loading your followed artists...
    </div>
</div>
</template>

<script>

import { getCookie, getApi } from '../api';
import Login from '../components/Login';

export default {
  name: 'Home',
  data: function() {
      return {
	  is_loading: false
      }
  },
  components: { Login },
  beforeMount: async function() {
    const spotifyParams = (new URL(location.href)).searchParams;
    // If "code" exists in query string assume this was a callback
    if( spotifyParams.has("code") ) {
	this.is_loading = true;
	if(!( spotifyParams.has("state") ) ||
	    spotifyParams.get("state") !== getCookie("spotifyState")) {
	    this.$store.commit('displayMessage', {'color' : 'is-danger', 'text' : 'Invalid State'})
	}
	// Send login request to backend
	var path = "/login?";
	for(const [key, value] of spotifyParams) {
	    path = path + `${key}=${value}&`;
	}
	const resp = await getApi(encodeURI(path));

	if(resp) {
	    // Remove the parameters from the URL using the history API
	    // This is definitely a hack but this is only for one page of the site so 
	    // I'm not too worried about it
	    history.replaceState && history.replaceState( null, '', location.pathname + location.search.replace(/[\?&]code=[^&]+/, '').replace(/^&/, '?')
	    );
	    history.replaceState && history.replaceState( null, '', location.pathname + location.search.replace(/[\?&]state=[^&]+/, '').replace(/^&/, '?')
	    );
	    let user = await getApi('/getUserInfo');
	    // Now we request to initialize this user
	    let result = await getApi('/initializeFollowing');

	    this.$store.commit('logIn', user);
	    this.$router.push('/dashboard');
	}
    } else if("error" in spotifyParams) {
	if(spotifyParams.error == "access_denied") {
	    this.$store.commit('displayMessage', {'color' : 'is-danger', 'text' : "Please allow Spotify access to use this app"});
	    // Don't go anywhere
	} else {
	    this.$store.commit('displayMessage', {'color' : 'is-danger', 'text' : "Unknown Spotify Error: " + spotifyParams.error});
	}
    }
  },
    created() {
	document.querySelector("body").classList.add("teal-gradient");
    },
    destroyed() {
	document.querySelector("body").classList.remove("teal-gradient");
    },
}
</script>

<style lang="scss">
body.teal-gradient {
    background: rgb(72,169,166);
    background: linear-gradient(150deg, rgba(72,169,166,1) 0%, rgba(66,129,164,1) 100%); 
    height: 100vh;
}
</style>
