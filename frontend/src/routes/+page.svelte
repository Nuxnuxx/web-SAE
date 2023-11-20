<script>
	import RecipeCard from "../components/recipeCard.svelte";
</script>

<svelte:head>
	<title>Accueil | PirateCook</title>
</svelte:head>

<h2>Test liste éléments : </h2>
<div class="list">
	<RecipeCard />
	<RecipeCard />
	<RecipeCard />
	<RecipeCard />
</div>

<h1>Welcome to SvelteKit</h1>
<p
	>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p
>

<style>
	.list {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		grid-gap: 10px;
		grid-auto-rows: minmax(100px, auto);
	}
</style>
