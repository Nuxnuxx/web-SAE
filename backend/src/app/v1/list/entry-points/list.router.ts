import { Router } from "express";
import { authenticate } from "./list.auth.js";
import {
	handleCreateList,
	handleDeleteList,
	handleGetList,
	handleModifyList,
} from "./list.controller.js";
import listRecipeRouter from "./list-recipe/list.recipe.router.js";
import { validate } from "./validators/list.validator.js";
import {
	schemaCreateListQuery,
	schemaDeleteListQuery,
	schemaModifyListQuery,
} from "./validators/list.schema.js";

const listRouter = Router();

listRouter.use(authenticate);

listRouter.post(
	"/createlist",
	validate(schemaCreateListQuery, "query"),
	handleCreateList
);
listRouter.put(
	"/modifylist",
	validate(schemaModifyListQuery, "query"),
	handleModifyList
);
listRouter.delete(
	"/deletelist",
	validate(schemaDeleteListQuery, "query"),
	handleDeleteList
);
listRouter.get("/displaylists", handleGetList);

listRouter.use("/recipe", listRecipeRouter);

export default listRouter;
