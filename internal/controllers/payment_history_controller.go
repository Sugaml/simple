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

// CreatePaymentHistory godoc
// @Summary Create a new PaymentHistory
// @Description Create a new PaymentHistory with the input paylod
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Param body body doc.PaymentHistory true "Create PaymentHistory"
// @Success 201 {object} doc.PaymentHistory
// @Router /payment/paymenthistory [post]
func (server *Server) CreatePaymentHistory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentHistory{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	/*data.Prepare()
	err = data.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}*/
	dataCreated, err := server.DB.CreatePaymentHistory(data)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

// GetPaymentHistory godoc
// @Summary Get PaymentHistory
// @Description Get list of PaymentHistory
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.PaymentHistory
// @Router /payment/paymenthistory [get]
func (server *Server) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	//data := []models.PaymentHistory{}
	data, err := server.DB.FindAllPaymentHistory()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, data)
}

// GetPaymentHistoryById godoc
// @Summary Get PaymentHistory by id
// @Description Get PaymentHistory by id
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentHistory id"
// @Success 200 {object} doc.PaymentHistory
// @Router /payment/paymenthistory/{id} [get]
func (server *Server) GetPaymentHistoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentHistory{}
	data, err = server.DB.FindByIdPaymentHistory(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, data)
}

// UpdatePaymentHistory godoc
// @Summary Update a PaymentHistory
// @Description Update a PaymentHistory with the input payload
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentHistory id"
// @Param body body doc.PaymentHistory true "Update PaymentHistory"
// @Success 200 {object} doc.PaymentHistory
// @Router /payment/paymenthistory/{id} [put]
func (server *Server) UpdatePaymentHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return
	}
	data := models.PaymentHistory{}
	data, err = server.DB.FindByIdPaymentHistory(uint(pid))
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
	dataUpdated, err := server.DB.UpdatePaymentHistory(data)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataUpdated)
}

// DeletePaymentHistory godoc
// @Summary Delete a PaymentHistory
// @Description Delete a PaymentHistory with the input payload
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Param id path int true "PaymentHistory id"
// @Success 204 {object} doc.PaymentHistory
// @Router /payment/paymenthistory/{id} [delete]
func (server *Server) DeletePaymentHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	pid, err := strconv.Atoi(aid)
	if err != nil {
		return
	}
	_, err = server.DB.DeletePaymentHistory(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
