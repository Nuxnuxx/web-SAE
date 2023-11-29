<script>
	import { goto } from "$app/navigation";

	let email = "";
	let password = "";

	function handleLogin() {
		// TODO:Ajoutez ici la logique de connexion
		console.log("Email:", email);
		console.log("Mot de passe:", password);
		goto("/");
	}
</script>

<h2>Quel plaisir de vous voir à nouveau !</h2>

<form on:submit|preventDefault={handleLogin}>
	<label for="email">
		<input
			id="email"
			type="email"
			bind:value={email}
			required
			placeholder="Email"
			autocomplete="username"
		/>
	</label>

	<label for="password">
		<input
			id="password"
			type="password"
			bind:value={password}
			required
			placeholder="Mot de passe"
			autocomplete="current-password"
		/>
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
