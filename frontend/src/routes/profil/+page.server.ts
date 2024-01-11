import type { Actions } from "@sveltejs/kit";
import { userDataStore, userStore } from "../../store";

export const actions: Actions = {
	logout: async ({ cookies }) => {
		cookies.set("token", "", {
			path: "/",
		});
		userStore.set(false);
		//@ts-ignore
		userDataStore.set({});
	},
};
