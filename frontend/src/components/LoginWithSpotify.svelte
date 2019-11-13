<button class="button is-primary" on:click={login}>
    Log in with Spotify
</button>

<script>
import {push, pop, replace} from 'svelte-spa-router'
import { onMount } from 'svelte';

let redirectURI = "";

let state = "";

onMount( async () => {
    const resp = await fetch("http://localhost:5000/api/getCallbackURL", 
	{ credentials: 'include' }
    );
    if(resp.ok) {
	const data = await resp.json();
	redirectURI = data.data.url;
	state = data.data.state;
    }
});

function login() {
    // Save state in a cookie
    document.cookie = "spotifyState=" + state;
    // Redirect to URI
    window.location.href = redirectURI;
}

</script>
