<template>
<div>
    <div class='title is-2'>
	Create a Review
    </div>
    <div class="columns">
	<div class="column">
	    <div class="level">
		<div class="level-left">
		    <h3 class="title is-3">Recently Followed</h3>
		</div>
		<div class="level-right">
		    <button class="button is-secondary" @click="refreshFollowed" :class="{'is-loading' : refreshingFollowed}">Refresh</button>
		</div>
	    </div>
	    <ArtistCard v-for="a in recentlyFollowed" v-bind="a" :key="a.SpotifyID"></ArtistCard>
	</div>
	<div class="column">
	    <div class="level">
		<div class="level-left">
		    <h3 class="title is-3">Recently Played</h3>
		</div>
		<div class="level-right">
		    <button class="button is-secondary" @click="refreshListened" :class="{'is-loading' : refreshingListened}">Refresh</button>
		</div>
	    </div>
	    <ArtistCard v-for="a in recentlyListened" v-bind="a" :key="a.SpotifyID"></ArtistCard>
	</div>
    </div>
</div>
</template>

<script>
import { getApi } from '../api';
import ArtistCard from '../components/ArtistCard';

function addImages(list) {
    for(const [index, artist] of list.entries()) {
	for(const image of artist.Images) {
	    if(image.Width === 160) {
		list[index].image_url = image.URL;
	    }
	}
    }
    return list;
}

export default {
    name: 'Dashboard',
    components: { ArtistCard },
    data: function() {
	return {
	    recentlyFollowed: [],
	    recentlyListened: [],
	    refreshingFollowed: false,
	    refreshingListened: false,
	}
    },
    mounted: function() {
	this.getFollowed()
	this.getListened()
    }, 
    methods: {
	getFollowed: async function() { 
	    let artists = await getApi('/recentlyFollowed');
	    if(artists) {
		artists = addImages(artists);
		this.recentlyFollowed = artists;
	    }
	},
	getListened: async function() { 
	    let artists = await getApi('/recentlyListened');
	    if(artists) {
		artists = addImages(artists);
		this.recentlyListened = artists;
	    }
	},
	refreshFollowed: async function() {
	    this.refreshingFollowed = true;
	    let result = await getApi('/recentlyFollowed/refresh');
	    console.log(result)
	    if(result) {
		this.getFollowed();
	    }
	    this.refreshingFollowed = false;
	},
	refreshListened: async function() {
	    this.refreshingListened = true;
	    let result = await getApi('/recentlyListened/refresh');
	    console.log(result)
	    if(result) {
		this.getListened();
	    }
	    this.refreshingListened = false;
	}
    }
}
</script>
