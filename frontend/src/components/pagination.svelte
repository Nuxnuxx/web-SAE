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

	//FIXME: a refacto bordel je chiale y'a des -1 partout + cette ternaire mama
	//FIXME2: j'ai tenté mettre une const pour rendre le tout plus claire
	const nbItemPerPage: number = 5;

	$: computedesenfer =
		pagination.totalPage % nbItemPerPage == 0
			? pagination.totalPage - nbItemPerPage + 1
			: pagination.totalPage - (pagination.totalPage % nbItemPerPage) - 1;

	$: start = Math.floor(current.page / nbItemPerPage) * nbItemPerPage;
	$: end = Math.min(start + nbItemPerPage, pagination.totalPage);
</script>

{#if pagination.totalPage > 1}
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
			<a
				href={current.page < pagination.totalPage - 10
					? `${$urlStore}page=${pagination.currentPage + 10}`
					: `${$urlStore}page=${current.page}`}>...</a
			>
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
{/if}

<style lang="scss">
	.pagination__wrapper {
		width: 90%;
		@media (min-width: 768px) {
			width: 70%;
		}
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
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
				&:first-child {
					margin-right: 2rem;
				}
				&:last-child {
					margin-left: 2rem;
				}
			}
		}
	}
</style>
