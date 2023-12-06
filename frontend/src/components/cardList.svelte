<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";
	import RecipeCard from "../components/recipeCard.svelte";
	export let listRecipe: RecipeDetail[];
</script>

<link
	rel="stylesheet"
	href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0"
/>

<div class="list">
	{#each listRecipe as recipe}
		<RecipeCard data={recipe} />
	{/each}
</div>

<style lang="scss">
	.list {
		margin: 1rem;
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		grid-gap: 1rem;
		grid-auto-rows: minmax(100px, auto);
	}
	@media (min-width: 768px) {
		.list {
			margin: 2rem 4rem;
			grid-template-columns: repeat(3, 1fr);
		}
	}
	@media (min-width: 1024px) {
		.list {
			grid-template-columns: repeat(4, 1fr);
		}
	}
</style>
