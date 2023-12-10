<script lang="ts">
	export let playlist;
</script>

<div class="playlist__wrapper">
	<img src="" alt="saucisse" />
	<div class="text__wrapper">
		{#if playlist.name == "liked"}
			<h3>Vos préférés</h3>
		{:else}
			<h3>{playlist.name}</h3>
		{/if}
		<p>{playlist.numberRecipes} Recettes</p>
	</div>
	<a href={`/playlist/${playlist.idPlaylist}`}>
		<span class="material-symbols-outlined"> arrow_forward </span>
	</a>
</div>

<style lang="scss">
	.playlist__wrapper {
		display: flex;
		flex-direction: row;
		align-items: center;
		margin: 0 auto;
		padding: 0 1rem;
		width: 80%;
		border: 1px solid #939393;
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
		}
	}
</style>
