<script lang="ts">
	import Filter from '../../components/filter.svelte';
</script>

<svelte:head>
	<title>Recherche | PirateCook</title>
</svelte:head>

<h1>Permet de chercher des recherche selon SearchBar / Filter</h1>

<div class="filter">
	<Filter />
</div>

<style lang="scss">
	.filter {
		width: 50%;
		margin: 0 auto;
	}
</style>