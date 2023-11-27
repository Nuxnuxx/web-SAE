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

export const createRecipeListInDatabase = async (
	idRecipe: number,
	idList: number,
	email: string
) => {
	const query = `
		MATCH (p:Playlist{idPlaylist:$idList})  
		MATCH (u:User) where u.mail=$email 
		MATCH (u)-[:A_UNE]->(p)   
		MATCH (r:Recipe{idRecipe:$idRecipe}) with p,u,r 
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
			`Error while fetching playlist from database with email : ${email}`,
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
		MATCH (p:Playlist{idPlaylist:$idList})-[r:EST_DANS]->(r:Recipe{idRecipe:$idRecipe})
		DELETE r
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
		if (result.length > 0) {
			return result;
		} else {
			throw new BaseError(
				`Error while getting recipe from list in database`,
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
