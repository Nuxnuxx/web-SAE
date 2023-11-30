<script lang="ts">
	import { getContext } from "svelte";
	import Show from "../../components/showProfile.svelte";

	const user: any = getContext("user");
	const food: string = user.food;
	const backgroundColor: string = user.backgroundColor;
</script>

<Show {food} {backgroundColor} />
