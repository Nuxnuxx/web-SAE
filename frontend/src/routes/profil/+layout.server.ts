import type { PageServerLoad } from "./$types";
import { getProfil } from "$lib/api/auth-request";
import { redirect } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (!token) {
		throw redirect(302, "/");
	}

	try {
		const profil = await getProfil(token);
		return {
			profil,
		};
	} catch (err) {
		throw err;
	}
};
