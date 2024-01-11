import pino from "pino";

const logger = pino({
	transport: {
		targets: [
			{
				level: "info",
				target: "pino-pretty", // must be installed separately
			},
			// {
			// 	level: "trace",
			// 	target: "pino/file",
			// 	options: { destination: "/path/to/store/logs" },
			// },
		],
	},
});

export class Logging {
	logError(error: Error) {
		logger.error(error);
	}
}
const logging = new Logging();

export default logging;
