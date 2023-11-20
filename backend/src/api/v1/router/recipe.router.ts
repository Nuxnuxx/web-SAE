import { Router } from "express";
import {
	getRecipeByIdController,
	getRecipesByKeyWordController,
	getRecipesController,
} from "../controllers/recipe.controller.js";
import { findRecipesByKeyWord } from "../../../services/recipe.js";

const recipeRouter = Router();

recipeRouter.get("/id/:id", getRecipeByIdController);
recipeRouter.get("/", getRecipesController);
recipeRouter.get("/search", getRecipesByKeyWordController);
recipeRouter.get("/filter", getRecipesByKeyWordController);

export default recipeRouter;
