import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "../$types";
import { getPlaylistRecipe } from "$lib/api/playlist-request";

export const load: PageServerLoad = async ({ parent, params }) => {
	const { token, playlists } = await parent();
	const id = Number(params.slug);

	if (!token) {
		throw redirect(302, "/auth");
	}

	return {
		playlist: await getPlaylistRecipe(token, id),
		playlists: playlists.result.filter((x) => x.idPlaylist == id),
	};
};
