
<div>Callback</div>

<script>

import { onMount } from 'svelte';
import queryString from 'query-string';
import {push, pop, replace} from 'svelte-spa-router';

onMount( async () => {
    const spotifyParams = queryString.parse(location.search);
    if("error" in spotifyParams) {
	if(spotifyParams.error === "access_denied") {
	    console.log("Please allow Spotify access to use this app");
	}
    }
    if(!("state" in spotifyParams) ||
	spotifyParams.state !== getCookie("spotifyState")) {
	    console.log("Invalid State");
    }

    const resp = await fetch("http://localhost:5000/api/login", {
	credentials: "include"
    });

    if(resp.ok) {
	push("/logged_in");
    }

});
</script>
