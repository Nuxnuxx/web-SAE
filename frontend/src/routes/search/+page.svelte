<script lang="ts">
	import CardList from "../../components/cardList.svelte";
	import Filter from "../../components/filter.svelte";
	import SearchBar from "../../components/searchBar.svelte";
	import type { PageLoad } from "../playlist/[slug]/$types";
	export let data: PageLoad;
</script>

<svelte:head>
	<title>Recherche | PirateCook</title>
</svelte:head>

<div class="filter">
	<h2>Recherche</h2>
	<SearchBar hiddenMobile={false} />
	<Filter />
	<CardList listRecipe={data.result.result} />
</div>

<style lang="scss">
	.filter {
		width: 90%;
		@media (min-width: 768px) {
			width: 50%;
		}
		margin: 0 auto;
		text-align: center;
		h2 {
			font-size: 3rem;
		}
	}
</style>
