import { getProfil } from "$lib/api/auth-request";
import type { LayoutServerLoad } from "./$types";
import { getPlaylist } from "$lib/api/playlist-request";

export const load: LayoutServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (token) {
		try {
			const playlists = await getPlaylist(token);
			const profil = await getProfil(token);
			return {
				token,
				playlists,
				profil,
			};
		} catch (err) {
			throw err;
		}
	}
};
