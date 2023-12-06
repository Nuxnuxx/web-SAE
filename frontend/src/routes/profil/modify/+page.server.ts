import { fail, type Actions } from "@sveltejs/kit";
import type { PageParentData, PageServerLoad } from "../$types";
import { sendModifyProfil } from "$lib/api/auth-request";
import type { ErrorsRegister } from "$lib/api/auth-types";

export const load: PageServerLoad = async ({ parent }: PageParentData) => {
	const { profil } = await parent();

	return {
		profil,
	};
};

export const actions: Actions = {
	modifyPassword: async ({ cookies, request }) => {
		const body = await request.formData();

		const newPassword: string = body.get("newPassword")?.toString() || "";
		const token: string = cookies.get("token") || "";

		try {
			if (newPassword.length < 8) {
				throw new Error(
					"le mot de passe doit contenir 8 characteres minimum"
				);
			}
			const result = await sendModifyProfil(token, newPassword);
			if (result) {
				cookies.set("token", result.result);
				return {
					location: "/profil",
				};
			}
		} catch (err) {
			let errors: ErrorsRegister = {};
			if (err instanceof Error) {
				errors = { server: err.message };
			}
			return fail(400, errors);
		}
	},
};
