<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";
	import readable from "readable-numbers";
	import AddToPlaylist from "./addToPlaylist.svelte";
	import AddToLiked from "./addToLiked.svelte";
	export let data: RecipeDetail;
	$: finalArrayImages = data.images.replace(/[\[\]"]+/g, "").split(", ");

	let idRecipe = data.idRecipe;
</script>

<div class="card">
	<a
		data-sveltekit-preload-data="tap"
		href={`/recipe/${data.idRecipe}`}
		class="card__img"
	>
		<img src={finalArrayImages[0]} alt={data.name} />
	</a>
	<div class="card__content">
		<a
			data-sveltekit-preload-data="tap"
			href={`/recipe/${data.idRecipe}`}
			class="card__title"
		>
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
				<!-- {/if} -->
				<AddToLiked {idRecipe} />
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
				<AddToPlaylist {idRecipe} />
				<!-- {/if} -->
			</span>
		</div>
	</div>
</div>

<style lang="scss">
	.card {
		position: relative;
		color: black;
		text-align: left;
		text-decoration: none;
		aspect-ratio: 1/1;
		display: flex;
		flex-direction: column;
		border-radius: 1.5vh;
		overflow: hidden;
		outline: 1px solid var(--white-color);

		.card__img {
			height: 60%;
			img {
				object-fit: cover;
				width: 100%;
			}
		}

		.card__content {
			padding: 5px;
			margin-top: auto;
			display: grid;
			// grid-auto-flow: column;
			grid-template-columns: 1fr auto;
			gap: 5px;
			height: 40%;
			background-color: var(--white-color);
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
					color: var(--very-light-secondary-color);
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

	// this class is for the like button
	.material-symbols-rounded.filled {
		font-variation-settings:
			"FILL" 1,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
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
