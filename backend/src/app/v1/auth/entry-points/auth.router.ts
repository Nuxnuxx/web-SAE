import { Router } from "express";
import {
	handleDeleteProfil,
	handleLogin,
	handleModifyProfil,
	// handleDeleteProfil,
	// handleLogin,
	// handleModifyProfil,
	handleRegister,
} from "./auth.controller.js";
import { validate } from "./validators/auth.validator.js";
import { schemaRegisterBody } from "./validators/auth.schema.js";

const authRouter = Router();

authRouter.post("/register", validate(schemaRegisterBody), handleRegister);
authRouter.get("/login", handleLogin);
authRouter.put("/profil", handleModifyProfil);
authRouter.delete("/profil", handleDeleteProfil);

export default authRouter;
