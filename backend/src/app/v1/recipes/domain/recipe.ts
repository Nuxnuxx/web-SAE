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
	//TODO: Impletement pagination
	const recipes = await getRecipesInDatabaseWithPagination(page);
	return recipes;
};

export const findRecipesByKeyWord = async (keyWord: string[]) => {
	//TODO: Implement query by keyword on name and ingredients to neo4j with pagination if more than X results
	const recipes = getRecipesByKeyWordInDatabase(keyWord);
	return recipes;
};

export const findRecipesByFilter = async (filter: string[]) => {
	//TODO: Implement query by filter to neo4j with pagination if more than X results
	const recipes = getRecipesByFilterInDatabase(filter);
	return recipes;
};
