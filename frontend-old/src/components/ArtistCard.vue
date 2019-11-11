<template>
    <div class="artist-card columns is-one-third">
	<div class="column is-one-fifth">
	    <img src="#" width="50" height="50" v-bind:src="artist.image_url" alt="Could not get image">
	</div>
	<div class="column is-size-5">
	    {{artist.Name}}
	</div>
	<div class="column">
	    <span class="icon is-large column">
		<i class="fas fa-arrow-right fa-2x"></i>
	    </span>
	</div>
    </div>
</template>

<script lang="ts">

import { Component, Vue } from "vue-property-decorator";

@Component({
    props: {
	artist: Object
    }
})
export default class ArtistCard extends Vue {}

</script>

<style>

.artist-card {
    border-radius: 5px;
    box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 3px 10px 0 rgba(0, 0, 0, 0.19);
    align-items: center;
    justify-content: space-between;
    margin: 5px 5px 5px 5px;
}

</style>
