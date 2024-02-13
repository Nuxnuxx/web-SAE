export const getSimilarRecipes = async (id: number, number: number) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/similarRecipes/${id}/${number}`,
			{
				method: "GET",
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

export const recipeTimeSpent = async (
	token: string,
	recipeId: number,
	timeSpent: number
) => {
	try {
		const result = await fetch(
			`${import.meta.env.VITE_API_URL}/recipeTimeSpent`,
			{
				method: "POST",
				headers: {
					"Content-type": "application/json",
					Authorization: `Bearer ${token}`,
				},
				body: JSON.stringify({ recipeId, timeSpent }),
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
