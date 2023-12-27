<script lang="ts">
	import { goto } from "$app/navigation";
	import { filterStore } from "../store";

	let name = "";
	let url = "";

	$: url = `/search${name ? `?name=${name}&` : "?"}page=0${
		$filterStore.price ? `&price=${$filterStore.price}` : ""
	}${
		$filterStore.difficulty ? `&difficulty=${$filterStore.difficulty}` : ""
	}`;

	export let hiddenMobile = false;
</script>

<div class={`nav__search ${hiddenMobile ? "hidden" : ""}`}>
	<input
		name="name"
		bind:value={name}
		placeholder="Rechercher..."
		type="search"
		on:keypress={(event) => (event.key === "Enter" ? goto(url) : null)}
	/>
	<a
		data-sveltekit-preload-data="tap"
		href={url}
		class="material-symbols-rounded">search</a
	>
</div>

<style lang="scss">
	.nav__search {
		display: flex;
		justify-content: center;
		align-items: center;
		text-align: center;

		input {
			background-color: var(--white-color);
			border: none;
			color: #646464;
			padding: 0.6rem 1rem;
			border-radius: 10px 0 0 10px;
			width: 40vw;
			transition: all 300ms;
			outline: 2px solid var(--light-secondary-color);
		}

		input:focus {
			outline: 2px solid var(--very-light-secondary-color);
		}

		input::-webkit-input-placeholder {
			font-weight: normal;
			color: var(--very-light-secondary-color);
		}

		input::-webkit-search-cancel-button {
			-webkit-appearance: none;
			height: 1.3em;
			width: 1.3em;
			border-radius: 50em;
			background: url("../lib/img/googlefont_icon_close.svg") no-repeat
				50% 50%;
			background-size: contain;
			opacity: 0;
			pointer-events: none;
		}

		input:focus::-webkit-search-cancel-button {
			opacity: 0.3;
			pointer-events: all;
			cursor: pointer;
		}

		a {
			padding: 0.47rem 1rem;
			text-decoration: none;
			border-radius: 0 10px 10px 0;
			border: none;
			color: var(--white-color);
			background-color: var(--primary-color);
			cursor: pointer;
		}
		display: flex;
	}

	.hidden {
		display: none;
	}

	@media (min-width: 768px) {
		.nav__search {
			display: flex;
		}
	}
</style>
