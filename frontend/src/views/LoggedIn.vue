<template>
    <div>
	Welcome {{username}}!
    </div>
</template>

<script lang="ts">
const axios = require('axios');
import { Component, Vue } from 'vue-property-decorator';
import { store, Message } from '../store';

@Component
export default class LoggedIn extends Vue {
    store = store
    username = ""
    mounted() {
	let component = this;
	axios.get('http://localhost:5000/api/getUserInfo', {
	    withCredentials: true
	}).then(function(response: any) {
	    const spotifyResp = response.data.data;
	    component.username = spotifyResp.display_name;
	}).catch(function(error: any) {
	    console.log("Recieved error:")
	    console.log(error.response)
	    if(error.response.status == 401) {
		store.addMessage(new Message("Please log in again", "is-error"));
		component.$router.push({name: "home"});
	    }
	})

    }
}
</script>
