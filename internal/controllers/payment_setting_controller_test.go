package controllers

import (
	"01cloud-payment/internal/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "01cloud-payment/internal/controllers/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPaymentSettingAPI(t *testing.T) {
	paymentsetting := models.PaymentSetting{
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "446600",
	}
	paymentsetting.ID = 1
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
					FindById(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(paymentsetting, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
			},
		},
		{
			name: "NotFound",
			PID:  paymentsetting.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindById(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(models.PaymentSetting{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, 404)
			},
		},
		{
			name: "InternalError",
			PID:  paymentsetting.ID,

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindById(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(models.PaymentSetting{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/payment/paymentsetting/%d", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreatePaymentSettingAPI(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"country":     "Nepal",
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "Lainchour",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				paymentsetting := models.PaymentSetting{
					Country:     "Nepal",
					State:       "Bagmati",
					City:        "Kathmandu",
					Street:      "Lainchour",
					Postal_Code: "446600",
				}
				store.EXPECT().
					Create(gomock.Any()).
					Times(1).
					Return(paymentsetting, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				//requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"country":     "",
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "Lainchour",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					Create(gomock.Any()).
					Times(1).
					Return(models.PaymentSetting{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Bad Request",
			body: gin.H{
				"country":     "",
				"state":       "",
				"city":        "",
				"street":      11,
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					Create(gomock.Any()).
					Times(0).
					Return(models.PaymentSetting{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
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

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			assert.NoError(t, err)

			url := "/payment/paymentsetting"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func requireBodyMatchPaymentSetting(t *testing.T, body *bytes.Buffer, paymentsetting models.PaymentSetting) {
	data, err := ioutil.ReadAll(body)
	assert.NoError(t, err)

	var gotPaymentsetting models.PaymentSetting
	err = json.Unmarshal(data, &gotPaymentsetting)
	assert.NoError(t, err)
	assert.Equal(t, paymentsetting, gotPaymentsetting)
}

func TestUpdatePaymentSettingAPI(t *testing.T) {
	paymentsetting := models.PaymentSetting{
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "446600",
	}
	paymentsetting.ID = 1
	args := models.PaymentSetting{
		Country: "India",
		State:   "Bagmati",
		City:    "Kathmandu",
		Street:  "Lainchour",
	}
	args.ID = 1
	testCases := []struct {
		name          string
		PID           uint
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  paymentsetting.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindById(gomock.Eq(paymentsetting.ID)).Times(1).Return(paymentsetting, nil)
				store.EXPECT().Update(gomock.Any()).Times(1).Return(paymentsetting, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
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

			url := fmt.Sprintf("/payment/paymentsetting/%d", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}

}

func TestDeletePaymentSettingAPI(t *testing.T) {
	paymentsetting := models.PaymentSetting{
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "446600",
	}
	paymentsetting.ID = 1
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
					Delete(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(int64(0), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
				//requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
			},
		},
		// {
		// 	name: "NotFound",
		// 	PID:  paymentsetting.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			FindById(gomock.Eq(paymentsetting.ID)).
		// 			Times(1).
		// 			Return(models.PaymentSetting{}, sql.ErrNoRows)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusNotFound, 404)
		// 	},
		// },
		// {
		// 	name: "InternalError",
		// 	PID:  paymentsetting.ID,

		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			FindById(gomock.Eq(paymentsetting.ID)).
		// 			Times(1).
		// 			Return(models.PaymentSetting{}, sql.ErrConnDone)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			//start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/payment/paymentsetting/%d", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}
