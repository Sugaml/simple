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

// GetInvoiceItems godoc
// @Summary Get InvoiceItems
// @Description Get list of InvoiceItems
// @Tags InvoiceItems
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.InvoiceItems
// @Router /payment/invoiceitems [get]
func (server *Server) GetInvoiceItems(w http.ResponseWriter, r *http.Request) {
	datas, err := server.DB.FindAllInvoiceItems()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}

// GetInvoiceItemsById godoc
// @Summary Get InvoiceItems by id
// @Description Get InvoiceItems by id
// @Tags InvoiceItems
// @Accept  json
// @Produce  json
// @Param id path int true "InvoiceItems id"
// @Success 200 {object} doc.InvoiceItems
// @Router /payment/invoiceitems/{id} [get]
func (server *Server) GetInvoiceItemsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataReceived, err := server.DB.FindByIdInvoiceItems(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

// UpdateInvoiceItems godoc
// @Summary Update a InvoiceItems
// @Description Update a InvoiceItems with the input payload
// @Tags InvoiceItems
// @Accept  json
// @Produce  json
// @Param id path int true "InvoiceItems id"
// @Param body body doc.InvoiceItems true "Update InvoiceItems"
// @Success 200 {object} doc.InvoiceItems
// @Router /payment/invoiceitems/{id} [put]
func (server *Server) UpdateInvoiceItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.InvoiceItems{}
	data, err = server.DB.FindByIdInvoiceItems(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate := models.InvoiceItems{}
	err = json.Unmarshal(body, &dataUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate.ID = data.ID
	dataUpdated, err := server.DB.UpdateInvoiceItems(dataUpdate)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataUpdated)
}

// DeleteInvoiceItems godoc
// @Summary Delete a InvoiceItems
// @Description Delete a InvoiceItems with the input payload
// @Tags InvoiceItems
// @Accept  json
// @Produce  json
// @Param id path int true "InvoiceItems id"
// @Success 204 {object} doc.InvoiceItems
// @Router /payment/invoiceitems/{id} [delete]
func (server *Server) DeleteInvoiceItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	//deduction.DeletedAt = time.Now()
	res, err := server.DB.DeleteInvoiceItems(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, res)
}
