export interface User {
	name?: string;
	email: string;
	password: string;
}

export interface ErrorsRegister {
	firstName?: string;
	lastName?: string;
	email?: string;
	password?: string;
	confirmPassword?: string;
	gender?: string;
	server?: string;
}
