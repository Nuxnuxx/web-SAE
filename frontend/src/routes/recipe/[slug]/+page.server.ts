import { getRecipe } from "$lib/api/recipe-request";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
	const id = Number(params.slug);
	const recipe = await getRecipe(id);
	if (recipe) {
		return {
			recipe: await getRecipe(id),
		};
	}

	error(404, "Not found");
};
