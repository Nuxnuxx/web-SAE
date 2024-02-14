import type { PlaylistDetail } from "$lib/api/playlist-types";
import { writable } from "svelte/store";
import type { User } from "$lib/api/auth-types";
import { History } from "swiper/modules";

export const filterStore = writable({
	name: "",
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

export const playlistStore = writable<PlaylistDetail[]>([]);

export const userStore = writable(false);
export const userDataStore = writable<User>();

export const coldstartLiked = writable<number[]>([]);

type Popup = {
	idRecipe: number | null;
	type: string | null;
};

export const popup = writable<Popup>({
	idRecipe: null,
	type: null,
});

type HistoryEntry = {
	idRecipe: number;
	name: string;
};

type ParcoursStore = {
	actualDay: number;
	history: HistoryEntry[];
};

export const parcours = writable<ParcoursStore>({
	actualDay: new Date().getDay(),
	history: [],
});
