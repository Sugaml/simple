package controllers

import (
	"01cloud-payment/internal/middleware"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (server *Server) setJSON(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.Router.HandleFunc(path, middleware.SetMiddlewareJSON(next)).Methods(method, "OPTIONS")
}
func (server *Server) setAdmin(path string, next func(http.ResponseWriter, *http.Request), method string) {
	server.setJSON(path, middleware.SetAdminMiddlewareAuthentication(next), method)
}

func (server *Server) initializeRoutes() {
	server.Router.PathPrefix("/payment/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("https://api.test.01cloud.dev/payment/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	// server.setJSON("/payment/invoice", server.CreateInvoice, "POST")
	// server.setJSON("/payment/invoice", server.GetInvoice, "GET")
	// server.setJSON("/payment/invoice/{id}", server.GetInvoiceById, "GET")
	// server.setJSON("/payment/invoice/{id}", server.UpdateInvoice, "PUT")
	// server.setJSON("/payment/invoice/{id}", server.DeleteInvoice, "DELETE")

	server.setAdmin("/payment/deduction", server.CreateDeduction, "POST")
	server.setAdmin("/payment/deduction", server.GetDeduction, "GET")
	server.setAdmin("/payment/deduction/{id}", server.GetDeductionById, "GET")
	server.setAdmin("/payment/deduction/{id}", server.UpdateDeduction, "PUT")
	server.setAdmin("/payment/deduction/{id}", server.DeleteDeduction, "DELETE")

	server.setJSON("/payment/threshold", server.CreateThreshold, "POST")
	server.setJSON("/payment/threshold", server.GetThreshold, "GET")
	server.setJSON("/payment/threshold/{id}", server.GetThresholdById, "GET")
	server.setJSON("/payment/threshold/{id}", server.UpdateThreshold, "PUT")
	server.setJSON("/payment/threshold/{id}", server.DeleteThreshold, "DELETE")

	server.setAdmin("/payment/promocode", server.CreatePromoCode, "POST")
	server.setAdmin("/payment/promocode", server.GetPromoCode, "GET")
	server.setAdmin("/payment/promocode/{id}", server.GetPromoCodeById, "GET")
	server.setAdmin("/payment/promocode/{id}", server.UpdatePromoCode, "PUT")
	server.setAdmin("/payment/promocode/{id}", server.DeletePromoCode, "DELETE")

	server.setJSON("/payment/gateway", server.CreateGateway, "POST")
	server.setJSON("/payment/gateway", server.GetGateway, "GET")
	server.setJSON("/payment/gateway/{id}", server.GetGatewayById, "GET")

	server.setJSON("/payment/paymenthistory", server.CreatePaymentHistory, "POST")
	server.setJSON("/payment/paymenthistory", server.GetPaymentHistory, "GET")
	server.setJSON("/payment/paymenthistory/{id}", server.GetPaymentHistoryById, "GET")
	server.setJSON("/payment/paymenthistory/{id}", server.UpdatePaymentHistory, "PUT")
	server.setJSON("/payment/paymenthistory/{id}", server.DeletePaymentHistory, "DELETE")

	// server.setJSON("/payment/invoiceitems", server.CreateInvoiceItems, "POST")
	// server.setJSON("/payment/invoiceitems", server.GetInvoiceItems, "GET")
	// server.setJSON("/payment/invoiceitems/{id}", server.GetInvoiceItemsById, "GET")
	// server.setJSON("/payment/invoiceitems/{id}", server.UpdateInvoiceItems, "PUT")
	// server.setJSON("/payment/invoiceitems/{id}", server.DeleteInvoiceItems, "DELETE")

	server.setJSON("/payment/transaction", server.CreateTransaction, "POST")
	server.setJSON("/payment/transaction", server.GetTransaction, "GET")
	server.setJSON("/payment/transaction/{id}", server.GetTransactionById, "GET")

	server.setJSON("/payment/paymentsetting", server.CreatePaymentSetting, "POST")
	server.setJSON("/payment/paymentsetting", server.GetPaymentSetting, "GET")
	server.setJSON("/payment/paymentsetting/{id}", server.GetPaymentSettingById, "GET")
	server.setJSON("/payment/paymentsetting/{id}", server.UpdatePaymentSetting, "PUT")
	server.setJSON("/payment/paymentsetting/{id}", server.DeletePaymentSetting, "DELETE")
}
