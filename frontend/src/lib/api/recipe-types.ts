export type RecipeDetail = {
	difficulty: string;
	//FIXME: should be array of string
	images: string;
	quantity: string;
	price: string;
	time: string;
	name: string;
	idRecipe: number;
};

export type RecipeStep = {
	[key: string]: {
		Step: string;
	};
};

export type RecipeIngredients = {
	[key: string]: {
		name: string;
		urlPicture: string;
	};
};

export type RecipeData = {
	recipeDetail: RecipeDetail;
	recipeStep: RecipeStep;
	recipeIngredients: RecipeIngredients;
};

export const Price = ["bon marché", "moyen", "assez cher"];

export const Difficulty = ["très facile", "facile", "moyen", "difficile"];

export type Pagination = {
	currentPage: number;
	totalPage: number;
	totalResult: number;
};

export type Current = {
	name: string;
	page: number;
	price: string;
	difficulty: string;
};
