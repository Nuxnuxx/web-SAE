<script lang="ts">
	import FoodImage from "$lib/img/homepage_food.png";
	import Swiper from "../components/Swiper.svelte";
	import { userStore, userDataStore } from "../store";

	let NAME = "";
	if ($userStore) {
		NAME =
			$userDataStore.firstName.charAt(0).toUpperCase() +
			$userDataStore.firstName.slice(1);
	}
</script>

<svelte:head>
	<title>Accueil | PirateCook</title>
</svelte:head>

<div class="background">
	<img
		src={FoodImage}
		class="foodimage foodimage__small"
		alt="blurred fun display"
	/>
	<img src={FoodImage} class="foodimage foodimage__big" alt="fun display" />
	<div class="container">
		{#if $userStore}
			<div class="title">
				<!-- I know having a h1 in a h2 is weird but I need to do this in 
					order to have at least one h1 in the page -->
				<h1 id="first">
					Bonjour <span class="red__highlight">{NAME}</span>,</h1
				>
				<h2 id="second"
					>Qu'est ce qu'on <span class="styled__font">cuisine</span
					></h2
				>
				<h2 id="third">Aujourd'hui ?</h2>
			</div>
		{:else}
			<div class="title">
				<!-- I know having a h1 in a h2 is weird but I need to do this in 
					order to have at least one h1 in the page -->
				<h1 id="first">
					<span class="styled__font">PirateCook</span>, l'inspiration</h1
				>
				<h2 id="second">culinaire qui vous</h2>
				<h2 id="third"
					><span class="red__highlight">ressemble</span>.</h2
				>
			</div>
		{/if}
	</div>
</div>

<style lang="scss">
	@function power($number, $exponent) {
		$value: 1;

		@if $exponent > 0 {
			@for $i from 1 through $exponent {
				$value: $value * $number;
			}
		}

		@return $value;
	}

	.background {
		background: linear-gradient(
			160deg,
			var(--secondary-color) 0%,
			rgba(130, 211, 230, 0) 70%
		);
		position: relative;
		overflow: hidden;
		height: 100vh;
	}

	img.foodimage {
		position: absolute;
		object-fit: cover;

		&.foodimage__small {
			top: -3rem;
			right: -12.5vw;
			height: 25vw;
			filter: blur(1.5px);
		}

		&.foodimage__big {
			top: -1rem;
			right: 4vw;
			height: 47vw;
			object-fit: cover;
		}
	}

	.container {
		padding: 2rem 0;
		margin-left: 5%;

		.title {
			position: relative;
			z-index: 1;
			h1,
			h2 {
				font-size: 4rem;
				letter-spacing: -1.8px;
				margin: 1.5rem 0;
			}
			#first {
				padding-left: 0;
			}
			#second {
				padding-left: 4rem;
			}
			#third {
				padding-left: 8rem;
			}
			.red__highlight {
				color: var(--primary-color);
				font-weight: bold;
			}
			.styled__font {
				font-family: "Leckerli One", cursive;
				font-weight: 400;
				letter-spacing: 0;
			}
		}
	}

	@media screen and (max-width: 1160px) {
		img.foodimage.foodimage__small {
			display: none;
		}

		img.foodimage.foodimage__big {
			height: 57%;
			right: 0;
		}
	}

	@media screen and (max-width: 1024px) {
		img.foodimage.foodimage__big {
			height: 40%;
			top: 7rem;
			right: 0;
		}
		.container {
			.title {
				#second {
					padding-left: 0;
				}
				#third {
					padding-left: 0;
				}
			}
		}
	}

	@media screen and (max-width: 768px) {
		img.foodimage.foodimage__big {
			height: 40%;
			top: 7rem;
			right: 0;
		}
		.container {
			.title {
				h1,
				h2 {
					font-size: 3rem;
					margin: 1rem 0;
				}
			}
		}
	}

	@media screen and (max-width: 550px) {
		img.foodimage.foodimage__big {
			top: 6rem;
			height: max(15%, 40vw);
		}

		.container {
			.title {
				h1,
				h2 {
					font-size: max(1.8rem, 8vw);
				}
			}
		}
	}
</style>
