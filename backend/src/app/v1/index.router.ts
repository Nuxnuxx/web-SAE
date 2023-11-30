import { Router } from "express";
import recipeRouter from "./recipes/entry-points/recipe.router.js";
import authRouter from "./auth/entry-points/auth.router.js";
import listRouter from "./list/entry-points/list.router.js";

const v1Router = Router();

v1Router.use("/recipe", recipeRouter);
v1Router.use("/auth", authRouter);
v1Router.use("/list", listRouter);

export default v1Router;
