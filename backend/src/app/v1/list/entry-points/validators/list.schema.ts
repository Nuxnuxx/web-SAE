import * as yup from "yup";

export const schemaCreateListQuery = yup.object().shape({
	namelist: yup.string().max(20).min(1).required(),
});

export const schemaModifyListQuery = yup.object().shape({
	idlist: yup.number().required(),
	namelist: yup.string().max(20).min(1).required(),
});

export const schemaDeleteListQuery = yup.object().shape({
	idlist: yup.number().required(),
});
