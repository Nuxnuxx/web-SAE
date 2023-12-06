export interface User {
	idUser?: string;
	firstName?: string;
	lastName?: string;
	mail: string;
	password: string;
	confirmPassword?: string;
	gender?: string;
}

export type ErrorsRegister = Record<string, string>;
