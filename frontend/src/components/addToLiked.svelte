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
	<div class="dropdown">
		<button
			on:click={() => {
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
				<p>apprenons a te connaitre d'abord</p>
				<button on:click={() => goto("/auth")} class="nav__login">
					<span class="material-symbols-rounded">person</span>
					connexion
				</button>
			</div>
		{/if}
	</div>
{/if}

<style lang="scss">
	button.material-symbols-rounded {
		/* reset all button properties */
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}

	.dropdown-content {
		display: block;
		position: absolute;
		background-color: var(--white-color);
		border: 3px solid var(--light-secondary-color);
		border-radius: 10px;
		width: 12rem;
		&::before {
			//TODO: tom fais ce truc
			content: "";
			position: absolute;
			top: -7%;
			left: 0%;
			width: 10%;
			height: 10%;
			border: 2.5px solid var(--light-secondary-color);
			border-bottom: none;
			border-right: none;
			background-color: var(--white-color);
			transform: rotate(45deg);
		}

		p {
			font-size: 0.8rem;
			font-weight: bold;
		}

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
			margin: 0 auto;
			margin-bottom: 1rem;

			&:hover {
				transform: scale(1.02);
			}
		}
	}
</style>
