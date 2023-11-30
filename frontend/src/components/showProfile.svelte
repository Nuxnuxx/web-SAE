<script>
	//TODO: get user data from backend
	import profileImg from "$lib/img/profile/food/coffee.png";
	let firstName = "John";
	let lastName = "Doe";
	let location = "France";
</script>

<main>
	<div class="card">
		<!--image de profil-->
		<div class="card__img">
			<img src={profileImg} alt="profile" />
		</div>

		<div class="card__content">
			<div class="content__name">
				<p>{firstName + " " + lastName}</p>
			</div>

			<div class="content__localisation">
				<span class="material-symbols-rounded"> location_on </span>
				<span>{location}</span>
			</div>

			<a type="button" class="content__button" href="/profil/modify">
				<span class="material-symbols-rounded button__icon">
					edit
				</span>
				<span class="button__text">Modifier votre profil</span>
			</a>

			<div class="content__deconnexion">
				<!-- TODO: link to disconnect -->
				<a href="auth">Déconnexion</a>
			</div>
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
		box-shadow: 0px 0px 10px 0px #dcdcdc;

		.card__img {
			width: 25%;
			height: 100%;
			border-radius: 100%;
			overflow: hidden;
			margin: 3rem auto 1rem auto;
			background: linear-gradient(
				160deg,
				#d282e6 0%,
				rgba(210, 130, 230, 0.25) 69.92%
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
					font-size: 2rem;
					font-weight: bold;
				}
			}

			.content__localisation {
				display: flex;
				align-items: center;
				margin-bottom: 1.5rem;
				color: #dcdcdc;
				font-size: 1rem;
			}

			.content__button {
				padding: 0.4rem;

				background-color: #de403e;
				color: #fff;
				border: none;
				border-radius: 10px;
				cursor: pointer;
				font-size: 15px;
				display: flex;
				align-items: center;
				justify-content: space-between;
				text-decoration: none;

				@media screen and (max-width: 1150px) {
					font-size: 16px;
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

				a {
					color: #dcdcdc;
				}
			}
		}
	}
</style>
