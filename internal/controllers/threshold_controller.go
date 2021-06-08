package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
	"errors"
	"strconv"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateThreshold godoc
// @Summary Create a new Threshold
// @Description Create a new Threshold with the input paylod
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Param body body doc.Threshold true "Create Threshold"
// @Success 201 {object} doc.Threshold
// @Router /payment/threshold [post]
func (server *Server) CreateThreshold(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Threshold{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataCreated, err := server.DB.CreateThreshold(data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

// GetThreshold godoc
// @Summary Get Threshold
// @Description Get list of threshold
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.Threshold
// @Router /payment/threshold [get]
func (server *Server) GetThreshold(w http.ResponseWriter, r *http.Request) {
	datas, err := server.DB.FindAllThreshold()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}

// GetThresholdById godoc
// @Summary Get Threshold by id
// @Description Get Threshold by id
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Param id path int true "Threshold id"
// @Success 200 {object} doc.Threshold
// @Router /payment/threshold/{id} [get]
func (server *Server) GetThresholdById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataReceived := models.Threshold{}
	dataReceived, err = server.DB.FindByIdThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

// UpdateThreshold godoc
// @Summary Update a Threshold
// @Description Update a Threshold with the input payload
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Param id path int true "Threshold id"
// @Param body body doc.Threshold true "Update Threshold"
// @Success 200 {object} doc.Threshold
// @Router /payment/threshold/{id} [put]
func (server *Server) UpdateThreshold(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Threshold{}
	data, err = server.DB.FindByIdThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("threshold not found"))
		return
	}
	dataUpdate := models.Threshold{}
	dataUpdate, err = server.DB.UpdateThreshold(data)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataUpdate)
}

// DeleteThreshold godoc
// @Summary Delete a Threshold
// @Description Delete a Threshold with the input payload
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Param id path int true "Threshold id"
// @Success 204 {object} doc.Threshold
// @Router /payment/threshold/{id} [delete]
func (server *Server) DeleteThreshold(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	pid, err := strconv.Atoi(aid)
	if err != nil {
		return
	}
	_, err = server.DB.DeleteThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
