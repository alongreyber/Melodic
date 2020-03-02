<template>
    <div>
	Hello! Welcome to Melodic
	<Login></Login>
    </div>
</template>

<script>

import { getCookie, getApi } from '../api';
import Login from '../components/Login';

export default {
  name: 'Home',
  components: { Login },
  beforeMount: async function() {
    const spotifyParams = (new URL(location.href)).searchParams;
    // If "code" exists in query string assume this was a callback
    if( spotifyParams.has("code") ) {
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
  }
}
</script>
