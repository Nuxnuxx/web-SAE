<script lang="ts">
	import { playlistStore } from "../store";
	export let idRecipe: number;

	let filter = "";
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

<style lang="scss">
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
</style>
