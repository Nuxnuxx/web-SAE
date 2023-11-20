<script>
	let data = {
		img: "https://assets.afcdn.com/recipe/20230608/143328_w600h600c1.webp",
		title: "Tarte aux pommes et chocolat Nestlé ®",
		nbLikes: 120,
		liked: true,
		saved: false,
		like: function () {
			data.liked = !data.liked;
		},
		save: function () {
			console.log("save");
			data.saved = !data.saved;
		},
	};
</script>

<!-- link need to be replaced by a local style-->
<link
	rel="stylesheet"
	href="https://fonts.googleapis.com/css2?family=Material+Symbols+Rounded:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
/>

<div class="card">
	<div class="card__img">
		<img src={data.img} alt={data.title} />
	</div>
	<div class="card__content">
		<div class="card__title">
			<h3>{data.title}</h3>
			<div class="card__likes">
				<span class="card__likes__number">{data.nbLikes}</span>
				<span class="material-symbols-rounded filled"> favorite </span>
			</div>
		</div>
		<div class="card__icons">
			<span class="card__likes__icon">
				{#if data.liked}
					<button
						on:click={() => data.like()}
						class="material-symbols-rounded filled red"
					>
						favorite
					</button>
				{:else}
					<button
						on:click={() => data.like()}
						class="material-symbols-rounded"
					>
						favorite
					</button>
				{/if}
			</span>
			<span class="card__saved__icon">
				{#if data.saved}
					<button
						on:click={() => data.save()}
						class="material-symbols-rounded green"
					>
						playlist_add_check
					</button>
				{:else}
					<button
						on:click={() => data.save()}
						class="material-symbols-rounded"
					>
						playlist_add
					</button>
				{/if}
			</span>
		</div>
	</div>
</div>

<style lang="scss">
	.card {
		aspect-ratio: 1/1;
		width: 20vh;
		display: flex;
		flex-direction: column;
		border-radius: 1.5vh;
		overflow: hidden;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);

		.card__img {
			width: 100%;
			height: 70%;
			border-radius: 5px;
			overflow: hidden;

			img {
				width: 100%;
				height: 100%;
				object-fit: cover;
			}
		}
		.card__content {
			padding: 10px;
			display: grid;
			grid-auto-flow: column;
			height: 30%;
			.card__title {
				h3 {
					font-size: 14px;
					margin: 0;
					width: 100%;
					line-height: normal;
					overflow: hidden;
					-webkit-box-orient: vertical;
					-webkit-line-clamp: 2;
					display: -webkit-box;
				}
				.card__likes {
					display: flex;
					align-items: center;
					height: 40%;
					color: #c4c4c4;
					font-size: 12px;
					.material-symbols-rounded {
						font-size: inherit;
						margin-left: 3px;
					}
				}
			}
			.card__icons {
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
			}
		}
	}

	button.material-symbols-rounded {
		/* reset all button properties */
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}

	.material-symbols-rounded {
		width: max-content;
		font-variation-settings:
			"FILL" 0,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
	}

	.material-symbols-rounded.filled {
		font-variation-settings:
			"FILL" 1,
			"wght" 400,
			"GRAD" 0,
			"opsz" 24;
	}

	.material-symbols-rounded.red {
		color: #f93362;
	}

	.material-symbols-rounded.green {
		color: #40d133;
	}
</style>
