<script lang="ts">
	// need to do it to register the modules
	import { register } from "swiper/element/bundle";
	register();

	import type { RecipeDetail } from "$lib/api/recipe-types";
	import RecipeCard from "./recipeCard.svelte";

	export let recipes: RecipeDetail[];
</script>

<swiper-container
	class="swiper"
	slides-per-view="2"
	slides-per-group="2"
	space-between="20"
	loop={true}
	css-mode={true}
	navigation={true}
	breakpoints={JSON.stringify({
		768: { slidesPerView: 4, slidesPerGroup: 4, spaceBetween: 30 },
	})}
>
	{#each recipes as recipe}
		<swiper-slide><RecipeCard data={recipe} /></swiper-slide>
	{/each}
</swiper-container>

<style lang="scss">
	.swiper {
		padding: 2rem 1rem;
	}
</style>
