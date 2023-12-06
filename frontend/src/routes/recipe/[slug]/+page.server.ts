import { getRecipe } from "$lib/api/recipe-request";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
	const id = Number(params.slug);
	return {
		recipe: await getRecipe(id),
	};
};
