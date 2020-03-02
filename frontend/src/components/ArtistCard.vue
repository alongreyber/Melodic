<template>
<div class="card add-spacing">
    <div class="card-content">

	<div class="columns is-one-third">
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

    </div>
</div>
</template>

<script>
export default {
    name: 'ArtistCard',
    props: ["image_url", "Name"]
}
</script>

<style>
.add-spacing {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>
