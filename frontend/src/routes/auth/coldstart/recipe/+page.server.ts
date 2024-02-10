import { coldstartLiked } from "../../../../store";
import { sendColdstart } from "$lib/api/auth-request";
import { schemaColdstart } from "$lib/api/auth-schema";
import { likeRecipe } from "$lib/api/recipe-request";
import { fail, type Actions, error } from "@sveltejs/kit";
import type { PageServerLoad } from "../$types";
import { redirect } from "@sveltejs/kit";
import { getRecipes } from "$lib/api/auth-request";
import { Price, Difficulty } from "$lib/api/recipe-types";

export const ssr = false;

export const load: PageServerLoad = async ({ cookies, url }) => {
	// vérifie si le coldstart a déjà été fait
	const coldstart = cookies.get("coldstart");

	// si pas de coldstart et que le token est présent, on redirige vers coldstart
	if (coldstart) {
		throw redirect(302, "/");
	}

	// récupérer prix et difficulté en get
	const price: number = +(url.searchParams.get("price") || "bon marché");
	const difficulty: number = +(
		url.searchParams.get("difficulty") || "très facile"
	);

	const priceEnum: string = Price[price];
	const difficultyEnum: string = Difficulty[difficulty];

	let recipes: any[] = [];
	let jsonRecipes: any;

	jsonRecipes = await getRecipes("", priceEnum, difficultyEnum, 0);
	recipes = jsonRecipes.result;
	if (recipes.length < 4) {
		jsonRecipes = await getRecipes("", priceEnum, "", 0);
		recipes = jsonRecipes.result;
	}
	return {
		price: price,
		difficulty: difficulty,
		recipes: recipes,
	};

	error(404, "No recipes found");
};

export const actions: Actions = {
	coldstart: async ({ cookies, request }) => {
		console.log("coldstart");
		const body = await request.formData();
		const coldstart = {
			price: body.get("price")?.toString() || "",
			difficulty: body.get("difficulty")?.toString() || "",
		};

		//const liked = body.get("liked")?.toString().split(",") || [];

		try {
			console.log("coldstart2");
			await schemaColdstart.validate(coldstart, {
				abortEarly: false,
			});
			console.log("coldstart3");
			const token = cookies.get("token");
			if (token === undefined || token === null) {
				throw new Error("token is not defined");
			}
			console.log("coldstart4");

			/*for (const id of liked) {
				await likeRecipe(token, Number(id));
			}*/
			const result = await sendColdstart(
				token,
				coldstart.price,
				coldstart.difficulty
			);
			if (result) {
				cookies.set("coldstart", "true", {
					path: "/",
				});
				return {
					location: "/",
				};
			}
		} catch (err) {
			return fail(400);
		}
	},
};
