<div class="artist-card columns is-one-third">
    <div class="column is-one-fifth">
	<img width="50" height="50" src="{artist.image_url}" alt="Could not get image">
    </div>
    <div class="column is-size-5">
	{artist.Name}
    </div>
    <div class="column is-one-fifth">
	<button class="button" on:click="{() => dispatch('sendToReview', artist )}" >
	    <span class="icon is-large">
		<i class="fas fa-lg fa-arrow-right"></i>
	    </span>
	</button>
    </div>
</div>

<script>
export let artist;

import { createEventDispatcher } from 'svelte';
const dispatch = createEventDispatcher();
</script>

<style>

.artist-card {
    border-radius: 5px;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2), 0 2px 5px 0 rgba(0, 0, 0, 0.19);
    align-items: center;
    justify-content: space-between;
    margin: 10px 5px 10px 5px;
}

</style>
