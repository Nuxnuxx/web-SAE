import type { Actions } from "@sveltejs/kit";

export const actions: Actions = {
	logout: async ({ cookies }) => {
		cookies.set("token", "");
	},
};
