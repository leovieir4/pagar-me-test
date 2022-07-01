package handler

import (
	"encoding/json"
	"pagar-me-test/domain/repository"
)

type PersonHandlerGenealogy struct {
	PersonRepository repository.PersonRepository
	PersonId         int64
}

func (p PersonHandlerGenealogy) Handle(decoder *json.Decoder) (interface{}, error) {

	tree := &repository.GenealogyTree{}

	p.PersonRepository.GenealogyTree = tree

	erro := p.PersonRepository.GetGenealogy(p.PersonId)

	return map[string]interface{}{"members": tree.Members}, erro

}
