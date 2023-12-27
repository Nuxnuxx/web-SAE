<script lang="ts">
	import type { User } from "$lib/api/auth-types";

	export let food: string = "";
	export let backgroundColor: string = "";
	export let data: User;

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

		<div class="card__content">
			<div class="content__name">
				<p>{data.firstName + " " + data.lastName}</p>
			</div>

			<div class="content__localisation">
				<span class="material-symbols-rounded"> location_on </span>
				<span>France</span>
			</div>

			<a type="button" class="content__button" href="/profil/modify">
				<span class="material-symbols-rounded button__icon">
					edit
				</span>
				<span class="button__text">Modifier votre profil</span>
			</a>

			<form
				method="post"
				action="?/logout"
				class="content__deconnexion"
			>
				<button type="submit">Déconnexion</button>
			</form>
		</div>
	</div>
</main>

<style lang="scss">
	p {
		margin: 0;
	}
	.card {
		margin-bottom: 7rem;
		width: auto;
		/*ne peux pas etre plus petit que les enfants*/
		overflow: auto;
		border-radius: 1rem;
		background: #fff;
		box-shadow: 0px 0px 10px 0px var(--light-secondary-color);

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

		.card__content {
			display: flex;
			flex-direction: column;
			align-items: center;
			margin-right: 20%;
			margin-left: 20%;

			.content__name {
				p {
					font-size: 1.5rem;
					text-align: center;
					font-weight: bold;
				}
			}

			.content__localisation {
				display: flex;
				align-items: center;
				margin-bottom: 1.5rem;
				color: var(--light-secondary-color);
				font-size: 1rem;
			}

			.content__button {
				padding: 0.4rem;

				background-color: var(--primary-color);
				color: #fff;
				border: none;
				border-radius: 10px;
				cursor: pointer;
				font-size: 1rem;
				display: flex;
				align-items: center;
				justify-content: space-between;
				text-decoration: none;

				@media screen and (max-width: 1150px) {
					font-size: 1rem;
				}

				.button__text {
					@media screen and (max-width: 1100px) {
						display: none;
					}
				}

				.button__icon {
					margin-left: auto;
					margin-right: auto;
				}
			}

			.content__deconnexion {
				margin-top: 0.2rem;
				margin-bottom: 3rem;

				button {
					all: unset;
					cursor: pointer;
					text-decoration: underline;
					color: var(--light-secondary-color);
				}
			}
		}
	}
</style>
