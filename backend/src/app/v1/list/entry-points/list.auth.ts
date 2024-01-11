import { NextFunction, Request, Response } from "express";
import { findUserByEmailInDatabase } from "../../auth/data-access/data-access.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";
import { compareSync } from "bcrypt";

export const authenticate = async (
	req: Request,
	res: Response,
	next: NextFunction
) => {
	const b64auth = (req.headers.authorization || "").split(" ")[1] || "";
	const [email, password] = Buffer.from(b64auth, "base64")
		.toString()
		.split(":");

	const userFound = await findUserByEmailInDatabase(email);

	if (userFound) {
		if (compareSync(password, userFound.password)) {
			//@ts-ignore
			req.user = userFound;
			next();
		} else {
			res.status(401).send({ message: "Invalid email or password" });
			throw new ApiError(
				"Unauthorized",
				"Invalid email or password",
				401
			);
		}
	} else {
		res.status(401).send({ message: "Invalid email or password" });
		throw new ApiError("Unauthorized", "Invalid email or password", 401);
	}
};
