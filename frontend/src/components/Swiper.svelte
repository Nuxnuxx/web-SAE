<script lang="ts">
	// need to do it to register the modules
	import { register } from "swiper/element/bundle";
	register();

	import type { RecipeDetail } from "$lib/api/recipe-types";
	import RecipeCard from "./recipeCard.svelte";
	import "swiper/scss";
	import { onMount } from "svelte";

	export let recipes: RecipeDetail[];
</script>

<swiper-container
	class="swiper"
	slides-per-view="2"
	slides-per-group="1"
	speed="500"
	loop="true"
	css-mode="true"
	navigation="true"
	breakpoints={{
		768: {
			slidesPerView: 4,
			slidesPerGroup: 1,
			spaceBetween: 20,
			loop: true,
		},
	}}
>
	{#each recipes as recipe}
		<swiper-slide>
			<RecipeCard data={recipe} />
		</swiper-slide>
	{/each}
</swiper-container>
<button id="my-prev-button" class="navigate material-symbols-rounded"
	>navigate_before</button
>
<button id="my-next-button" class="navigate material-symbols-rounded"
	>navigate_next</button
>

<style lang="scss">
	.swiper__wrapper {
		position: relative;
	}

	.swiper {
		padding: 2rem 3rem;

		@media screen and (max-width: 768px) {
			padding: 2rem 1rem;
		}
	}

	:root {
		--swiper-navigation-top-offset: 35%;
		--swiper-navigation-sides-offset: 1.5rem;
		@media screen and (max-width: 768px) {
			--swiper-navigation-sides-offset: 0;
		}
	}

	.navigate {
		text-align: center;
		position: absolute;
		color: var(--black-color);
		font-size: 2rem;
		background: var(--white-color);
		aspect-ratio: 1;
		z-index: 20;
		border: none;
		border-radius: 50%;
		&#my-prev-button {
			top: var(--swiper-navigation-top-offset);
			left: var(--swiper-navigation-sides-offset);
		}
		&#my-next-button {
			top: var(--swiper-navigation-top-offset);
			right: var(--swiper-navigation-sides-offset);
		}
	}
</style>
