<script lang="ts">
	import { setContext } from "svelte";

	let array = [
		"burger",
		"buritto",
		"cake",
		"carrot",
		"coffee",
		"cookies",
		"croissant",
		"cup cake",
		"donut",
		"egg and bacon",
		"eggplant",
		"float",
		"fries",
		"hot dog",
		"ice cream",
		"lolipop round",
		"lolipop swirl",
		"macaroon",
		"meat",
		"pancake",
		"pizza",
		"popsicle",
		"pretzel",
		"steak",
		"sushi caviar",
		"taco",
	];

	let random = Math.floor(Math.random() * array.length);
	let food = array[random];
	//FIXME: this is a workaround to get the image path, it may do nothing on a server but it throws an error on localhost
	const FoodImage = new URL(`./food/${food}.png`, import.meta.url).href;
	let backgroundColor: string;
	// change backgroundColor randomly while keeping same hue
	random = Math.floor(Math.random() * 100);
	backgroundColor = `${random}, 100%, 80%`;

	setContext("user", { food, backgroundColor });
</script>

<svelte:head>
	<title>Profil | PirateCook</title>
</svelte:head>

<div class="background" style="--theme-color: {backgroundColor}">
	<img src={FoodImage} class="foodimage foodimage__small" alt="fun display" />
	<img src={FoodImage} class="foodimage foodimage__big" alt="fun display" />
	<div class="container">
		<slot />
	</div>
</div>

<style lang="scss">
	.background {
		background: linear-gradient(
			160deg,
			hsl(var(--theme-color)) 0%,
			rgba(210, 130, 230, 0) 60%
		);
	}
	img.foodimage {
		position: absolute;
		object-fit: cover;

		&.foodimage__small {
			top: 4rem;
			right: 0rem;
			width: 15%;
			transform: rotateY(180deg);
			filter: blur(2.5px);
		}

		&.foodimage__big {
			top: 7rem;
			right: 5rem;
			height: 70%;
			object-fit: cover;
		}
	}

	.container {
		padding: 2rem 0;
		margin-left: 5%;
		width: 40%;
	}

	@media screen and (max-width: 1160px) {
		img.foodimage.foodimage__small {
			top: 6rem;
			width: 10%;
		}

		img.foodimage.foodimage__big {
			height: 60%;
		}
	}

	@media screen and (max-width: 768px) {
		.container {
			margin: auto;
			width: 80%;
		}
		img.foodimage {
			display: none;
		}
	}
</style>
