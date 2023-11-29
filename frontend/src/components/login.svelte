<script lang="ts">
	import { goto } from "$app/navigation";
	import { sendLogin } from "$lib/api/auth-request";
	import { schemaLogin } from "$lib/api/auth-schema";
	import type { ErrorsRegister } from "$lib/api/auth-types";
	import * as yup from "yup";

	let errors: ErrorsRegister = {};

	let loginValues = {
		email: "",
		password: "",
	};

	async function handleLogin() {
		try {
			await schemaLogin.validate(loginValues, {
				abortEarly: false,
			});
			const result = await sendLogin(loginValues);
			if (result) {
				localStorage.setItem("user", JSON.stringify(result));
				goto("/");
			}
		} catch (err) {
			if (err instanceof yup.ValidationError) {
				errors = err.inner.reduce((acc, err) => {
					const key = String(err.path);
					return { ...acc, [key]: err.message };
				}, {});
			} else {
				if (err instanceof Error) {
					errors = { server: err.message };
				}
			}
		}
	}
</script>

<h2>Quel plaisir de vous voir à nouveau !</h2>

{#if errors.server}
	<span class="error">{errors.server}</span>
{/if}

<form on:submit|preventDefault={handleLogin}>
	<label for="email">
		<input
			id="email"
			type="text"
			bind:value={loginValues.email}
			placeholder="Email"
			autocomplete="username"
		/>
		{#if errors.email}
			<span class="error">{errors.email}</span>
		{/if}
	</label>

	<label for="password">
		<input
			id="password"
			type="password"
			bind:value={loginValues.password}
			placeholder="Mot de passe"
			autocomplete="current-password"
		/>
		{#if errors.password}
			<span class="error">{errors.password}</span>
		{/if}
	</label>

	<!-- TODO: Ajouter un lien vers la page de réinitialisation du mot de passe -->
	<a href="/">Mot de passe oublié ?</a>

	<button type="submit"
		>Connexion
		<span class="material-symbols-rounded"> arrow_forward </span>
	</button>
</form>

<style lang="scss">
	h2 {
		text-align: center;
		color: #333;
		font-weight: normal;
		font-size: 22px;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 20px;

		label {
			min-width: 350px;
			max-width: 700px;
			width: 40%;
			display: flex;
			flex-direction: column;
			margin-bottom: 10px;
			padding: 20px;

			input {
				border: 1px solid #ccc;
				border-radius: 3%;
				border-top: none;
				border-left: none;
				border-right: none;

				/*padding du texte*/
				padding: 5px 15px;
				font-size: 16px;
				outline: none;
				&::placeholder {
					opacity: 0.4;
					transition: 0.5s;
				}

				&:focus::placeholder {
					color: transparent;
				}
			}
		}
		a {
			text-decoration: none;
			color: #9f9f9f;
			text-decoration: underline;
			font-size: 14px;
			margin-top: -20px;
			margin-bottom: 25px;
			margin-right: -40%;
		}

		button {
			padding: 10px 20px;
			background-color: #de403e;
			color: #fff;
			border: none;
			border-radius: 20px;
			cursor: pointer;
			font-size: 16px;
			display: flex;
			align-items: center;
		}

		.material-symbols-rounded {
			font-size: 20px;
			font-weight: bold;
			margin-left: 10px;
			margin-right: -10px;
		}
	}
</style>
