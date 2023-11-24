import { Router } from "express";
import {
	handleRecipeById,
	handleRecipeByFilter,
	handleRecipeByKeyWord,
	handleRecipes,
} from "./recipe.controller.js";
import { validate } from "./validators/recipe.validator.js";
import {
	schemaByIdParams,
	schemaFilterParams,
	schemaFilterQuery,
	schemaKeywordParams,
	schemaPageParams,
} from "./validators/recipe.schema.js";

const recipeRouter = Router();

recipeRouter.get(
	"/id/:id",
	validate(schemaByIdParams, "params"),
	handleRecipeById
);

recipeRouter.get(
	"/page/:page",
	validate(schemaPageParams, "params"),
	handleRecipes
);

recipeRouter.get(
	"/search/:page/:keyword",
	validate(schemaKeywordParams, "params"),
	handleRecipeByKeyWord
);

recipeRouter.get(
	"/filter/:page",
	validate(schemaFilterParams, "params"),
	validate(schemaFilterQuery, "query"),
	handleRecipeByFilter
);

export default recipeRouter;
