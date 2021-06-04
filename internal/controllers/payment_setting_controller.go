package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePaymentSetting godoc
// @Summary Create a new PaymentSetting
// @Description Create a new PaymentSetting with the input paylod
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Param body body doc.PaymentSetting true "Create PaymentSetting"
// @Success 201 {object} doc.PaymentSetting
// @Router /payment/paymentsetting [post]
func (server *Server) CreatePaymentSetting(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusNoContent, err)
		return
	}
	data := models.PaymentSetting{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusNoContent, err)
		return
	}
	paymentsetting, err := server.DB.Create(data)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, paymentsetting)
}

// GetPaymentSetting godoc
// @Summary Get PaymentSetting
// @Description Get list of PaymentSetting
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.PaymentSetting
// @Router /payment/paymentsetting [get]
func (server *Server) GetPaymentSetting(w http.ResponseWriter, r *http.Request) {
	paymentsetting, err := server.DB.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, paymentsetting)
}

// GetPaymentSettingById godoc
// @Summary Get PaymentSetting by id
// @Description Get PaymentSetting by id
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentSetting id"
// @Success 200 {object} doc.PaymentSetting
// @Router /payment/paymentsetting/{id} [get]
func (server *Server) GetPaymentSettingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	data := models.PaymentSetting{}
	data, err = server.DB.FindById(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, data)
}

// UpdatePaymentSetting godoc
// @Summary Update a PaymentSetting
// @Description Update a PaymentSetting with the input payload
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentSetting id"
// @Param body body doc.PaymentSetting true "Update PaymentSetting"
// @Success 200 {object} doc.PaymentSetting
// @Router /payment/paymentsetting/{id} [put]
func (server *Server) UpdatePaymentSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return
	}
	data := models.PaymentSetting{}
	data, err = server.DB.FindById(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	dataUpdated, err := server.DB.Update(data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataUpdated)
}

// DeletePaymentSetting godoc
// @Summary Delete a PaymentSetting
// @Description Delete a PaymentSetting with the input payload
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentSetting id"
// @Success 204 {object} doc.PaymentSetting
// @Router /payment/paymentsetting/{id} [delete]
func (server *Server) DeletePaymentSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	pid, err := strconv.Atoi(aid)
	if err != nil {
		return
	}
	//paymentsetting := models.PaymentSetting{}
	_, err = server.DB.Delete(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
