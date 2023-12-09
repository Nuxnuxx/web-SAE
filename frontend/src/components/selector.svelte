<script lang="ts">
	export let data: string[];
	export let rating_like: boolean = false;
	export let selected: number = 0;
	const length = data.length;
</script>

<div class={rating_like ? "selector rating_like" : "selector"}>
	{#each data as item, i (i)}
		<input
			value={length - i}
			id={`selector-${item}${length - i}`}
			type="radio"
			bind:group={selected}
		/>
		<label
			for="selector-{item}{length - i}"
			class="material-symbols-rounded">{item}</label
		>
	{/each}
</div>

<style lang="scss">
	.rating_like {
		& input:checked ~ label,
		& label:hover,
		& label:hover ~ label {
			color: var(--secondary-color);
			transition: color 0.3s;
		}
	}
	.selector {
		display: inline-block;
		margin: 3rem 0;
		& > * {
			margin-right: 2vw;
		}

		& input {
			display: none;
		}

		& label {
			float: right;
			cursor: pointer;
			color: #dcdcdc;
			transition: color 0.3s;
		}

		& label {
			font-size: 10vw;
		}

		& input:checked + label,
		& label:hover {
			color: var(--secondary-color);
			transition: color 0.3s;
		}
	}

	@media (min-width: 768px) {
		.selector {
			& label {
				font-size: 5vw;
			}
		}
	}

	@media (min-width: 1024px) {
		.selector {
			& label {
				font-size: 4vw;
			}
		}
	}
</style>
