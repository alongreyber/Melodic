<template>
<div class="card add-spacing">
    <div class="card-content">

	<div class="columns is-one-third is-mobile is-vcentered">
	    <div class="column is-one-fifth is-square">
		<img :src="image_url_or_placeholder" alt="Could not get image" style="max-width: 50%">
	    </div>
	    <div class="column is-size-5">
		{{ Name }}
	    </div>
	    <div class="column is-one-fifth">
		<button class="button is-primary" @click="$emit('makeReview')">
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
    props: ["image_url", "Name"],
    computed: {
	image_url_or_placeholder: function() {
	    if(this.image_url)
		return this.image_url
	    else
		return "https://d1nhio0ox7pgb.cloudfront.net/_img/o_collection_png/green_dark_grey/256x256/plain/user.png"

	}
    },

}
</script>

<style>
.add-spacing {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>
