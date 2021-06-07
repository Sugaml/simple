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
			name: "InternalError",
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
					CreatePaymentSetting(gomock.Any()).
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
					CreatePaymentSetting(gomock.Any()).
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
				store.EXPECT().FindByIdPaymentSetting(gomock.Eq(uint(paymentsetting.ID))).Times(1).Return(paymentsetting, nil)
				store.EXPECT().UpdatePaymentSetting(paymentsetting.ID).Times(1).Return(paymentsetting, nil)
				fmt.Println(paymentsetting)
				//store.Update(paymentsetting)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, 201)
				//requireBodyMatchPaymentSetting(t, recorder.Body, paymentsetting)
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
			//store.Update(paymentsetting)
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
			name: "Delete",
			PID:  1,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPaymentSetting(gomock.Eq(uint(paymentsetting.ID))).Times(1).Return(paymentsetting, nil)
				store.EXPECT().DeletePaymentSetting(uint(paymentsetting.ID)).Times(1).Return(int64(1), nil)
				//store.Delete(uint(paymentsetting.ID))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, 204)
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
func TestFindAllPaymentSettingAPI(t *testing.T) {
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

// func TestPostApi(t *testing.T) {
// 	var jsonReq = []byte(`{"country":"nepal","state":"state","city":"ktm","street":"street"}`)
// 	req, _ := http.NewRequest("POST", "/paymentsetting/paymentsetting", bytes.NewBuffer(jsonReq))
// 	handler := http.HandlerFunc(tserver.CreatePaymentSetting)
// 	response := httptest.NewRecorder()
// 	handler.ServeHTTP(response, req)
// 	status := response.Code
// 	if status != http.StatusOK {
// 		t.Errorf("Handler Return a wrong status code : got %v want %v", status, http.StatusOK)
// 	}
// 	var paymentsetting models.PaymentSetting
// 	paymentsetting.ID = 1
// 	json.NewDecoder(io.Reader(response.Body)).Decode(&paymentsetting)
// 	assert.NotNil(t, paymentsetting.ID)
// 	assert.Equal(t, "Nepal", paymentsetting.Country)

// }

// func TestGetAllPaymentsetting(t *testing.T) {
// 	ps := models.PaymentSetting{
// 		Country:     "Nepal",
// 		State:       "Bagmati",
// 		City:        "Kathmandu",
// 		Street:      "Lainchour",
// 		Postal_Code: "134543",
// 		Promocode:   "NEWYEAR2021",
// 	}
// 	tserver.DB.Create(ps)
// 	req, _ := http.NewRequest("GET", "/payment/paymentsetting", nil)
// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(tserver.GetPaymentSetting)
// 	handler.ServeHTTP(r, req)
// 	checkStatusCode(r.Code, http.StatusOK, t)

// }

// func checkStatusCode(code int, want int, t *testing.T) {
// 	if code != want {
// 		t.Errorf("Wrong Content Type : got %v want %v", code, want)
// 	}
// }
