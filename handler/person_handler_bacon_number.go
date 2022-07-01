package handler

import (
	"encoding/json"
	"pagar-me-test/delivery/parameter"
	"pagar-me-test/domain/repository"

	log "github.com/sirupsen/logrus"
)

type PersonHandlerBaconNumber struct {
	PersonRepository repository.PersonRepository
}

func (p PersonHandlerBaconNumber) Handle(decoder *json.Decoder) (interface{}, error) {

	var parameter parameter.BaconNumberAndKinshipParameter

	if err := decoder.Decode(&parameter); err != nil {

		log.WithFields(
			log.Fields{
				"error": err,
			},
		).Error("err: ")
		return nil, err

	}

	return p.PersonRepository.BaconNumber(parameter.FirstPersonId, parameter.SecondPersonId)

}
