import {
	getRecipesByFilterInDatabase,
	getRecipesByIdInDatabase,
	getRecipesByKeyWordInDatabase,
	getRecipesInDatabaseWithPagination,
} from "../data-access.js";

export const getRecipeById = async (id: number) => {
	const recipe = await getRecipesByIdInDatabase(id);
	return recipe;
};

export const getRecipes = async (page: number) => {
	const recipes = await getRecipesInDatabaseWithPagination(page);
	return recipes;
};

export const findRecipesByKeyWord = async (keyWordArray: string) => {
	//TODO: implement with multiple keywords and pagination
	const recipes = await getRecipesByKeyWordInDatabase(keyWordArray);
	return recipes;
};

export const findRecipesByFilter = async (filter: string[]) => {
	//TODO: Implement pagination
	const recipes = getRecipesByFilterInDatabase(filter);
	return recipes;
};
