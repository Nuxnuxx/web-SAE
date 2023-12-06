<script lang="ts">
	import Recipe from "../../../components/recipe.svelte";
	import type { PageData } from "./$types";

	export let data: PageData;
</script>

<svelte:head>
	<title>{data.recipe.result.recipeDetail.name} | PirateCook</title>
</svelte:head>

<Recipe recipe={data.recipe.result} />
