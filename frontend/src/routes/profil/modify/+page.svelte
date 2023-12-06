<script lang="ts">
	import { getContext, onMount } from "svelte";
	import Modify from "../../../components/modifyProfil.svelte";
	import type { PageLoad } from "../../playlist/[slug]/$types";
	import { goto } from "$app/navigation";

	const user: any = getContext("user");
	const food: string = user.food;
	const backgroundColor: string = user.backgroundColor;

	export let data: PageLoad;

</script>

<Modify data={data.profil} {food} {backgroundColor} />
