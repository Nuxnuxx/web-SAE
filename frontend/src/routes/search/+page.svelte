<script lang="ts">
	import CardList from "../../components/cardList.svelte";
	import Filter from "../../components/filter.svelte";
	import type { PageLoad } from "../playlist/[slug]/$types";
	export let data: PageLoad;
</script>

<svelte:head>
	<title>Recherche | PirateCook</title>
</svelte:head>

<div class="filter">
	<Filter />
	<CardList listRecipe={data.result.result} />
</div>

<style lang="scss">
	.filter {
		width: 50%;
		margin: 0 auto;
	}
</style>
