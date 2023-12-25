export const getRecipe = async (id: number) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/recipe/${id}`,
			{
				method: "GET",
				headers: {
					"Content-type": "application/json",
				},
			}
		);

		console.log("id recipe", id)
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

export const likeRecipe = async (token: string, id: number) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/list/recipe?id=${id}`,
			{
				method: "POST",
				headers: {
					"Content-type": "application/json",
					Authorization: `Bearer ${token}`,
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
