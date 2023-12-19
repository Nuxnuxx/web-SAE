import { getRecipes } from "$lib/api/auth-request";
import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { likeRecipe } from "$lib/api/recipe-request";
import { addPlaylistRecipe } from "$lib/api/recipe-request";

export const load: PageServerLoad = async ({ url }) => {
	const href = url.href;
	const urlString = new URL(href);
	const searchParams = urlString.searchParams;
	const name = searchParams.get("name") || "";
	const page: number = Number(searchParams.get("page"));
	const price = searchParams.get("price") || "";
	const difficulty = searchParams.get("difficulty") || "";
	const oldData = {
		name,
		page,
		price,
		difficulty,
	};

	try {
		const result = await getRecipes(name, price, difficulty, page);
		return {
			result,
			oldData,
		};
	} catch (err) {
		throw err;
	}
};

export const actions: Actions = {
	likeRecipe: async ({ request, cookies }) => {
		try {
			const token = cookies.get("token") || "";
			const body = await request.formData();
			const id = body.get("id");
			const result = await likeRecipe(token, id);
			return {
				result,
			};
		} catch (err) {
			throw err;
		}
	},

	addPlaylistRecipe: async ({ request, cookies }) => {
		try {
			const token = cookies.get("token") || "";
			const body = await request.formData();
			const id = body.get("id");
			const idlist = body.get("idlist");
			const result = await addPlaylistRecipe(token, id, idlist);
			return {
				result,
			};
		} catch (err) {
			throw err;
		}
	},
};
