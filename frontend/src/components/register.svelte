<script lang="ts">
	let email = "";
	let firstName = "";
	let lastName = "";
	let password = "";
	let confirmPassword = "";
	let gender = "";

	function handleSubmit() {
		// Handle form submission here
		console.log("Form submitted");
		console.log("Email:", email);
		console.log("First Name:", firstName);
		console.log("Last Name:", lastName);
		console.log("Password:", password);
		console.log("Confirm Password:", confirmPassword);
		console.log("Gender:", gender);
	}
</script>

<h2
	>Rejoignez nous, et trouver l'inspiration<br />culinaire qui
	<span class="highlight">vous</span> ressemble</h2
>

<form on:submit|preventDefault={handleSubmit}>
	<div>
		<label for="email"></label>
		<input
			type="email"
			id="email"
			bind:value={email}
			required
			placeholder="Email*"
			autocomplete="username"
		/>
	</div>

	<div class="name">
		<div>
			<label for="lastName"></label>
			<input
				type="text"
				id="lastName"
				bind:value={lastName}
				required
				placeholder="Nom*"
				autocomplete="family-name"
			/>
		</div>

		<div>
			<label for="firstName"></label>
			<input
				type="text"
				id="firstName"
				bind:value={firstName}
				required
				placeholder="Prénom*"
				autocomplete="given-name"
			/>
		</div>
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="password"
			bind:value={password}
			required
			placeholder="Mot de passe*"
			autocomplete="new-password"
		/>
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="confirmPassword"
			bind:value={confirmPassword}
			required
			placeholder="Confirmer mot de passe*"
			autocomplete="new-password"
		/>
	</div>

	<div>
		<p class="genre">Je suis* :</p>
	</div>

	<div class="radio-container">
		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="male" name="radio-group" type="radio" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Un pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="female" name="radio-group" type="radio" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Une pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="other" name="radio-group" type="radio" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Autre</span>
			</label>
		</div>
	</div>

	<button type="submit">
		Connexion
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
		width: 40%;
		margin: 0 auto;
		min-width: 350px;
		max-width: 700px;

		span {
			.material-symbols-rounded {
				font-size: 1.2rem;
				font-weight: bold;
			}
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

		.genre {
			color: var(--light-black-color);
		}

		div {
			width: 100%;
			display: flex;
			flex-direction: column;
			align-items: center;
		}

		input {
			&:not([type="checkbox"]) {
				border: 1px solid #ccc;
				border-top: none;
				border-left: none;
				border-right: none;

				padding: 0.5rem 1rem;
				outline: none;

				width: 80%;
			}
			&::placeholder {
				opacity: 0.4;
				transition: 0.5s;
			}

			&:focus::placeholder {
				color: transparent;
			}
		}

		.name {
			display: flex;
			flex-direction: row;
			justify-content: space-between;
		}

		.radio-container {
			display: flex;
			flex-direction: row;
			justify-content: space-around;
			color: var(--light-black-color);

			.radio-wrapper {
				margin: 20px 0 20px 0;
				min-width: fit-content;
				display: flex;
				justify-content: space-around;

				.radio-button input[type="radio"] {
					display: none;
				}

				.radio-button input[type="radio"]:checked ~ .radio-checkmark {
					border: 2px solid var(--primary-color);
				}

				.radio-button
					input[type="radio"]:checked
					~ .radio-checkmark:before {
					transform: translate(-50%, -50%) scale(1);
				}

				.radio-button {
					display: flex;
					align-items: center;
					cursor: pointer;
					transition: all 0.2s ease-in-out;

					.radio-checkmark {
						display: inline-block;
						position: relative;
						width: 16px;
						height: 16px;
						margin-right: 10px;
						border: 2px solid #ccc;
						border-radius: 30%;
					}
					.radio-checkmark:before {
						content: "";
						position: absolute;
						top: 50%;
						left: 50%;
						transform: translate(-50%, -50%) scale(0);
						width: 8px;
						height: 8px;
						border-radius: 30%;
						background-color: var(--primary-color);
						transition: all 0.2s ease-in-out;
					}
				}
				.radio-button:hover {
					transform: translateY(-2px);
				}
			}
		}
	}
</style>
