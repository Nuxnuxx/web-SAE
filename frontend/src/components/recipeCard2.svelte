<script lang="ts">
	import type { RecipeDetail } from "$lib/api/recipe-types";
	import readable from "readable-numbers";
	import AddToPlaylist from "./addToPlaylist.svelte";
	import AddToLiked from "./addToLiked.svelte";
	import DEFAULT from "$lib/img/sample.png";
	import { coldstartLiked } from "../store";
	import Coldstart from "./coldstart.svelte";

	export let data: RecipeDetail;
	$: finalArrayImages = data.images.replace(/[\[\]"]+/g, "").split(", ");

	let idRecipe = data.idRecipe;
	let isIn = false;

	const handleLike = () => {
		$coldstartLiked.forEach((element) => {
			if (element == idRecipe) {
				isIn = true;
			}
		});

		if (!isIn) {
			$coldstartLiked = [...$coldstartLiked, idRecipe];
		} else {
			$coldstartLiked.splice($coldstartLiked.indexOf(idRecipe), 1);
		}
	};
</script>

<div class="card">
	<!--faire une bordure bleu si dans la liste-->

	<div class="card__img">
		<img
			src={finalArrayImages[0] == "" ? DEFAULT : finalArrayImages[0]}
			alt={data.name}
		/>
	</div>

	<div class="card__content">
		<h3>{data.name}</h3>
		<input type="checkbox" on:click={handleLike} />
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
		outline: 1px solid var(--white-color);

		.card__img {
			height: 80%;
			border-radius: 1.5vh 1.5vh 0.7vh 0.7vh;
			overflow: hidden;
			img {
				object-fit: cover;
				width: 100%;
			}
		}

		.card__content {
			padding: 5px;
			margin-top: auto;
			display: flex;
			flex-direction: row;
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
