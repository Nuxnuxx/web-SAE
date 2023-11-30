<script lang="ts">
	// let array = ["burger","buritto","cake","carrot","coffee","cookies","croissant","cup cake","donut",
	//             "egg and bacon","eggplant","float","fries","hot dog","ice cream","lolipop round",
	//             "lolipop swirl","macaroon","meat","pancake","pizza","popsicle","pretzel","steak",
	//             "sushi caviar","taco"];
	// // select randomly one item in array
	// let random = Math.floor(Math.random() * array.length);
	// let food = array[random] + ".png";
	// let FoodImage: string;
	// import foodFolder from "$lib/img/profile/food/";
	// import(`${foodFolder}${food}`).then(module => {
	//     FoodImage = module.default;
	// });

	import FoodImage from "$lib/img/profile/food/coffee.png";
</script>

<svelte:head>
	<title>Profil | PirateCook</title>
</svelte:head>

<div class="background">
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
			#d282e6 0%,
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
			right: 4rem;
			height: 80%;
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
