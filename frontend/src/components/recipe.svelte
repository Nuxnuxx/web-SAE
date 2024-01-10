<script lang="ts">
	import type { RecipeData } from "$lib/api/recipe-types";

	export let recipe: RecipeData;
	let finalArrayImages = recipe.recipeDetail.images
		.replace(/[\[\]"]+/g, "")
		.split(", ");
</script>

<div class="recipe">
	<div class="recipe__header">
		<div class="recipe__header__img">
			<img src={finalArrayImages[0]} alt={recipe.recipeDetail.name} />
		</div>
		<h2>{recipe.recipeDetail.name}</h2>
		<div class="recipe__header__infos">
			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">alarm</span>
				<p>{recipe.recipeDetail.time}</p>
			</div>

			<span class="circle"></span>

			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">sentiment_satisfied</span
				>
				<p>{recipe.recipeDetail.difficulty}</p>
			</div>

			<span class="circle"></span>

			<div class="recipe__header__infos__item">
				<span class="material-symbols-rounded">euro_symbol</span>
				<p>{recipe.recipeDetail.price}</p>
			</div>
		</div>
	</div>

	<div class="recipe__line"> </div>

	<div class="recipe__ingredient__wrapper">
		<ul class="recipe__ingredient__list">
			{#each Object.entries(recipe.recipeIngredients) as [_, ingredient]}
				<li>
					<img src={ingredient.urlPicture} alt="" />
					<!-- //FIXME : this is not working properly the split is not working | EDIT TOM : care to explain what should be happening ? -->
					{#if ingredient.name.split(" ").length > 1}
						<span class="recipe__ingredient__text">
							<span>{ingredient.name.split(" ")[0]}</span>
							{ingredient.name.slice(
								ingredient.name.indexOf(" ") + 1
							)}
						</span>
					{:else}
						<span class="recipe__ingredient__text"
							>{ingredient.name}</span
						>
					{/if}
				</li>
			{/each}
		</ul>

		<img
			class="recipe__img"
			src={finalArrayImages[0]}
			alt={recipe.recipeDetail.name}
		/>
	</div>

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
				height: 15rem;
				width: 100%;
				position: relative;
				&::after {
					content: "";
					position: absolute;
					top: 35%;
					left: 0;
					background: linear-gradient(to top, #fff, transparent);
					width: 100%;
					height: 65%;
				}
				img {
					object-fit: cover;
					width: 100%;
					height: 100%;
				}
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
				align-items: center;
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

				.recipe__ingredient__text {
					span {
						font-weight: bold;
					}
				}
			}

			img {
				width: 3rem;
				aspect-ratio: 1;
				margin-right: 1rem;
				border: 1.5px solid var(--light-secondary-color);
				border-radius: 10%;
			}
		}

		.recipe__img {
			display: none;
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
	}

	@media (min-width: 768px) {
		.recipe {
			.recipe__header {
				h2 {
					font-size: 3rem;
				}

				.recipe__header__infos {
					width: 60%;
					.recipe__header__infos__item {
						flex-direction: row;
						p {
							margin-left: 0.5rem;
						}
					}

					.circle {
						width: 0.5rem;
						height: 0.5rem;
						border-radius: 50%;
						background-color: var(--light-secondary-color);
					}
				}
			}

			.recipe__line {
				width: 70%;
			}

			.recipe__ingredient__list {
				display: grid;
				grid-template-columns: repeat(2, 1fr);
				padding-inline-start: 4rem;

				li {
					font-size: 1.5rem;
				}

				img {
					width: 4rem;
				}
			}

			.recipe__instruction {
				padding: 0 3rem;

				p {
					font-size: 1.9rem;
				}

				.recipe__instruction__text {
					font-family: Inter;
					font-size: 1.6rem;
				}
			}
		}
	}

	@media (min-width: 1024px) {
		.recipe {
			.recipe__header {
				.recipe__header__img {
					display: none;
				}
				h2 {
					font-size: max(5vw);
				}

				.recipe__header__infos {
					width: 60%;
					.recipe__header__infos__item {
						p {
							font-size: 1.2rem;
						}
					}
				}
			}

			.recipe__line {
				width: 70%;
			}

			.recipe__ingredient__wrapper {
				display: flex;
				flex-direction: row;
				align-items: center;
			}

			.recipe__ingredient__list {
				width: 50%;
			}

			.recipe__line__wrapper {
				font-size: 3rem;
			}

			.recipe__img {
				display: unset;
				object-fit: cover;
				width: 50%;
				aspect-ratio: 1/1;
				border-radius: 15px;
				margin: 0 2rem;
			}
		}
	}
</style>
