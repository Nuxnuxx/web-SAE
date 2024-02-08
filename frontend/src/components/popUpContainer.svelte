<script lang="ts">
	import { goto } from "$app/navigation";
	import { onDestroy } from "svelte";
	import { userStore, popup } from "../store";
	import FormLike from "./formLike.svelte";
	import FormPlaylist from "./formPlaylist.svelte";

	export let idRecipe: number;
	export let type: string;

	let button: HTMLButtonElement;

	$: isNull = $popup.idRecipe == null || $popup.type == null;
	$: SamePopUp = idRecipe == $popup.idRecipe && type == $popup.type;
	$: mouseEnter = false;

	// pour gerer le cas de changement de page a la destruction du composant on reset le store
	onDestroy(() => {
		popup.set({
			idRecipe: null,
			type: null,
		});
	});
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class={`dropdown ${SamePopUp ? "selected" : ""}`}>
	{#if type == "like" && $userStore}
		<FormLike {idRecipe} />
	{:else}
		<button
			bind:this={button}
			class="material-symbols-rounded"
			on:blur={() => {
				// si on sort de la pop up on la ferme
				if (!mouseEnter) {
					popup.set({
						idRecipe: null,
						type: null,
					});
				}
			}}
			on:click={() => {
				// au premier click pour ouvrir la premiere pop up
				if (isNull) {
					popup.set({
						idRecipe: idRecipe,
						type: type,
					});
				}
				// si c'est pas la meme on lui met le nouvelle id pour qu'elle souvre au bonne endoirt
				else if (!SamePopUp) {
					popup.set({
						idRecipe: idRecipe,
						type: type,
					});
				}
				// sinon cest la meme et on la ferme (en gros le double click)
				else {
					popup.set({
						idRecipe: null,
						type: null,
					});
				}
			}}>{type === "playlist" ? "playlist_add" : "favorite"}</button
		>
	{/if}
	{#if $userStore && SamePopUp}
		<div
			class="dropdown-content"
			on:mouseenter={() => {
				// on met mouseEnter a true pour pas que la pop up se ferme
				mouseEnter = true;
			}}
			on:mouseleave={() => {
				mouseEnter = !mouseEnter;
				// pour que quand on sort on revient a letat initial ou le button est focus
				button.focus();
			}}
		>
			<FormPlaylist {idRecipe} />
		</div>
	{:else if SamePopUp}
		<div
			on:mouseenter={() => {
				// on met mouseEnter a true pour pas que la pop up se ferme
				mouseEnter = true;
			}}
			on:mouseleave={() => {
				mouseEnter = !mouseEnter;
				// pour que quand on sort on revient a letat initial ou le button est focus
				button.focus();
			}}
			class="dropdown-content"
		>
			<button on:click={() => goto("/auth")} class="nav__login">
				<span class="material-symbols-rounded">person</span>
				connexion
			</button>
		</div>
	{/if}
</div>

<style lang="scss">
	.dropdown {
		position: relative;
		&.selected::before {
			content: "";
			position: absolute;
			z-index: 2;
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

		&.like {
			color: var(--primary-color);
			transition: all 0.2s ease-out;
			animation: scaled 0.5s ease-in-out;

			@keyframes scaled {
				0% {
					transform: scale(1);
				}
				50% {
					transform: scale(1.2);
				}
				100% {
					transform: scale(1);
				}
			}
		}
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
