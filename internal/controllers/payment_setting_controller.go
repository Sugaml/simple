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
// @Param x-user-id header integer true "x-user-id"
// @Param body body doc.PaymentSetting true "Create PaymentSetting"
// @Success 201 {object} doc.PaymentSetting
// @Router /payment/paymentsetting [post]
func (server *Server) CreatePaymentSetting(w http.ResponseWriter, r *http.Request) {
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
	data := models.PaymentSetting{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data.UserID = uint(userID)
	dataCreated, err := server.DB.CreatePaymentSetting(data)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

/*func (server *Server) SavePromoCode(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentSetting{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataPromo := models.PromoCode{}
	dataCreated, err := dataPromo.FindByPromoCode(server.DB, data.Promocode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}*/

// GetPaymentSetting godoc
// @Summary Get PaymentSetting
// @Description Get list of PaymentSetting
// @Tags PaymentSetting
// @Param x-user-id header integer true "x-user-id"
// @Accept  json
// @Produce  json
// @Success 200 {array} doc.PaymentSetting
// @Router /payment/paymentsetting [get]
func (server *Server) GetPaymentSetting(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return
	}
	datas, err := server.DB.FindByUserIDPaymentSetting(uint(userID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
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
	/*	mm := r.Header.Add("xx", "ssss")
		// var s *http.Header

		// s.Add("x", "sss")
		// s = map[string][]string{

		// 	"x-user-id": {"1"},
		// }
		ss:= r.Header.Get("xx")
		fmt.Println("dd", mm)*/
	_, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	dataReceived, err := server.DB.FindByIdPaymentSetting(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

// UpdatePaymentSetting godoc
// @Summary Update a PaymentSetting
// @Description Update a PaymentSetting with the input payload
// @Tags PaymentSetting
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "PaymentSetting id"
// @Param body body doc.PaymentSetting true "Update PaymentSetting"
// @Success 200 {object} doc.PaymentSetting
// @Router /payment/paymentsetting/{id} [put]
func (server *Server) UpdatePaymentSetting(w http.ResponseWriter, r *http.Request) {
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
	data, err := server.DB.FindByIdPaymentSetting(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&data); err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	// defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate := models.PaymentSetting{}
	err = json.Unmarshal(body, &dataUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate.ID = data.ID
	dataCreated, err := server.DB.UpdatePaymentSetting(dataUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
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
	_, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.ParseUint(aid, 10, 64)
	_, err = server.DB.DeletePaymentSetting(uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
