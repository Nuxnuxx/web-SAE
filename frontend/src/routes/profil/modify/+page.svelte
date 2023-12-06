<script lang="ts">
	import { getContext } from "svelte";
	import Modify from "../../../components/modifyProfil.svelte";
	import type { PageLoad } from "../../playlist/[slug]/$types";

	const user: any = getContext("user");
	const food: string = user.food;
	const backgroundColor: string = user.backgroundColor;

	export let data: PageLoad;
</script>

<Modify data={data.profil} {food} {backgroundColor} />
