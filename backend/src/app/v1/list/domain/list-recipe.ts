import {
	createRecipeInListInDatabase,
	deleteRecipeListInDatabase,
	createRecipeInLikedInDatabase,
	getRecipesListInDatabase,
} from "../data-access/list-recipe.js";
import { checkListBelongToUser } from "../data-access/list.js";

export const createRecipeLiked = async (id: number, email: string) => {
	const result = createRecipeInLikedInDatabase(id, email);
	return result;
};

export const createRecipeList = async (
	idRecipe: number,
	idList: number,
	email: string
) => {
	await checkListBelongToUser(idList, email);
	const result = await createRecipeInListInDatabase(idRecipe, idList, email);
	return result;
};

export const getRecipesList = async (id: number, email: string) => {
	await checkListBelongToUser(id, email);
	const result = getRecipesListInDatabase(id);
	return result;
};

export const deleteRecipeList = async (
	idList: number,
	idRecipe: number,
	email: string
) => {
	await checkListBelongToUser(idList, email);
	const result = deleteRecipeListInDatabase(idList, idRecipe);
	return result;
};
