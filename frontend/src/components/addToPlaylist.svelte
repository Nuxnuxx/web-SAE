<script lang="ts">
	import { enhance } from "$app/forms";
	import { playlistStore } from "../store";
	export let idRecipe: number;

	let isOpen = false;
</script>

<div class="dropdown">
	<button on:click={() => (isOpen = !isOpen)} class="material-symbols-rounded"
		>playlist_add</button
	>

	{#if isOpen}
		<div class="dropdown-content">
			<form use:enhance method="post" action="?/addPlaylistRecipe">
				<input hidden name="id" value={idRecipe} type="text" />

				{#each $playlistStore as playlist}
					<div>
						<button
							type="submit"
							name="idlist"
							value={playlist.idPlaylist}
							class="dropdown-item"
						>
							{playlist.name}
						</button>
					</div>
				{/each}
			</form>
		</div>
	{/if}
</div>

<style lang="scss">
	.dropdown-content {
		display: block;
		position: absolute;
		background-color: #f9f9f9;
		box-shadow: 0px rgba(0, 0, 0, 0.2);
		z-index: 1;
	}

	.dropdown-item {
		cursor: pointer;
		/* reset all button properties */
		background: none;
		border: none;
		padding: 1rem;
		cursor: pointer;
	}

	.dropdown-item:hover {
		background-color: #ddd;
	}

	button.material-symbols-rounded {
		/* reset all button properties */
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}
</style>
