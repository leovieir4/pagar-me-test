package model

import (
	"fmt"

	"pagar-me-test/delivery/parameter"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type PersonModel struct {
	Driver neo4j.Driver
}

type PolicNotFoundError struct {
	Name string
	Id   int64
}

func (p PolicNotFoundError) Error() string {
	return fmt.Sprintf("Policy with name = %s and id = %d not found", p.Name, p.Id)
}

func (p *PersonModel) Get(parameter parameter.Parameter, schema string) {

}
