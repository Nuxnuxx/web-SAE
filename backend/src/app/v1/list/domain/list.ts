import {
	checkListBelongToUser,
	checkListInDatabase,
	createListInDatabase,
	deleteListInDatabase,
	getListInDatabase,
	modifyListInDatabase,
} from "../data-access/list.js";

export const createList = async (name: string, email: string) => {
	await checkListInDatabase(name, email);
	const result = await createListInDatabase(name, email);
	return result;
};

export const modifyList = async (id: number, name: string, email: string) => {
	await checkListBelongToUser(id, email);
	const result = await modifyListInDatabase(id, name);
	return result;
};

export const deleteList = async (id: number, email: string) => {
	await checkListBelongToUser(id, email);
	const result = await deleteListInDatabase(id);
	return result;
};

export const getList = async (email: string) => {
	const result = await getListInDatabase(email);
	return result;
};
