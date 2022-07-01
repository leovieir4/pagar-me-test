package handler

import (
	"encoding/json"
	"pagar-me-test/delivery/parameter"
	"pagar-me-test/domain/repository"

	log "github.com/sirupsen/logrus"
)

type PersonHandlerCreate struct {
	PersonRepository repository.PersonRepository
}

func (p PersonHandlerCreate) Handle(decoder *json.Decoder) (interface{}, error) {

	err, parameterValid := ValidParameters(decoder)

	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err
	}

	currentParameter := parameterValid.(parameter.Parameter)

	return p.PersonRepository.Create(currentParameter.Person)

}
