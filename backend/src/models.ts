import neo4j from "neo4j-driver";
import config from "./config/index.js";

var driver = neo4j.driver(
	"neo4j://localhost",
	neo4j.auth.basic(config.secrets.DB_USER, config.secrets.DB_PASSWORD),
	{ maxTransactionRetryTime: 3000, maxConnectionPoolSize: 100 }
);

const database = driver.session({ database: "neo4j" });

export default database;
