<script lang="ts">
	import { goto } from "$app/navigation";
	import { sendRegister } from "$lib/api/auth-request";

	let email = "";
	let firstName = "";
	let lastName = "";
	let password = "";
	let confirmPassword = "";
	let gender = "";

	async function handleRegister() {
		try {
			const result = await sendRegister(
				firstName,
				lastName,
				email,
				password
			);
			if (result) {
				localStorage.setItem("user", JSON.stringify(result));
				goto("/");
			}
		} catch (err) {
			console.log(err);
		}
	}
</script>

<h2
	>Rejoignez nous, et trouver l'inspiration<br />culinaire qui
	<span class="text-highlight">vous</span> ressemble</h2
>

<form on:submit|preventDefault={handleRegister}>
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
		<label for="lastName"></label>
		<input
			type="text"
			id="lastName"
			bind:value={lastName}
			required
			placeholder="Nom*"
			autocomplete="family-name"
		/>

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
		color: #333;
		font-weight: normal;
		font-size: 22px;
		.text-highlight {
			color: #de403e;
			font-weight: bold;
		}
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 20px;

		span {
			.material-symbols-rounded {
				font-size: 20px;
				font-weight: bold;
				margin-left: 10px;
				margin-right: -10px;
			}
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

		div {
			margin-bottom: 10px;
			width: 40%;
			min-width: 350px;
			max-width: 700px;
			padding-bottom: 20px;

			.genre {
				color: #939393;

				margin-left: 1%;
				margin-bottom: -20px;
			}
		}

		input {
			&:not([type="checkbox"]) {
				border: 1px solid #ccc;
				border-radius: 3%;
				border-top: none;
				border-left: none;
				border-right: none;

				/*padding du texte*/
				padding: 5px 15px;
				font-size: 16px;
				outline: none;

				width: 100%;
			}
			&:not([type="checkbox"]) {
				border: 1px solid #ccc;
				border-radius: 3%;
				border-top: none;
				border-left: none;
				border-right: none;

				/*padding du texte*/
				padding: 5px 15px;
				font-size: 16px;
				outline: none;

				width: 100%;
			}
			&::placeholder {
				opacity: 0.4;
			}

			&::placeholder {
				transition: 0.5s;
			}

			&:focus::placeholder {
				color: transparent;
			}
		}

		.radio-container {
			display: flex;
			justify-content: space-around;
			color: #939393;

			.radio-wrapper {
				margin: 20px 0 20px 0;
				min-width: fit-content;
				display: flex;
				justify-content: space-around;

				.radio-button input[type="radio"] {
					display: none;
				}

				.radio-button input[type="radio"]:checked ~ .radio-checkmark {
					border: 2px solid #de403e;
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
						background-color: #de403e;
						transition: all 0.2s ease-in-out;
					}
				}
				.radio-button:hover {
					transform: translateY(-2px);
				}
			}
		}
	}

	@media screen and (max-width: 1000px) {
		#lastName {
			margin-bottom: 30px;
		}
	}

	@media screen and (min-width: 1000px) {
		.name {
			display: flex;
			flex-flow: row;
			align-items: stretch;
		}

		#lastName {
			flex: 0 1 auto;
			width: 100%;

			margin-right: 5%;
		}

		#firstName {
			flex: 1 1 auto;
			width: 100%;
		}
	}
</style>
