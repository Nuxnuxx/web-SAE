<script lang="ts">
	import { enhance } from "$app/forms";
	import { playlistStore, isPLaylistAddButtonOpen } from "../store";
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

	{#if $isPLaylistAddButtonOpen.open && $isPLaylistAddButtonOpen.id == idRecipe}
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
			content: "";
			position: absolute;
			top: -10%;
			left: 0%;
			width: 10%;
			height: 10%;
			border: 3px solid var(--light-secondary-color);
			clip-path: polygon(50% 0%, 0 100%, 100% 100%);
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
	}

	button.material-symbols-rounded {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}
</style>
