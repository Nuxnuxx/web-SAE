<script lang="ts">
	import type { ErrorsRegister, User } from "$lib/api/auth-types";

	export let errors: ErrorsRegister;
	export let value: User;
</script>

<h2
	>Rejoignez nous, et trouver l'inspiration<br />culinaire qui
	<span class="highlight">vous</span> ressemble</h2
>

<form method="post" action="?/register">
	{#if errors?.server}
		<span class="error">{errors?.server}</span>
	{/if}
	<div class="input__wrapper">
		<div>
			<label for="mail"></label>
			<input
				type="text"
				id="mail"
				name="mail"
				placeholder="Mail*"
				value={value?.mail || ""}
				autocomplete="username"
			/>
			{#if errors?.mail}
				<span class="error">{errors?.mail}</span>
			{/if}
		</div>

		<div class="input">
			<div class="name__wrapper">
				<div>
					<label for="firstName"></label>
					<input
						type="text"
						id="firstName"
						name="firstName"
						placeholder="Prénom*"
						value={value?.firstName || ""}
						autocomplete="given-name"
					/>
					{#if errors?.firstName}
						<span class="error">{errors?.firstName}</span>
					{/if}
				</div>

				<div>
					<label for="lastName"></label>
					<input
						type="text"
						id="lastName"
						name="lastName"
						placeholder="Nom*"
						value={value?.lastName || ""}
						autocomplete="family-name"
					/>
					{#if errors?.lastName}
						<span class="error">{errors?.lastName}</span>
					{/if}
				</div>
			</div>
		</div>

		<div>
			<label for="password"> </label>
			<input
				type="password"
				id="password"
				name="password"
				placeholder="Mot de passe*"
				value={value?.password || ""}
				autocomplete="new-password"
			/>
			{#if errors?.password}
				<span class="error">{errors?.password}</span>
			{/if}
		</div>

		<div>
			<label for="password"> </label>
			<input
				type="password"
				id="confirmPassword"
				name="confirmPassword"
				placeholder="Confirmer mot de passe*"
				value={value?.confirmPassword || ""}
				autocomplete="new-password"
			/>
			{#if errors?.confirmPassword}
				<span class="error">{errors?.confirmPassword}</span>
			{/if}
		</div>
	</div>

	<div>
		<p class="genre">Je suis* :</p>
	</div>

	<div class="radio-container">
		<div class="radio-wrapper">
			<label class="radio-button">
				<input
					id="male"
					checked={value?.gender == "male" ? true : false}
					name="gender"
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
					checked={value?.gender == "female" ? true : false}
					name="gender"
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
					checked={value?.gender == "other" ? true : false}
					name="gender"
					type="radio"
					value="other"
				/>
				<span class="radio-checkmark"></span>
				<span class="radio-label">Autre</span>
			</label>
		</div>
	</div>
	{#if errors?.gender}
		<span class="error">{errors?.gender}</span>
	{/if}

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
		width: 80%;
		margin: 0 auto;
		margin-bottom: 2rem;
	}

	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 40%;
		margin: 0 auto;
		min-width: 350px;
		max-width: 700px;

		.input__wrapper {
			display: flex;
			flex-direction: column;
			align-items: center;
			gap: 1.5rem;
			width: 100%;
		}

		.error {
			color: var(--primary-color);
			font-size: 0.8rem;
			font-weight: bold;
			text-align: center;
		}

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
			margin: 1rem 0 0 0;
			border-radius: 2rem;
			background-color: var(--primary-color);
			color: var(--white-color);
			font-size: 1rem;
			border: none;
			cursor: pointer;
		}

		.genre {
			color: var(--light-black-color);
			margin: 1.5rem 0 0 0;

			width: 85%;
			text-align: start;
		}

		div {
			width: 100%;
			display: flex;
			flex-direction: column;
			align-items: center;
		}

		input {
			&:not([type="checkbox"]) {
				border-bottom: 1px solid var(--light-secondary-color);
				border-top: none;
				border-left: none;
				border-right: none;

				padding: 0.5rem 1rem;
				outline: none;

				width: 80%;
			}
			&::placeholder {
				opacity: 0.4;
				transition: 300ms ease-in-out;
			}

			&:focus::placeholder {
				transform: translateY(-2rem);
			}
		}
		.input {
			width: 100%;
			.name__wrapper {
				width: 88%;
				flex-direction: row;
				justify-content: space-between;

				div {
					width: 100%;
				}
			}
		}

		.radio-container {
			display: flex;
			flex-direction: row;
			justify-content: space-around;
			color: var(--light-black-color);

			.radio-wrapper {
				margin: 1rem 0 1rem 0;
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
