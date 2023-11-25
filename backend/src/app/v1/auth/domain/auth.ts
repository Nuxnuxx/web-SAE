import { compareSync } from "bcrypt";
import { ImodifyUser, User } from "../auth.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";
import {
	createUserInDatabase,
	deleteUserInDatabase,
	findUserByEmailInDatabase,
	modifyUserInDatabase,
} from "../data-access/data-access.js";
import { hashPassword } from "./util.js";

export const registerUser = async (user: User) => {
	const hashedPass = hashPassword(user.password);

	const userFound = await findUserByEmailInDatabase(user.email);

	if (userFound) {
		throw new ApiError(
			`User with email ${user.email} already exists`,
			"CONFLICT",
			409,
			true
		);
	} else {
		user.password = hashedPass;
		await createUserInDatabase(user);
		const message = `User with email ${user.email} has been created`;
		return { message };
	}
};

export const loginUser = async (user: User) => {
	const userFound = await findUserByEmailInDatabase(user.email);

	if (userFound) {
		if (compareSync(user.password, userFound.password)) {
			return userFound;
		} else {
			throw new ApiError(
				`User with email ${user.email} has wrong credentials`,
				"UNAUTHORIZED",
				401,
				true
			);
		}
	} else {
		throw new ApiError(
			`User with email ${user.email} does not exist`,
			"NOT FOUND",
			404,
			true
		);
	}
};

export const modifyUser = async (user: ImodifyUser) => {
	const userFound = await findUserByEmailInDatabase(user.email);

	if (userFound) {
		if (compareSync(user.password, userFound.password)) {
			const hashedPass = hashPassword(user.newPassword);

			user.newPassword = hashedPass;
			await modifyUserInDatabase(user);
			const message = `User with email ${user.email} has been modified`;
			return { message };
		} else {
			throw new ApiError(
				`User with email ${user.email} has wrong credentials`,
				"UNAUTHORIZED",
				401,
				true
			);
		}
	} else {
		throw new ApiError(
			`User with email ${user.email} does not exist`,
			"NOT FOUND",
			404,
			true
		);
	}
};

export const deleteUser = async (user: User) => {
	const userFound = await findUserByEmailInDatabase(user.email);

	if (userFound) {
		if (compareSync(user.password, userFound.password)) {
			await deleteUserInDatabase(user.email);
			const message = `User with email ${user.email} has been deleted`;
			return { message };
		} else {
			throw new ApiError(
				`User with email ${user.email} has wrong credentials`,
				"UNAUTHORIZED",
				401,
				true
			);
		}
	} else {
		throw new ApiError(
			`User with email ${user.email} does not exist`,
			"NOT FOUND",
			404,
			true
		);
	}
};
