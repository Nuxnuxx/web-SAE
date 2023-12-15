<script lang="ts">
	import type { RecipeData } from "$lib/api/recipe-types";

	export let recipe: RecipeData;
	//FIXME: image is not displayed maybe the path is wrong
	// the path is wrong for sure
</script>

<div class="recipe">
	<div class="recipe__header">
		<img
			class="recipe__header__img"
			src={recipe.recipeDetail.images[0]}
			alt={recipe.recipeDetail.name}
		/>
		<h2>{recipe.recipeDetail.name}</h2>
		<div class="recipe__header__infos">
			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">alarm</span>
				<p>{recipe.recipeDetail.time}</p>
			</div>

			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">sentiment_satisfied</span
				>
				<p>{recipe.recipeDetail.difficulty}</p>
			</div>
			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">euro_symbol</span>
				<p>{recipe.recipeDetail.price}</p>
			</div>
		</div>
	</div>

	<div class="recipe__line"> </div>

	<ul class="recipe__ingredient__list">
		{#each Object.entries(recipe.recipeIngredients) as [_, ingredient]}
			<li>
				<img src={ingredient.urlPicture} alt="" />
				<span class="recipe__ingredient__text">{ingredient.name}</span>
			</li>
		{/each}
	</ul>

	<div class="recipe__line__wrapper">
		<span></span>
		Préparation
		<span></span>
	</div>

	<div class="recipe__instruction">
		{#each Object.entries(recipe.recipeStep) as [i, instruction]}
			<p>Étape {i}</p>
			<p class="recipe__instruction__text">{instruction.Step}</p>
		{/each}
	</div>

	<div class="recipe__line"> </div>
</div>

<style lang="scss">
	.recipe {
		display: flex;
		flex-direction: column;
		row-gap: 2rem;

		.recipe__header {
			display: flex;
			flex-direction: column;
			.recipe__header__img {
				width: 100%;
				height: 15rem;
			}

			h2 {
				color: var(--primary-color);
				text-align: center;
				font-family: Leckerli One;
				font-size: 2rem;
				font-style: normal;
				font-weight: 400;
				line-height: normal;
			}

			.recipe__header__infos {
				display: flex;
				justify-content: space-around;
				margin: 0 auto;
				width: 90%;
				.recipe__header__infos__item {
					display: flex;
					flex-direction: column;
					align-items: center;

					.material-symbols-rounded {
						font-size: 2rem;
						color: var(--primary-color);
					}
				}
			}
		}

		.recipe__ingredient__list {
			display: flex;
			flex-direction: column;
			row-gap: 1.5rem;
			padding-inline-start: 1.5rem;

			li {
				display: flex;
				flex-direction: row;
				align-items: center;
				font-size: 1rem;
			}

			img {
				width: 3rem;
				aspect-ratio: 1;
				margin-right: 1rem;
				border: 1.5px solid var(--light-secondary-color);
				border-radius: 10%;
			}
		}

		.recipe__instruction {
			padding: 0 1.5rem;
			font-family: Leckerli One;

			p {
				color: var(--black-color);
				font-family: Leckerli One;
				font-size: 1.6rem;
			}

			.recipe__instruction__text {
				font-family: Inter;
				font-size: 1.2rem;
			}
		}

		.recipe__line {
			width: 80%;
			height: 0px;
			margin: 0 auto;
			border-bottom: 1.5px solid var(--light-secondary-color);
			text-align: center;
		}

		.recipe__line__wrapper {
			color: var(--black-color);
			font-family: Leckerli One;
			font-size: 1.5rem;
			text-align: center;

			span {
				width: 20%;
				height: 1.5px;
				background-color: var(--light-secondary-color);
				display: inline-block;
				vertical-align: middle;
				margin: 0 1rem;
				@media (min-width: 768px) {
					margin: 0 2rem;
					width: 30%;
				}
			}
		}

		@media screen and (min-width: 768px) {
			.recipe {
				.recipe__header {
					.recipe__header__img {
						display: none;
					}
				}
			}
		}
	}
</style>
