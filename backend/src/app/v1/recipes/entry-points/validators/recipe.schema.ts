import * as yup from "yup";
import { Recipe } from "../../recipe";
import { Difficulty, Price } from "./recipe.enum.js";

export const schemaByIdParams = yup.object().shape({
	id: yup.number().positive().integer().required(),
});

export const schemaPageParams = yup.object().shape({
	page: yup.number().integer().moreThan(-1).required(),
});

export const schemaKeywordParams = yup.object().shape({
	page: yup.number().integer().moreThan(-1).required(),
	keyword: yup.string().required().required(),
});

export const schemaFilterParams = yup.object().shape({
	page: yup.number().integer().moreThan(-1).required(),
});

const recipeKeys: (keyof Recipe)[] = [
	"idRecipe",
	"name",
	"price",
	"quantity",
	"difficulty",
];

export const schemaFilterQuery = yup.object().shape({
	...recipeKeys.reduce(
		(acc, key) => ({
			...acc,
			[key]:
				key === "price"
					? yup.string().oneOf(Object.values(Price)).notRequired()
					: key === "difficulty"
					  ? yup
								.string()
								.oneOf(Object.values(Difficulty))
								.notRequired()
					  : yup.string().notRequired(),
		}),
		{}
	),
});
