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

		// href to /profil
		window.location.href = "/profil";
	}
</script>

<link
	rel="stylesheet"
	href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
/>

<main>
	<div class="card">
		<!--image de profil-->
		<div class="card__img">
			<img src="https://picsum.photos/200" alt="profile picture" />
		</div>

		<form on:submit|preventDefault={handleSubmit} class="card__form" action="/profil">
			<div class="form__div">
				<fieldset class="form__field">
					<legend class="form__legend">Prénom</legend>
					<input
						type="text"
						id="firstName"
						bind:value={firstName}
						placeholder="Damn"
						class="form__input"
					/>
					<label for="firstName" class="form__label"> </label>
				</fieldset>
			</div>

			<div class="form__div">
				<fieldset class="form__field">
					<legend class="form__legend">Nom</legend>
					<input
						type="text"
						id="lastName"
						bind:value={lastName}
						placeholder="Mister P."
						class="form__input"
					/>
					<label for="lastName" class="form__label"></label>
				</fieldset>
			</div>

			<div class="form__div">
				<fieldset class="form__field">
					<legend class="form__legend">Mail</legend>
					<input
						type="email"
						name="email"
						id="email"
						bind:value={email}
						placeholder="lasauce@gmail.com"
						class="form__input"
					/>
					<label for="email" class="form__label"></label>
				</fieldset>
			</div>

			<div class="form__div">
				<fieldset class="form__field">
					<legend class="form__legend">Mot de passe</legend>
					<input
						type="password"
						id="password"
						bind:value={password}
						placeholder="oui"
						class="form__input"
					/>
					<label for="password" class="form__label"></label>
				</fieldset>
			</div>

			<button type="submit" class="form__button">
				<span class="button__text">Sauvegarder</span>
				<span class="material-symbols-rounded"> done </span>
			</button>
		</form>
	</div>
</main>

<style lang="scss">
	.card {
		border: 1px solid #ccc;
		border-radius: 20px;
		width: auto;
		/*ne peux pas etre plus petit que les enfants*/
		overflow: auto;
		border-radius: 20px;
		background: #f5f5f5;
		box-shadow: 0px 0px 10px 0px #dcdcdc;

		.card__img {
			width: 25%;
			height: 100%;
			border-radius: 100%;
			overflow: hidden;
			margin-top: 40px;
			/*centrer l'image*/
			margin-left: auto;
			margin-right: auto;

			img {
				width: 100%;
				height: 100%;
				object-fit: cover;
				vertical-align: middle;
			}
		}

		.card__form {
			display: flex;
			flex-direction: column;
			margin-right: 20%;
			margin-left: 20%;
			margin-top: 20px;

			.form__div {
				padding-bottom: 10px;
			}

			.form__field {
				border: 1px solid #ccc; /* Couleur de la bordure */
				border-radius: 20px; /* Rayon des coins arrondis */
			}

			.form__legend {
				/*décaler vers la droite*/
				margin-left: 10px;
				opacity: 0.3;
			}

			.form__label {
				display: flex;
				flex-direction: column;
				margin-bottom: 10px;
			}

			.form__input {
				border: none;
				background: #f5f5f5;
				
			}

			.form__input::placeholder {
				opacity: 0.4;
			}

			.form__button {
				margin-bottom: 30px;
				margin-left: 20%;
				margin-right: 20%;

				padding: 10px 20px;
				background-color: #de403e;
				color: #fff;
				border: none;
				border-radius: 20px;
				cursor: pointer;
				font-size: 20px;
				display: flex;
				align-items: center;
				justify-content: space-between;

				@media screen and (max-width: 1150px) {
					font-size: 16px;
				}

				.button__text {
					@media screen and (max-width: 1100px) {
						display: none;
					}
				}
			}
		}
	}

	.material-symbols-rounded {
		font-variation-settings:
			"FILL" 0,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;

		font-weight: bold;
		margin-left: auto;
		margin-right: auto;
	}
</style>
