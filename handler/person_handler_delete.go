package handler

import (
	"encoding/json"
	"pagar-me-test/domain/repository"
)

type PersonHandlerDelete struct {
	PersonRepository repository.PersonRepository
	PersonId         int64
}

func (p PersonHandlerDelete) Handle(decoder *json.Decoder) (interface{}, error) {

	return p.PersonRepository.Delete(p.PersonId)

}
