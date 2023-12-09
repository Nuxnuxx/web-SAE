<script lang="ts">
	import { goto } from "$app/navigation";
	import LOGO from "$lib/img/LOGO.png";
	import { onMount } from "svelte";
	import SearchBar from "./searchBar.svelte";

	export let user = false;

	let didScroll: boolean = false;
	let lastScrollTop = 0;
	let delta = 5;

	// get nav element
	let nav: HTMLElement;
	onMount(() => {
		if (nav.style) {
			// get header height
			const headerHeight = nav.offsetHeight;
			// get scroll position
			let scrollPos = 0;

			setInterval(function () {
				if (didScroll && nav.style) {
					hasScrolled();
					didScroll = false;
				}
			}, 250);

			// add scroll event listener
			window.addEventListener("scroll", function () {
				didScroll = true;
			});
			function hasScrolled() {
				scrollPos = window.scrollY;

				// Make sure they scroll more than delta
				if (Math.abs(lastScrollTop - scrollPos) <= delta) return;

				scrollPos > lastScrollTop && scrollPos > headerHeight
					? (nav.style.cssText = "top : -" + headerHeight + "px;")
					: (nav.style.cssText = "top : 0;");
				lastScrollTop = scrollPos;
			}
		}
	});
</script>

<nav bind:this={nav}>
	<a href="/"><img alt="PirateCook" src={LOGO} /></a>
	<SearchBar hiddenMobile={true} />
	{#if user}
		<div class="nav__iconwrapper">
			<a href="/playlist" class="nav__iconwrapper__icon"
				><span class="material-symbols-rounded">book_2</span></a
			>
			<a href="/profil" class="nav__iconwrapper__icon"
				><span class="material-symbols-rounded user">person</span></a
			>
		</div>
	{:else}
		<button on:click={() => goto("/auth")} class="nav__login">
			<span class="material-symbols-rounded">person</span>
			Connexion
		</button>
	{/if}
</nav>

<style lang="scss">
	nav {
		height: 5rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 0 1rem;
		background-color: var(--white-color);
		box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.25);

		// In-out header
		position: fixed;
		top: 0;
		width: fill-available;
		transition: top 0.2s ease-in-out;
		z-index: 10;

		a {
			img {
				height: 2.5rem;
			}
		}

		.nav__iconwrapper {
			display: flex;
			justify-content: space-around;
			align-items: center;
			gap: 2rem;
			.nav__iconwrapper__icon {
				display: contents;
				color: var(--black-color);
				:hover {
					color: var(--primary-color);
				}
				.material-symbols-rounded {
					font-size: 2rem;
				}
				.material-symbols-rounded.user {
					font-size: 2.4rem;
				}
			}
		}

		.nav__login {
			display: flex;
			align-items: center;
			gap: 0.5rem;
			padding: 0.5rem 1rem;
			border-radius: 1rem;
			background-color: var(--primary-color);
			color: var(--white-color);
			font-weight: medium;
			font-size: 1rem;
			border: none;
			cursor: pointer;
			transition: all 0.2s ease-out;

			&:hover {
				transform: scale(1.02);
			}
		}
	}
	@media (min-width: 768px) {
		nav {
			a {
				img {
					height: 3.5rem;
				}
			}
		}
	}
</style>
