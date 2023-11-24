export class BaseError extends Error {
	public readonly log: string;
	public readonly httpCode: HttpStatusCode;
	public readonly methodName?: string;
	public readonly isOperational: boolean;

	constructor(
		log: string,
		httpCode: HttpStatusCode = HttpStatusCode.INTERNAL_SERVER,
		description: string,
		isOperational: boolean,
		methodName?: string
	) {
		super(<string>description);
		Object.setPrototypeOf(this, new.target.prototype);

		this.log = log;
		this.httpCode = httpCode;
		this.isOperational = isOperational;
		if (methodName) this.methodName = methodName;

		Error.captureStackTrace(this);
	}
}

export enum HttpStatusCode {
	OK = 200,
	BAD_REQUEST = 400,
	NOT_FOUND = 404,
	INTERNAL_SERVER = 500,
}
