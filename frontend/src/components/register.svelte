<script lang="ts">
	import type { ActionData } from "../routes/auth/$types";

	export let form: ActionData;
</script>

<h2
	>Rejoignez nous, et trouver l'inspiration<br />culinaire qui
	<span class="highlight">vous</span> ressemble</h2
>

<form method="post" action="?/register">
	{#if form?.server}
		<span class="error">{form?.server}</span>
	{/if}
	<div>
		<label for="mail"></label>
		<input
			type="text"
			id="mail"
			name="mail"
			placeholder="Mail*"
			autocomplete="username"
		/>
		{#if form?.mail}
			<span class="error">{form?.mail}</span>
		{/if}
	</div>

	<div class="name">
		<div>
			<label for="firstName"></label>
			<input
				type="text"
				id="firstName"
				name="firstName"
				placeholder="Prénom*"
				autocomplete="given-name"
			/>
			{#if form?.firstName}
				<span class="error">{form?.firstName}</span>
			{/if}
		</div>

		<div>
			<label for="lastName"></label>
			<input
				type="text"
				id="lastName"
				name="lastName"
				placeholder="Nom*"
				autocomplete="family-name"
			/>
			{#if form?.lastName}
				<span class="error">{form?.lastName}</span>
			{/if}
		</div>
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="password"
			name="password"
			placeholder="Mot de passe*"
			autocomplete="new-password"
		/>
		{#if form?.password}
			<span class="error">{form?.password}</span>
		{/if}
	</div>

	<div>
		<label for="password"> </label>
		<input
			type="password"
			id="confirmPassword"
			name="confirmPassword"
			placeholder="Confirmer mot de passe*"
			autocomplete="new-password"
		/>
		{#if form?.confirmPassword}
			<span class="error">{form?.confirmPassword}</span>
		{/if}
	</div>

	<div>
		<p class="genre">Je suis* :</p>
	</div>

	<div class="radio-container">
		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="male" name="gender" type="radio" value="male" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Un pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="female" name="gender" type="radio" value="female" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Une pirate</span>
			</label>
		</div>

		<div class="radio-wrapper">
			<label class="radio-button">
				<input id="other" name="gender" type="radio" value="other" />
				<span class="radio-checkmark"></span>
				<span class="radio-label">Autre</span>
			</label>
		</div>
	</div>
	{#if form?.gender}
		<span class="error">{form?.gender}</span>
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
			margin: 0 auto;
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

			div {
				width: fit-content;
				align-items: inherit;
			}
			input {
				width: 55%;
			}
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
