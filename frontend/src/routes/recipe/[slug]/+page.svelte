<script>
	export let data;

	import Recipe from '../../../components/recipe.svelte';
</script>

<svelte:head>
	<title>Recette {data.slug} | PirateCook</title>
</svelte:head>

<Recipe />