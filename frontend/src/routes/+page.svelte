<script lang="ts">
	import type { Recipe } from "../app";
	import CardList from "../components/cardList.svelte";

	const recipeList: Recipe[] = [
		{
			img: "https://assets.afcdn.com/recipe/20230608/143328_w600h600c1.webp",
			title: "Tarte aux pommes et chocolat Nestlé ®",
			nbLikes: 1245,
			liked: false,
			saved: false,
		},
		{
			img: "https://assets.afcdn.com/recipe/20130910/1418_w380h380c1cx1728cy2352.webp",
			title: "Coq au Vin",
			nbLikes: 789,
			liked: true,
			saved: false,
		},
		{
			img: "https://assets.afcdn.com/recipe/20161201/35212_w380h380c1cx2144cy1424.webp",
			title: "Ratatouille",
			nbLikes: 567,
			liked: false,
			saved: true,
		},
		{
			img: "https://assets.afcdn.com/recipe/20130910/6624_w380h380c1cx1872cy2808.webp",
			title: "Boeuf Bourguignon",
			nbLikes: 1023,
			liked: true,
			saved: true,
		},
		{
			img: "https://assets.afcdn.com/recipe/20130819/24146_w380h380c1cx2592cy1728.webp",
			title: "Crêpes",
			nbLikes: 876,
			liked: false,
			saved: false,
		},
		{
			img: "https://assets.afcdn.com/recipe/20150522/21358_w380h380c1cx1250cy1630.webp",
			title: "Salade Niçoise",
			nbLikes: 453,
			liked: true,
			saved: true,
		},
	];
</script>

<svelte:head>
	<title>Accueil | PirateCook</title>
</svelte:head>

<CardList listRecipe={recipeList} />

<h1>Welcome to SvelteKit</h1>
<p
	>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p
>
