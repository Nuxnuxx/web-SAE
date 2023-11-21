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
	const query = "MATCH (n) WHERE id(n) = $nodeId RETURN n";
	const result = await database.run(query, { nodeId: id });
	return result;
};

export const getRecipesByKeyWordInDatabase = async (keyWord: string[]) => {
	const query = "";
	const result = await database.run(query, { keyWord: keyWord });
	return result;
};

export const getRecipesByFilterInDatabase = async (filter: string[]) => {
	const query = "";
	const result = await database.run(query, { filter: filter });
	return result;
};
