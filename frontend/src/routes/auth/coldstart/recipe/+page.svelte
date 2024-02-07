<script lang="ts">
	import Swiper from "../../../../components/Swiper.svelte";
	import { getRecipes } from "$lib/api/auth-request";
	import { Price, Difficulty } from "$lib/api/recipe-types";
	import { coldstartLiked } from "../../../../store";
	import { onMount } from "svelte";
	import type { PageData } from "../$types";

	export let data: PageData;
	console.log(data);
</script>

<form method="post" action="?/coldstart">
	<h2>
		Pour finir, fait nous savoir ce que tu aimes en choisissant
		<span class="highlight">au moins 3</span>
		recettes.
	</h2>

	<!-- faire le swiper selon les recipesbypreferences -->
	<Swiper recipes={data.recipes} />

	<!-- en invisible, mettre les réponses pour les avoir dans le body -->
	<input type="hidden" name="price" value={data.price} />
	<input type="hidden" name="difficulty" value={data.difficulty} />
	<input type="hidden" name="liked" value={$coldstartLiked} />

	<button class="button" on:click={() => {}} type="submit">
		<span>Terminer</span>
	</button>
</form>
