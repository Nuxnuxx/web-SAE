import { QueryStatistics } from "neo4j-driver";
import database from "../../../../models.js";
import { BaseError } from "../../../../libraries/error-handling/app-error.js";

export const createRecipeInLikedInDatabase = async (
	id: number,
	email: string
) => {
	const query = `
		MATCH (p:Playlist{name:'liked'})
		MATCH (u:User) WHERE u.mail = $email
		MATCH (u)-[:A_UNE]->(p)
		MATCH (r:Recipe{idRecipe: $id})
		CREATE (r)-[:est_dans]->(p)
	`;

	try {
		const raw = await database.run(query, {
			id,
			email,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: "Recipe has been created in liked" };
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
			`Error while fetching recipe from database with id : ${id}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const createRecipeInListInDatabase = async (
	idRecipe: number,
	idList: number,
	email: string
) => {
	const query = `
		MATCH (p:Playlist{idPlaylist:$idList})
		MATCH (u:User) WHERE u.mail = $email
		MATCH (u)-[:A_UNE]->(p)
		MATCH (r:Recipe{idRecipe: $idRecipe})
		CREATE (r)-[:est_dans]->(p)

	`;

	try {
		const raw = await database.run(query, {
			idList,
			idRecipe,
			email,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;
		if (stats.containsUpdates()) {
			return { message: `Recipe has been created in list ${idList}` };
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
			`Error while fetching recipe from database with idList : ${idList}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const deleteRecipeListInDatabase = async (
	idList: number,
	idRecipe: number
) => {
	const query = `
		MATCH (r:Recipe{idRecipe:$idRecipe})-[l:est_dans]->(p:Playlist{idPlaylist:$idList}) delete l
	`;

	try {
		const raw = await database.run(query, {
			idList,
			idRecipe,
		});

		const stats: QueryStatistics = raw.summary.updateStatistics;

		if (stats.containsUpdates()) {
			return { message: "Recipe has been deleted from playlist" };
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
			`Error while fetching playlist from database with id : ${idList}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};

export const getRecipesListInDatabase = async (id: number) => {
	const query = `
  MATCH (r:Recipe)-[:est_dans]->(p:Playlist{idPlaylist:$id}) return r
	`;

	try {
		const raw = await database.run(query, {
			id,
		});

		const result = raw.records.map((record) => {
			return record.get("r").properties;
		});
		return result;
	} catch (err) {
		throw new BaseError(
			`Error while fetching playlist from database with id : ${id}`,
			500,
			"INTERNAL SERVER ERROR",
			true
		);
	}
};
