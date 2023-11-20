import { Recipe } from "..";
import models from "../models/models";

//FIXME: delete this when neo4j setup
const recipes: Recipe[] = [
	{ name: "Poulet a la creme", id: 1, numberLike: 0, urlImage: "" },
	{ name: "Sanglier au miel", id: 4, numberLike: 0, urlImage: "" },
	{ name: "Saumon au curry", id: 3, numberLike: 0, urlImage: "" },
	{ name: "Poulet au curry", id: 2, numberLike: 0, urlImage: "" },
];

export const getRecipeById = async (id: number) => {
	//TODO: Implement query to neo4j
	const recipe: Recipe = {
		name: "Poulet a la creme",
		id: 1,
		numberLike: 0,
		urlImage: "",
	};
	return recipe;
};

export const getRecipes = async () => {
	//TODO: Impletement query to neo4j with pagination
	return recipes;
};

export const findRecipesByKeyWord = async (keyWord: string[]) => {
	//TODO: Implement query by keyword on name and ingredients to neo4j with pagination if more than X results
	return recipes;
};

export const findRecipesByFilter = async (filter: string[]) => {
	//TODO: Implement query by filter to neo4j with pagination if more than X results
	return recipes;
};
