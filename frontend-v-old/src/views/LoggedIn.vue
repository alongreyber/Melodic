<template>
    <div>
	<h1 class="title is-1">
	    Welcome {{username}}!
	</h1>
	<div class="columns">
	    <div class="column">
		<h3 class="title is-3">Artists to Listen To</h3>
		<ArtistList v-bind:listenTo="listenTo" v-bind:class="{ isLoadingEl: loadingListenTo }"></ArtistList>
	    </div>
	    <div class="column">
	    </div>
	    <div class="column">
		<h3 class="title is-3">Artists to Review</h3>
	    </div>
	</div>
    </div>
</template>

<script lang="ts">
import axios from "axios";
import { Component, Vue } from "vue-property-decorator";
import store from "@/utils/store";
import Message from "@/utils/message";

import ArtistList from "@/components/ArtistList.vue";


@Component({
    components: {
	ArtistList,
    }
})
export default class LoggedIn extends Vue {
    public store = store;
    public username: string = "";
    public listenTo: any[] = [];

    public loadingListenTo: boolean = true;

    public mounted() {
	const component = this;
	axios.get("http://localhost:5000/api/getUserInfo", {
	    withCredentials: true,
	}).then((response: any) => {
	    const spotifyResp = response.data.data;
	    component.username = spotifyResp.display_name;
	}).catch((error: any) => {
	    if(error.response.status === 401) {
		store.addMessage(new Message("Please log in again", "is-error"));
		component.$router.push({name: "home"});
	    }
	});

	// Get artists following
	axios.get("http://localhost:5000/api/getListenTo", {
	    withCredentials: true,
	}).then((response: any) => {
	    component.listenTo = response.data.data;
	    for(const [index, artist] of component.listenTo.entries()) {
		for(const image of artist.Images) {
		    console.log("Image");
		    if(image.Width === 160) {
			console.log("Setting image URL");
			component.listenTo[index].image_url = image.URL;
		    }
		}
	    }
	    component.loadingListenTo = false;
	    // Fetch more info for artist
	}).catch((error: any) => {
	});

    }
}
</script>

<style>
</style>
