package handler

import (
	"encoding/json"
	"errors"
	"pagar-me-test/delivery/parameter"
)

const (
	invalidParameter = "invalid parameters, check if all mandatory parameters have been sent!"
)

type Handler interface {
	Handle(*json.Decoder) (interface{}, error)
}

func ValidParameters(decoder *json.Decoder) (error, interface{}) {
	var parameter parameter.Parameter

	if err := decoder.Decode(&parameter); err != nil {
		return err, nil
	}

	if valid := parameter.Validator(); !valid {
		return errors.New(invalidParameter), nil
	}

	return nil, parameter

}
