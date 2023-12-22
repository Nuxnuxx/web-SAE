import { readable, writable } from "svelte/store";

export const filterStore = writable({
	price: "",
	difficulty: "",
});

export const urlStore = writable("search?");

export const userStore = writable(false);
export const userDataStore = writable({
	idUser: 0,
	firstName: "",
	lastName: "",
	gender: "",
	price: "",
	mail: "",
});
