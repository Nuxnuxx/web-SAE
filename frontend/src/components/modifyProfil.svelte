<script lang="ts">
	import type { ErrorsRegister, User } from "$lib/api/auth-types";
	import { enhance } from "$app/forms";
	import Input from "./input.svelte";

	export let data: User;
	export let food: string = "";
	export let backgroundColor: string = "";
	export let errors: ErrorsRegister;

	//FIXME: this is a workaround to get the image path, it may do nothing on a server but it throws an error on localhost
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

			<button type="submit" class="form__button">
				<span class="button__text">Sauvegarder</span>
				<span class="material-symbols-rounded"> done </span>
			</button>
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
		box-shadow: 0px 0px 10px 0px #dcdcdc;

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
		.form__button {
			margin: 1.5rem auto;

			padding: 10px 20px;
			background-color: #de403e;
			color: #fff;
			border: none;
			border-radius: 20px;
			cursor: pointer;
			font-size: 1.2rem;
			display: flex;
			align-items: center;
			justify-content: space-between;

			@media screen and (max-width: 1150px) {
				font-size: 1rem;
				margin-right: auto;
				margin-left: auto;
			}
		}
	}
</style>
