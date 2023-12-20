import { writable } from "svelte/store";
import type { PlaylistDetail } from "$lib/api/playlist-types";

export const filterStore = writable({
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

export const playlistStore = writable<PlaylistDetail[]>([]);
