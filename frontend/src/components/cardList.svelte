<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";
	import RecipeCard from "../components/recipeCard.svelte";
	export let listRecipe: RecipeDetail[];
</script>

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
			margin: 2rem 0rem;
			grid-template-columns: repeat(3, 1fr);
		}
	}
</style>
