<template>
    <div>Callback</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import store from '../store';

const axios = require('axios');

@Component
export default class SpotifyCallback extends Vue {
    data: {
	store: store
    }
    mounted() {
	var spotifyParams = this.$route.query;
	if("error" in spotifyParams) {
	    if(spotifyParams.error == "access_denied") {
		store.addMesssage({text: "Please allow Spotify access to use this app", color: "is-error"});
		this.$router.push({name: "home"});
	    }
	}
	if(!("state" in spotifyParams) ||
	    spotifyParams.state != getCookie("spotifyState")) {
		store.addMesssage({text: "Invalid State", color: "is-error"});
		this.$router.push({name: "home"});

	}

	axios.get('http://localhost:5000/api/login', {
	    params: spotifyParams,
	    withCredentials: true,
	}).then(function(response) {

	    this.$router.push({name: "logged_in"});
	});

    }
}

function getCookie(cname: string) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}
</script>
