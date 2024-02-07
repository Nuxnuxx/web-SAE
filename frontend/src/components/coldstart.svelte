<script lang="ts">
	import { getRecipes } from "$lib/api/auth-request";
	import Selector from "./selector.svelte";
	import { enhance } from "$app/forms";
	import { Price, Difficulty } from "$lib/api/recipe-types";
	import Swiper2 from "./Swiper2.svelte";
	import { coldstartLiked } from "../store";
	import { redirect } from "@sveltejs/kit";

	let currentQuestion = 0;
	let answers: number[] = [];
	let recipes: any[] = [];
	let jsonRecipes: any;

	/*	async function getRecipesByPreferences() {
		jsonRecipes = await getRecipes(
			"",
			Price[answers[1]],
			Difficulty[answers[2]],
			0
		);
		recipes = jsonRecipes.result;
		if (recipes.length < 4) {
			jsonRecipes = await getRecipes("", Price[answers[1]], "", 0);

			recipes = jsonRecipes.result;
		}
	}*/
</script>

<main>
	<div class="container">
		<form method="post" action="?/endColdstart">
			{#if currentQuestion === 0}
				<h2
					>Nous allons maintenant vous poser quelque
					<span class="highlight">questions</span>
					pour apprendre à vous connaître.
				</h2>
				<button
					class="button"
					on:click={() => {
						currentQuestion++;
					}}
				>
					<span class="buttonnext__text">Suivant</span><span
						class="material-symbols-rounded"
					>
						arrow_forward
					</span>
				</button>
				<a class="container__buttonpass" href="/">Passez cette étape</a>
			{:else if currentQuestion === 1}
				<h2>
					Quelle tranche de
					<span class="highlight"> prix </span>
					préférez vous ?
				</h2>
				<Selector
					data={["euro_symbol", "euro_symbol", "euro_symbol"]}
					rating_like={true}
					bind:selected={answers[1]}
				/>
				<button
					class="button"
					on:click={() => {
						currentQuestion++;
					}}
				>
					<span class="buttonnext__text">Suivant</span><span
						class="material-symbols-rounded"
					>
						arrow_forward
					</span>
				</button>
			{:else if currentQuestion === 2}
				<h2>
					Avec quel niveau de
					<span class="highlight">difficulté</span>
					êtes vous familier ?
				</h2>
				<Selector
					data={[
						"mood_bad",
						"sentiment_neutral",
						"sentiment_content",
						"sentiment_calm",
					]}
					bind:selected={answers[2]}
				/>
				<button class="button" on:click={() => {}} type="submit">
					<span class="buttonnext__text">Suivant</span><span
						class="material-symbols-rounded"
					>
						arrow_forward
					</span>
				</button>

				<input type="hidden" name="price" value={answers[1]} />
				<input type="hidden" name="difficulty" value={answers[2]} />
			{/if}
		</form>
	</div>
</main>

<style lang="scss">
	h2 {
		font-size: 2rem;
		text-align: center;
		font-weight: 400;
	}

	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

		margin: 10%;

		form {
			display: flex;
			flex-direction: column;
			align-items: center;
		}

		.container__buttonpass {
			margin-top: 10px;
			color: var(--light-black-color);
			text-decoration: underline;
		}

		.button {
			padding: 10px 50px;
			background-color: var(--primary-color);
			color: #fff;
			border: none;
			border-radius: 20px;
			cursor: pointer;
			font-size: 16px;
			display: flex;
			align-items: center;

			.buttonnext__text {
				margin-right: 10px;
			}

			&:disabled {
				background-color: var(--light-secondary-color);
				border: 3px solid var(--light-secondary-color);
				cursor: not-allowed;
			}
		}
	}
</style>
