import { Integer } from "neo4j-driver";

export const getRecipeById = async (id: Integer) => {
	// const recipe = await models.run("recipes", "get", { id })
	const recipe = "Poulet a la creme";
	return recipe;
};
