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

// CreateDeduction godoc
// @Summary Create a new Deduction
// @Description Create a new Deduction with the input payload
// @Tags Deduction
// @Accept  json
// @Produce  json
// @Param body body doc.Deduction true "Create Deduction"
// @Success 201 {object} doc.Deduction
// @Router /payment/deduction [post]
func (server *Server) CreateDeduction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Deduction{}
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
	dataCreated, err := server.DB.CreateDeduction(data)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

// GetDeduction godoc
// @Summary Get Deduction
// @Description Get list of deductions
// @Tags Deductions
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.Deduction
// @Router /payment/deduction [get]
func (server *Server) GetDeduction(w http.ResponseWriter, r *http.Request) {
	datas, err := server.DB.FindAllDeduction()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}

// GetDeductionById godoc
// @Summary Get Deduction by id
// @Description Get Deduction by id
// @Tags Deduction
// @Accept  json
// @Produce  json
// @Param id path int true "Deduction id"
// @Success 200 {object} doc.Deduction
// @Router /payment/deduction/{id} [get]
func (server *Server) GetDeductionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataReceived, err := server.DB.FindByIdDeduction(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

// UpdateDeduction godoc
// @Summary Update a Deduction
// @Description Update a Deduction with the input payload
// @Tags Deduction
// @Accept  json
// @Produce  json
// @Param id path int true "Deduction id"
// @Param body body doc.Deduction true "Update Deduction"
// @Success 200 {object} doc.Deduction
// @Router /payment/deduction/{id} [put]
func (server *Server) UpdateDeduction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	dataReceived, err := server.DB.FindByIdDeduction(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	d := models.Deduction{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	dataReceived.Name = d.Name
	dataReceived.Value = d.Value
	dataReceived.IsPercent = d.IsPercent
	dataReceived.Country = d.Country
	dataReceived.Attributes = d.Attributes
	dataReceived.Description = d.Description
	dataReceived.EffectiveDate = d.EffectiveDate
	dataReceived.UpdatedAt = time.Now()
	dataCreated, err := server.DB.UpdateDeduction(dataReceived)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

//vars := mux.Vars(r)
//pid, err := strconv.ParseUint(vars["id"], 10, 64)
//if err != nil {
//responses.ERROR(w, http.StatusNotFound, err)
//return
//}
//data := models.PaymentSetting{}
//dataReceived, err := data.Find(server.DB, pid)
//if err != nil {
//responses.ERROR(w, http.StatusNotFound, err)
//return
//}
//decoder := json.NewDecoder(r.Body)
//if err := decoder.Decode(&data); err != nil {
//responses.ERROR(w, http.StatusBadRequest, err)
//return
//}
//defer r.Body.Close()
//dataCreated, err := dataReceived.Update(server.DB)
//if err != nil {
//responses.ERROR(w, http.StatusInternalServerError, err)
//return
//}
//responses.JSON(w, http.StatusCreated, dataCreated)

// DeleteDeduction godoc
// @Summary Delete a Deduction
// @Description Delete a Deduction with the input payload
// @Tags Deduction
// @Accept  json
// @Produce  json
// @Param id path int true "Deduction id"
// @Success 204 {object} doc.Deduction
// @Router /payment/deduction/{id} [delete]
func (server *Server) DeleteDeduction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	res, err := server.DB.DeleteDeduction(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, res)
}
