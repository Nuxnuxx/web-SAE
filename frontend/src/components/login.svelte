<script lang="ts">
	import type { ActionData } from "../routes/auth/$types";

	export let form: ActionData;
</script>

<h2>Quel plaisir de vous voir à nouveau !</h2>

<form method="post" action="?/login">
	{#if form?.server}
		<span class="error">{form?.server}</span>
	{/if}
	<label for="mail">
		<input
			id="mail"
			name="mail"
			type="text"
			placeholder="mail*"
			autocomplete="username"
		/>
		{#if form?.mail}
			<span class="error">{form?.mail}</span>
		{/if}
	</label>

	<label for="password">
		<input
			id="password"
			name="password"
			type="password"
			placeholder="Mot de passe*"
			autocomplete="current-password"
		/>
		{#if form?.password}
			<span class="error">{form?.password}</span>
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
