<script lang="ts">
	import { enhance } from "$app/forms";

	export let idRecipe: number;
</script>

<form use:enhance method="post" action="?/likeRecipe">
	<input hidden name="id" value={idRecipe} type="text" />
	<button type="submit" class="material-symbols-rounded">
		favorite
	</button></form
>

<style lang="scss">
	button.material-symbols-rounded {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
		outline: inherit;
	}
</style>
