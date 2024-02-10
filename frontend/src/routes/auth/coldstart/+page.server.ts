import { sendColdstart } from "$lib/api/auth-request";
import { schemaColdstart } from "$lib/api/auth-schema";
import { likeRecipe } from "$lib/api/recipe-request";
import { fail, type Actions, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { get } from "svelte/store";
import { coldstartLiked } from "../../../store";

export const load: PageServerLoad = async ({ cookies }) => {
	// vérifie si le coldstart a déjà été fait
	const coldstart = cookies.get("coldstart");

	// si pas de coldstart et que le token est présent, on redirige vers coldstart
	if (coldstart) {
		throw redirect(302, "/");
	}
};

export const actions: Actions = {
	// redirect to the /recipe page to finish the coldstart
	endColdstart: async ({ request }) => {
		//transmettre le prix et difficulté
		const body = await request.formData();
		const coldstart = {
			price: body.get("price")?.toString() || "",
			difficulty: body.get("difficulty")?.toString() || "",
		};

		//redirect to the /recipe page with the coldstart data
		throw redirect(
			302,
			"/auth/coldstart/recipe?" +
				new URLSearchParams(coldstart).toString()
		);
	},
};
