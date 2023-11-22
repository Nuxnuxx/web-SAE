import { Router } from "express";
import recipeRouter from "./recipes/entry-points/recipe.router.js";

const v1Router = Router();

v1Router.use("/recipe", recipeRouter);

export default v1Router;
