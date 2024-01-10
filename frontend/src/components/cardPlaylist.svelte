<script lang="ts">
	import type { PlaylistDetail } from "$lib/api/playlist-types";
	import DEFAULT from "$lib/img/default_playlist.png";

	export let playlist: PlaylistDetail;
	let finalArrayImages = playlist.images.replace(/[\[\]"]+/g, "").split(", ");
</script>

<div class="playlist__wrapper">
	<!--INFO: need to get a image of the first recipe to put in the ternary -->
	<img
		src={playlist.numberRecipes >= 1 ? finalArrayImages[0] : DEFAULT}
		alt={playlist.name}
	/>
	<div class="text__wrapper">
		<h3>{playlist.name}</h3>
		<p>{playlist.numberRecipes} Recettes</p>
	</div>
	<a href={`/playlist/${playlist.idPlaylist}`}>
		<span class="material-symbols-rounded"> arrow_forward </span>
	</a>
</div>

{#if playlist.name == "liked"}
	<div class="lines"></div>
{/if}

<style lang="scss">
	.playlist__wrapper {
		display: flex;
		flex-direction: row;
		gap: 1rem;
		align-items: center;
		margin: 0 auto;
		width: 80%;
		border: 1px solid var(--light-secondary-color);
		border-radius: 10px;
		overflow: hidden;
		padding: 0.5rem;

		img {
			height: 5rem;
			object-fit: cover;
			aspect-ratio: 1/1;
			border-radius: 0.5rem;
		}

		.text__wrapper {
			display: flex;
			flex-direction: column;

			h3 {
				margin: 0;
			}
			p {
				color: var(--light-black-color);
				margin: 0;
			}
		}
		a {
			margin-left: auto;
			text-decoration: none;
			color: var(--black-color);
			span {
				font-size: 3rem;
				font-weight: 600;
				padding-right: 1rem;
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
