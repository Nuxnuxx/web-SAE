<script lang="ts">
	import type { PlaylistDetail } from "$lib/api/playlist-types";

	export let playlist: PlaylistDetail;
</script>

<div class="playlist__wrapper">
	<img src="" alt="image" />
	<div class="text__wrapper">
		<h3>{playlist.name}</h3>
		<p>{playlist.numberRecipes} Recettes</p>
	</div>
	<a href={`/playlist/${playlist.idPlaylist}`}>
		<span class="material-symbols-rounded filled"> arrow_forward </span>
	</a>
</div>

{#if playlist.name == "liked"}
	<div class="lines"></div>
{/if}

<style lang="scss">
	.playlist__wrapper {
		display: flex;
		flex-direction: row;
		align-items: center;
		margin: 0 auto;
		padding: 0 1rem;
		width: 80%;
		border: 1px solid var(--light-secondary-color);
		border-radius: 10px;
		.text__wrapper {
			margin-left: 2rem;
			display: flex;
			flex-direction: column;

			p {
				color: var(--light-black-color);
			}
		}
		a {
			margin-left: auto;
			text-decoration: none;
			color: var(--black-color);
			span {
				font-size: 3rem;
				font-weight: 600;
			}
		}
	}

	.lines {
		height: 1.5px;
		width: 60%;
		display: block;
		margin: 1rem auto;
		background-color: var(--light-secondary-color);
	}
</style>
