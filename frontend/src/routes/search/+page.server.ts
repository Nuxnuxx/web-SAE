import { getRecipes } from "$lib/api/auth-request";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ url }) => {
	const href = url.href;
	const urlString = new URL(href);
	const searchParams = urlString.searchParams;
	const name = searchParams.get("name") || "";
	const page: number = Number(searchParams.get("page"));
	const price = searchParams.get("price") || "";
	const difficulty = searchParams.get("difficulty") || "";

	try {
		const result = await getRecipes(name, price, difficulty, page);
		return {
			result,
		};
	} catch (err) {
		throw err;
	}
};
