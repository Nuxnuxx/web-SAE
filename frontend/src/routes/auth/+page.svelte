<script lang="ts">
	import Register from "../../components/register.svelte";
	import Login from "../../components/login.svelte";
	import type { ActionData } from "./$types";
	import type { ErrorsRegister, User } from "$lib/api/auth-types";

	let selectedButton = "login";

	function OnClick(event: any) {
		if (event.target.classList.contains("login")) {
			selectedButton = "login";
		} else if (event.target.classList.contains("register")) {
			selectedButton = "register";
		}
	}

	export let form: ActionData;
	if (form == undefined) {
		form = {
			id: "",
			errors: {},
			values: {
				mail: "",
				password: "",
			},
		};
	}

	let registerErrors: ErrorsRegister = {};
	let registerValue: User;
	let loginErrors: ErrorsRegister = {};
	let loginValue: User;
	if (form.id == "register") {
		selectedButton = "register";
		registerErrors = form.errors;
		registerValue = form.values;
	} else if (form.id == "login") {
		selectedButton = "login";
		loginErrors = form.errors;
		loginValue = form.values;
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
	<Login value={loginValue} errors={loginErrors} />
{:else}
	<Register value={registerValue} errors={registerErrors} />
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
