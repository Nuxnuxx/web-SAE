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
		res.status(200).type("json").send(result);
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
	const nameList = req.query.namelist;
	const idList = req.query.idlist;

	//@ts-ignore
	const parsedId = parseInt(idList, 10);

	try {
		//@ts-ignore
		const result = await modifyList(parsedId, nameList, req.user.email);
		res.status(200).type("json").send(result);
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
	const idList = req.query.idlist;

	//@ts-ignore
	const parsedId = parseInt(idList, 10);

	try {
		//@ts-ignore
		const result = await deleteList(parsedId, req.user.email);
		res.status(200).type("json").send(result);
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
	try {
		//@ts-ignore
		const result = await getList(req.user.email);
		res.status(200).type("json").send(result);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};
