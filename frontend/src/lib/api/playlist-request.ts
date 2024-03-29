export const getPlaylist = async (token: string) => {
	try {
		const result = await fetch(`${import.meta.env.VITE_API_URL}/list`, {
			method: "GET",
			headers: {
				"Content-type": "application/json",
				Authorization: "Bearer " + token,
			},
		});

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

export const getPlaylistRecipe = async (token: string, id: number) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/list/recipe?id=${id}`,
			{
				method: "GET",
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

export const createPlaylist = async (token: string, name: string) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/list?name=${name}`,
			{
				method: "POST",
				headers: {
					"Content-type": "application/json",
					Authorization: "Bearer " + token,
				},
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
