import { Router } from "express";
import { getProduct } from "../controllers/recipe.controller.js";

const recipeRouter = Router()

recipeRouter.get("/", getProduct)

export default recipeRouter
