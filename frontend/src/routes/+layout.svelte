<script lang="ts">
	import Footer from "../components/footer.svelte";
	import Header from "../components/header.svelte";
	import "$lib/global.scss";
	import type { LayoutData } from "./$types";
	import { userDataStore, userStore } from "../store";

	export let data: LayoutData;

	let user = false;
	$: if (data.token) {
		userStore.set(true);
		userDataStore.set({ ...data.profil });
	}
</script>

<Header />
<main>
	<slot />
</main>

<Footer />

<style lang="scss">
	main {
		min-height: 100vh;
	}
</style>
