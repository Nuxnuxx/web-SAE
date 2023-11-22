export interface Recipe {
	id: number;
	name: string;
	urlImage: string;
	numberLike: bigInt;
	userLike?: boolean;
	userList?: boolean;
}

export interface Filter {
	difficulty?: string;
	quantity?: string;
	price?: string;
	name?: string;
	idRecipe?: number;
}
