<script>
	let email = "";
	let firstName = "";
	let lastName = "";
	let password = "";

	function handleSubmit() {
		// Handle form submission here
		console.log("Form submitted");
		console.log("First Name:", firstName);
		console.log("Last Name:", lastName);
		console.log("Email:", email);
		console.log("Password:", password);
	}
</script>

<link
	rel="stylesheet"
	href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
/>

<main>
	<form on:submit|preventDefault={handleSubmit}>
		<div class="form-container">
			<fieldset>
				<legend>Nom</legend>
				<input
					type="text"
					id="lastName"
					bind:value={lastName}
					required
					placeholder="Nom*"
					autocomplete="family-name"
				/>
				<label for="lastName"></label>
			</fieldset>
		</div>

		<div>
			<label for="firstName"> </label>
			<input
				type="text"
				id="firstName"
				bind:value={firstName}
				required
				placeholder="Prénom*"
				autocomplete="given-name"
			/>
		</div>

		<div class="form-container">
			<fieldset>
				<legend>Mail</legend>
				<input
					type="email"
					name="email"
					id="email"
					bind:value={email}
					required
					placeholder="La sauce"
				/>
				<label for="email"></label>
			</fieldset>
		</div>

		<div>
			<fieldset>
				<legend>Mot de passe</legend>
				<input
					type="password"
					id="password"
					bind:value={password}
					required
					placeholder="Mot de passe*"
					autocomplete="new-password"
				/>
				<label for="password"></label>
			</fieldset>
		</div>

		<button type="submit">
			Sauvegarder
			<span class="material-symbols-rounded"> done </span>
		</button>
	</form>
</main>

<style>
	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 20px;
	}

	fieldset {
		border: 1px solid #ccc; /* Couleur de la bordure */
		border-radius: 40px; /* Rayon des coins arrondis */
		box-sizing: border-box; /* Garantit que la largeur inclut la bordure et le rembourrage */
	}

	legend {
		/*décaler vers la droite*/
		margin-left: 20px;
		opacity: 0.4;
	}

	label {
		display: flex;
		flex-direction: column;
		margin-bottom: 10px;
		/*min-width: 80%;*/
		/*width selon taille de l'écran*/
		width: 50%;
		min-width: 350px;
	}

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
	}

	/*rend le placeholder un peu plus transparent par défaut*/
	input::-webkit-input-placeholder {
		opacity: 0.4;
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
		font-variation-settings:
			"FILL" 0,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
	}
</style>
