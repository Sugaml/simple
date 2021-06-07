package controllers

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (server *Server) initializeRoutes() {
	server.Router.PathPrefix("/payment/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/payment/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	// server.Router.HandleFunc("/payment/invoice/{id}", server.CreateInvoice).Methods("POST")
	// server.Router.HandleFunc("/payment/invoice", server.GetInvoice).Methods("GET")
	// server.Router.HandleFunc("/payment/invoice/{id}", server.GetInvoiceById).Methods("GET")
	// server.Router.HandleFunc("/payment/invoice/{id}", server.UpdateInvoice).Methods("PUT")
	// server.Router.HandleFunc("/payment/invoice/{id}", server.DeleteInvoice).Methods("DELETE")

	// server.Router.HandleFunc("/payment/deduction", server.CreateDeduction).Methods("POST")
	// server.Router.HandleFunc("/payment/deduction", server.GetDeduction).Methods("GET")
	// server.Router.HandleFunc("/payment/deduction/{id}", server.GetDeductionById).Methods("GET")
	// server.Router.HandleFunc("/payment/deduction/{id}", server.UpdateDeduction).Methods("PUT")
	// server.Router.HandleFunc("/payment/deduction/{id}", server.DeleteDeduction).Methods("DELETE")

	// server.Router.HandleFunc("/payment/threshold", server.CreateThreshold).Methods("POST")
	// server.Router.HandleFunc("/payment/threshold", server.GetThreshold).Methods("GET")
	// server.Router.HandleFunc("/payment/threshold/{id}", server.GetThresholdById).Methods("GET")
	// server.Router.HandleFunc("/payment/threshold/{id}", server.UpdateThreshold).Methods("PUT")
	// server.Router.HandleFunc("/payment/threshold/{id}", server.DeleteThreshold).Methods("DELETE")

	server.Router.HandleFunc("/payment/promocode", server.CreatePromoCode).Methods("POST")
	server.Router.HandleFunc("/payment/promocode", server.GetPromoCode).Methods("GET")
	server.Router.HandleFunc("/payment/promocode/{id}", server.GetPromoCodeById).Methods("GET")
	server.Router.HandleFunc("/payment/promocode/{id}", server.UpdatePromoCode).Methods("PUT")
	server.Router.HandleFunc("/payment/promocode/{id}", server.DeletePromoCode).Methods("DELETE")

	// server.Router.HandleFunc("/payment/gateway", server.CreateGateway).Methods("POST")
	// server.Router.HandleFunc("/payment/gateway", server.GetGateway).Methods("GET")
	// server.Router.HandleFunc("/payment/gateway/{id}", server.GetGatewayById).Methods("GET")

	server.Router.HandleFunc("/payment/paymenthistory", server.CreatePaymentHistory).Methods("POST")
	server.Router.HandleFunc("/payment/paymenthistory", server.GetPaymentHistory).Methods("GET")
	server.Router.HandleFunc("/payment/paymenthistory/{id}", server.GetPaymentHistoryById).Methods("GET")
	server.Router.HandleFunc("/payment/paymenthistory/{id}", server.UpdatePaymentHistory).Methods("PUT")
	server.Router.HandleFunc("/payment/paymenthistory/{id}", server.DeletePaymentHistory).Methods("DELETE")

	// server.Router.HandleFunc("/payment/invoiceitems", server.CreateInvoiceItems).Methods("POST")
	// server.Router.HandleFunc("/payment/invoiceitems", server.GetInvoiceItems).Methods("GET")
	// server.Router.HandleFunc("/payment/invoiceitems/{id}", server.GetInvoiceItemsById).Methods("GET")
	// server.Router.HandleFunc("/payment/invoiceitems/{id}", server.UpdateInvoiceItems).Methods("PUT")
	// server.Router.HandleFunc("/payment/invoiceitems/{id}", server.DeleteInvoiceItems).Methods("DELETE")

	// server.Router.HandleFunc("/payment/transaction", server.CreateTransaction).Methods("POST")
	// server.Router.HandleFunc("/payment/transaction", server.GetTransaction).Methods("GET")
	// server.Router.HandleFunc("/payment/transaction/{id}", server.GetTransactionById).Methods("GET")

	server.Router.HandleFunc("/payment/paymentsetting", server.CreatePaymentSetting).Methods("POST")
	server.Router.HandleFunc("/payment/paymentsetting", server.GetPaymentSetting).Methods("GET")
	server.Router.HandleFunc("/payment/paymentsetting/{id}", server.GetPaymentSettingById).Methods("GET")
	server.Router.HandleFunc("/payment/paymentsetting/{id}", server.UpdatePaymentSetting).Methods("PUT")
	server.Router.HandleFunc("/payment/paymentsetting/{id}", server.DeletePaymentSetting).Methods("DELETE")
}
