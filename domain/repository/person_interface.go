package repository

import (
	"pagar-me-test/delivery/parameter"
	"pagar-me-test/domain/entity"
)

type Repository interface {
	Create(person parameter.Parameter) (entity.Person, error)
	Delete(personId int64) (interface{}, error)
	Relashion(parentId, child int64) (interface{}, error)
	BaconNumber(first_person_id, second_person_id int64) (interface{}, error)
	Kinship(first_person_id, second_person_id int64) (interface{}, error)
}
