import { Router } from "express";
import {
	handleRecipeById,
	handleRecipeByFilter,
	handleRecipeByKeyWord,
	handleRecipes,
} from "./recipe.controller.js";

const recipeRouter = Router();

recipeRouter.get("/id/:id", handleRecipeById);
recipeRouter.get("/page/:page", handleRecipes);
recipeRouter.get("/search/:keyword", handleRecipeByKeyWord);
recipeRouter.get("/filter", handleRecipeByFilter);

export default recipeRouter;
