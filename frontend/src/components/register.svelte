<script lang="ts">
	import { goto } from "$app/navigation";
	import { sendRegister } from "$lib/api/auth-request";
	import { schemaRegister } from "$lib/api/auth-schema";
	import type { ErrorsRegister } from "$lib/api/auth-types";
	import * as yup from "yup";

	let errors: ErrorsRegister = {};

	let registerValues = {
		firstName: "",
		lastName: "",
		email: "",
		password: "",
		confirmPassword: "",
		gender: "",
	};

	//TODO: Gestion du gender dans la bd
	async function handleRegister() {
		try {
			await schemaRegister.validate(registerValues, {
				abortEarly: false,
			});
			const finalValues = {
				name: `${registerValues.firstName} ${registerValues.lastName}`,
				email: registerValues.email,
				password: registerValues.password,
			};
			const result = await sendRegister(finalValues);
			errors = {};
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
					console.log(errors);
				}
			}
		}
	}
</script>

<h2
	>Rejoignez nous, et trouver l'inspiration<br />culinaire qui
	<span class="text-highlight">vous</span> ressemble</h2
>

{#if errors.server}
	<span class="error">{errors.server}</span>
{/if}

<form on:submit|preventDefault={handleRegister}>
	<div>
		<label for="email"></label>
		<input
			type="text"
			id="email"
			bind:value={registerValues.email}
			placeholder="Email*"
			autocomplete="username"
		/>
		{#if errors.email}
			<span class="error">{errors.email}</span>
		{/if}
	</div>

	<div class="name">
		<label for="firstName"></label>
		<input
			type="text"
			id="firstName"
			bind:value={registerValues.firstName}
			placeholder="Prénom*"
			autocomplete="given-name"
		/>
		{#if errors.firstName}
			<span class="error">{errors.firstName}</span>
		{/if}

		<label for="lastName"></label>
		<input
			type="text"
			id="lastName"
			bind:value={registerValues.lastName}
			placeholder="Nom*"
			autocomplete="family-name"
		/>
		{#if errors.lastName}
			<span class="error">{errors.lastName}</span>
		{/if}
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="password"
			bind:value={registerValues.password}
			placeholder="Mot de passe*"
			autocomplete="new-password"
		/>
		{#if errors.password}
			<span class="error">{errors.password}</span>
		{/if}
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="confirmPassword"
			bind:value={registerValues.confirmPassword}
			placeholder="Confirmer mot de passe*"
			autocomplete="new-password"
		/>
		{#if errors.confirmPassword}
			<span class="error">{errors.confirmPassword}</span>
		{/if}
	</div>

	<div>
		<p class="genre">Je suis* :</p>
	</div>

	<div class="radio-container">
		<div class="radio-wrapper">
			<label class="radio-button">
				<input
					id="male"
					name="radio-group"
					bind:group={registerValues.gender}
					type="radio"
					value="male"
				/>
				<span class="radio-checkmark"></span>
				<span class="radio-label">Un pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input
					id="female"
					name="radio-group"
					bind:group={registerValues.gender}
					type="radio"
					value="female"
				/>
				<span class="radio-checkmark"></span>
				<span class="radio-label">Une pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input
					id="other"
					name="radio-group"
					bind:group={registerValues.gender}
					type="radio"
					value="other"
				/>
				<span class="radio-checkmark"></span>
				<span class="radio-label">Autre</span>
			</label>
		</div>
	</div>
	{#if errors.gender}
		<span class="error">{errors.gender}</span>
	{/if}

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
