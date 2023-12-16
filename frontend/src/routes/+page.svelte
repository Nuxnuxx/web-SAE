<script lang="ts">
	import RecipeCard from "../components/recipeCard.svelte";
	import type { RecipeDetail } from "$lib/api/recipe-types";


	 // Mock data for demonstration
	const mockRecipe: RecipeDetail = {
		difficulty: 'easy',
		images: ['image-url-1', 'image-url-2'],
		quantity: '4 servings',
		price: 'medium',
		time: '30 minutes',
		name: 'Delicious Recipe',
		idRecipe: 1,
	};
</script>

<svelte:head>
	<title>Accueil | PirateCook</title>
</svelte:head>

<h1>Wow that the homepage</h1>

<div class="recipe__container">
	<RecipeCard data={mockRecipe} />
</div>

<style lang='scss'>
	.recipe__container {
		display: flex;
		justify-content: center;
		width: 20%;
	}
</style>