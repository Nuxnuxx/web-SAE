<script lang="ts">
	import type { Pagination, Current } from "$lib/api/recipe-types";
	import { filterStore, urlStore } from "../store";
	export let pagination: Pagination;
	export let current: Current;

	$: {
		let newUrl = "search?";
		if (current.name) {
			newUrl += `name=${current.name}&`;
		}

		if ($filterStore.difficulty) {
			newUrl += `difficulty=${$filterStore.difficulty}&`;
		}

		if ($filterStore.price) {
			newUrl += `price=${$filterStore.price}&`;
		}

		urlStore.set(newUrl);
	}
	//TODO: condition if there is one page dont render the components

	//FIXME: a refacto bordel je chiale y'a des -1 partout + cette ternaire
	$: computedesenfer =
		pagination.totalPage % 5 == 0
			? pagination.totalPage - 6
			: pagination.totalPage - (pagination.totalPage % 5) - 1;

	$: start = Math.floor(current.page / 5) * 5;
	$: end = Math.min(start + 5, pagination.totalPage);
</script>

<div class="pagination__wrapper">
	<a
		class="material-symbols-rounded"
		class:active={true}
		href={`${$urlStore}page=${
			current.page === 0 ? pagination.totalPage - 1 : current.page - 1
		}`}
	>
		chevron_left
	</a>
	{#each Array.from({ length: end - start }, (_, i) => i + start) as i}
		<a href={`${$urlStore}page=${i}`} class:active={i === current.page}
			>{i + 1}</a
		>
	{/each}
	{#if current.page <= computedesenfer}
		<a>...</a>
		<a href={`${$urlStore}page=${pagination.totalPage - 1}`}
			>{pagination.totalPage}</a
		>
	{/if}
	<a
		class="material-symbols-rounded"
		class:active={true}
		href={`${$urlStore}page=${
			current.page === pagination.totalPage - 1 ? 0 : current.page + 1
		}`}
	>
		chevron_right
	</a>
</div>

<style lang="scss">
	.pagination__wrapper {
		width: 90%;
		@media (min-width: 768px) {
			width: 70%;
		}
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		margin: 4rem auto;
		a {
			padding: auto;
			font-size: 1rem;
			font-weight: bold;
			text-decoration: none;
			color: black;
			aspect-ratio: 1;
			height: 2rem;
			line-height: 2rem;
			@media (min-width: 768px) {
				height: 3rem;
				line-height: 3rem;
				font-size: 1.2rem;
			}
		}
		.active {
			color: white;
			background-color: var(--secondary-color);
			border-radius: 100%;
		}
		&:first-child,
		&:last-child {
			.material-symbols-rounded {
				font-size: 1.7rem;
			}
		}
	}
</style>
