package controllers

import (
	mockdb "01cloud-payment/internal/controllers/mock"
	"01cloud-payment/internal/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetByIDPaymentHistoryAPI(t *testing.T) {
	paymenthistories := []models.PaymentHistory{}
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
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByUserIDPaymentHistory(gomock.Eq(paymenthistory.UserID)).
					Times(1).
					Return(paymenthistories, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentHistory(t, recorder.Body, paymenthistories)
			},
		},
		// {
		// 	name: "NotFound",
		// 	PID:  paymenthistory.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			FindByUserIDPaymentHistory(gomock.Eq(paymenthistory.UserID)).
		// 			Times(1).
		// 			Return(&models.PaymentHistory{}, sql.ErrNoRows)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusNotFound, 404)
		// 	},
		// },
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByUserIDPaymentHistory(gomock.Eq(paymenthistory.UserID)).
					Times(1).
					Return([]models.PaymentHistory{}, sql.ErrConnDone)
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

			url := "/payment/paymenthistory"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			request.Header.Add("x-user-id", "1")
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchPaymentHistory(t *testing.T, body *bytes.Buffer, paymenthistory []models.PaymentHistory) {
	data, err := ioutil.ReadAll(body)
	assert.NoError(t, err)

	var gotPaymenthistory []models.PaymentHistory
	err = json.Unmarshal(data, &gotPaymenthistory)
	assert.NoError(t, err)
	assert.Equal(t, len(paymenthistory), len(gotPaymenthistory))
}
