import { NextFunction, Request, Response } from "express";
import {
	loginUser,
	registerUser,
	modifyUser,
	deleteUser,
} from "../domain/auth.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";

export const handleRegister = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const user = {
		name: req.body.name,
		email: req.body.email,
		password: req.body.password,
	};

	try {
		const register = await registerUser(user);
		res.status(200).send(register);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleLogin = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const user = req.body;

	try {
		const login = await loginUser(user);
		res.status(200).send(login);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleModifyProfil = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const user = req.body;

	try {
		const modify = await modifyUser(user);
		res.status(200).send(modify);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};

export const handleDeleteProfil = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const user = req.body;

	try {
		const deleteProfil = await deleteUser(user);
		res.status(200).send(deleteProfil);
	} catch (err) {
		const message =
			err instanceof ApiError ? err.message : "Internal Server Error";
		res.status(err instanceof ApiError ? err.httpCode : 500).send({
			message,
		});
		next(err);
	}
};
