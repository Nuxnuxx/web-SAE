import express from "express";
import morgan from "morgan";
import cors from "cors";
import v1Router from "./app/v1/index.router.js";

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

export default app;
