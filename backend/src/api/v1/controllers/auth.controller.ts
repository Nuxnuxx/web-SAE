import { Request, Response } from "express";
import { compareSync, genSaltSync, hashSync } from "bcrypt";
import { findUser } from "../../../services/auth.js";

//TODO: Try catch tout ca

export const registerController = async (req: Request, res: Response) => {
	const salt = genSaltSync(10);
	const hashedPass = hashSync(req.body.password, salt);

	const user = {
		name: req.body.name,
		email: req.body.email,
		password: hashedPass,
	};

	//TODO: findUser method to find user

	const userFound = findUser(user.email);
	if (userFound) {
		res.send({ ok: false, message: "User already exists" });
		return;
	} else {
		//TODO: send to database
		res.send({ ok: true, message: "User created" });
	}
};

export const loginController = async (req: Request, res: Response) => {
	const userFound = findUser(req.body.email);

	if (userFound) {
		if (compareSync(req.body.password, userFound.password)) {
			res.send({
				ok: true,
				name: userFound.name,
				email: userFound.email,
			});
		} else {
			res.send({ ok: false, message: "Credential are wrong" });
		}
	} else {
		res.send({ ok: false, message: "Credential are wrong" });
	}
};

export const modifyProfilController = async (req: Request, res: Response) => {
	const userFound = findUser(req.body.email);

	if (userFound) {
		if (compareSync(req.body.password, userFound.password)) {
			const salt = genSaltSync(10);
			const hashedPass = hashSync(req.body.newPassword, salt);
			//TODO: write to database (create service function)
		}
	}
};

export const deleteProfilController = async (req: Request, res: Response) => {
	const userFound = findUser(req.body.email);

	if (userFound) {
		if (compareSync(req.body.password, userFound.password)) {
			//TODO: write to database (create service function)
		}
	}
};
