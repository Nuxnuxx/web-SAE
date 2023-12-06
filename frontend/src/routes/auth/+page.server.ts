import { sendLogin, sendRegister } from "$lib/api/auth-request";
import { schemaLogin, schemaRegister } from "$lib/api/auth-schema";
import type { ErrorsRegister, User } from "$lib/api/auth-types";
import { fail, type Actions, redirect } from "@sveltejs/kit";
import * as yup from "yup";
import type { PageServerLoad } from "./$types";

//TODO: TOM fait ca
// separer les resultat des deux form par un id different pour les afficher separement
// ainsi que renvoyer le formulaire qui etait actif pour changer la valeur du selectedButton
// les errorrs stylise parceque la c'est moche

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	if (token) {
		throw redirect(302, "/profil");
	}
};

export const actions: Actions = {
	login: async ({ cookies, request }) => {
		const body = await request.formData();
		const user: User = {
			mail: body.get("mail")?.toString() || "",
			password: body.get("password")?.toString() || "",
		};

		try {
			await schemaLogin.validate(user, {
				abortEarly: false,
			});
			const result = await sendLogin(user);
			if (result) {
				cookies.delete("token");
				cookies.set("token", result.result);
				return {
					location: "/",
				};
			}
		} catch (err) {
			let errors: ErrorsRegister = {};
			if (err instanceof yup.ValidationError) {
				errors = err.inner.reduce((acc, err) => {
					const key = String(err.path);
					return { ...acc, [key]: err.message };
				}, {});
			} else {
				if (err instanceof Error) {
					errors = { server: err.message };
				}
			}
			return fail(400, errors);
		}
	},

	register: async ({ cookies, request }) => {
		const body = await request.formData();
		const user: User = {
			mail: body.get("mail")?.toString() || "",
			firstName: body.get("firstName")?.toString() || "",
			lastName: body.get("lastName")?.toString() || "",
			password: body.get("password")?.toString() || "",
			confirmPassword: body.get("confirmPassword")?.toString() || "",
			gender: body.get("gender")?.toString() || "",
		};

		const finalUser: User = {
			mail: user.mail,
			firstName: user.firstName,
			lastName: user.lastName,
			password: user.password,
		};

		try {
			await schemaRegister.validate(user, {
				abortEarly: false,
			});
			const result = await sendRegister(finalUser);
			if (result) {
				cookies.delete("token");
				cookies.set("token", result.result);
				return {
					location: "/",
				};
			}
		} catch (err) {
			let errors: ErrorsRegister = {};
			if (err instanceof yup.ValidationError) {
				errors = err.inner.reduce((acc, err) => {
					const key = String(err.path);
					return { ...acc, [key]: err.message };
				}, {});
			} else {
				if (err instanceof Error) {
					errors = { server: err.message };
				}
			}
			return fail(400, errors);
		}
	},
};
