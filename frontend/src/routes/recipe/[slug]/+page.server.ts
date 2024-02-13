import { getRecipe } from "$lib/api/recipe-request";
import { error } from "@sveltejs/kit";
import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import {
	getSimilarRecipes,
	recipeTimeSpent,
} from "$lib/api/recommendation-request";

export const load: PageServerLoad = async ({ params, cookies }) => {
	const token = cookies.get("token") || "";
	const id = Number(params.slug);
	const recipe = await getRecipe(id);
	const numberOfRecipesSimilar = 8;
	if (recipe) {
		return {
			token: token,
			recipe: await getRecipe(id),
			similarRecipe: await getSimilarRecipes(id, numberOfRecipesSimilar),
		};
	}

	error(404, "Not found");
};

export const actions: Actions = {
	timeSpent: async ({ request, cookies }) => {
		try {
			const token = cookies.get("token");

			if (token === undefined || token === null) {
				throw new Error("token is not defined");
			}

			const body = await request.formData();
			const timeSpent = Number(body.get("timeSpent"));
			const recipeId = Number(body.get("recipeId"));

			if (
				recipeId === undefined ||
				recipeId === null ||
				isNaN(recipeId)
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

			const result = await recipeTimeSpent(token, recipeId, timeSpent);
			return {
				result,
			};
		} catch (err) {
			throw err;
		}
	},
};
