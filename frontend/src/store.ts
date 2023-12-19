import { writable } from "svelte/store";
import { getPlaylist } from "$lib/api/playlist-request";

export const filterStore = writable({
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

// Permet d'obtenir toutes les playlists de l'utilisateur
export function getPlaylists(token: string) {
	return getPlaylist(token);
}