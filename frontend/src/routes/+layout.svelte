<script lang="ts">
	import Footer from "../components/footer.svelte";
	import Header from "../components/header.svelte";
	import type { LayoutData } from "./$types";

	export let data: LayoutData;

	let user = false;
	if (data.token) {
		user = data.token.length > 0 ? false : true;
	}
</script>

<Header user={data.token} />
<main>
	<slot />
</main>

<Footer />

<style lang="scss">
	main {
		min-height: 100vh;
		padding: 2rem 0;
	}
</style>
