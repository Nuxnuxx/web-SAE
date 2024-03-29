import type { User } from "./auth-types";

export const getRecipes = async (
	name: string,
	price: string,
	difficulty: string,
	page: number
) => {
	try {
		let url;
		//INFO: if no name is given, we don't want to add the name query param
		if (name) {
			url = `${
				import.meta.env.VITE_API_URL
			}/recipe/page/${page}?name=${name}`;
		} else {
			url = `${import.meta.env.VITE_API_URL}/recipe/page/${page}`;
		}

		// Add price and difficulty to the URL if they exist
		if (price && name) {
			url += `&price=${price}`;
		} else if (price) {
			url += `?price=${price}`;
		}

		if (difficulty && price) {
			url += `&difficulty=${difficulty}`;
		} else if (difficulty) {
			url += `?difficulty=${difficulty}`;
		}

		const result = await fetch(url, {
			method: "GET",
		});

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};

export const sendModifyProfil = async (
	token: string,
	firstName: string,
	lastName: string,
	newPassword: string
) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/auth/profil`,
			{
				method: "PUT",
				body: JSON.stringify({ firstName, lastName, newPassword }),
				headers: {
					"Content-type": "application/json",
					Authorization: "Bearer " + token,
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};

export const getProfil = async (data: string) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/auth/profil`,
			{
				method: "GET",
				headers: {
					"Content-type": "application/json",
					Authorization: "Bearer " + data,
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};

export const sendLogin = async (data: User) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/auth/login`,
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
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};

export const sendRegister = async (data: User) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/auth/register`,
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
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};

export const sendColdstart = async (
	token: string,
	price: string,
	difficulty: string
) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/auth/coldstart`,
			{
				method: "POST",
				body: JSON.stringify({ price, difficulty }),
				headers: {
					"Content-type": "application/json",
					Authorization: "Bearer " + token,
				},
			}
		);

		if (result.ok) {
			const data = await result.json();
			return data;
		} else {
			const { error } = await result.json();
			throw new Error(error);
		}
	} catch (err) {
		throw err;
	}
};
