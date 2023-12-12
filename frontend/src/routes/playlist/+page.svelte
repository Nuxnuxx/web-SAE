<script lang="ts">
	import type { PageData } from "../$types";
	import CardPlaylist from "../../components/cardPlaylist.svelte";

	export let data: PageData;
	const playlistList = data.playlists.result.sort((a, b) => {
		return a.name === "liked" ? -1 : b.name === "liked" ? 1 : 0;
	});
</script>

<svelte:head>
	<title>Playlists | PirateCook</title>
</svelte:head>

<div class="wrapper">
	<h2>Vos livres de <span class="">recettes</span></h2>

	<form method="post" action="?/createRecipe">
		<input placeholder="Nom livre de recettes..." type="text" name="name" />
		<button type="submit"
			>Ajout
			<span class="material-symbols-outlined"> add </span>
		</button>
	</form>

	<div class="favoris">
		<span></span>
		Favoris
		<span></span>
	</div>

	<div class="playlist__list__wrapper">
		{#each playlistList as playlist}
			<CardPlaylist {playlist} />
		{/each}
	</div>
</div>

<style lang="scss">
	.wrapper {
		padding: 0.1rem 0;
		display: flex;
		flex-direction: column;
		row-gap: 2rem;
	}

	form {
		display: flex;
		flex-direction: row;
		width: 40%;
		margin: 0 auto;
		input {
			border-radius: 10px;
			border: none;
			outline: 2px solid #dcdcdc;
			margin-right: 1rem;
			padding: 0.5rem 1rem;
			width: 70%;

			&::placeholder {
				color: var(--light-secondary-color);
				font-size: 1.1rem;
			}
		}

		button {
			width: 20%;
			padding: 0 1rem;
			text-decoration: none;
			border-radius: 10px;
			border: 3px solid var(--primary-color);
			color: var(--white-color);
			background-color: var(--primary-color);
			cursor: pointer;
			display: flex;
			flex-direction: row;
			align-items: center;
			font-weight: bold;
			font-size: 1rem;
			span {
				font-size: 2rem;
			}
		}
	}

	.playlist__list__wrapper {
		display: flex;
		flex-direction: column;
		row-gap: 1.5rem;
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

	.favoris {
		color: var(--black-color);
		font-family: Leckerli One;
		font-size: 2.5rem;
		text-align: center;

		span {
			width: 20%;
			height: 1.5px;
			background-color: var(--light-secondary-color);
			display: inline-block;
			vertical-align: middle;
			margin: 0 2rem;
		}
	}
</style>
