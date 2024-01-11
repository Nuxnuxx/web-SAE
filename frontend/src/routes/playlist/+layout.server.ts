import { redirect } from "@sveltejs/kit";
import { getPlaylist } from "$lib/api/playlist-request";
import type { LayoutServerLoad } from "../$types";

export const load: LayoutServerLoad = async ({ parent }) => {
	const { token } = await parent();

	if (!token) {
		throw redirect(302, "/auth");
	}
};
