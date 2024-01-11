export interface User {
	email: string;
	password: string;
}

declare global {
	namespace Express {
		interface Request {
			user?: User;
		}
	}
}
