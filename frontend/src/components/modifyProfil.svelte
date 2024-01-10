<script lang="ts">
	import type { ErrorsRegister, User } from "$lib/api/auth-types";
	import { enhance } from "$app/forms";
	import Input from "./input.svelte";

	export let data: User;
	export let food: string = "";
	export let backgroundColor: string = "";
	export let errors: ErrorsRegister;

	const profileImg = new URL(
		`../routes/profil/food/${food}.png`,
		import.meta.url
	).href;
</script>

<main>
	<div class="card">
		<!--image de profil-->
		<div class="card__img" style="--theme-color: {backgroundColor}">
			<img src={profileImg} alt="profile" />
		</div>

		<form
			use:enhance
			method="post"
			class="card__form"
			action="?/modifyPassword"
		>
			<Input
				name="firstName"
				placeholder="First name"
				secure={false}
				value={data.firstName || ""}
			/>
			<Input
				name="lastName"
				placeholder="Last name"
				secure={false}
				value={data.lastName || ""}
			/>
			<Input
				name="mail"
				placeholder="Mail"
				secure={false}
				value={data.mail}
			/>
			<Input
				name="newPassword"
				placeholder="Password"
				secure={true}
				value={""}
			/>
			{#if errors?.server}
				<span class="error">{errors?.server}</span>
			{/if}

			<div class="button__wrapper">
				<a class="btn button__secondary" href="/profil">
					Annuler
					<span class="material-symbols-rounded"> close </span>
				</a>

				<button type="submit" class="btn button__primary">
					Sauvegarder
					<span class="material-symbols-rounded"> done </span>
				</button>
			</div>
		</form>
	</div>
</main>

<style lang="scss">
	.card {
		margin-bottom: 1rem;
		width: auto;
		/*ne peux pas etre plus petit que les enfants*/
		overflow: auto;
		border-radius: 1rem;
		background: #fff;
		box-shadow: 0px 0px 10px 0px var(--light-secondary-color);

		.error {
			color: var(--primary-color);
			margin-top: 0.1rem;
			font-size: 0.8rem;
			font-weight: bold;
			text-align: center;
		}

		.card__form {
			margin: 0 4rem;
		}

		.card__img {
			width: 25%;
			height: 100%;
			border-radius: 100%;
			overflow: hidden;
			margin: 3rem auto 1rem auto;
			background: linear-gradient(
				160deg,
				hsl(var(--theme-color)) 0%,
				hsla(var(--theme-color), 0.25) 69.92%
			);
			img {
				width: 100%;
				height: 100%;
				object-fit: cover;
				vertical-align: middle;
			}
		}
		.button__wrapper {
			display: flex;
			gap: 2rem;
			padding: 0 1rem;
			.btn {
				width: fill-available;
				margin: 1.5rem auto;

				padding: 10px 20px;
				border-radius: 20px;
				cursor: pointer;
				font-size: 1.2rem;
				display: flex;
				align-items: center;
				justify-content: space-between;
				text-decoration: none;

				&.button__primary {
					background-color: var(--primary-color);
					color: var(--white-color);
					border: none;
				}

				&.button__secondary {
					background-color: var(--white-color);
					color: var(--light-black-color);
					border: 1px solid var(--light-black-color);
				}

				@media screen and (max-width: 1150px) {
					font-size: 1rem;
					margin-right: auto;
					margin-left: auto;
				}
			}
		}
	}
</style>
