import database from "../../../models.js";
import neo4j from "neo4j-driver";

const PRODUCT_PER_PAGE: number = 10;

export const getRecipesInDatabaseWithPagination = async (page_p: number) => {
	const query = "MATCH (n:Recipe) RETURN n SKIP $page LIMIT $limit";
	// float by default if not using .int on params
	const currentPage = neo4j.int(page_p * PRODUCT_PER_PAGE);

	const raw = await database.run(query, {
		page: currentPage,
		limit: neo4j.int(PRODUCT_PER_PAGE),
	});

	const result = raw.records.map((record) => record.get(0).properties);

	return result;
};

export const getRecipesByIdInDatabase = async (id: number) => {
	const query = "MATCH (n) WHERE n.idRecipe = $idRecipe RETURN n";
	// need a string to match the data in the database
	const idString: string = String(id);
	const raw = await database.run(query, { idRecipe: idString });
	const result = raw.records.map((record) => record.get(0).properties);

	return result;
};

export const getRecipesByKeyWordInDatabase = async (keyWordArray: string) => {
	const query = `
        MATCH (n:Recipe)
        WHERE n.name CONTAINS $keyWord
        RETURN n
				LIMIT 10
    `;
	const raw = await database.run(query, { keyWord: keyWordArray });
	const result = raw.records.map((record) => record.get(0).properties);
	return result;
};

export const getRecipesByFilterInDatabase = async (filter: string[]) => {
	let query = "MATCH (r:Recipe) WHERE ";

	const filters = Object.keys(filter).map((key) => {
		if (key == "name") {
			return `r.${key} CONTAINS $${key}`;
		}
		return `r.${key} = $${key}`;
	});

	query += filters.join(" AND ");
	query += " RETURN r";
	const raw = await database.run(query, filter);
	const result = raw.records.map((record) => record.get(0).properties);
	return result;
};
