<script lang="ts">
	import Selector from "./selector.svelte";
	import { Price, Difficulty } from "$lib/api/recipe-types";
	import { filterStore } from "../store";

	let answers: number[] = [];

	let filter = false;

	$: filterStore.set({
		price: Price[answers[0]],
		difficulty: Difficulty[answers[1]],
	});
</script>

<div class="filter">
	<div class="filter__title">
		<button on:click={() => (filter = !filter)}>
			<span class="material-symbols-outlined"> tune </span>
			<span class="filter__title__text">Filtres</span>
		</button>
	</div>

	{#if !filter}
		<div class="filter__card">
			<div class="filter__card__price">
				<span class="title__text">Prix : </span>
				<Selector
					data={["euro_symbol", "euro_symbol", "euro_symbol"]}
					rating_like={true}
					bind:selected={answers[0]}
					coldstart={false}
				/>
			</div>

			<div class="filter__card__difficulty">
				<span class="title__text">Difficulté : </span>
				<Selector
					data={[
						"mood_bad",
						"sentiment_neutral",
						"sentiment_content",
						"sentiment_calm",
					]}
					rating_like={false}
					bind:selected={answers[1]}
					coldstart={false}
				/>
			</div>

			<!-- TODO: tag filtering -->
			<!-- <div class="filter__card__tag"> -->
			<!-- 	<span class="title__text">Tags</span> -->
			<!-- 	<Selector -->
			<!-- 		data={[]} -->
			<!-- 		rating_like={true} -->
			<!-- 		bind:selected={answers[3]} -->
			<!-- 	/> -->
			<!-- </div> -->
		</div>
	{/if}
</div>

<style lang="scss">
	.filter {
		width: 88%;
		display: flex;
		flex-direction: column;
		margin: 0 auto;
		.filter__title {
			button {
				all: unset;
				display: flex;
				align-items: center;
				cursor: pointer;
				margin: 1rem 0 1rem 1rem;
			}

			.material-symbols-outlined {
				font-size: 1.5rem;
				margin-right: 1rem;
				opacity: 0.2;
			}

			.filter__title__text {
				color: rgba(0, 0, 0, 0.2);
				font-size: 1.5rem;
			}
		}

		.filter__card {
			display: flex;
			flex-direction: column;
			outline: 2px solid #e0e0e0;
			border-radius: 20px;

			.filter__card__price,
			.filter__card__difficulty {
				display: flex;
				flex-direction: row;
				align-items: center;

				padding-left: 5%;
			}
			.title__text {
				color: var(--black-color);
				font-weight: bold;
				font-size: 1.2rem;
				@media screen and (min-width: 768px) {
					font-size: 1.5rem;
				}
			}
		}
	}
</style>
