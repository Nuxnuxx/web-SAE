import { Router } from "express";
import { getProduct } from "./recipe.controller.js";

const recipeRouter = Router();

recipeRouter.get("/:id", getProduct);

export default recipeRouter;
