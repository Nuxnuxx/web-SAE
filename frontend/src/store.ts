import { writable } from "svelte/store";

export const filterStore = writable({
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

// Permet d'obtenir toutes les playlists de l'utilisateur
export const playlistStore = writable();