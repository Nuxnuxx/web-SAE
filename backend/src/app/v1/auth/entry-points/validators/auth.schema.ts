import * as yup from "yup";

export const schemaRegisterBody = yup.object().shape({
	name: yup.string().max(10).min(3).required(),
	email: yup.string().email().required(),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum.")
		.matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
});

export const schemaLoginBody = yup.object().shape({
	email: yup.string().email().required(),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum.")
		.matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
});

export const schemaModifyBody = yup.object().shape({
	name: yup.string().max(10).min(3).required(),
	email: yup.string().email().required(),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum.")
		.matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
	newPassword: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum.")
		.matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
});

export const schemaDeleteBody = yup.object().shape({
	name: yup.string().max(10).min(3).required(),
	email: yup.string().email().required(),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum.")
		.matches(/[a-zA-Z]/, "Password can only contain Latin letters."),
});
