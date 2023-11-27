import { Router } from "express";
import { handleAddRecipeLiked, handleAddRecipeList, handleDeleteRecipeList, handleGetRecipeList } from "./list.recipe.controller.js";

const listRecipeRouter = Router();

listRecipeRouter.post("/addlike", handleAddRecipeLiked);
listRecipeRouter.post("/addlist", handleAddRecipeList);
listRecipeRouter.get("/list", handleGetRecipeList);
listRecipeRouter.delete("/list", handleDeleteRecipeList);

export default listRecipeRouter;
