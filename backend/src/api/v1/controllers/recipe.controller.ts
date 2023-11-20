import { Request, Response } from "express";
import {
	findRecipesByFilter,
	findRecipesByKeyWord,
	getRecipeById,
	getRecipes,
} from "../../../services/recipe.js";

//TODO: Try catch tout ca tout ca

export const getRecipeByIdController = async (req: Request, res: Response) => {
	const { id } = req.params;
	const parsedInt = parseInt(id, 10);
	const result = await getRecipeById(parsedInt);
	res.send(result);
};

export const getRecipesController = async (req: Request, res: Response) => {
	const result = await getRecipes();
	res.send(result);
};

export const getRecipesByKeyWordController = async (
	req: Request,
	res: Response
) => {
	const { searchKey } = req.query;

	const keywords: string[] = Array.isArray(searchKey)
		? searchKey.map((key: any) => (typeof key === "string" ? key : ""))
		: typeof searchKey === "string"
		  ? [searchKey]
		  : [];

	const result = await findRecipesByKeyWord(keywords);
	res.send(result);
};

export const getRecipesByFilterController = async (
	req: Request,
	res: Response
) => {
	const { searchKey: filter } = req.query;

	const keywords: string[] = Array.isArray(filter)
		? filter.map((key: any) => (typeof key === "string" ? key : ""))
		: typeof filter === "string"
		  ? [filter]
		  : [];

	const result = await findRecipesByFilter(keywords);
	res.send(result);
};
