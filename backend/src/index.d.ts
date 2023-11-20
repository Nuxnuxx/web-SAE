export interface Recipe {
	id: number;
	name: string;
	urlImage: string;
	numberLike: bigInt;
	userLike?: boolean;
	userList?: boolean;
}
