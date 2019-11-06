<template>
    <div>Callback</div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { store } from "../utils/store";
import { Message } from "../utils/message";

import axios from "axios";

@Component
export default class SpotifyCallback extends Vue {
    public store = store;
    public mounted() {
	var spotifyParams = this.$route.query;
	if("error" in spotifyParams) {
	    if(spotifyParams.error === "access_denied") {
		store.addMessage(new Message("Please allow Spotify access to use this app", "is-error"));
		this.$router.push({name: "home"});
	    }
	}
	if(!("state" in spotifyParams) ||
	    spotifyParams.state != getCookie("spotifyState")) {
		store.addMessage(new Message("Invalid State", "is-error"));
		this.$router.push({name: "home"});

	}

	const component = this;
	axios.get("http://localhost:5000/api/login", {
	    params: spotifyParams,
	    withCredentials: true,
	}).then(function(response: any) {

	    component.$router.push({name: "logged_in"});
	}).catch(function(error: any) {
	});

    }
}

function getCookie(cname: string) {
  const name = cname + "=";
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(";");
  for(const c of ca) {
    while (c.charAt(0) === " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) === 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}
</script>
