<script lang="ts">
	import { playlistStore } from "../../store";
	import CardPlaylist from "../../components/cardPlaylist.svelte";

	let name = "";

	// function to know if the new name of the playlist is already taken in the store
	function isNameAlreadyTaken(name: string) {
		return $playlistStore.some((playlist) => playlist.name == name);
	}
</script>

<svelte:head>
	<title>Playlists | PirateCook</title>
</svelte:head>

<div class="wrapper">
	<h2>Vos livres de <span class="">recettes</span></h2>

	<form method="post" action="?/createList">
		<input
			placeholder="Nom livre de recettes..."
			type="text"
			name="name"
			bind:value={name}
		/>
		<button
			type="submit"
			disabled={isNameAlreadyTaken(name) || name.length < 1}
		>
			<span class="text"> Ajout </span>
			<span class="material-symbols-outlined"> add </span>
		</button>
	</form>

	<div class="favoris">
		<span></span>
		Favoris
		<span></span>
	</div>

	<div class="playlist__list__wrapper">
		{#each $playlistStore as playlist}
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
		width: 80%;
		@media (min-width: 768px) {
			width: 40%;
		}
		margin: 0 auto;
		input {
			border-radius: 10px;
			border: none;
			outline: 2px solid var(--light-secondary-color);
			margin-right: 1rem;
			padding: 0.5rem 1rem;
			width: fill-available;

			&::placeholder {
				color: var(--light-secondary-color);
				font-size: 1rem;
			}
		}

		button {
			padding: 0 1rem;
			text-decoration: none;
			border-radius: 10px;
			border: 3px solid var(--primary-color);
			color: var(--white-color);
			background-color: var(--primary-color);
			cursor: pointer;
			display: flex;
			flex-direction: row;
			gap: 0.5rem;
			align-items: center;
			font-weight: bold;
			font-size: 1rem;

			@media screen and (max-width: 425px) {
				padding: 0;
			}

			span {
				&.material-symbols-outlined {
					font-size: 2rem;
				}

				@media screen and (max-width: 425px) {
					&.text {
						display: none;
					}
				}
			}

			&:disabled {
				background-color: var(--light-secondary-color);
				border: 3px solid var(--light-secondary-color);
				cursor: not-allowed;
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
		font-size: 2rem;
		text-align: center;

		span {
			width: 20%;
			height: 1.5px;
			background-color: var(--light-secondary-color);
			display: inline-block;
			vertical-align: middle;
			margin: 0 1rem;
			@media (min-width: 768px) {
				margin: 0 2rem;
				width: 30%;
			}
		}
	}
</style>
