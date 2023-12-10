import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getPlaylist } from "$lib/api/playlist-request";

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (!token) {
		throw redirect(302, "/");
	}

	try {
		const playlists = await getPlaylist(token);
		return {
			playlists,
		};
	} catch (err) {
		throw err;
	}
};
