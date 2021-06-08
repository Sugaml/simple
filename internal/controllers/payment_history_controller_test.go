package controllers

import (
	mockdb "01cloud-payment/internal/controllers/mock"
	"01cloud-payment/internal/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByIDPaymentHistoryAPI(t *testing.T) {
	paymenthistory := models.PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		TransactionID: 1,
		InvoiceID:     1,
	}
	paymenthistory.ID = 1
	testCases := []struct {
		name          string
		PID           uint
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPaymentHistory(gomock.Eq(paymenthistory.ID)).
					Times(1).
					Return(&paymenthistory, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentHistory(t, recorder.Body, paymenthistory)
			},
		},
		{
			name: "NotFound",
			PID:  paymenthistory.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPaymentHistory(gomock.Eq(paymenthistory.ID)).
					Times(1).
					Return(&models.PaymentHistory{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, 404)
			},
		},
		{
			name: "InternalError",
			PID:  paymenthistory.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPaymentHistory(gomock.Eq(paymenthistory.ID)).
					Times(1).
					Return(&models.PaymentHistory{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			//start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/payment/paymenthistory/%d", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreatePaymentHistoryAPI(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"userid":        1,
				"credit":        500,
				"debit":         500,
				"balance":       0,
				"transactionid": 1,
				"invoiceid":     1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				paymenthistory := models.PaymentHistory{
					UserID:        1,
					Credit:        500,
					Debit:         500,
					Balance:       0,
					TransactionID: 1,
					InvoiceID:     1,
				}
				store.EXPECT().
					CreatePaymentHistory(gomock.Any()).
					Times(1).
					Return(&paymenthistory, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"userid":        0,
				"credit":        500,
				"debit":         500,
				"balance":       0,
				"transactionid": 1,
				"invoiceid":     1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePaymentHistory(gomock.Any()).
					Times(1).
					Return(&models.PaymentHistory{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		// {
		// 	name: "Bad Request",
		// 	body: gin.H{
		// 		"userid":        "hello",
		// 		"credit":        500,
		// 		"debit":         500,
		// 		"balance":       0,
		// 		"transactionid": 1,
		// 		"invoiceid":     1,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			CreatePaymentHistory(gomock.Any()).
		// 			Times(0).
		// 			Return(&models.PaymentHistory{}, nil)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			assert.NoError(t, err)

			url := "/payment/paymenthistory"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func TestFindAllPaymentHistoryAPI(t *testing.T) {
	paymenthistory := []models.PaymentHistory{
		{
			UserID:        1,
			Credit:        500,
			Debit:         500,
			Balance:       0,
			TransactionID: 1,
			InvoiceID:     1,
		},
		{
			UserID:        2,
			Credit:        300,
			Debit:         300,
			Balance:       0,
			TransactionID: 2,
			InvoiceID:     2,
		},
	}
	type Query struct {
		pageID   int
		pageSize int
	}
	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindAllPaymentHistory(gomock.Any()).
					Times(1).
					Return(&paymenthistory, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: 5,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindAllPaymentHistory(gomock.Any()).
					Times(1).
					Return(&[]models.PaymentHistory{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/payment/paymenthistory"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeletePaymentHistoryAPI(t *testing.T) {
	paymenthistory := models.PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		TransactionID: 1,
		InvoiceID:     1,
	}
	paymenthistory.ID = 1
	testCases := []struct {
		name          string
		PID           uint
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeletePaymentHistory(uint(paymenthistory.ID)).Times(1).Return(int64(1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeletePaymentHistory(gomock.Any()).
					Times(1).
					Return(int64(0), sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			//start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/payment/paymenthistory/%d", tc.PID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchPaymentHistory(t *testing.T, body *bytes.Buffer, paymenthistory models.PaymentHistory) {
	data, err := ioutil.ReadAll(body)
	assert.NoError(t, err)

	var gotPaymenthistory models.PaymentHistory
	err = json.Unmarshal(data, &gotPaymenthistory)
	assert.NoError(t, err)
	assert.Equal(t, paymenthistory, gotPaymenthistory)
}
