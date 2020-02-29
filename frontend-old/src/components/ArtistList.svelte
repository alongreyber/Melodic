<div class="artist-list {loading ? 'isLoadingEl' : '' }">
    {#each artists as artist}
	<!-- Syntax below automatically propogates event upwards -->
	<ArtistCard {artist} on:sendToReview ></ArtistCard>
    {/each}
</div>

<script>
import ArtistCard from "./ArtistCard.svelte";

export let artists;
export let loading;

</script>

<style>
.artist-list {
    border: 1px solid gray;
    padding: 2px 2px 2px 2px;
    min-height: 50vh;
}
</style>
