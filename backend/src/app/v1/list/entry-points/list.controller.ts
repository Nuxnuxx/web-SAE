import { NextFunction, Request, Response } from "express";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";
import { createList, deleteList, getList, modifyList } from "../domain/list.js";

export const handleCreateList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const nameList = req.query.namelist;

	try {
		//@ts-ignore
		const result = await createList(nameList, req.user.email);
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

export const handleModifyList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { nameList, idList } = req.query;

	try {
		//@ts-ignore
		const result = await modifyList(idList, nameList, req.user.mail);
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

export const handleDeleteList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { idList } = req.query;

	try {
		//@ts-ignore
		const result = await deleteList(idList, req.user.email);
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

export const handleGetList = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const { idList } = req.query;

	try {
		//@ts-ignore
		const result = await getList(idList, req.user.email);
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
