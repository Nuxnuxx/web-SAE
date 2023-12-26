import { fail, type Actions, redirect, Redirect_1 } from "@sveltejs/kit";
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

		const firstName: string = body.get("firstName")?.toString() || "";
		const lastName: string = body.get("lastName")?.toString() || "";
		const newPassword: string = body.get("newPassword")?.toString() || "";
		const token: string = cookies.get("token") || "";

		try {
			if (newPassword.length < 8) {
				throw new Error(
					"Password is too short - should be 8 chars minimum."
				);
			}
			const result = await sendModifyProfil(token, firstName, lastName, newPassword);
			if (result) {
				cookies.set("token", result.result, {
					path: "/",
				});
				throw redirect(303, "/profil");
			}
		} catch (err) {
			let errors: ErrorsRegister = {};
			if (err instanceof Error) {
				errors = { server: err.message };
			}
			if (err?.location) {
				throw err;
			}
			return fail(400, errors);
		}
	},
};
