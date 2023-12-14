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
			: `${data.playlists[0].name}`}</h2
	>

	<div class="lines"></div>

	{#if data.playlist.result == ""}
		<div class="no__recipe">Pas de recettes pour le moment</div>
	{:else}
		<div class="recipe__wrapper">
			{#each data.playlist.result as recipe}
				<CardPlaylistRecipe {recipe} />
			{/each}
		</div>
	{/if}
</div>

<style lang="scss">
	.wrapper {
		padding: 0.1rem 0;
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
		h2 {
			color: var(--primary-color);
			text-align: center;
			font-family: Leckerli One;
			font-size: 4rem;
			font-weight: 400;
		}

		.recipe__wrapper {
			display: flex;
			flex-direction: column;
			row-gap: 1rem;
		}

		.no__recipe {
			color: var(--black-color);
			font-size: 2.5rem;
			font-weight: bold;
			margin: 2rem auto;
		}

		.lines {
			height: 1.5px;
			width: 60%;
			display: block;
			margin: 1rem auto;
			background-color: var(--light-secondary-color);
		}
	}
</style>
