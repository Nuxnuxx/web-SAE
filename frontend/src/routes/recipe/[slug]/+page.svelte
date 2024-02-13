<script lang="ts">
	import Swiper from "../../../components/Swiper.svelte";
	import Recipe from "../../../components/recipe.svelte";
	import type { PageData } from "./$types";

	export let data: PageData;
</script>

<svelte:head>
	<title>{data.recipe.result.recipeDetail.name} | PirateCook</title>
</svelte:head>

<Recipe recipe={data.recipe.result} />
<Swiper recipes={data.similarRecipe.result} />
