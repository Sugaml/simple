package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
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
// @Param x-user-id header integer true "x-user-id"
// @Param body body doc.Threshold true "Create Threshold"
// @Success 201 {object} doc.Threshold
// @Router /payment/threshold [post]
func (server *Server) CreateThreshold(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentThreshold{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data.UserID = uint(userID)
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
// @Param x-user-id header integer true "x-user-id"
// @Success 200 {array} doc.Threshold
// @Router /payment/threshold [get]
func (server *Server) GetThreshold(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	datas, err := server.DB.FindByUserIDThreshold(uint(userID))
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
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "Threshold id"
// @Success 200 {object} doc.Threshold
// @Router /payment/threshold/{id} [get]
func (server *Server) GetThresholdById(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataReceived, err := server.DB.FindByIdThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
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
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "Threshold id"
// @Param body body doc.Threshold true "Update Threshold"
// @Success 200 {object} doc.Threshold
// @Router /payment/threshold/{id} [put]
func (server *Server) UpdateThreshold(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentThreshold{}
	data, err = server.DB.FindByIdThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate := models.PaymentThreshold{}
	err = json.Unmarshal(body, &dataUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate.ID = data.ID
	dataUpdated, err := server.DB.UpdateThreshold(dataUpdate)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataUpdated)
}

// DeleteThreshold godoc
// @Summary Delete a Threshold
// @Description Delete a Threshold with the input payload
// @Tags Threshold
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "Threshold id"
// @Success 204 {object} doc.Threshold
// @Router /payment/threshold/{id} [delete]
func (server *Server) DeleteThreshold(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = server.DB.DeleteThreshold(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "successfully deleted")
}
