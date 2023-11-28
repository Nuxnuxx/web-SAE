import * as yup from "yup";

export const schemaCreateRecipeLikedQuery = yup.object().shape({
	idrecipe: yup.number().required(),
});

export const schemaCreateRecipeListQuery = yup.object().shape({
	idrecipe: yup.number().required(),
	idlist: yup.number().required(),
});

export const schemaGetRecipeListQuery = yup.object().shape({
	id: yup.number().required(),
});

export const schemaDeleteRecipeListQuery = yup.object().shape({
	idlist: yup.number().required(),
	idrecipe: yup.number().required(),
});
