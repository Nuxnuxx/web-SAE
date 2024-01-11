import { getRecipe } from "$lib/api/recipe-request";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { getSimilarRecipes } from "$lib/api/recommendation-request";

export const load: PageServerLoad = async ({ params }) => {
	const id = Number(params.slug);
	const recipe = await getRecipe(id);
	const numberOfRecipesSimilar = 8;
	if (recipe) {
		return {
			recipe: await getRecipe(id),
			similarRecipe: await getSimilarRecipes(id, numberOfRecipesSimilar),
		};
	}

	error(404, "Not found");
};
