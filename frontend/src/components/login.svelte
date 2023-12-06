<script lang="ts">
	import type { ErrorsRegister, User } from "$lib/api/auth-types";

	export let errors: ErrorsRegister;
	export let value: User;
</script>

<h2>Quel plaisir de vous voir à nouveau !</h2>

<form method="post" action="?/login">
	{#if errors?.server}
		<span class="error">{errors?.server}</span>
	{/if}
	<label for="mail">
		<input
			id="mail"
			name="mail"
			type="text"
			placeholder="Mail*"
			value={value?.mail || ""}
			autocomplete="username"
		/>
		{#if errors?.mail}
			<span class="error">{errors?.mail}</span>
		{/if}
	</label>

	<label for="password">
		<input
			id="password"
			name="password"
			type="password"
			value={value?.password || ""}
			placeholder="Mot de passe*"
			autocomplete="current-password"
		/>
		{#if errors?.password}
			<span class="error">{errors?.password}</span>
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
		width: 80%;
		margin: 0 auto;
		margin-bottom: 2.8rem;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin: 0 auto;
		row-gap: 1.5rem;
		width: 40%;
		min-width: 350px;
		max-width: 700px;

		.error {
			color: var(--primary-color);
			margin-top: 0.1rem;
			font-size: 0.8rem;
			font-weight: bold;
			text-align: center;
		}

		label {
			width: 100%;
			display: flex;
			flex-direction: column;
			align-items: center;
			input {
				border-bottom: 1px solid var(--light-secondary-color);
				border-top: none;
				border-left: none;
				border-right: none;
				width: 80%;

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
