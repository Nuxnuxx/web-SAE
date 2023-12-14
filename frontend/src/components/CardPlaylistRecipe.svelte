<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";

	export let recipe: RecipeDetail;
</script>

<div class="recipe__wrapper">
	<img src={recipe.images[0]} alt={recipe.name} />
	<div class="text__wrapper">
		{#if recipe.name == "liked"}
			<h3>Vos préférés</h3>
		{:else}
			<h3>{recipe.name}</h3>
		{/if}
		<p>{recipe.time} Recettes</p>
	</div>
	<a href={`/recipe/${recipe.idRecipe}`}>
		<span class="material-symbols-outlined"> arrow_forward </span>
	</a>
</div>

<style lang="scss">
	.recipe__wrapper {
		display: flex;
		text-align: left;
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
		}
	}
</style>
