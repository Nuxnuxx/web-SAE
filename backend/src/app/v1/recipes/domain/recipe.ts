import {
	getRecipesByFilterInDatabase,
	getRecipesByIdInDatabase,
	getRecipesByKeyWordInDatabase,
	getRecipesInDatabaseWithPagination,
} from "../data-access/data-access.js";
import { Filter } from "../recipe.js";


export const getRecipeById = async (id: number) => {
	const recipe = await getRecipesByIdInDatabase(id);
	return recipe;
};

export const getRecipes = async (page: number) => {
	const recipes = await getRecipesInDatabaseWithPagination(page);
	return recipes;
};

export const findRecipesByKeyWord = async (keyWordArray: string, page: number) => {
	const recipes = await getRecipesByKeyWordInDatabase(keyWordArray, page);
	return recipes;
};

export const findRecipesByFilter = async (filter: Filter, page: number) => {
	const recipes = await getRecipesByFilterInDatabase(filter, page);
	return recipes;
};
