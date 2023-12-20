<script lang="ts">
	import Footer from "../components/footer.svelte";
	import Header from "../components/header.svelte";
	import "$lib/global.scss";
	import type { LayoutData } from "./$types";
	import { playlistStore, userDataStore, userStore } from "../store";
	import type { PlaylistDetail } from "$lib/api/playlist-types";

	export let data: LayoutData;

	let user = false;
	$: if (data.token) {
		user = data.token.length > 0 ? false : true;
		const playlistList: PlaylistDetail[] = data.playlists.result.sort(
			(a: PlaylistDetail, b: PlaylistDetail) => {
				return a.name === "liked" ? -1 : b.name === "liked" ? 1 : 0;
			}
		);
		userStore.set(true);
		userDataStore.set({ ...data.profil });

		playlistStore.set(playlistList);
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
