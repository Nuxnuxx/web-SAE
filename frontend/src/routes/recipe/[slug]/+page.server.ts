import { getRecipe } from "$lib/api/recipe-request";
import { error } from "@sveltejs/kit";
import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import {
	getSimilarRecipes,
	recipeTimeSpent,
} from "$lib/api/recommendation-request";

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

export const actions: Actions = {
	timeSpent: async ({ request, cookies }) => {
		try {
			console.log("request");
			const token = cookies.get("token");

			if (token === undefined || token === null) {
				throw new Error("token is not defined");
			}

			const body = await request.formData();
			const timeSpent = Number(body.get("timeSpent"));
			const idRecipe = Number(body.get("idRecipe"));

			if (
				idRecipe === undefined ||
				idRecipe === null ||
				isNaN(idRecipe)
			) {
				throw new Error("recipeId is not defined");
			}

			if (
				timeSpent === undefined ||
				timeSpent === null ||
				isNaN(timeSpent)
			) {
				throw new Error("timeSpent is not defined");
			}

			console.log("timeSpent", timeSpent);
			const result = await recipeTimeSpent(token, idRecipe, timeSpent);
			return {
				result,
			};
		} catch (err) {
			throw err;
		}
	},
};
