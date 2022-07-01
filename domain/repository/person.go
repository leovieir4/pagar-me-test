package repository

import (
	"errors"
	"pagar-me-test/delivery/parameter"
	"pagar-me-test/domain/entity"
	"pagar-me-test/infrastructure/connectors"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	log "github.com/sirupsen/logrus"
)

const (
	IncestuousError = "relashion not allowed because is incestuous"
	QueryNotFound   = "not found check if all nodes exists"
)

type PersonRepository struct {
	GenealogyTree *GenealogyTree
}

type GenealogyTree struct {
	Members      []interface{}
	IteratedList []int64
}

func (p PersonRepository) BaconNumber(first_person_id, second_person_id int64) (interface{}, error) {

	driver, err := connectors.NewConnector(connectors.Neo4jDriver)

	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, errResponse := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (p1:Person) WHERE ID(p1) = $first_person_id MATCH (p2:Person) WHERE ID(p2) = $second_person_id return {bacon_number:length(shortestPath((p1)-[*..15]-(p2))), first_person:p1, second_person:p2}",
			map[string]interface{}{"first_person_id": first_person_id, "second_person_id": second_person_id})

		if err != nil {

			if err != nil {
				log.WithFields(
					log.Fields{
						"error": err,
					},
				).Error("err: ")
				return nil, err
			}
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if errResponse != nil {

		log.WithFields(
			log.Fields{
				"error": errResponse,
			},
		).Error("err: ")
		return nil, errResponse
	}

	return result, nil

}

func (p PersonRepository) validRelashion(parentId, childId int64) (bool, error) {

	childTree, _ := iterateTree(childId)
	candidateParentTree, err := iterateTree(parentId)

	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return false, err
	}

	if childTree == nil || candidateParentTree == nil {
		return false, errors.New(QueryNotFound)
	}

	mappedParent := childTree.(map[string]interface{})["relationships"].([]interface{})
	mappedCandidate := candidateParentTree.(map[string]interface{})["relationships"].([]interface{})

	for _, childItem := range mappedParent {
		childId := childItem.(map[string]interface{})["id"].(int64)

		for _, cadidateItem := range mappedCandidate {
			candidateId := cadidateItem.(map[string]interface{})["id"].(int64)

			if childId == candidateId {
				return true, nil
			}

		}

	}

	return false, nil

}
func (p PersonRepository) Kinship(first_person_id, second_person_id int64) (interface{}, error) {
	driver, err := connectors.NewConnector(connectors.Neo4jDriver)

	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, errorResponse := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (p1:Person) WHERE ID(p1) = $first_person_id MATCH (p2:Person) WHERE ID(p2) = $second_person_id return shortestPath((p1)-[*..15]-(p2))",
			map[string]interface{}{"first_person_id": first_person_id, "second_person_id": second_person_id})

		if err != nil {

			log.WithFields(
				log.Fields{
					"error": err,
				},
			).Error("err: ")

			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if errorResponse != nil {
		log.WithFields(
			log.Fields{
				"error": errorResponse,
			},
		).Error("err: ")
		return nil, errorResponse
	}

	relationString := ""

	distance := 0

	if result == nil {
		return nil, errors.New(QueryNotFound)
	}
	for i, item := range result.(dbtype.Path).Relationships {

		distance++

		if i+1 < len(result.(dbtype.Path).Relationships) {
			if item.EndId == result.(dbtype.Path).Relationships[i+1].StartId {
				relationString += "-PARENT"
			} else {
				relationString += "-SON"
			}
			continue
		}
		if first_person_id == item.StartId {
			relationString += "-PARENT"
		} else {
			relationString += "-SON"
		}

	}

	return getKinship(relationString, int64(distance)), nil
}

