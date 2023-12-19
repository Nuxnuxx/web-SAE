<script lang="ts">
	import Footer from "../components/footer.svelte";
	import Header from "../components/header.svelte";
	import "$lib/global.scss"
	import type { LayoutData } from "./$types";
	import { playlistStore } from "../store";

	export let data: LayoutData;

	let user = false;
	if (data.token) {
		user = data.token.length > 0 ? false : true;
		const playlistList = data.playlists.result.sort((a, b) => {
			return a.name === "liked" ? -1 : b.name === "liked" ? 1 : 0;
		});

		playlistStore.set(playlistList);
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
	}
</style>
