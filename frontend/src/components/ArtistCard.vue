<template>
<div class="artist-card columns is-one-third">
    <div class="column is-one-fifth">
	<img width="50" height="50" :src="image_url" alt="Could not get image">
    </div>
    <div class="column is-size-5">
	{{ Name }}
    </div>
    <div class="column is-one-fifth">
	<button class="button is-primary">
	    Start
	</button>
    </div>
</div>
</template>

<script>
export default {
    name: 'ArtistCard',
    props: ["image_url", "Name"]
}
</script>
