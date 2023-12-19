import type { LayoutServerLoad } from "./$types";
import { getPlaylist } from "$lib/api/playlist-request";

export const load: LayoutServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (token){
		try {
			const playlists = await getPlaylist(token);
			return {
				token,
				playlists,
			};
		} catch (err) {
			throw err;
		}
	}
	return {
		token,
	};
};
