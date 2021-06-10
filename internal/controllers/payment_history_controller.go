package controllers

import (
	"01cloud-payment/pkg/responses"
	"strconv"

	"net/http"
)

// GetPaymentHistory godoc
// @Summary Get PaymentHistory
// @Description Get list of PaymentHistory
// @Tags PaymentHistory
// @Accept  json
// @Produce  json
// @Param x-user-id header integer true "x-user-id"
// @Success 200 {array} doc.PaymentHistory
// @Router /payment/paymenthistory [get]
func (server *Server) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("x-user-id"), 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	datas, err := server.DB.FindByUserIDPaymentHistory(uint(userID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}
