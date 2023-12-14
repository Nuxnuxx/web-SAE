import { createPlaylist } from "$lib/api/playlist-request";
import { error } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
	createList: async ({ cookies, request }) => {
		try {
			const token = cookies.get("token") || "";
			const body = await request.formData();
			const name = body.get("name")?.toString() || "";
			const result = await createPlaylist(token, name);
			return {
				result,
			};
		} catch (err) {
			throw error(404, "Nop");
		}
	},
};
