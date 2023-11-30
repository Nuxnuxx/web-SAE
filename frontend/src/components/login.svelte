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
		color: var(--black-color);
		font-weight: normal;
		font-size: 1.4rem;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		row-gap: 1rem;

		label {
			min-width: 350px;
			max-width: 700px;
			width: 40%;
			display: flex;
			flex-direction: column;
			padding: 1rem 0;

			input {
				border-bottom: 1px solid var(--light-secondary-color);
				border-top: none;
				border-left: none;
				border-right: none;

				padding: 0.5rem 1rem;
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
			color: var(--light-black-color);
			text-decoration: underline;
			font-size: 0.9rem;
		}

		button {
			display: flex;
			align-items: center;
			gap: 0.5rem;
			padding: 0.5rem 1rem;
			border-radius: 2rem;
			background-color: var(--primary-color);
			color: var(--white-color);
			font-size: 1rem;
			border: none;
			cursor: pointer;
		}

		.material-symbols-rounded {
			font-size: 1.2rem;
			font-weight: bold;
		}
	}
</style>
