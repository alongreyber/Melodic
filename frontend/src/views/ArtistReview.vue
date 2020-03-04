<template>
    <div>
	<div class="title is-3">
	    Write a Review!
	</div>
    </div>
</template>

<script>
import { getApi } from '../api';

export default {
    props: ['id'],
    data: function() {
	return {
	    review: null
	}
    },
    mounted: async function() {
	this.review = getApi(`/review/${this.id}`);
    },
}
</script>
