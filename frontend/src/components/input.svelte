<script lang="ts">
	export let name: string;
	export let placeholder: string;
	export let secure: boolean;

	export let value: string;
</script>

{#if secure}
	<div>
		<input
			required
			type="password"
			{name}
			autocomplete="off"
			class="input"
			bind:value
		/>
		<label for={name} class="user-label">{placeholder}</label>
	</div>
{:else}
	<div>
		<input
			required
			type="text"
			{name}
			autocomplete="off"
			class="input"
			bind:value
		/>
		<label for={name} class="user-label">{placeholder}</label>
	</div>
{/if}

<style lang="scss">
	div {
		position: relative;
		margin: 1rem 0;
	}

	.input {
		width: fill-available;
		border: solid 1.5px var(--very-light-secondary-color);
		border-radius: 0.5rem;
		background: #fff;
		padding: 0.5rem 1rem;
		font-size: 1rem;
		color: #000;
		transition: border 150ms cubic-bezier(0.4, 0, 0.2, 1);
	}

	.input:focus,
	input:valid {
		outline: none;
		border: 1.5px solid var(--secondary-color);
	}

	.input:focus ~ label,
	input:valid ~ label {
		transform: translateY(-50%) scale(0.8);
		background-color: #fff;
		padding: 0 0.2em;
		color: var(--secondary-color);
	}
	.user-label {
		position: absolute;
		left: 15px;
		color: var(--very-light-secondary-color);
		pointer-events: none;
		transform: translateY(0.5rem);
		transition: 150ms cubic-bezier(0.4, 0, 0.2, 1);
	}
</style>
