import { BaseError } from "../../../../libraries/error-handling/app-error.js";
import { QueryStatistics } from "neo4j-driver";
import database from "../../../../models.js";
import { ImodifyUser, User } from "../auth.js";

export const findUserByEmailInDatabase = async (
	email: string
): Promise<User | null> => {
	const query = "MATCH (n:User) WHERE n.mail = $email return n";

	try {
		const raw = await database.run(query, { email });

		const result = raw.records.map((record) => record.get(0).properties);

		if (result == null || result == undefined || result.length == 0) {
			return null;
		}

		const user: User = {
			name: result[0].name,
			email: result[0].mail,
			password: result[0].password,
		};

		return user;
	} catch (err) {
		throw new BaseError(
			`Error while fetching user from database with email : ${email}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const createUserInDatabase = async (user: User): Promise<User> => {
	const query =
		"CREATE (n:User {name: $name, mail: $mail, password: $password})";

	try {
		const created = await database.run(query, {
			name: user.name,
			mail: user.email,
			password: user.password,
		});

		const stats: QueryStatistics = created.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return user;
		} else {
			throw new BaseError(
				`Error while creating user in database`,
				500,
				"INTERNAL SERVER ERROR",
				true
			);
		}
	} catch (err) {
		throw new BaseError(
			`Error while creating user in database`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const modifyUserInDatabase = async (user: ImodifyUser) => {
	const query =
		"MATCH (n:User) WHERE n.mail = $mail SET n.password = $password";

	try {
		const modified = await database.run(query, {
			mail: user.email,
			password: user.newPassword,
		});

		const stats: QueryStatistics = modified.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: "User has been modified" };
		} else {
			throw new BaseError(
				`Error while modifying user in database`,
				500,
				"INTERNAL SERVER ERROR",
				true
			);
		}
	} catch (err) {
		throw new BaseError(
			`Error while modifying user in database`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const deleteUserInDatabase = async (email: string) => {
	const query = "MATCH (n:User) WHERE n.mail = $mail DETACH DELETE n";

	try {
		const erased = await database.run(query, { mail: email });

		const stats: QueryStatistics = erased.summary.updateStatistics;

		if (stats.containsUpdates()) {
			return { message: "User has been deleted" };
		} else {
			throw new BaseError(
				`Error while deleting user in database`,
				500,
				"INTERNAL SERVER ERROR",
				true
			);
		}
	} catch (err) {
		throw new BaseError(
			`Error while deleting user in database`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};
