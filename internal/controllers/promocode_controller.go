package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
	"strconv"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// CreatePromoCode godoc
// @Summary Create a new PromoCode
// @Description Create a new PromoCode with the input payload
// @Tags PromoCode
// @Accept  json
// @Produce  json
// @Param body body doc.PromoCode true "Create PromoCode"
// @Success 201 {object} doc.PromoCode
// @Router /payment/promocode [post]
func (server *Server) CreatePromoCode(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PromoCode{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataCreated, err := server.DB.CreatePromocode(data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

// GetPromoCode godoc
// @Summary Get PromoCode
// @Description Get list of PromoCode
// @Tags PromoCode
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.PromoCode
// @Router /payment/promocode [get]
func (server *Server) GetPromoCode(w http.ResponseWriter, r *http.Request) {
	data, err := server.DB.FindAllPromocode()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, data)
}

// GetPromoCodeById godoc
// @Summary Get PromoCode by id
// @Description Get PromoCode by id
// @Tags PromoCode
// @Accept  json
// @Produce  json
// @Param id path int true "PromoCode id"
// @Success 200 {object} doc.PromoCode
// @Router /payment/promocode/{id} [get]
func (server *Server) GetPromoCodeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PromoCode{}
	data, err = server.DB.FindByIdPromocode(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, data)
}

// UpdatePromoCode godoc
// @Summary Update a PromoCode
// @Description Update a PromoCode with the input payload
// @Tags PromoCode
// @Accept  json
// @Produce  json
// @Param id path int true "PromoCode id"
// @Param body body doc.PromoCode true "Update PromoCode"
// @Success 200 {object} doc.PromoCode
// @Router /payment/promocode/{id} [put]
func (server *Server) UpdatePromoCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	dataReceived := models.PromoCode{}
	dataReceived, err = server.DB.FindByIdPromocode(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	d := models.PromoCode{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	dataReceived.Title = d.Title
	dataReceived.Code = d.Code
	dataReceived.IsPercent = d.IsPercent
	dataReceived.Discount = d.Discount
	dataReceived.Limit = d.Limit
	dataReceived.Count = d.Count
	dataReceived.Active = d.Active
	dataReceived.UpdatedAt = time.Now()

	dataCreated, err := server.DB.UpdatePromocode(dataReceived)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

// DeletePromoCode godoc
// @Summary Delete a PromoCode
// @Description Delete a PromoCode with the input payload
// @Tags PromoCode
// @Accept  json
// @Produce  json
// @Param id path int true "PromoCode id"
// @Success 204 {object} doc.PromoCode
// @Router /payment/promocode/{id} [delete]
func (server *Server) DeletePromoCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	pid, err := strconv.Atoi(aid)
	if err != nil {
		return
	}
	_, err = server.DB.DeletePromocode(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
