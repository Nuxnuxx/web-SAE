import { NextFunction, Request, Response } from "express";
import { ApiError } from "../../../../../libraries/error-handling/api-error.js";
import {
	createRecipeLiked,
	createRecipeList,
	deleteRecipeList,
	getRecipesList,
} from "../../domain/list-recipe.js";

export const handleAddRecipeLiked = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { idRecipe } = req.query;

	try {
		//@ts-ignore
		const result = await createRecipeLiked(idRecipe, req.user.email);
		res.status(200).send(result);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleAddRecipeList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { idRecipe, idList } = req.query;

	try {
		//@ts-ignore
		const result = await createRecipeList(idRecipe, idList, req.user.email);
		res.status(200).send(result);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleGetRecipeList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { id } = req.query;

	try {
		//@ts-ignore
		const result = await getRecipesList(id, req.user.email);
		res.status(200).send(result);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleDeleteRecipeList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { id } = req.query;

	try {
		//@ts-ignore
		const result = await deleteRecipeList(id, req.user.email);
		res.status(200).send(result);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};
