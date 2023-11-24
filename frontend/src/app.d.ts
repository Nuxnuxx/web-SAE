// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}
}

export interface Recipe {
	img: string;
	title: string;
	nbLikes: number;
	liked: boolean;
	saved: boolean;
}

export {};
