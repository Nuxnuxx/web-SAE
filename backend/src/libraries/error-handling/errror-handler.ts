import logging from "../logger/index.js";
import { BaseError } from "./app-error.js";

class ErrorHandler {
	async handleError(
		error: Error /* responseStream: Response */
	): Promise<void> {
		logging.logError(error);
		// fireMonitoringMetric(error);
		// await crashIfUntrustedErrorOrSendResponse(error, responseStream);
	}

	isTrustedError(error: Error) {
		if (error instanceof BaseError) {
			return error.isOperational;
		}
		return false;
	}
}

export const handler = new ErrorHandler();
