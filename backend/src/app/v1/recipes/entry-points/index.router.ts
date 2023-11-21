import express from "express";
import recipeRouter from "./recipe.router.js";
import authRouter from "./auth.router.js";

const v1Router = express.Router();

v1Router.use("/recipe", recipeRouter);
v1Router.use("/auth", authRouter);

export default v1Router;
