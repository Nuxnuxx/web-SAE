<script lang="ts">
	import { goto } from "$app/navigation";
	import {
		playlistStore,
		isPlaylistAddButtonOpen,
		userStore,
	} from "../store";
	export let idRecipe: number;

	let filter = "";
	// dont include Vos préférés playlist
	$: playlistWithoutFav = $playlistStore.filter((a) => {
		if (a.name != "Vos préférés") {
			return a;
		}
	});
	$: playlistList = playlistWithoutFav.filter((a) => {
		if (filter != "") {
			return a.name.toLowerCase().includes(filter.toLowerCase());
		} else if (a.name != "Vos préférés") {
			return a;
		}
	});
</script>

<div
	class="dropdown {$isPLaylistAddButtonOpen.open &&
	$isPLaylistAddButtonOpen.id == idRecipe
		? 'slected'
		: ''}"
>
	<button
		on:click={() => {
			if ($isPlaylistAddButtonOpen.id != idRecipe) {
				isPlaylistAddButtonOpen.set({
					id: idRecipe,
					open: true,
				});
			} else {
				isPlaylistAddButtonOpen.set({
					id: idRecipe,
					open: !$isPlaylistAddButtonOpen.open,
				});
			}
		}}
		class="material-symbols-rounded">playlist_add</button
	>

	{#if $isPlaylistAddButtonOpen.open && $isPlaylistAddButtonOpen.id == idRecipe && $userStore}
		<div class="dropdown-content">
			<form method="post" action="?/addPlaylistRecipe">
				<input hidden name="idRecipe" value={idRecipe} type="text" />
				<div class="playlist-search-wrapper">
					<div class="material-symbols-rounded">search</div>
					<input
						type="search"
						autocomplete="off"
						bind:value={filter}
						placeholder="Livre de recette"
					/>
				</div>

				<div class="row"></div>

				{#each playlistList as playlist}
					<button
						type="submit"
						name="idPlaylist"
						value={playlist.idPlaylist}
						class="playlist-choice"
					>
						{playlist.name}
					</button>
				{/each}
			</form>
		</div>
	{:else if $isPlaylistAddButtonOpen.open && $isPlaylistAddButtonOpen.id == idRecipe && !$userStore}
		<div class="dropdown-content">
			<p>Apprenons a te connaitre d'abord</p>
			<button on:click={() => goto("/auth")} class="nav__login">
				<span class="material-symbols-rounded">person</span>
				Connexion
			</button>
		</div>
	{/if}
</div>

<style lang="scss">
	.dropdown {
		position: relative;
		&.slected::before {
			z-index: 2;
			content: "";
			position: absolute;
			top: 1.4rem;
			left: 5%;
			width: 0.7rem;
			height: 0.7rem;
			border: 2.5px solid var(--light-secondary-color);
			border-bottom: none;
			border-right: none;
			background-color: var(--white-color);
			transform: rotate(45deg);
		}
	}
	.dropdown-content {
		display: block;
		position: absolute;
		z-index: 1;
		background-color: var(--white-color);
		border: 3px solid var(--light-secondary-color);
		border-radius: 10px;
		width: 12rem;
		left: -100%;

		form {
			display: flex;
			flex-flow: column nowrap;
			margin: 1rem 0.4rem;
			text-align: left;

			.playlist-search-wrapper {
				display: flex;
				flex-flow: row nowrap;
				border: 2.5px solid var(--light-secondary-color);
				border-radius: 10px;
				padding-left: 0.4rem;

				div {
					color: var(--very-light-secondary-color);
					font-size: 1.4rem;
				}

				input {
					width: 75%;
					border: none;
					margin-left: 0.2rem;
					outline: none;
					&::placeholder {
						color: var(--very-light-secondary-color);
					}
				}
			}

			.row {
				width: 95%;
				border: 1px solid var(--light-secondary-color);
				margin: 0.7rem auto;
			}

			.playlist-choice {
				all: unset;
				cursor: pointer;
				margin: 0 0.5rem;
			}
		}

		p {
			font-size: 0.8rem;
			font-weight: bold;
		}

		.nav__login {
			display: flex;
			align-items: center;
			gap: 0.5rem;
			padding: 0.5rem 1rem;
			border-radius: 1rem;
			background-color: var(--primary-color);
			color: var(--white-color);
			font-weight: medium;
			font-size: 1rem;
			border: none;
			cursor: pointer;
			transition: all 0.2s ease-out;
			margin: 0 auto;
			margin-bottom: 1rem;

			&:hover {
				transform: scale(1.02);
			}
		}
	}

	button.material-symbols-rounded {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}

	@media (max-width: 768px) {
		.dropdown {
			&.slected::before {
				display: none;
			}
		}
		.dropdown-content {
			position: fixed;
			left: 50%;
			transform: translate(-50%);
		}
	}
</style>
