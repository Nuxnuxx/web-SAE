import { getProfil } from "$lib/api/auth-request";
import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "../$types";

export const load: LayoutServerLoad = async ({ parent }) => {
	const { token } = await parent();

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
