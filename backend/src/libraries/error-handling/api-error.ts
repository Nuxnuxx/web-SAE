import { BaseError, HttpStatusCode } from "./app-error.js";

export class ApiError extends BaseError {
	constructor(
		message: string,
		methodName = "",
		httpCode: HttpStatusCode = HttpStatusCode.INTERNAL_SERVER,
		isOperational = true
	) {
		super("", httpCode, message, isOperational, methodName);
	}
}
