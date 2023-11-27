<script>
	import Modify from "../../../components/modifyProfil.svelte";
</script>

<svelte:head>
	<title>Profil | PirateCook</title>
</svelte:head>

<div class="background">
	<div class="container">
		<Modify />
	</div>
</div>

<style>
	.background {
		margin-top: -3%;
		background-image: url("https://www.referenseo.com/wp-content/uploads/2019/03/image-attractive-960x540.jpg");
		background-size: cover;
	}

	.container {
		padding-top: 1px;
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