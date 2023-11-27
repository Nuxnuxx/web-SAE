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
	const idRecipe = req.query.idrecipe;

	// INFO: A enlever quand le schema de verification sera la
	if (typeof idRecipe !== "string") {
		throw new ApiError(
			"Id is not a string",
			"ID_IS_NOT_A_STRING",
			400,
			true
		);
	}

	const parsedId = parseInt(idRecipe, 10);

	try {
		//@ts-ignore
		const result = await createRecipeLiked(parsedId, req.user.email);
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
	const idRecipe = req.query.idrecipe;
	const idList = req.query.idlist;

	//@ts-ignore
	const parsedIdRecipe = parseInt(idRecipe, 10);
	//@ts-ignore
	const parsedIdList = parseInt(idList, 10);

	try {
		//@ts-ignore
		const result = await createRecipeList(
			parsedIdRecipe,
			parsedIdList,
			req.user.email
		);
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
	const id = req.query.id;

	const parsedId = parseInt(id, 10);

	try {
		//@ts-ignore
		const result = await getRecipesList(parsedId, req.user.email);
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
	const idList = req.query.idlist;
	const idRecipe = req.query.idrecipe;

	//INFO: A enlever quand le schema de verification sera la
	if (typeof idList !== "string") {
		throw new ApiError(
			"Id is not a string",
			"ID_IS_NOT_A_STRING",
			400,
			true
		);
	}

	//INFO: A enlever quand le schema de verification sera la
	if (typeof idRecipe !== "string") {
		throw new ApiError(
			"Id is not a string",
			"ID_IS_NOT_A_STRING",
			400,
			true
		);
	}

	const parsedIdList = parseInt(idList);
	const parsedIdRecipe = parseInt(idRecipe);

	try {
		const result = await deleteRecipeList(
			parsedIdList,
			parsedIdRecipe,
			req.user.email
		);
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
