import { Router } from "express";
import { authenticate } from "./list.auth.js";
import { handleCreateList, handleDeleteList, handleGetList, handleModifyList } from "./list.controller.js";
import listRecipeRouter from "./list-recipe/list.recipe.router.js";

const listRouter = Router();

listRouter.use(authenticate);

listRouter.post("/createlist", handleCreateList);
listRouter.put("/modifylist", handleModifyList);
listRouter.delete("/deletelist", handleDeleteList);
listRouter.get("/displaylists", handleGetList);

listRouter.use("/recipe", listRecipeRouter);

export default listRouter;
