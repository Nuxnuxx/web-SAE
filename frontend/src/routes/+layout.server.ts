import { getProfil } from "$lib/api/auth-request";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (token) {
		try {
			const profil = await getProfil(token);
			return {
				token,
				profil,
			};
		} catch (e) {
			console.log(e);
		}
	}
	return {
		token,
	};
};
