<script lang="ts">
	import type { PageData } from "../$types";
	import CardPlaylistRecipe from "../../../components/CardPlaylistRecipe.svelte";

	export let data: PageData;
</script>

<svelte:head>
	<title
		>{data.playlists[0].name == "liked"
			? "Vos préférés"
			: `Playlist ${data.playlists[0].name}`} | PirateCook</title
	>
</svelte:head>

<div class="wrapper">
	<h2
		>{data.playlists[0].name == "liked"
			? "Vos préférés"
			: `${data.playlist[0].name}`}</h2
	>

	<div class="recipe__wrapper">
		{#each data.playlist.result as recipe}
			<CardPlaylistRecipe {recipe} />
		{/each}
	</div>
</div>

<style lang="scss">
	.wrapper {
		text-align: center;
		h2 {
			color: var(--primary-color);
			text-align: center;
			font-family: Leckerli One;
			font-size: 5rem;
			font-style: normal;
			font-weight: 400;
			line-height: normal;
		}

		.recipe__wrapper {
			display: flex;
			flex-direction: column;
			row-gap: 1rem;
		}
	}
</style>
