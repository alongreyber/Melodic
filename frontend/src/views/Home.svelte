<div>
<h1 class="title is-1">Home</h1>
    <LoginWithSpotify></LoginWithSpotify>
</div>

<script>
import LoginWithSpotify from '../components/LoginWithSpotify.svelte';

import { onMount } from 'svelte';
import { push, pop, replace } from 'svelte-spa-router';

function getCookie(cname) {
  const name = cname + "=";
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(";");
  for(let c of ca) {
    while (c.charAt(0) === " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) === 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

onMount( async () => {
    // Need to handle the spotify callback 
    const spotifyParams = (new URL(location.href)).searchParams;
    // If "code" exists in query string assume this was a callback
    if( spotifyParams.has("code") ) {
	if(!( spotifyParams.has("state") ) ||
	    spotifyParams.get("state") !== getCookie("spotifyState")) {
		console.log("Invalid State");
	}
	// Forward to backend
	var url = new URL("http://localhost:5000/api/login")
	for(const [key, value] of spotifyParams) {
	    url.searchParams.append(key, value);
	}
	const resp = await fetch(url, {
	    credentials: "include"
	});

	if(resp.ok) {
	    // Remove the parameters from the URL using the history API
	    // This is definitely a hack but this is only for one page of the site so 
	    // I'm not too worried about it
	    history.replaceState && history.replaceState( null, '', location.pathname + location.search.replace(/[\?&]code=[^&]+/, '').replace(/^&/, '?')
	    );
	    history.replaceState && history.replaceState( null, '', location.pathname + location.search.replace(/[\?&]state=[^&]+/, '').replace(/^&/, '?')
	    );
	    push("/logged_in");
	}
    } else if("error" in spotifyParams) {
	if(spotifyParams.error === "access_denied") {
	    console.log("Please allow Spotify access to use this app");
	}
    }
})

</script>
