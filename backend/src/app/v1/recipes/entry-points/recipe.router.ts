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
recipeRouter.get("/search/:page/:keyword", handleRecipeByKeyWord);
recipeRouter.get("/filter/:page", handleRecipeByFilter);

export default recipeRouter;
