<script lang="ts">
	import readable from "readable-numbers";
	import type { Recipe } from "../app";
	export let data: Recipe;

	const like = () => {
		data.liked = !data.liked;
	};
	const save = () => {
		data.saved = !data.saved;
	};
</script>

<div class="card">
	<div class="card__img">
		<img src={data.img} alt={data.title} />
	</div>
	<div class="card__content">
		<div class="card__title">
			<h3>{data.title}</h3>
			<div class="card__likes">
				<span class="card__likes__number"
					>{readable(data.nbLikes, 1)}</span
				>
				<span class="material-symbols-rounded filled"> favorite </span>
			</div>
		</div>
		<div class="card__icons">
			<span class="card__likes__icon">
				{#if data.liked}
					<button
						on:click={() => like()}
						class="material-symbols-rounded filled red"
					>
						favorite
					</button>
				{:else}
					<button
						on:click={() => like()}
						class="material-symbols-rounded"
					>
						favorite
					</button>
				{/if}
			</span>
			<span class="card__saved__icon">
				{#if data.saved}
					<button
						on:click={() => save()}
						class="material-symbols-rounded green"
					>
						playlist_add_check
					</button>
				{:else}
					<button
						on:click={() => save()}
						class="material-symbols-rounded"
					>
						playlist_add
					</button>
				{/if}
			</span>
		</div>
	</div>
</div>

<style lang="scss">
	.card {
		aspect-ratio: 1/1;
		display: flex;
		flex-direction: column;
		border-radius: 1.5vh;
		overflow: hidden;
		// box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);

		.card__img {
			width: 100%;
			height: 70%;
			border-radius: 5px;
			overflow: hidden;

			img {
				width: 100%;
				height: 100%;
				object-fit: cover;
			}
		}
		.card__content {
			padding: 5px;
			display: grid;
			// grid-auto-flow: column;
			grid-template-columns: 1fr auto; /* The first column takes all available space, the second takes only necessary space */
			gap: 5px; /* Adjust the gap as needed */
			height: 30%;
			.card__title {
				h3 {
					font-size: 14px;
					margin: 0;
					width: 100%;
					line-height: normal;
					overflow: hidden;
					-webkit-box-orient: vertical;
					-webkit-line-clamp: 2;
					display: -webkit-box;
				}
				.card__likes {
					display: flex;
					align-items: center;
					height: 40%;
					color: #c4c4c4;
					font-size: 12px;
					.material-symbols-rounded {
						font-size: inherit;
						margin-left: 3px;
					}
				}
			}
			.card__icons {
				display: flex;
				flex-direction: column;
				align-items: center;
			}
		}
	}

	button.material-symbols-rounded {
		/* reset all button properties */
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}

	.material-symbols-rounded {
		width: max-content;
		font-variation-settings:
			"FILL" 0,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
	}

	.material-symbols-rounded.filled {
		font-variation-settings:
			"FILL" 1,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
	}

	.material-symbols-rounded.red {
		color: #f93362;
	}

	.material-symbols-rounded.green {
		color: #40d133;
	}

	@media (min-width: 768px) {
		.card {
			.card__content {
				padding: 10px;
				.card__title {
					h3 {
						font-size: 16px;
					}
				}
				.card__likes {
					font-size: 14px;
				}
			}
		}
	}
</style>
