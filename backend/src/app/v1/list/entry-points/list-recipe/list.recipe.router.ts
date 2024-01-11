import { Router } from "express";
import {
	handleAddRecipeLiked,
	handleAddRecipeList,
	handleDeleteRecipeList,
	handleGetRecipeList,
} from "./list.recipe.controller.js";
import { validate } from "./validators/list-recipe.validator.js";
import {
	schemaCreateRecipeLikedQuery,
	schemaCreateRecipeListQuery,
	schemaDeleteRecipeListQuery,
	schemaGetRecipeListQuery,
} from "./validators/list-recipe.schema.js";

const listRecipeRouter = Router();

listRecipeRouter.post(
	"/addlike",
	validate(schemaCreateRecipeLikedQuery, "query"),
	handleAddRecipeLiked
);

listRecipeRouter.post(
	"/addlist",
	validate(schemaCreateRecipeListQuery, "query"),
	handleAddRecipeList
);

listRecipeRouter.get(
	"/list",
	validate(schemaGetRecipeListQuery, "query"),
	handleGetRecipeList
);

listRecipeRouter.delete(
	"/list",
	validate(schemaDeleteRecipeListQuery, "query"),
	handleDeleteRecipeList
);

export default listRecipeRouter;
