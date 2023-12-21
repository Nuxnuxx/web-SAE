<script lang="ts">
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
	import {
		playlistStore,
		isPLaylistAddButtonOpen,
		userStore,
	} from "../store";
	export let idRecipe: number;

	let filter = "";
	$: playlistList = $playlistStore.filter((a) => {
		if (filter != "") {
			return a.name.toLowerCase().includes(filter.toLowerCase());
		} else {
			return a.name != "Vos préférés";
		}
	});
</script>

<div class="dropdown">
	<button
		on:click={() => {
			if ($isPLaylistAddButtonOpen.id != idRecipe) {
				isPLaylistAddButtonOpen.set({
					id: idRecipe,
					open: true,
				});
			} else {
				isPLaylistAddButtonOpen.set({
					id: idRecipe,
					open: !$isPLaylistAddButtonOpen.open,
				});
			}
		}}
		class="material-symbols-rounded">playlist_add</button
	>

	{#if $isPLaylistAddButtonOpen.open && $isPLaylistAddButtonOpen.id == idRecipe && $userStore}
		<div class="dropdown-content">
			<form use:enhance method="post" action="?/addPlaylistRecipe">
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
	{:else if $isPLaylistAddButtonOpen.open && $isPLaylistAddButtonOpen.id == idRecipe && !$userStore}
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
	.dropdown-content {
		display: block;
		position: absolute;
		background-color: var(--white-color);
		border: 3px solid var(--light-secondary-color);
		border-radius: 10px;
		width: 12rem;
		&::before {
			//TODO: tom fais ce truc
			content: "";
			position: absolute;
			top: -7%;
			left: 0%;
			width: 10%;
			height: 10%;
			border: 2.5px solid var(--light-secondary-color);
			border-bottom: none;
			border-right: none;
			background-color: var(--white-color);
			transform: rotate(45deg);
		}

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
</style>
