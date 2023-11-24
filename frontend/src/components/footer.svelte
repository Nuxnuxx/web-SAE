<script>
</script>

<footer>
	<div class="bubble__wrapper">
		<a
			href="https://github.com/Nuxnuxx/web-SAE"
			class="small__bubble github__bubble"
		>
			<span class="github__icon"></span>
		</a>
		<button
			on:click={() => window.scrollTo({ top: 0, behavior: "smooth" })}
			class="big__bubble"
			aria-label="Scroll to top"
		>
			<span class="material-symbols-rounded"> arrow_upward </span>
		</button>
		<a href="https://dribbble.com/" class="small__bubble dribbble__bubble">
			<span class="dribbble__icon"></span>
		</a>
	</div>
	<p class="flavor__text"
		>Site fait avec <span class="red__highlight">amour</span> par : <br />
		<span class="styled__font">Une équipe épuisée</span></p
	>
	<p class="bottom__text">Tous droits réservés PirateCook.fr - 2023-2023</p>
</footer>

<style lang="scss">
	// MOBILE FIRST
	footer {
		position: relative;
		width: 100%;
		aspect-ratio: 1920/250;
		background: url("../lib/img/footer_shape.svg") no-repeat top left 50%;
		background-size: cover;
		background-origin: border-box;

		.bubble__wrapper {
			position: relative;
			top: -1rem;
			height: 50%;
			display: flex;
			justify-content: center;

			.big__bubble {
				border: none;
				outline: #82d3e6 0.2rem solid;
				outline-offset: -0.1rem; // needed to have seamless border
				width: 3rem;
				height: 3rem;
				aspect-ratio: 1/1;
				background: #82d3e6;
				border-radius: 50%;
				display: flex;
				justify-content: center;
				align-items: center;
				cursor: pointer;
				span {
					font-size: 2rem;
					color: #f5f5f5;
				}
			}

			.small__bubble {
				width: 2.25rem;
				height: 2.25rem;
				aspect-ratio: 1/1;
				background: #82d3e6;
				border-radius: 50%;
				display: flex;
				justify-content: center;
				align-items: center;
				cursor: pointer;

				&.github__bubble {
					margin-top: 0.75rem;
				}
				&.dribbble__bubble {
					margin-top: 0.5rem;
				}
				span {
					color: #f5f5f5;
					&.dribbble__icon {
						margin: auto;
						background: url("../lib/img/dribbble_icon.svg")
							no-repeat 50% 50%;
						background-size: cover;
						width: 1.75rem;
						height: 1.75rem;
					}
					&.github__icon {
						margin: auto;
						background: url("../lib/img/github_icon.svg") no-repeat
							50% 50%;
						background-size: cover;
						width: 1.75rem;
						height: 1.75rem;
					}
				}
			}
		}

		p {
			margin: 0.2rem;
			color: #f5f5f5;
			&.flavor__text {
				margin-left: 2rem;
				font-size: 0.6rem;
			}
			&.bottom__text {
				font-size: 0.5rem;
				text-align: center;
				padding-bottom: 0.2rem;
				margin: 0;
			}
			.red__highlight {
				color: #de403e;
				font-weight: bold;
			}
			.styled__font {
				font-family: "Leckerli One", cursive;
			}
		}
	}
	@media (min-width: 768px) {
		// TABLET
		footer {
			.bubble__wrapper {
				.big__bubble {
					outline: #82d3e6 0.3rem solid;
					outline-offset: -0.1rem; // needed to have seamless border
					width: 4.5rem;
					height: 4.5rem;
					span {
						font-size: 3rem;
					}
				}
				.small__bubble {
					width: 3.25rem;
					height: 3.25rem;

					&.github__bubble {
						margin-top: 1rem;
					}
					&.dribbble__bubble {
						margin-top: 0.75rem;
					}
					span {
						font-size: 3rem;

						&.dribbble__icon {
							width: 2.75rem;
							height: 2.75rem;
						}
						&.github__icon {
							width: 2.75rem;
							height: 2.75rem;
						}
					}
				}
			}
			p {
				color: #f5f5f5;
				&.flavor__text {
					margin-left: 2rem;
					font-size: 1.5rem;
				}
				&.bottom__text {
					font-size: 1rem;
					text-align: center;
					padding-bottom: 1rem;
					margin: 0;
				}
				.red__highlight {
					color: #de403e;
					font-weight: bold;
				}
				.styled__font {
					font-family: "Leckerli One", cursive;
				}
			}
		}
	}
	@media (min-width: 1024px) {
		// LAPTOP

		footer {
			.bubble__wrapper {
				.big__bubble {
					outline: #82d3e6 0.4rem solid;
					outline-offset: -0.1rem; // needed to have seamless border
					width: 6rem;
					height: 6rem;
					span {
						font-size: 4rem;
					}
				}
				.small__bubble {
					width: 4.5rem;
					height: 4.5rem;

					&.github__bubble {
						margin-top: 1.5rem;
					}
					&.dribbble__bubble {
						margin-top: 1rem;
					}
					span {
						font-size: 4rem;

						&.dribbble__icon {
							width: 3.5rem;
							height: 3.5rem;
						}
						&.github__icon {
							width: 3.5rem;
							height: 3.5rem;
						}
					}
				}
			}
		}
	}
</style>
