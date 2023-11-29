<script lang="ts">
	import Selector from "./selector.svelte";

	let currentQuestion = 0;
	let answers: number[] = [];

	function SubmitEvent() {
		// Send answers to the server
	}
</script>

<main>
	<div class="container">
		{#if currentQuestion === 0}
			<h2
				>Nous allons maintenant vous poser quelque
				<span class="highlight">questions</span>
				pour apprendre à vous connaître.
			</h2>
			<button
				class="container__buttonNext"
				on:click={() => {
					currentQuestion++;
				}}
			>
				<span class="buttonNext__text">Suivant</span><span
					class="material-symbols-rounded"
				>
					arrow_forward
				</span>
			</button>
			<a class="container__buttonPass" href="/">Passez cette étape</a>
		{:else if currentQuestion === 1}
			<h2>
				Quelle tranche de
				<span class="highlight"> prix </span>
				préférez vous ?
			</h2>
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
			<button
				class="container__buttonNext"
				on:click={() => {
					currentQuestion++;
				}}
			>
				<span class="buttonNext__text">Suivant</span><span
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
					"sentiment_satisfied",
				]}
				bind:selected={answers[2]}
			/>
			<button
				class="container__buttonNext"
				on:click={() => {
					currentQuestion++;
				}}
			>
				<span class="buttonNext__text">Suivant</span><span
					class="material-symbols-rounded"
				>
					arrow_forward
				</span>
			</button>
		{:else if currentQuestion === 3}
			<h2>
				Et pour les <span class="highlight">épices</span> ?
			</h2>
			<Selector
				data={["fire_truck", "local_fire_department", "block"]}
				bind:selected={answers[3]}
			/>
			<button
				class="container__buttonNext"
				on:click={() => {
					currentQuestion++;
				}}
			>
				<span class="buttonNext__text">Suivant</span><span
					class="material-symbols-rounded"
				>
					arrow_forward
				</span>
			</button>
		{:else}
			<h2>
				Pour finir, fait nous savoir ce que tu aimes en choisissant
				<span class="highlight">au moins 3</span>
				recettes.
			</h2>

			<button
				class="container__buttonNext"
				on:click={() => {
					SubmitEvent;
				}}
			>
				<span>Terminer</span>
			</button>
		{/if}
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

		.container__buttonNext {
			padding: 10px 50px;
			background-color: #de403e;
			color: #fff;
			border: none;
			border-radius: 20px;
			cursor: pointer;
			font-size: 16px;
			display: flex;
			align-items: center;

			.buttonNext__text {
				margin-right: 10px;
			}
		}

		.container__buttonPass {
			margin-top: 10px;
			color: #939393;
			text-decoration: underline;
		}
	}

	.highlight {
		color: #de403e;
		font-weight: bold;
	}
</style>
