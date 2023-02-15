package neo4jtester

type Neo4jDatabase struct {
	uri      string
	username string
	password string
}

// NewNeo4jDatabase returns a reference to the Neo4jDatabase struct
func NewNeo4jDatabase(uri, username, password string) *Neo4jDatabase {
	return &Neo4jDatabase{
		uri:      uri,
		username: username,
		password: password,
	}
}
