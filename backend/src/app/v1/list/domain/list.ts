import {
    checkListInDatabase,
	createListInDatabase,
	deleteListInDatabase,
	getListInDatabase,
	modifyListInDatabase,
} from "../data-access/list.js";

export const createList = async (name: string, email: string) => {
	await checkListInDatabase(name, email)
	const result = await createListInDatabase(name, email);
	return result;
};

export const modifyList = async (id: number, name: string, email: string) => {
	const result = await modifyListInDatabase(id, name, email);
	return result;
};

export const deleteList = async (id: number, email: string) => {
	const result = await deleteListInDatabase(id, email);
	return result;
};

export const getList = async (id: number, email: string) => {
	const result = await getListInDatabase(id, email);
	return result;
};
