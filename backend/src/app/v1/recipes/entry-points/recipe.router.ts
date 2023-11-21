import { Router } from "express";
import {
	getRecipeByIdController,
	getRecipesByKeyWordController,
	getRecipesController,
} from "./recipe.controller.js";

const recipeRouter = Router();

recipeRouter.get("/id/:id", getRecipeByIdController);
recipeRouter.get("/:page", getRecipesController);
recipeRouter.get("/search", getRecipesByKeyWordController);
recipeRouter.get("/filter", getRecipesByKeyWordController);

export default recipeRouter;
