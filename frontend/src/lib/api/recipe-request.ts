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
