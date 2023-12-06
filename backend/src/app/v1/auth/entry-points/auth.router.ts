import { Router } from "express";
import {
	handleDeleteProfil,
	handleLogin,
	handleModifyProfil,
	handleRegister,
} from "./auth.controller.js";
import { validate } from "./validators/auth.validator.js";
import {
	schemaDeleteBody,
	schemaLoginBody,
	schemaModifyBody,
	schemaRegisterBody,
} from "./validators/auth.schema.js";

const authRouter = Router();

authRouter.post("/register", validate(schemaRegisterBody), handleRegister);
authRouter.post("/login", validate(schemaLoginBody), handleLogin);
authRouter.put("/profil", validate(schemaModifyBody), handleModifyProfil);
authRouter.delete("/profil", validate(schemaDeleteBody), handleDeleteProfil);

export default authRouter;
