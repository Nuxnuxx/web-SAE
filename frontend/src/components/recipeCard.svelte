<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";
	import readable from "readable-numbers";
	export let data: RecipeDetail;
</script>

<div class="card">
	<a href={`/recipe/${data.idRecipe}`} class="card__img">
		<!-- <img src={data.images.toString()} alt={data.name} /> -->
	</a>
	<div class="card__content">
		<a href={`/recipe/${data.idRecipe}`} class="card__title">
			<h3>{data.name}</h3>
			<div class="card__likes">
				<span class="card__likes__number">{readable(100000, 1)}</span>
				<span class="material-symbols-rounded filled"> favorite </span>
			</div>
		</a>
		<div class="card__icons">
			<span class="card__likes__icon">
				<!-- {#if data.liked} -->
				<!-- <form action="?/likeRecipe"> -->
				<!-- 	<button class="material-symbols-rounded filled red"> -->
				<!-- 		favorite -->
				<!-- 	</button> -->
				<!-- </form> -->
				<!-- {:else} -->
				<form method="post" action="?/likeRecipe">
					<input hidden name="id" value={data.idRecipe} type="text" />
					<button
						type="submit"
						class={`material-symbols-rounded like`}
					>
						favorite
					</button>
				</form>
				<!-- {/if} -->
			</span>
			<span class="card__saved__icon">
				<!-- {#if data.saved} -->
				<!-- 	<button -->
				<!-- 		on:click={() => save()} -->
				<!-- 		class="material-symbols-rounded green" -->
				<!-- 	> -->
				<!-- 		playlist_add_check -->
				<!-- 	</button> -->
				<!-- {:else} -->
				<button class="material-symbols-rounded"> playlist_add </button>
				<!-- {/if} -->
			</span>
		</div>
	</div>
</div>

<style lang="scss">
	.card {
		color: black;
		text-decoration: none;
		aspect-ratio: 1/1;
		display: flex;
		flex-direction: column;
		border-radius: 1.5vh;
		overflow: hidden;

		.card__img {
			width: 100%;
			height: 70%;
			border-radius: 0.3rem;
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
			grid-template-columns: 1fr auto;
			gap: 5px;
			height: 30%;
			.card__title {
				text-decoration: none;
				color: var(--black-color);
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
				z-index: 2;
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
		color: var(--light-primary-color);
	}

	.material-symbols-rounded.green {
		color: var(--last-color);
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
