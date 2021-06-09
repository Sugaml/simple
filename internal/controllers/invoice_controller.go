package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetInvoice godoc
// @Summary Get Invoice
// @Description Get list of invoice
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Param datestart path string true "Invoice datestart"
// @Param dateend path string true "Invoice dateend"
// @Param save path string true "Invoice save"
// @Success 201 {object} doc.Invoice
// @Router /payment/invoice [get]
func (server *Server) GetInvoice(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	firstday := r.URL.Query().Get("datestart")
	lastday := r.URL.Query().Get("dateend")
	t := time.Now()
	var datestart time.Time
	var dateend time.Time
	if firstday == "" {
		datestart = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)

	} else {
		datestart, err = time.Parse("2006-01-02", firstday)
		if err != nil {
			panic(err)
		}
	}
	if lastday == "" {
		dateend = time.Now()
	} else {
		dateend, err = time.Parse("2006-01-02", lastday)
		if err != nil {
			panic(err)
		}
	}

	//GetProject
	projectReceived, err := server.DB.FindAllByUser(uint(userID), datestart, dateend)
	if err != nil {
		panic(err)
	}
	var total float64
	invoiceItems := []models.InvoiceItems{}
	for _, project := range projectReceived {
		createdyear, createdmonth, _ := project.CreatedAt.Date()
		yearnow, monthnow, _ := dateend.Date()
		subscriptionReceived, err := server.DB.FindSubscription(uint(project.SubscriptionID))
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		projectPrice := subscriptionReceived.Price
		projectStartDate := datestart.Day()
		projectEndDate := dateend.Day()

		//GetTotal
		if projectStartDate < project.CreatedAt.Day() && createdyear == yearnow && createdmonth == monthnow {
			if project.DeletedAt != nil && projectEndDate > project.DeletedAt.Day() {
				projectStartDate = project.CreatedAt.Day()
				projectEndDate = project.DeletedAt.Day()
			} else {
				projectStartDate = project.CreatedAt.Day()
			}
		}
		effectiveTime := projectEndDate - projectStartDate + 1
		cost := float64(projectPrice) / 30 * float64(effectiveTime)
		invoiceItem := models.InvoiceItems{}
		invoiceItem.UserID = uint(userID)
		invoiceItem.Particular = project.Name
		invoiceItem.Days = uint(effectiveTime)
		invoiceItem.Rate = uint(projectPrice)
		invoiceItem.Total = math.Round(cost*100) / 100
		invoiceItems = append(invoiceItems, invoiceItem)
		total = total + cost

	}
	//deduction
	paymentsettings, _ := server.DB.FindByUserIDPaymentSetting(uint(userID))
	deductions, _ := server.DB.FindByCountryDeduction(paymentsettings.Country)
	VAT := int(deductions.Value)
	AppliedVAT := float64(VAT) / float64(100) * float64(total)
	//promocode
	promocodeData, _ := server.DB.FindByPromoCode(paymentsettings.Promocode)
	discount := int(promocodeData.Discount)
	appliedDiscount := float64(discount) / float64(100) * float64(total)
	finalTotal := total + AppliedVAT - appliedDiscount
	//threshold
	thresholdGet, err := server.DB.FindByUserIDThreshold(uint(userID))
	if err != nil {
		log.Print(err)
	}
	if thresholdGet != (models.PaymentThreshold{}) {
		if thresholdGet.ThresholdLimit < uint(total) {
			//send notification
			log.Println(total)
		}

	}

	invoice := models.Invoice{}
	invoice.UserID = uint(userID)
	invoice.TotalCost = math.Round(finalTotal*100) / 100
	invoice.StartDate = datestart
	invoice.EndDate = dateend
	invoice.DeductionID = deductions.ID
	save := r.URL.Query().Get("save")
	if save == "true" {
		invoiceSaved, err := server.DB.CreateInvoice(invoice)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		for _, invoiceItem := range invoiceItems {
			invoiceItem.InvoiceID = invoiceSaved.ID
			_, error := server.DB.CreateInvoiceItems(invoiceItem)
			if error != nil {
				responses.ERROR(w, http.StatusNotFound, err)
				return
			}
		}
		invoice.InvoiceItems = &invoiceItems
		paymentHistory := models.PaymentHistory{}
		paymentHistory.Credit = invoiceSaved.TotalCost
		paymentHistory.InvoiceID = invoiceSaved.ID
		paymentHistory.Balance = invoiceSaved.TotalCost
		paymentHistory.UserID = invoiceSaved.UserID
		paymentHistory.Date = invoice.EndDate
		_, error := server.DB.CreatePaymentHistory(paymentHistory)
		if error != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
	}
	invoice.InvoiceItems = &invoiceItems
	responses.JSON(w, http.StatusOK, invoice)
}

// func (server *Server) GetAllInvoice(w http.ResponseWriter, r *http.Request) {
// 	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	datas, err := server.DB.FindAllInvoice( uint(userID))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	fmt.Println("ddd", datas)
// 	responses.JSON(w, http.StatusOK, datas)
// }

// GetInvoiceById godoc
// @Summary Get Invoice by id
// @Description Get Invoice by id
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "Invoice id"
// @Success 200 {object} doc.Invoice
// @Router /payment/invoice/{id} [get]
func (server *Server) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
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
	dataReceived, err := server.DB.FindByIdInvoice(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

// UpdateInvoice godoc
// @Summary Update a Invoice
// @Description Update a Invoice with the input payload
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Param id path int true "Invoice id"
// @Param body body doc.Invoice true "Update Invoice"
// @Success 200 {object} doc.Invoice
// @Router /payment/invoice/{id} [put]
func (server *Server) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
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
	data, err := server.DB.FindByIdInvoice(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate := models.Invoice{}
	err = json.Unmarshal(body, &dataUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	dataUpdate.ID = data.ID
	dataUpdated, err := server.DB.UpdateInvoice(dataUpdate)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataUpdated)
}

// DeleteInvoice godoc
// @Summary Delete a Invoice
// @Description Delete a Invoice with the input payload
// @Param x-user-id header integer true "x-user-id"
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param id path int true "Invoice id"
// @Success 204 {object} doc.Invoice
// @Router /payment/invoice/{id} [delete]
func (server *Server) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
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
	_, err = server.DB.DeleteInvoice(uint(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
