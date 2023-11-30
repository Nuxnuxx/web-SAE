import neo4j from "neo4j-driver";
import config from "./config/index.js";

var driver = neo4j.driver(
	"neo4j://localhost",
	neo4j.auth.basic(config.secrets.DB_USER, config.secrets.DB_PASSWORD)
);

const database = driver.session();

export default database;
