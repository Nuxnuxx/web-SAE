import { QueryStatistics } from "neo4j-driver";
import { BaseError } from "../../../../libraries/error-handling/app-error.js";
import database from "../../../../models.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";

export const checkListInDatabase = async (name: string, email: string) => {
	const query = `MATCH(p: Playlist { name: $name })<-[:A_UNE]-(u: User)
		 RETURN u, p`;

	try {
		const raw = await database.run(query, {
			name,
		});

		if (raw.records) {
			throw new ApiError(
				"Playlist already exists",
				"PLAYLIST_ALREADY_EXISTS",
				400
			);
		}
	} catch (err) {
		if (err instanceof ApiError) {
			throw err;
		} else {
			console.log(err);
			throw new BaseError(
				`Error while fetching playlist from database with email: ${email} and name: ${name}`,
				500,
				"INTERNAL_SERVER_ERROR",
				true
			);
		}
	}
};

export const createListInDatabase = async (name: string, email: string) => {
	const query = `
		MERGE(id: UniqueId { name: 'Playlist' })
		ON CREATE SET id.count = 100
		ON MATCH SET id.count = id.count + 1

		WITH id.count AS uid
			CREATE(p: Playlist { idPlaylist: uid, name: $name })


		WITH p, $email AS userMailValue
			MATCH(u: User { mail: userMailValue})

		// Créer un lien entre le nœud Playlist et le nœud User
		CREATE(u) - [: A_UNE] -> (p)
	`;

	try {
		const raw = await database.run(query, {
			name,
			email,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: "Playlist has been created" };
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
			`Error while fetching playlist from database with email : ${email} and name : ${name}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const modifyListInDatabase = async (
	id: number,
	name: string,
	email: string
) => {
	//TODO: do the query paulo
	const query = "";
	return "modify";
};

export const deleteListInDatabase = async (id: number, email: string) => {
	//TODO: do the query paulo
	const query = "";
	return "delete";
};

export const getListInDatabase = async (id: number, email: string) => {
	//TODO: do the query paulo
	const query = "";
	return "get playlist";
};
