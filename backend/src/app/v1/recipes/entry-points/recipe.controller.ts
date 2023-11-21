import { Request, Response } from "express";
import {
	findRecipesByFilter,
	findRecipesByKeyWord,
	getRecipeById,
	getRecipes,
} from "../domain/recipe.js";

//TODO: Try catch tout ca tout ca

export const handleRecipeById = async (req: Request, res: Response) => {
	const { id } = req.params;
	const parsedInt = parseInt(id, 10);

	const result = await getRecipeById(parsedInt);

	res.send(result);
};

export const handleRecipes = async (req: Request, res: Response) => {
	const page = parseInt(req.params.page, 10);

	const result = await getRecipes(page);

	res.send(result);
};

export const handleRecipeByKeyWord = async (req: Request, res: Response) => {
	const keyWord = req.params.keyword;

	const result = await findRecipesByKeyWord(keyWord);

	res.send(result);
};

export const handleRecipeByFilter = async (req: Request, res: Response) => {
	//FIXME: filter not typed a string[]
	const filter: any = req.query;

	const result = await findRecipesByFilter(filter);

	res.send(result);
};
