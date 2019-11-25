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

async function sendToReview( event ) {
    // Remove from listenTo
    const artist = event.detail;
    listenTo = listenTo.filter(obj => obj.ID !== artist.ID)
    // Add to start of toReview
    toReview.unshift(artist);
    toReview = toReview;
    // Update in DB

    const resp = await fetch("http://localhost:5000/api/moveToReview/" + artist.ID, 
	{ credentials: "include" } 
    );
    if(!resp.ok) {
	// Do something
	console.log("Request failed (badly)");
    }
}

onMount( async () => {
    const userResp = await fetch("http://localhost:5000/api/getUserInfo", 
	{ credentials: "include" } 
    );
    if(!userResp.ok) {
	// Do something
    }
    const userData = await userResp.json();
    username = userData.data.display_name;

    // Get artists on listenTo
    const listenResp = await fetch("http://localhost:5000/api/listenTo", {
	credentials: "include"
    });
    if(!listenResp.ok) {
	console.log("Response failed:");
	console.log(listenResp);
	// Do something
    }
    const listenData = await listenResp.json();
    listenTo = listenData.data;
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
    const reviewResp = await fetch("http://localhost:5000/api/toReview", {
	credentials: "include"
    });
    if(!reviewResp.ok) {
	console.log("Response failed:");
	console.log(listenResp);
	// Do something
    }
    const reviewData = await reviewResp.json();
    toReview = toReview.data;
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