func getKinship(relationString string, distance int64) string {
	if distance == 1 {
		switch relationString {
		case "-SON":
			return "Dad/Mother"
		default:
			return "Son/Daughter"
		}
	}

	if distance == 2 {
		switch relationString {
		case "-SON-SON":
			return "Brother/Sister"
		case "-PARENT-SON":
			return "Grandfather/Grandmother"
		default:
			return "Grandson"
		}
	}

	if distance == 3 {
		switch relationString {
		case "-PARENT-SON-SON":
			return "Nephew/Niece"
		case "-PARENT-PARENT-SON":
			return "Great-grandmother"
		default:
			return "Uncle/Aunt"
		}
	}
	if distance == 4 {
		switch relationString {
		case "-PARENT-SON-SON-SON":
			return "Cousin"
		}
	}

	if distance == 5 {
		switch relationString {
		case "-PARENT-SON-SON-SON-SON":
			return "Uncle/Aunt"
		case "-SON-PARENT-SON-SON-SON":
			return "Nephew/Niece"
		}

	}

	return relationString
}

func (p PersonRepository) Relashion(parentId, childId int64) (interface{}, error) {

	valid, errorQuery := p.validRelashion(parentId, childId)

	if errorQuery != nil {
		log.WithFields(
			log.Fields{
				"error": errorQuery,
			},
		).Error("err: ")
		return nil, errorQuery
	}

	if valid {
		return nil, errors.New(IncestuousError)
	}

	driver, errorResponse := connectors.NewConnector(connectors.Neo4jDriver)

	if errorResponse != nil {
		log.WithFields(
			log.Fields{
				"error": errorResponse,
			},
		).Error("err: ")
		return nil, errorResponse
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, errQuery := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (child:Person) WHERE ID(child)=$child_id MATCH (parent:Person) WHERE ID(parent)=$parent_id CREATE (child)-[:CHILD_OF]->(parent)",
			map[string]interface{}{"child_id": childId, "parent_id": parentId})

		if err != nil {
			log.WithFields(
				log.Fields{
					"error": err,
				},
			).Error("err: ")
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if errQuery != nil {

		log.WithFields(
			log.Fields{
				"error": errQuery,
			},
		).Error("err: ")
		return nil, errQuery

	}

	return map[string]string{"message": "relashionship create with success!"}, nil
}
func iterateTree(personId int64) (interface{}, error) {

	driver, err := connectors.NewConnector(connectors.Neo4jDriver)

	if err != nil {

		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err

	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	return session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (p:Person) WHERE ID(p)=$id WITH p, [(p)<-[*]-(x) | { name: x.name, relationship:'sons', id: ID(x)}] as descendants,  [(x)<-[*]-(p) | {name: x.name, relationship:'parent',id: ID(x)}] as ancestors RETURN {id:ID(p), name: p.name, relationships:ancestors + descendants} ",
			map[string]interface{}{"id": personId})

		if err != nil {

			log.WithFields(
				log.Fields{
					"error": err,
				},
			).Error("err: ")
			return nil, err

		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

}
func (p PersonRepository) GetGenealogy(personId int64) error {

	p.GenealogyTree.IteratedList = append(p.GenealogyTree.IteratedList, personId)

	tree, err := iterateTree(personId)

	if err != nil {

		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return err

	}

	p.GenealogyTree.Members = append(p.GenealogyTree.Members, tree)

	if tree == nil {
		return errors.New(QueryNotFound)
	}

	items := tree.(map[string]interface{})["relationships"].([]interface{})

	for _, item := range items {

		itemMapped := item.(map[string]interface{})
		id, _ := itemMapped["id"].(int64)
		if p.checkItemInList(id) {

			p.GetGenealogy(id)
		}
	}

	return nil

}

func (p PersonRepository) checkItemInList(id int64) bool {
	for _, item := range p.GenealogyTree.IteratedList {
		if item == id {
			return false
		}
	}
	return true
}

func (p PersonRepository) Delete(personId int64) (interface{}, error) {

	driver, _ := connectors.NewConnector(connectors.Neo4jDriver)

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (p:Person) WHERE ID(p)=$id DETACH DELETE p",
			map[string]interface{}{"id": personId})

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "node deleted with success"}, nil

}

func (p PersonRepository) Create(person parameter.Person) (entity.Person, error) {
	driver, err := connectors.NewConnector(connectors.Neo4jDriver)

	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
	}

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	id, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (p:Person) SET p.name = $name, p.age = $age RETURN id(p)",
			map[string]interface{}{"name": person.Name, "age": person.Age})

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return entity.Person{}, err
	}
	createdPerson := entity.Person{
		Name: person.Name,
		Id:   id.(int64),
		Age:  person.Age,
	}
	return createdPerson, nil

}
