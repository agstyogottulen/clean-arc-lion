package http

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/agstyogottulen/clean-arc-lion/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/agstyogottulen/clean-arc-lion/common"
	"github.com/agstyogottulen/clean-arc-lion/courier/service"
)

type CourierHandler struct {
	CourierService service.Service
}

func NewCourierHandler(r *mux.Router, courierService service.Service) {
	handler := &CourierHandler{
		CourierService: courierService,
	}

	v1 := r.PathPrefix("/v1/courier").Subrouter()

	v1.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.Create))).Methods(http.MethodPost)
	v1.Handle("/all", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.ReadAll))).Methods(http.MethodGet)
	v1.Handle("/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.Read))).Methods(http.MethodGet)
	v1.Handle("/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.Update))).Methods(http.MethodPut)
	v1.Handle("/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handler.Delete))).Methods(http.MethodDelete)
	v1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		common.Response(w, common.Message(false, "url not found"))
		return
	})
}

func (c *CourierHandler) Create(w http.ResponseWriter, r *http.Request) {
	courier := new(models.Courier)

	if err := json.NewDecoder(r.Body).Decode(&courier); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "Invalid Request "+err.Error()))
		return
	}

	response, err := c.CourierService.Create(courier)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}

func (c *CourierHandler) Read(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "please provide valid id"))
		return
	}

	response, err := c.CourierService.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}

func (c *CourierHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.CourierService.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}

func (c *CourierHandler) Update(w http.ResponseWriter, r *http.Request) {
	courier := new(models.Courier)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "please provide valid id"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&courier); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "invalid request "+err.Error()))
		return
	}

	response, err := c.CourierService.Update(id, courier)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}

func (c *CourierHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, common.Message(false, "please provide valid id"))
		return
	}

	response, err := c.CourierService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		common.Response(w, response)
		return
	}

	common.Response(w, response)
	return
}
