import {
	createRecipeListInDatabase,
	deleteRecipeListInDatabase,
	getRecipesListInDatabase,
	createRecipeInLikedInDatabase,
} from "../data-access/list-recipe.js";

export const createRecipeLiked = async (id: number, name: string) => {
	const result = createRecipeInLikedInDatabase(id, name);
	return result;
};

export const createRecipeList = async (
	name: string,
	nameList: string,
	email: string
) => {
	const result = await createRecipeListInDatabase(name, nameList, email);
	return result;
};

export const getRecipesList = async (id: number) => {
	const result = getRecipesListInDatabase(id);
	return result;
};

export const deleteRecipeList = async (id: number) => {
	const result = deleteRecipeListInDatabase(id);
	return result;
};
