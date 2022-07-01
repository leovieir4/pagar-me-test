package delivery

import (
	"encoding/json"
	"net/http"
	"pagar-me-test/domain/repository"
	"pagar-me-test/handler"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HttpDelivery struct {
	Router *mux.Router
}

const (
	internalServerError    = "internal server error"
	errorHandlerNotDefined = "handler not defined!"
	defaultPort            = ":80"
)

func (h *HttpDelivery) executeHandlerPersonCreate(w http.ResponseWriter, r *http.Request) {
	h.executeHandler(w, r, handler.PersonHandlerCreate{PersonRepository: repository.PersonRepository{}})
}

func (h *HttpDelivery) executeHandlerPersonRelashion(w http.ResponseWriter, r *http.Request) {
	h.executeHandler(w, r, handler.PersonHandlerRelashion{PersonRepository: repository.PersonRepository{}})
}

func (h *HttpDelivery) executeHandlerPersonBaconNumber(w http.ResponseWriter, r *http.Request) {
	h.executeHandler(w, r, handler.PersonHandlerBaconNumber{PersonRepository: repository.PersonRepository{}})
}

func (h *HttpDelivery) executeHandlerPersonKiship(w http.ResponseWriter, r *http.Request) {
	h.executeHandler(w, r, handler.PersonHandlerKinship{PersonRepository: repository.PersonRepository{}})
}

func (h *HttpDelivery) executeHandlerPersonDelete(w http.ResponseWriter, r *http.Request) {

	PersonId, _ := strconv.ParseInt(mux.Vars(r)["person_id"], 10, 64)

	h.executeHandler(w, r, handler.PersonHandlerDelete{
		PersonRepository: repository.PersonRepository{},
		PersonId:         PersonId,
	})
}

func (h *HttpDelivery) executeHandlerPersonGenealogy(w http.ResponseWriter, r *http.Request) {

	PersonId, _ := strconv.ParseInt(mux.Vars(r)["person_id"], 10, 64)

	h.executeHandler(w, r, handler.PersonHandlerGenealogy{
		PersonRepository: repository.PersonRepository{},
		PersonId:         PersonId,
	})
}

func (h *HttpDelivery) executeHandler(w http.ResponseWriter, r *http.Request, handlerExc handler.Handler) {

	log.WithFields(
		log.Fields{
			"method": r.Method,
			"uri":    r.URL.String(),
		},
	).Info("Starting request...")

	if handlerExc == nil {

		log.WithFields(
			log.Fields{
				"error": errorHandlerNotDefined,
			},
		).Error("err: ")

		ErrorResponse(w, http.StatusBadRequest, internalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	response, errorResponse := handlerExc.Handle(decoder)

	if errorResponse != nil {
		log.WithFields(
			log.Fields{
				"error": errorResponse,
			},
		).Error("err: ")
		ErrorResponse(w, http.StatusBadRequest, errorResponse.Error())

	} else {
		JSONResponse(w, http.StatusOK, response)
	}
}

func (h *HttpDelivery) Execute() error {

	h.Router = mux.NewRouter()
	h.initializeRoutes()

	if errorResponse := http.ListenAndServe(defaultPort, h.Router); errorResponse != nil {
		log.WithFields(
			log.Fields{
				"error": errorResponse,
			},
		).Error("err: ")

		return errorResponse
	}

	return nil
}

func (h *HttpDelivery) initializeRoutes() {
	h.Router.HandleFunc("/person/create", h.executeHandlerPersonCreate).Methods("POST")
	h.Router.HandleFunc("/person/delete/{person_id}", h.executeHandlerPersonDelete).Methods("DELETE")
	h.Router.HandleFunc("/person/genealogy/{person_id}", h.executeHandlerPersonGenealogy).Methods("GET")
	h.Router.HandleFunc("/person/relashion", h.executeHandlerPersonRelashion).Methods("POST")
	h.Router.HandleFunc("/person/bacon_number", h.executeHandlerPersonBaconNumber).Methods("POST")
	h.Router.HandleFunc("/person/kinship", h.executeHandlerPersonKiship).Methods("POST")
}

func JSONResponse(w http.ResponseWriter, code int, response interface{}) {
	json, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(json)
}

func ErrorResponse(w http.ResponseWriter, code int, errorResponse string) {
	JSONResponse(w, code, map[string]string{"error": errorResponse})
}
