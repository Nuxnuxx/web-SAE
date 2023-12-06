import { getRecipes } from "$lib/api/auth-request";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ url }) => {
	const href = url.href;
	const queryString = url.toString();

	const name = queryString.split("?")[1]
	const finalName = name.split("=")[1];
	try {
		const result = await getRecipes(finalName);
		return {
			result,
		};
	} catch (err) {
		throw err;
	}
};
