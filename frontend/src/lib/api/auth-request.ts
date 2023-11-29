import type { User } from "./auth-types";

export const sendLogin = async (data: User) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/api/v1/auth/login`,
			{
				method: "POST",
				body: JSON.stringify({ ...data }),
				headers: {
					"Content-type": "application/json",
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { message } = await result.json();
			throw new Error(message);
		}
	} catch (err) {
		throw err;
	}
};

export const sendRegister = async (data: User) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/api/v1/auth/register`,
			{
				method: "POST",
				body: JSON.stringify({ ...data }),
				headers: { "Content-type": "application/json" },
			}
		);

		if (result.status == 200) {
			const data = await result.json();
			return data;
		} else {
			const { message } = await result.json();
			throw new Error(message);
		}
	} catch (err) {
		throw err;
	}
};
