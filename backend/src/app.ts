import express, { NextFunction, Request, Response } from "express";
import morgan from "morgan";
import cors from "cors";
import v1Router from "./app/v1/index.router.js";
import { handler } from "./libraries/error-handling/errror-handler.js";

const app = express();

app.disable("x-powered-by");

app.use(morgan("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use(
	cors({
		credentials: true,
		origin: "http://localhost:5173",
	})
);

app.use("/api/v1", v1Router);

app.use(async (err: Error, req: Request, res: Response, next: NextFunction) => {
	if (!handler.isTrustedError(err)) {
		next(err);
	}

	await handler.handleError(err);
});

export default app;
