import * as yup from "yup";
import { NextFunction, Request, Response } from "express";

export const validate = <T>(
	schema: yup.Schema<T>,
	reqType: "params" | "query"
): any => {
	return async (req: Request, res: Response, next: NextFunction) => {
		try {
			await schema.validate(req[reqType]);
			next();
		} catch (err) {
			if (err instanceof yup.ValidationError) {
				res.status(400).json({
					message: err.errors.toString(),
				});
			}
			next(err);
		}
	};
};
