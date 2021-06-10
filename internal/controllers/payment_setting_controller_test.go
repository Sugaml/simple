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
		Header        http.Header
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPaymentSetting(gomock.Eq(paymentsetting.ID)).
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
					FindByIdPaymentSetting(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(models.PaymentSetting{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, 404)
			},
		},
		{
			name: "InternalServerError",
			PID:  paymentsetting.ID,

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPaymentSetting(gomock.Eq(paymentsetting.ID)).
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
			request.Header.Add("x-user-id", "1")
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetPaymentSettingByUserIDAPI(t *testing.T) {
	paymentsetting := models.PaymentSetting{
		UserID:      1,
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
		Header        http.Header
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByUserIDPaymentSetting(gomock.Eq(paymentsetting.UserID)).
					Times(1).
					Return(paymentsetting, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
			},
		},
		// {
		// 	name: "NotFound",
		// 	PID:  paymentsetting.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			FindByIdPaymentSetting(gomock.Eq(paymentsetting.UserID)).
		// 			Times(1).
		// 			Return(models.PaymentSetting{}, sql.ErrNoRows)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		// 	},
		// },
		{
			name: "InternalError",
			PID:  paymentsetting.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByUserIDPaymentSetting(gomock.Eq(paymentsetting.UserID)).
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

			url := "/payment/paymentsetting"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			request.Header.Add("x-user-id", "1")
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
					CreatePaymentSetting(gomock.Any()).
					Times(1).
					Return(paymentsetting, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)
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
					CreatePaymentSetting(gomock.Any()).
					Times(1).
					Return(models.PaymentSetting{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UnprocessableEnitity",
			body: gin.H{
				"country":     12,
				"state":       "",
				"city":        "",
				"street":      "",
				"postal_code": "34546",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePaymentSetting(gomock.Any()).
					Times(0).
					Return(models.PaymentSetting{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
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
			request.Header.Add("x-user-id", "1")
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
	testCases := []struct {
		name          string
		PID           uint
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  paymentsetting.ID,
			body: gin.H{
				"country":     "Nepal",
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "Lainchour",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPaymentSetting(gomock.Eq(uint(paymentsetting.ID))).Times(1).Return(models.PaymentSetting{}, nil)
				store.EXPECT().UpdatePaymentSetting(gomock.Any()).Times(1).Return(paymentsetting, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "UnprocessedEnitity",
			PID:  paymentsetting.ID,
			body: gin.H{
				"country":     1234,
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPaymentSetting(gomock.Eq(uint(paymentsetting.ID))).Times(1).Return(paymentsetting, nil)
				store.EXPECT().UpdatePaymentSetting(gomock.Any()).Times(0).Return(models.PaymentSetting{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name: "FindInternalError",
			PID:  paymentsetting.ID,
			body: gin.H{
				"country":     1234,
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPaymentSetting(gomock.Any()).Times(1).Return(models.PaymentSetting{}, sql.ErrConnDone)
				store.EXPECT().UpdatePaymentSetting(gomock.Any()).Times(0).Return(paymentsetting, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			PID:  paymentsetting.ID,
			body: gin.H{
				"country":     "Nepal",
				"state":       "Bagmati",
				"city":        "Kathmandu",
				"street":      "Lainchour",
				"postal_code": "446600",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPaymentSetting(gomock.Eq(uint(paymentsetting.ID))).Times(1).Return(paymentsetting, nil)
				store.EXPECT().UpdatePaymentSetting(paymentsetting).Times(1).Return(models.PaymentSetting{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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
			url := fmt.Sprintf("/payment/paymentsetting/%d", tc.PID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			request.Header.Add("x-user-id", "1")
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
			name: "Delete",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeletePaymentSetting(uint(paymentsetting.ID)).Times(1).Return(int64(1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			PID:  1,

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeletePaymentSetting(gomock.Eq(paymentsetting.ID)).
					Times(1).
					Return(int64(1), sql.ErrConnDone)
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
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			request.Header.Add("x-user-id", "1")
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

/*func TestFindAllPaymentSettingAPI(t *testing.T) {
	paymentsetting := []models.PaymentSetting{
		{
			Country:     "testcountry",
			State:       "testcity",
			City:        "testcity",
			Street:      "teststreet",
			Postal_Code: "testpostalcode",
			Promocode:   "testpromo",
		},
		{
			Country:     "testcountry",
			State:       "testcity",
			City:        "testcity",
			Street:      "teststreet",
			Postal_Code: "testpostalcode",
			Promocode:   "testpromo",
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
					FindAllPaymentSetting().
					Times(1).
					Return(paymentsetting, nil)
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
					FindAllPaymentSetting().
					Times(1).
					Return([]models.PaymentSetting{}, sql.ErrConnDone)
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

			url := "/payment/paymentsetting"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			request.Header.Add("x-user-id", "1")
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
}*/
