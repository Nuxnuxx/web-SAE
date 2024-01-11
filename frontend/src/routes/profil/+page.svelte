<script lang="ts">
	import { getContext } from "svelte";
	import Show from "../../components/showProfile.svelte";
	import type { PageData } from "../$types";

	const user: any = getContext("user");
	const food: string = user.food;
	const backgroundColor: string = user.backgroundColor;

	export let data: PageData;
</script>

<Show data={data.profil} {food} {backgroundColor} />
