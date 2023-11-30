<script lang="ts">
	import Register from "../../components/register.svelte";
	import Login from "../../components/login.svelte";

	let selectedButton = "login";

	let bool = true;

	function OnClick(event: any) {
		console.log(event.target.className);
		if (event.target.classList.contains("login")) {
			bool = true;
			selectedButton = "login";
		} else if (event.target.classList.contains("register")) {
			bool = false;
			selectedButton = "register";
		}
	}
</script>

<svelte:head>
	<title>Authentification | PirateCook</title>
</svelte:head>

<div class="auth">
	<div class="button-wrapper">
		<button
			on:click={OnClick}
			class:active={selectedButton === "login"}
			class="login">Se connecter</button
		>
		<button
			on:click={OnClick}
			class:active={selectedButton === "register"}
			class="register">S'inscrire</button
		>
	</div>
</div>

{#if selectedButton == "login"}
	<Login />
{:else}
	<Register />
{/if}

<style lang="scss">
	.auth {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin: 20px 0 20px 0;
	}

	button {
		/*border uniquement en bas*/
		border-top: none;
		border-left: none;
		border-right: none;
		background-color: transparent;
		cursor: pointer;

		border-bottom: 2px solid #000;

		opacity: 0.3;

		font-size: 18px;
	}

	.login {
		margin-right: 100px;
	}

	.active {
		border-bottom-color: red;
		color: red;
		font-weight: bold;
		opacity: 1;
	}
</style>
