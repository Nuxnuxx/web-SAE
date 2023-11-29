<script lang="ts">
	import LOGO from "$lib/img/LOGO.png";
	import { onMount } from "svelte";

	let user = false;

	/////////////////////
	// Header behavior //
	/////////////////////
	// altern between nav-up and nav-down classes on nav element on scroll down in svelte

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
	<div class="nav__search">
		<input placeholder="Rechercher..." type="search" />
		<button class="material-symbols-rounded">search</button>
	</div>
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
		<button class="nav__login">
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

		a {
			img {
				height: 2.5rem;
			}
		}

		.nav__search {
			display: flex;
			align-items: center;
			justify-content: space-between;
			text-align: center;
			input {
				background-color: var(--white-color);
				border: none;
				color: #646464;
				padding: 0.6rem 1rem;
				border-radius: 10px 0 0 10px;
				width: 40vw;
				transition: all 300ms;
				outline: 2px solid #dcdcdc;
			}

			input:focus {
				outline: 2px solid #c4c4c4;
			}

			input::-webkit-input-placeholder {
				font-weight: normal;
				color: #c4c4c4;
			}

			input::-webkit-search-cancel-button {
				-webkit-appearance: none;
				height: 1.3em;
				width: 1.3em;
				border-radius: 50em;
				background: url("../lib/img/googlefont_icon_close.svg")
					no-repeat 50% 50%;
				background-size: contain;
				opacity: 0;
				pointer-events: none;
			}

			input:focus::-webkit-search-cancel-button {
				opacity: 0.3;
				pointer-events: all;
				cursor: pointer;
			}

			button {
				padding: 0.47rem 1rem;
				border-radius: 0 10px 10px 0;
				border: none;
				color: var(--white-color);
				background-color: var(--primary-color);
				cursor: pointer;
			}
			display: none;
		}

		.nav__iconwrapper {
			display: flex;
			justify-content: space-around;
			align-items: center;
			gap: 2rem;
			.nav__iconwrapper__icon {
				display: contents;
				color: #000;
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
			.nav__search {
				display: flex;
			}

			a {
				img {
					height: 3.5rem;
				}
			}
		}
	}
	@media (min-width: 1024px) {
		nav {
			.nav__search {
				display: flex;
			}
		}
	}
</style>
