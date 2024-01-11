import * as yup from "yup";

enum Gender {
	MALE = "male",
	FEMALE = "female",
	OTHER = "other",
}

export const schemaRegister = yup.object().shape({
	firstName: yup
		.string()
		.required("firstname is required")
		.max(20, "firstname is too long")
		.min(3, "firstname is too short"),
	lastName: yup
		.string()
		.required("lastname is required")
		.max(20, "lastname is too long")
		.min(3, "lastname is too short"),
	mail: yup
		.string()
		.required("Mail is required")
		.email("It is not a valid email"),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum."),
	confirmPassword: yup
		.string()
		.required("Please confirm your password")
		.oneOf([yup.ref("password")], "Passwords do not match"),
	gender: yup
		.string()
		.required("Gender is required")
		.oneOf(Object.values(Gender), "Please choose a gender"),
});

export const schemaLogin = yup.object().shape({
	mail: yup
		.string()
		.required("Mail is required")
		.email("It is not a valid email"),
	password: yup
		.string()
		.required("No password provided.")
		.min(8, "Password is too short - should be 8 chars minimum."),
});

enum Price {
	// 3 price "bon marché", "moyen", "assez cher"
	BON_MARCHE = "0",
	MOYEN = "1",
	ASSEZ_CHER = "2",
}

enum Difficulty {
	// 4 difficulty "très facile", "facile", "moyen", "difficile"
	TRES_FACILE = "0",
	FACILE = "1",
	MOYEN = "2",
	DIFFICILE = "3",
}

export const schemaColdstart = yup.object().shape({
	price: yup
		.string()
		.required("Price is required")
		.oneOf(Object.values(Price), "Please choose a price"),
	difficulty: yup
		.string()
		.required("Difficulty is required")
		.oneOf(Object.values(Difficulty), "Please choose a difficulty"),
});
