import { NextFunction, Request, Response } from "express";
import {
	findRecipesByFilter,
	findRecipesByKeyWord,
	getRecipeById,
	getRecipes,
} from "../domain/recipe.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";
import { Filter } from "../recipe.js";

export const handleRecipeById = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	try {
		const { id } = req.params;
		const parsedInt = parseInt(id, 10);

		try {
			const result = await getRecipeById(parsedInt);
			res.status(200).send(result);
		} catch (err) {
			const message =
				err instanceof ApiError ? err.message : "Internal Server Error";
			res.status(err instanceof ApiError ? err.httpCode : 500).send({
				message,
			});
			next(err);
		}
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleRecipes = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const page = parseInt(req.params.page, 10);

	try {
		const result = await getRecipes(page);
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

export const handleRecipeByKeyWord = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const keyWord = req.params.keyword;
	const page = parseInt(req.params.page, 10);

	try {
		const result = await findRecipesByKeyWord(keyWord, page);
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

export const handleRecipeByFilter = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const filter: Filter = req.query;
	const page: number = parseInt(req.params.page, 10);

	try {
		const result = await findRecipesByFilter(filter, page);
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
