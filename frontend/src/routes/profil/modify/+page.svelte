<script lang="ts">
	import { getContext } from "svelte";
	import Modify from "../../../components/modifyProfil.svelte";

	const user: any = getContext("user");
	const food: string = user.food;
	const backgroundColor: string = user.backgroundColor;
</script>

<Modify {food} {backgroundColor} />
