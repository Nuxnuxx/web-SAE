import type { PlaylistDetail } from "$lib/api/playlist-types";
import { writable } from "svelte/store";
import type { User } from "$lib/api/auth-types";

export const filterStore = writable({
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

export const playlistStore = writable<PlaylistDetail[]>([]);

export const userStore = writable(false);
export const userDataStore = writable<User>();

export const isPLaylistAddButtonOpen = writable({
	id: 0,
	open: false,
});
