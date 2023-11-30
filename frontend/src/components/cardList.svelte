<script lang="ts">
	import type { Recipe } from "../app";
	import RecipeCard from "../components/recipeCard.svelte";
	export let listRecipe: Recipe[];
</script>

<h2>Test liste éléments : </h2>
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
