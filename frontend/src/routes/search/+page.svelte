<script lang="ts">
	import type { PageData } from "../$types";
	import CardList from "../../components/cardList.svelte";
	import Filter from "../../components/filter.svelte";
	import Pagination from "../../components/pagination.svelte";
	import SearchBar from "../../components/searchBar.svelte";
	export let data: PageData;
</script>

<svelte:head>
	<title>Recherche | PirateCook</title>
</svelte:head>

<div class="filter">
	<h2>Recherche</h2>
	<SearchBar hiddenMobile={false} />
	<Filter />
	<CardList listRecipe={data.result.result} />
	<Pagination current={data.oldData} pagination={data.result.metadata} />
</div>

<style lang="scss">
	.filter {
		padding: 2rem 0;
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
