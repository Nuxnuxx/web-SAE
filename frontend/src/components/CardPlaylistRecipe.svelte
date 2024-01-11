<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";

	export let recipe: RecipeDetail;
	$: finalArrayImages = recipe.images.replace(/[\[\]"]+/g, "").split(", ");
</script>

<div class="recipe__wrapper">
	<img src={finalArrayImages[0]} alt={recipe.name} />
	<div class="text__wrapper">
		<h3>{recipe.name}</h3>
		<p>{recipe.time}</p>
	</div>
	<a href={`/recipe/${recipe.idRecipe}`}>
		<span class="material-symbols-rounded"> arrow_forward </span>
	</a>
</div>

<style lang="scss">
	.recipe__wrapper {
		display: flex;
		text-align: left;
		flex-direction: row;
		gap: 1rem;
		align-items: center;
		margin: 0 auto;
		padding-right: 0.5rem;
		width: 80%;
		border: 1px solid var(--light-secondary-color);
		border-radius: 10px;
		overflow: hidden;

		img {
			height: 7rem;
			object-fit: cover;
			aspect-ratio: 1/1;
		}

		.text__wrapper {
			display: flex;
			flex-direction: column;

			h3 {
				margin: 0 0 0.5rem 0;
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
</style>
