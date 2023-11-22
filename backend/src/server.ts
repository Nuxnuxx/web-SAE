import app from "./app.js";
import config from "./config/index.js";
import { handler } from "./libraries/error-handling/errror-handler.js";

process.on("uncaughtException", (error: Error) => {
	handler.handleError(error);
	if (!handler.isTrustedError(error)) {
		process.exit(1);
	}
});

process.on("unhandledRejection", (reason: Error, promise: Promise<any>) => {
	throw reason;
});

app.listen(config.port, () => console.log(`Listening on port ${config.port}`));
