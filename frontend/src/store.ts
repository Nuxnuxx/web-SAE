import type { PlaylistDetail } from "$lib/api/playlist-types";
import { writable } from "svelte/store";
import type { User } from "$lib/api/auth-types";

export const filterStore = writable({
	name: "",
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

export const playlistStore = writable<PlaylistDetail[]>([]);

export const userStore = writable(false);
export const userDataStore = writable<User>();

export const isPlaylistAddButtonOpen = writable({
	id: 0,
	open: false,
});

export const isLikedButtonOpen = writable({
	id: 0,
	open: false,
});

export const coldstartLiked = writable<number[]>([]);
