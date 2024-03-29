import { sendLogin, sendRegister, sendColdstart } from "$lib/api/auth-request";
import {
	schemaLogin,
	schemaRegister,
	schemaColdstart,
} from "$lib/api/auth-schema";
import type { ErrorWithId, User, Coldstart } from "$lib/api/auth-types";
import { fail, type Actions, redirect } from "@sveltejs/kit";
import * as yup from "yup";
import type { PageServerLoad } from "./$types";
import { createPlaylist } from "$lib/api/playlist-request";

export const load: PageServerLoad = async ({ cookies }) => {
	// vérifie si le coldstart a déjà été fait
	const coldstart = cookies.get("coldstart");
	const token = cookies.get("token");

	// si pas de coldstart et que le token est présent, on redirige vers coldstart
	if (!coldstart && token) {
		throw redirect(302, "/auth/coldstart");
	}

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
				cookies.set("token", result.result, {
					path: "/",
				});
				cookies.delete("coldstart");
				cookies.set("coldstart", "true", {
					path: "/",
				});
				return {
					location: "/",
				};
			}
		} catch (err) {
			let errors: ErrorWithId = {
				id: "login",
				errors: {},
				values: user,
			};
			if (err instanceof yup.ValidationError) {
				errors.errors = err.inner.reduce((acc, err) => {
					const key = String(err.path);
					return { ...acc, [key]: err.message };
				}, {});
			} else {
				if (err instanceof Error) {
					errors.errors = { server: err.message };
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
			gender: user.gender,
		};

		try {
			await schemaRegister.validate(user, {
				abortEarly: false,
			});
			const result = await sendRegister(finalUser);
			if (result) {
				cookies.delete("token");
				cookies.set("token", result.result, {
					path: "/",
				});
				cookies.delete("coldstart");
				const token = cookies.get("token");

				if (token === undefined || token === null) {
					throw new Error("token is not defined");
				}
				const playlist = await createPlaylist(token, "liked");
				return {
					location: "/",
				};
			}
		} catch (err) {
			let errors: ErrorWithId = {
				id: "register",
				errors: {},
				values: user,
			};
			if (err instanceof yup.ValidationError) {
				errors.errors = err.inner.reduce((acc, err) => {
					const key = String(err.path);
					return { ...acc, [key]: err.message };
				}, {});
			} else {
				if (err instanceof Error) {
					errors.errors = { server: err.message };
				}
			}
			return fail(400, errors);
		}
	},
};
