package handler

import (
	"encoding/json"
	"pagar-me-test/delivery/parameter"
	"pagar-me-test/domain/repository"

	log "github.com/sirupsen/logrus"
)

type PersonHandlerRelashion struct {
	PersonRepository repository.PersonRepository
}

func (p PersonHandlerRelashion) Handle(decoder *json.Decoder) (interface{}, error) {

	var parameter parameter.RelashionParameter

	if err := decoder.Decode(&parameter); err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err
	}

	return p.PersonRepository.Relashion(parameter.ParentId, int64(parameter.ChildId))

}
