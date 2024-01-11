<script lang="ts">
	import { enhance } from "$app/forms";
	import { goto } from "$app/navigation";
	import { userStore, isLikedButtonOpen } from "../store";

	export let idRecipe: number;
</script>

{#if $userStore}
	<form use:enhance method="post" action="?/likeRecipe">
		<input hidden name="id" value={idRecipe} type="text" />
		<button type="submit" class={`material-symbols-rounded like`}>
			favorite
		</button>
	</form>
{:else}
	<div
		class="dropdown {$isLikedButtonOpen.open &&
		$isLikedButtonOpen.id == idRecipe
			? 'selected'
			: ''}"
	>
		<button
			on:click={() => {
				// Check if the pop up id open on store is the same as the idRecipe
				// if not we set the store data to the new recipe clicked and open it
				// else we can just close the one already open
				if ($isLikedButtonOpen.id != idRecipe) {
					isLikedButtonOpen.set({
						id: idRecipe,
						open: true,
					});
				} else {
					isLikedButtonOpen.set({
						id: idRecipe,
						open: !$isLikedButtonOpen.open,
					});
				}
			}}
			class="material-symbols-rounded">favorite</button
		>
		{#if $isLikedButtonOpen.open && $isLikedButtonOpen.id == idRecipe && !$userStore}
			<div class="dropdown-content">
				<button on:click={() => goto("/auth")} class="nav__login">
					<span class="material-symbols-rounded">person</span>
					connexion
				</button>
			</div>
		{/if}
	</div>
{/if}

<style lang="scss">
	.dropdown {
		position: relative;
		&.selected::before {
			z-index: 2;
			content: "";
			position: absolute;
			top: 1.4rem;
			left: 5%;
			margin-left: 0.25rem;
			width: 0.7rem;
			height: 0.7rem;
			border: 2.5px solid var(--light-secondary-color);
			border-bottom: none;
			border-right: none;
			background-color: var(--white-color);
			transform: rotate(45deg);
		}
	}
	.dropdown-content {
		display: block;
		position: absolute;
		z-index: 1;
		background-color: var(--white-color);
		border: 3px solid var(--light-secondary-color);
		border-radius: 10px;
		width: 12rem;
		left: -100%;

		.nav__login {
			display: flex;
			align-items: center;
			gap: 0.5rem;
			padding: 0.5rem 1rem;
			border-radius: 1rem;
			background-color: var(--primary-color);
			color: var(--white-color);
			font-weight: medium;
			font-size: 1rem;
			border: none;
			cursor: pointer;
			transition: all 0.2s ease-out;
			margin: 1rem auto;

			&:hover {
				transform: scale(1.02);
			}
		}
	}

	button.material-symbols-rounded {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}

	@media (max-width: 768px) {
		.dropdown {
			&.selected::before {
				display: none;
			}
		}
		.dropdown-content {
			position: fixed;
			left: 50%;
			transform: translate(-50%);
		}
	}
</style>
