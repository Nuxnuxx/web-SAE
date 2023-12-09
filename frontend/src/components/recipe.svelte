<script lang="ts">
	import type { RecipeData } from "$lib/api/recipe-types";

	export let recipe: RecipeData;
	//FIXME: image is not displayed maybe the path is wrong
</script>

<div>
	<div class="recipe__header">
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

	<div class="recipe__incredient">
		<ul class="recipe__ingredient__list">
			{#each Object.entries(recipe.recipeIngredients) as [_, ingredient]}
				<li>
					<img src={ingredient.urlPicture} alt="" />
					<span class="recipe__ingredient__text"
						>{ingredient.name}</span
					>
				</li>
			{/each}
		</ul>

		<div>
			<img
				class="recipe__img"
				src={recipe.recipeDetail.images[0]}
				alt="Image de la recette"
				aria-hidden="true"
			/>
		</div>
	</div>

	<div class="recipe__line">
		<span class="recipe__line__text"> Préparation </span>
	</div>

	<div class="recipe__instruction">
		{#each Object.entries(recipe.recipeStep) as [i, instruction]}
			<p>{i}.</p>
			<p class="recipe__instruction__text">{instruction.Step}</p>
		{/each}
	</div>

	<div class="recipe__line"> </div>
</div>

<style lang="scss">
	.recipe__header {
		h2 {
			color: #de403e;
			text-align: center;
			font-family: Leckerli One;
			font-size: 96px;
			font-style: normal;
			font-weight: 400;
			line-height: normal;
			margin-bottom: 2%;
		}

		.recipe__header__infos {
			display: flex;
			justify-content: space-around;
			margin: 0 auto;
			width: 50%;
			margin-bottom: 50px;
			.recipe__header__infos__item {
				display: flex;
				flex-direction: row;
				align-items: center;

				.material-symbols-rounded {
					font-size: 35px;
					color: #de403e;
					margin-right: 5px;
				}
			}
		}
	}

	.recipe__incredient {
		@media screen and (min-width: 600px) {
			display: flex;
			flex-direction: row;
			justify-content: space-around;
		}

		.recipe__ingredient__list {
			@media screen and (min-width: 500px) {
				columns: 2;
				-webkit-columns: 2;
				-moz-columns: 2;
			}
			list-style-type: none;

			/*mettre le texte au centre de la colonne*/
			text-align: center;

			li {
				display: flex;
				flex-direction: row;
				align-items: center;
				margin-bottom: 5px;
			}

			img {
				width: 100px;
				height: 100px;
				margin-right: 5px;

				border: 1px solid black;
				border-radius: 10%;
			}
		}
	}

	.recipe__instruction {
		padding: 3% 10%;
		font-size: 20px;
		font-family: Leckerli One;
		font-style: normal;
		font-weight: 400;
		line-height: normal;
		text-align: justify;

		p {
			color: #000;
			font-family: Leckerli One;
			font-size: 36px;
			font-style: normal;
			font-weight: 400;
			line-height: normal;
		}

		.recipe__instruction__text {
			color: #000;
			font-family: Inter;
			font-size: 24px;
			font-style: normal;
			font-weight: 400;
			line-height: normal;
		}
	}

	.recipe__img {
		width: 250px;
		height: 250px;
	}

	.recipe__line {
		width: 60%;
		height: 0px;
		margin: 0 auto;
		border-bottom: 1px solid black;
		text-align: center;
		opacity: 0.5;
		padding-bottom: 2%;

		@media screen and (max-width: 800px) {
			margin-bottom: 5%;
		}

		.recipe__line__text {
			color: #000;
			font-family: Leckerli One;
			font-size: 40px;
			padding: 0 30px;
			background-color: #fff;
		}
	}
</style>
