<div>
    <h1 class="title is-1">
	Welcome {username}!
    </h1>
    <div class="columns">
	<div class="column">
	    <h3 class="title is-3">Artists to Listen To</h3>
	    <!-- Stuff -->
	</div>
	<div class="column">
	</div>
	<div class="column">
	    <h3 class="title is-3">Artists to Review</h3>
	</div>
    </div>
</div>

<script>
let username = "";
let listenTo = [];

let loadingListenTo = true;

import { onMount } from 'svelte';

onMount( async () => {
    const resp = await fetch("http://localhost:5000/api/getUserInfo", 
	{ credentials: "any" } 
    );
    if(!resp.ok) {
	// Do something
    }
    const spotifyResp = await resp.json();
    username = spotifyResp.display_name;

    // Get artists following
    resp = await fetch("http://localhost:5000/api/getListenTo", {
	credentials: "any"
    });
    listenTo = await resp.json();
    console.log("ListenTo:");
    console.log(listenTo);
    for(const [index, artist] of listenTo.entries()) {
	for(const image of artist.Images) {
	    if(image.Width === 160) {
		listenTo[index].image_url = image.URL;
	    }
	}
    }
    loadingListenTo = false;
});
</script>
