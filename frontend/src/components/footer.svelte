<script>
</script>

<footer>
	<div class="bubble__wrapper">
		<a
			href="https://github.com/Nuxnuxx/web-SAE"
			target="_blank"
			rel="noopener noreferrer"
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
		<a
			href="https://www.figma.com/file/UiACSX9ROK89nP9Di5RmkH/PirateCook?type=design&node-id=0%3A1&mode=design&t=ofeghSqgFnjGOAdl-1"
			target="_blank"
			rel="noopener noreferrer"
			class="small__bubble figma__bubble"
		>
			<span class="figma__icon"></span>
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
				outline: var(--secondary-color) 0.2rem solid;
				outline-offset: -0.1rem; // needed to have seamless border
				width: 3rem;
				height: 3rem;
				aspect-ratio: 1/1;
				background: var(--secondary-color);
				border-radius: 50%;
				display: flex;
				justify-content: center;
				align-items: center;
				cursor: pointer;
				span {
					font-size: 2rem;
					color: var(--white-color);
				}
			}

			.small__bubble {
				width: 2.25rem;
				height: 2.25rem;
				aspect-ratio: 1/1;
				background: var(--secondary-color);
				border-radius: 50%;
				display: flex;
				justify-content: center;
				align-items: center;
				cursor: pointer;

				&.github__bubble {
					margin-top: 0.75rem;
				}
				&.figma__bubble {
					margin-top: 0.5rem;
				}
				span {
					color: var(--white-color);
					&.figma__icon {
						margin: auto;
						background: var(--white-color);
						mask-image: url("../lib/img/figma_icon.svg");
						mask-size: 100% 100%;
						mask-repeat: no-repeat;
						background-size: cover;
						width: 1.75rem;
						height: 1.75rem;
					}
					&.github__icon {
						margin: auto;
						background: var(--white-color);
						mask-image: url("../lib/img/github_icon.svg");
						mask-size: 100% 100%;
						mask-repeat: no-repeat;
						background-size: cover;
						width: 1.75rem;
						height: 1.75rem;
					}
				}
			}
		}

		p {
			margin: 0.2rem;
			color: var(--white-color);
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
				color: var(--primary-color);
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
					outline: var(--secondary-color) 0.3rem solid;
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
					&.figma__bubble {
						margin-top: 0.75rem;
					}
					span {
						font-size: 3rem;

						&.figma__icon {
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
				color: var(--white-color);
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
					color: var(--primary-color);
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
			margin-top: 4rem;
			.bubble__wrapper {
				.big__bubble {
					outline: var(--secondary-color) 0.4rem solid;
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
					&.figma__bubble {
						margin-top: 1rem;
					}
					span {
						font-size: 4rem;

						&.figma__icon {
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
