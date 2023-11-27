import { QueryStatistics } from "neo4j-driver";
import { BaseError } from "../../../../libraries/error-handling/app-error.js";
import database from "../../../../models.js";
import { ApiError } from "../../../../libraries/error-handling/api-error.js";

export const checkListBelongToUser = async (id: number, email: string) => {
	const query = `
	Match (p:Playlist) where p.idPlaylist = $id match (u:User) where u.mail=$email with p,u match (u)-[:A_UNE]->(p) return u,p
	`;

	try {
		const raw = await database.run(query, {
			id,
			email,
		});

		if (raw.records == null || raw.records.length == 0) {
			throw new ApiError(
				"Playlist doesn't belong to user",
				"PLAYLIST_DOESNT_BELONG_TO_USER",
				400,
				true
			);
		}
	} catch (err) {
		if (err instanceof ApiError) {
			throw err;
		} else {
			throw new BaseError(
				`Error while fetching playlist from database with email: ${email} and id: ${id}`,
				500,
				"INTERNAL_SERVER_ERROR",
				true
			);
		}
	}
};

export const checkListInDatabase = async (name: string, email: string) => {
	const query = `
 MATCH (:User {mail: $email})-[:A_UNE]->(p:Playlist {name: $name})
    RETURN p
	`;

	try {
		const raw = await database.run(query, {
			name,
			email,
		});

		if (raw.records && raw.records.length > 0) {
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

export const modifyListInDatabase = async (id: number, name: string) => {
	const query = `
		match (p:Playlist{idPlaylist:$id}) SET p.name = $name
	`;

	try {
		const raw = await database.run(query, {
			id,
			name,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: "Playlist has been modified" };
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
			`Error while fetching playlist from database with id : ${id}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const deleteListInDatabase = async (id: number) => {
	//FIXME: Paulo fixing this query tonight at 11/27/2023
	const query = `
		MATCH (p:Playlist{idPlaylist:$id}) DELETE p
	`;

	try {
		const raw = await database.run(query, {
			id,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: "Playlist has been deleted" };
		} else {
			throw new BaseError(
				`Error while deleting playlist in database`,
				500,
				"INTERNAL SERVER ERROR",
				true
			);
		}
	} catch (err) {
		throw new BaseError(
			`Error while deleting playlist from database with id : ${id}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const getListInDatabase = async (email: string) => {
	const query = `
		MATCH (u:User{mail:$email}) MATCH (p:Playlist) MATCH (u)-[:A_UNE]->(p) return u,p
	`;

	try {
		const raw = await database.run(query, {
			email,
		});

		const result = raw.records.map((record) => record.get(1).properties);

		return result;
	} catch (err) {
		throw new BaseError(
			`Error while fetching playlist from database from the user ${email}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};
