import neo4j from "neo4j-driver";

var driver = neo4j.driver(
	"neo4j://localhost",
	neo4j.auth.basic("neo4j", "Wawa02290")
);

const database = driver.session();

export default database;
