import { Router } from "express";
import {
	deleteProfilController,
	loginController,
	modifyProfilController,
	registerController,
} from "./auth.controller.js";

const authRouter = Router();

authRouter.post("/register", registerController);
authRouter.get("/login", loginController);
authRouter.put("/profil", modifyProfilController);
authRouter.delete("/profil", deleteProfilController);

export default authRouter;
