import { writable } from "svelte/store";

export const filterStore = writable({
  price: '',
  difficulty: ''
});
