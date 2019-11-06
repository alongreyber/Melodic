<template>
    <div>
	<h1 class="title is-1">
	    Welcome {{username}}!
	</h1>
	<div class="columns">
	    <div class="column">
		<h3 class="title is-3">Artists to Listen To</h3>
		<div class="card">
		    <div class="card-content">
			<ul>
			    <li v-for="artist in listenTo">
				{{artist.Name}}
			    </li>
			</ul>
		    </div>
		</div>
	    </div>
	    <div class="column">
	    </div>
	</div>
    </div>
</template>

<script lang="ts">
import axios from "axios";
import { Component, Vue } from "vue-property-decorator";
import { store } from "../utils/store";
import { Message } from "../utils/message";

@Component
export default class LoggedIn extends Vue {
    public store = store;
    public username: string = "";
    public listenTo: any[] = [];
    public mounted() {
	const component = this;
	axios.get("http://localhost:5000/api/getUserInfo", {
	    withCredentials: true
	}).then(function(response: any) {
	    const spotifyResp = response.data.data;
	    component.username = spotifyResp.display_name;
	}).catch(function(error: any) {
	    if(error.response.status === 401) {
		store.addMessage(new Message("Please log in again", "is-error"));
		component.$router.push({name: "home"});
	    }
	});

	// Get artists following
	axios.get("http://localhost:5000/api/getListenTo", {
	    withCredentials: true
	}).then(function(response: any) {
	    component.listenTo = response.data.data;
	    for(var i = 0; i < component.listenTo.length; i++) {
		for(var j = 0; j < component.listenTo[i].Images.length; j++) {
		    if(component.listenTo[i].Images[j].Width === "160") {
			const image_url = component.listenTo[i].Images[j].URL;
			component.listenTo[i].image_url = image_url;
		    }
		}
	    }
	    // Fetch more info for artist
	}).catch(function(error: any) {
	});

    }
}
</script>
