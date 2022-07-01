package connectors

import (
	"pagar-me-test/configs"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	Neo4jDriver = "neo4j"
)

func NewConnector(connectorName string) (neo4j.Driver, error) {

	switch connectorName {

	case Neo4jDriver:
		return neo4j.NewDriver(configs.GetNeo4JURI(), neo4j.BasicAuth(configs.GetNeo4jUserName(), configs.GetNeo4JPassword(), ""))
	default:
		return nil, nil
	}
}
