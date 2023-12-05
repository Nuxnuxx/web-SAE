import { redirect, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (!token) {
		throw redirect(302, "/");
	}
};

export const actions: Actions = {
	logout: async ({ cookies }) => {
		cookies.set("token", "");
	},
};
