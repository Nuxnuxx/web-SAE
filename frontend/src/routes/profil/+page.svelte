<script>
	import Modify from "../../components/modifyProfil.svelte";
	import Show from "../../components/showProfile.svelte";

	let modify = false;
</script>

<svelte:head>
	<title>Profil | PirateCook</title>
</svelte:head>

<div class="background">
	<div class="container">
		{#if modify}
			<Modify />
		{:else}
			<Show />
		{/if}
	</div>
</div>

<style>
	.background {
		background-image: url("https://www.referenseo.com/wp-content/uploads/2019/03/image-attractive-960x540.jpg");
		background-size: cover;
	}

	.container {
		padding-top: 1%;
		padding-bottom: 3%;
		margin-left: 5%;
		width: 40%;
	}

	@media screen and (max-width: 768px) {
		.container {
			width: 80%;
		}
	}
</style>
