import { genSaltSync, hashSync } from "bcrypt";

export const hashPassword = (password: string) => {
	const salt = genSaltSync(10);
	const hashedPass = hashSync(password, salt);
	return hashedPass;
};
