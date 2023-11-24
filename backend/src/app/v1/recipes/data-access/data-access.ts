import { ApiError } from "../../../../libraries/error-handling/api-error.js";
import {
	BaseError,
	HttpStatusCode,
} from "../../../../libraries/error-handling/app-error.js";
import database from "../../../../models.js";
import neo4j from "neo4j-driver";
import { Filter } from "../recipe.js";

const PRODUCT_PER_PAGE: number = 10;

export const getRecipesInDatabaseWithPagination = async (page_p: number) => {
	const query = "MATCH (n:Recipe) RETURN n SKIP $page LIMIT $limit";
	// float by default if not using .int on params
	const currentPage = neo4j.int(page_p * PRODUCT_PER_PAGE);

	let raw;
	try {
		raw = await database.run(query, {
			page: currentPage,
			limit: neo4j.int(PRODUCT_PER_PAGE),
		});
	} catch (err) {
		throw new BaseError(
			"INTERNAL SERVER ERROR",
			HttpStatusCode.INTERNAL_SERVER,
			`Error while fetching recipes from database`,
			true
		);
	}

	const result = raw.records.map((record) => record.get(0).properties);

	if (result.length == 0) {
		throw new ApiError(
			`No recipes found on page ${currentPage}`,
			"NOT FOUND",
			HttpStatusCode.NOT_FOUND,
			true
		);
	}

	const metadata = await getMetaDataInDatabase(page_p, query);

	result.push({ metadata });

	return result;
};

export const getRecipesByIdInDatabase = async (id: number) => {
	const query = "MATCH (n:Recipe) WHERE n.idRecipe = $idRecipe RETURN n";
	// need a string to match the data in the database
	const idString: string = String(id);

	let raw;
	try {
		raw = await database.run(query, { idRecipe: idString });
	} catch (err) {
		throw new BaseError(
			`Error while fetching recipe from database with id : ${idString}`,
			HttpStatusCode.INTERNAL_SERVER,
			"INTERNAL SERVER ERROR",
			true
		);
	}

	const result = raw.records.map((record) => record.get(0).properties);

	if (result == null || result == undefined || result.length == 0) {
		throw new ApiError(
			`Recipe with id ${idString} not found`,
			"NOT FOUND",
			HttpStatusCode.NOT_FOUND,
			true
		);
	}

	return result;
};

export const getRecipesByKeyWordInDatabase = async (
	keyWordArray: string,
	page_p: number
) => {
	const keyword = keyWordArray.split(" ");

	const query = `
		MATCH (r:Recipe)
		WHERE ANY(keyword IN $keyWords WHERE r.name CONTAINS keyword )
		RETURN r
		SKIP $page
		LIMIT $limit
		UNION
		MATCH (r:Recipe)-[:INGREDIENTS]->(i:Ingredient)
		WHERE ANY(keyword IN $keyWords WHERE i.name CONTAINS keyword )
		RETURN r
		SKIP $page
		LIMIT $limit
	`;

	const currentPage = neo4j.int(page_p * PRODUCT_PER_PAGE);

	let raw;
	try {
		raw = await database.run(query, {
			keyWords: keyword,
			page: currentPage,
			limit: neo4j.int(PRODUCT_PER_PAGE),
		});
	} catch (err) {
		throw new BaseError(
			"INTERNAL SERVER ERROR",
			HttpStatusCode.INTERNAL_SERVER,
			`Error while fetching recipes from database with keyword : ${keyWordArray}`,
			true
		);
	}

	const result = raw.records.map((record) => record.get(0).properties);

	if (result == null || result == undefined || result.length == 0) {
		throw new ApiError(
			`No recipes found with keyword ${keyWordArray}`,
			"NOT FOUND",
			HttpStatusCode.NOT_FOUND,
			true
		);
	}

	const metadata = await getMetaDataInDatabase(page_p, query, {
		keyWords: keyword,
	});

	result.push({ metadata });

	return result;
};

export const getRecipesByFilterInDatabase = async (
	filter: Filter,
	page_p: number
) => {
	let query = "MATCH (r:Recipe) WHERE ";

	const filters = Object.keys(filter).map((key) => {
		if (key == "name") {
			return `r.${key} CONTAINS $${key}`;
		}
		return `r.${key} = $${key}`;
	});

	query += filters.join(" AND ");
	query += " RETURN r SKIP $page LIMIT $limit";

	const currentPage = neo4j.int(page_p * PRODUCT_PER_PAGE);

	let raw;
	try {
		raw = await database.run(query, {
			...filter,
			page: currentPage,
			limit: neo4j.int(PRODUCT_PER_PAGE),
		});
	} catch (err) {
		throw new BaseError(
			"INTERNAL SERVER ERROR",
			HttpStatusCode.INTERNAL_SERVER,
			`Error while fetching recipes from database with filters : ${filter}`,
			true
		);
	}
	const result = raw.records.map((record) => record.get(0).properties);

	if (result == null || result == undefined || result.length == 0) {
		const errorMessage = `No recipes found with filters ${Object.keys(
			filter
		)
			.map((k) => {
				return k + " : " + (filter as Record<string, Filter>)[k];
			})
			.join(", ")}`;
		throw new ApiError(
			errorMessage,
			"NOT FOUND",
			HttpStatusCode.NOT_FOUND,
			true
		);
	}

	const metadata = await getMetaDataInDatabase(page_p, query, {
		...filter,
		page: currentPage,
	});

	result.push({ metadata });

	return result;
};

export const getMetaDataInDatabase = async (
	page_p: number,
	query: string,
	params?: Object
) => {
	let finalQuery = query.replace(
		/RETURN (\w+)/g,
		"RETURN count($1) as total"
	);

	finalQuery = finalQuery.replace(/\s*SKIP \$page\s*LIMIT \$limit\s*/gm, "");

	finalQuery = finalQuery.replace(/(UNION)/g, "\n$1");

	let totalRecords;
	try {
		totalRecords = await database.run(finalQuery, params);
	} catch (err) {
		throw new BaseError(
			"INTERNAL SERVER ERROR",
			HttpStatusCode.INTERNAL_SERVER,
			`Error while fetching total recipes from database`,
			true
		);
	}

	const totalRecipes = totalRecords.records[0].get("total").toNumber();

	const totalPages = Math.ceil(totalRecipes / PRODUCT_PER_PAGE);

	const metadata = {
		currentPage: page_p,
		totalPages: totalPages,
		totalRecipes: totalRecipes,
	};

	return metadata;
};
