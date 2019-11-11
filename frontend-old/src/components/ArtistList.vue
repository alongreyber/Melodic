<template>
    <div class="artist-list" v-bind:class="{ isLoadingEl:loadingListenTo }">
	<ArtistCard v-for="artist in listenTo" v-bind:artist="artist">
	</ArtistCard>
    </div>
</template>

<script lang="ts">
import ArtistCard from "./ArtistCard.vue";
import { Component, Vue } from "vue-property-decorator";

@Component({
    components: {
	ArtistCard,
    },
    props: {
	listenTo: Object,
    },
})
export default class ArtistList extends Vue { }
</script>

<style>
.artist-list {
    border: 1px solid gray;
    padding: 2px 2px 2px 2px;
    min-height: 50vh;
}
</style>
