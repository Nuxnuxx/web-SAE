import { sendColdstart } from "$lib/api/auth-request";
import { schemaColdstart } from "$lib/api/auth-schema";
import { fail, type Actions, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ cookies }) => {
	// vérifie si le coldstart a déjà été fait
	const coldstart = cookies.get("coldstart");

	// si pas de coldstart et que le token est présent, on redirige vers coldstart
	if (coldstart) {
		throw redirect(302, "/profil");
	}
};

export const actions: Actions = {
	coldstart: async ({ cookies, request }) => {
		const body = await request.formData();
		const coldstart = {
			price: body.get("price")?.toString() || "",
			difficulty: body.get("difficulty")?.toString() || "",
		};

		try {
			await schemaColdstart.validate(coldstart, {
				abortEarly: false,
			});
			const token = cookies.get("token");
			if (token === undefined || token === null) {
				throw new Error("token is not defined");
			}
			const result = await sendColdstart(
				token,
				coldstart.price,
				coldstart.difficulty
			);
			if (result) {
				cookies.set("coldstart", "true", {
					path: "/",
				});
				return {
					location: "/",
				};
			}
		} catch (err) {
			return fail(400);
		}
	},
};
