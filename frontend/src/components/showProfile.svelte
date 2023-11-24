<script>
</script>

<link 
    rel="stylesheet" 
    href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
/>

<main>
	<div class="card">
		<!--image de profil-->
		<div class="card__img">
			<img src="https://picsum.photos/200" alt="profile picture"/>
		</div>
        
        <div class="card__content">
            <div class="content__nom">
                <p>Mister P.</p>
            </div>

            <div class="content__localisation">
                <span class="material-symbols-rounded">
                    location_on
                    </span>
                <p>France</p>
            </div>

            <button type="submit" class="content__button">
                <span class="material-symbols-rounded">
                    edit
                    </span>
                <span class="button__text">Modifier votre profil</span>
            </button>

            <div class="content__deconnexion">
                <p>Déconnexion</p>
            </div>
        </div>
	</div>
</main>

<style lang="scss">
	.card {
		border : 1px solid #ccc;
		border-radius: 20px;
		width: auto;
		/*ne peux pas etre plus petit que les enfants*/
		overflow: auto;
		border-radius: 20px;
		background: #F5F5F5;
		box-shadow: 0px 0px 10px 0px #DCDCDC;


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
	
		.card__content {
			display: flex;
			flex-direction: column;
			margin-right: 20%;
			margin-left: 20%;
			margin-top: 20px;

			.content__nom {
				padding-bottom: 10px;
			}

            .content__localisation {
                display: flex;
                align-items: center;
                padding-bottom: 10px;

                p {
                    margin-left: 10px;
                }
            }

			.content__button {
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

            .content__deconnexion {
                padding-bottom: 10px;
            }
		}
	}

    .material-symbols-rounded {
        font-variation-settings:
        'FILL' 0,
        'wght' 400,
        'GRAD' 0,
        'opsz' 24
    }
</style>
