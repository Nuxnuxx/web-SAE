<script lang="ts">
	import Selector from "./selector.svelte";

	let answers: number[] = [];

	let url = "search/?name=saucisse";

	let filter = true;
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
					data={[
						"euro_symbol",
						"euro_symbol",
						"euro_symbol",
						"euro_symbol",
						"euro_symbol",
					]}
					rating_like={true}
					bind:selected={answers[1]}
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
						"sentiment_satisfied",
					]}
					rating_like={true}
					bind:selected={answers[2]}
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

			<a href={url}>submit</a>
		</div>
	{/if}
</div>

<style lang="scss">
	.filter {
		.filter__title {
			margin-bottom: 1rem;

			button {
				all: unset;

				display: flex;
				align-items: center;
				cursor: pointer;
			}

			.material-symbols-outlined {
				font-size: 1.5rem;
				margin-right: 1rem;
				opacity: 0.2;
			}

			.filter__title__text {
				color: rgba(0, 0, 0, 0.2);
				font-family: Inter;
				font-size: 24px;
				font-style: normal;
				font-weight: 400;
				line-height: normal;
			}
		}

		.filter__card {
			display: flex;
			flex-direction: column;
			border: 1px solid #e0e0e0;
			border-radius: 20px;

			.filter__card__price,
			.filter__card__difficulty,
			.filter__card__tag {
				display: flex;
				flex-direction: row;
				align-items: center;

				padding-left: 10%;
			}
			.title__text {
				color: #000;
				font-family: Inter;
				font-size: 32px;
				font-style: normal;
				font-weight: 400;
				line-height: normal;

				@media screen and (max-width: 850px) {
					display: none;
				}
			}
		}
	}
</style>
