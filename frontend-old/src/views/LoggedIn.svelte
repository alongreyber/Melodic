<div>
    <h1 class="title is-1">
	Welcome {username}!
    </h1>
    <div class="columns">
	<div class="column">
	    <h3 class="title is-3">Artists to Listen To</h3>
	    <ArtistList on:sendToReview="{sendToReview}" artists={listenTo} loading={loadingListenTo} ></ArtistList>
	</div>
	<div class="column">
	</div>
	<div class="column">
	    <h3 class="title is-3">Artists to Review</h3>
	    <ArtistList artists={toReview} loading={loadingToReview} ></ArtistList>
	</div>
    </div>
</div>

<script>
let username = "";
let listenTo = [];
let toReview = [];

let loadingListenTo = true;
let loadingToReview = true;

import { onMount } from 'svelte';
import ArtistList from '../components/ArtistList.svelte';
import * as backend from "../utils/backend.js";

async function sendToReview( event ) {
    // Remove from listenTo
    const artist = event.detail;
    listenTo = listenTo.filter(obj => obj.ID !== artist.ID)
    // Add to start of toReview
    toReview.unshift(artist);
    toReview = toReview;
    // Update in DB

    const data = backend.request("http://localhost:5000/api/moveToReview/" + artist.ID);
}

onMount( async () => {
    const userData = await backend.request("http://localhost:5000/api/getUserInfo"); 
    username = userData.display_name;

    // Get artists on listenTo
    listenTo = await backend.request("http://localhost:5000/api/listenTo");
    for(const [index, artist] of listenTo.entries()) {
	for(const image of artist.Images) {
	    if(image.Width === 160) {
		listenTo[index].image_url = image.URL;
	    }
	}
    }
    loadingListenTo = false;

    // TODO this should run async with above
    // Get artists on listenTo
    toReview = await backend.request("http://localhost:5000/api/toReview");
    console.log("toReview:");
    console.log(toReview);
    for(const [index, artist] of toReview.entries()) {
	for(const image of artist.Images) {
	    if(image.Width === 160) {
		toReview[index].image_url = image.URL;
	    }
	}
    }
    //listenTo = listenTo;
    loadingToReview = false;
});
</script>
