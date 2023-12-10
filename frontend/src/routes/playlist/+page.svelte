<script lang="ts">
	import CardPlaylist from "../../components/cardPlaylist.svelte";
	import type { PageLoad } from "./[slug]/$types";

	export let data: PageLoad;
</script>

<svelte:head>
	<title>Playlists | PirateCook</title>
</svelte:head>

<div class="wrapper">
	<h2>Vos livres de <span class="">recettes</span></h2>

	<div class="playlist__list__wrapper">
		{#each data.playlists.result as playlist}
			<CardPlaylist {playlist} />
		{/each}
	</div>
</div>

<style lang="scss">
	.wrapper {
		padding: 0.2rem 0;
	}

	.playlist__list__wrapper {
		display: flex;
		flex-direction: column;
		row-gap: 1rem;
	}

	h2 {
		color: var(--black-color);
		text-align: center;
		font-size: 2rem;
		@media (min-width: 400px) {
			font-size: 3rem;
		}
		@media (min-width: 768px) {
			font-size: 4rem;
		}
		font-weight: bold;
		margin-bottom: 2%;
		span {
			font-family: Leckerli One;
			color: var(--primary-color);
		}
	}
</style>
