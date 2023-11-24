export interface Recipe {
	idRecipe: number;
	name: string;
	price: string;
	quantity: string;
	difficulty: string;
}

export interface Filter {
	difficulty?: string;
	quantity?: string;
	price?: string;
	name?: string;
	idRecipe?: number;
}
