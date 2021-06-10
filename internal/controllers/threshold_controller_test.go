package controllers

import (
	mockdb "01cloud-payment/internal/controllers/mock"
	"01cloud-payment/internal/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentThresholdAPI(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"userid":         1,
				"thresholdlimit": 2000,
				"email":          "test@gmail.com",
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				threshold := models.PaymentThreshold{
					UserID:         1,
					ThresholdLimit: 2000,
					Email:          "test@gmail.com",
					Active:         true,
				}
				store.EXPECT().
					CreateThreshold(gomock.Any()).
					Times(1).
					Return(threshold, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"userid":         1,
				"thresholdlimit": 2000,
				"email":          "test@gmail.com",
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateThreshold(gomock.Any()).
					Times(1).
					Return(models.PaymentThreshold{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UnprocessableEnitity",
			body: gin.H{
				"userid":         1,
				"thresholdlimit": "2000",
				"email":          "test@gmail.com",
				"active":         "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateThreshold(gomock.Any()).
					Times(0).
					Return(models.PaymentThreshold{}, nil)
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

			url := "/payment/threshold"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("x-user-id", "1")
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func TestGetByIDThresholdAPI(t *testing.T) {
	threshold := models.PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	threshold.ID = 1
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
					FindByIdThreshold(gomock.Eq(threshold.ID)).
					Times(1).
					Return(threshold, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NotFound",
			PID:  threshold.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdThreshold(gomock.Eq(threshold.ID)).
					Times(1).
					Return(models.PaymentThreshold{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, 404)
			},
		},
		{
			name: "InternalServerError",
			PID:  threshold.ID,

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdThreshold(gomock.Eq(threshold.ID)).
					Times(1).
					Return(models.PaymentThreshold{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/payment/threshold/%d", tc.PID)
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
func TestGetThresholdByUserIDAPI(t *testing.T) {
	threshold := models.PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	threshold.ID = 1
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
					FindByUserIDThreshold(gomock.Eq(threshold.UserID)).
					Times(1).
					Return(threshold, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		// {
		// 	name: "NotFound",
		// 	PID:  threshold.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			FindByUserIDThreshold(gomock.Eq(threshold.UserID)).
		// 			Times(1).
		// 			Return(models.PaymentThreshold{}, sql.ErrNoRows)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		// 	},
		// },
		{
			name: "InternalError",
			PID:  threshold.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByUserIDThreshold(gomock.Eq(threshold.UserID)).
					Times(1).
					Return(models.PaymentThreshold{}, sql.ErrConnDone)
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

			url := "/payment/threshold"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			request.Header.Add("x-user-id", "1")
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdatePaymentThreshholdAPI(t *testing.T) {
	threshold := models.PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	threshold.ID = 1
	testCases := []struct {
		name          string
		PID           uint
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  threshold.ID,
			body: gin.H{
				"userid":         1,
				"thresholdlimit": 2000,
				"email":          "test@gmail.com",
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdThreshold(gomock.Eq(uint(threshold.ID))).Times(1).Return(models.PaymentThreshold{}, nil)
				store.EXPECT().UpdateThreshold(gomock.Any()).Times(1).Return(threshold, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "UnprocessedEnitity",
			PID:  threshold.ID,
			body: gin.H{
				"userid":         1,
				"thresholdlimit": "2000",
				"email":          1234,
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdThreshold(gomock.Eq(uint(threshold.ID))).Times(1).Return(threshold, nil)
				store.EXPECT().UpdateThreshold(gomock.Any()).Times(0).Return(models.PaymentThreshold{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name: "FindInternalError",
			PID:  threshold.ID,
			body: gin.H{
				"userid":         1,
				"thresholdlimit": 2000,
				"email":          "test@gmail.com",
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdThreshold(gomock.Any()).Times(1).Return(models.PaymentThreshold{}, sql.ErrConnDone)
				store.EXPECT().UpdateThreshold(gomock.Any()).Times(0).Return(threshold, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			PID:  threshold.ID,
			body: gin.H{
				"userid":         1,
				"thresholdlimit": 2000,
				"email":          "test@gmail.com",
				"active":         true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdThreshold(gomock.Eq(uint(threshold.ID))).Times(1).Return(threshold, nil)
				store.EXPECT().UpdateThreshold(gomock.Any()).Times(1).Return(models.PaymentThreshold{}, sql.ErrConnDone)
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
			url := fmt.Sprintf("/payment/threshold/%d", tc.PID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			request.Header.Add("x-user-id", "1")
			require.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}

}
func TestDeleteThresholdAPI(t *testing.T) {
	threshold := models.PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	threshold.ID = 1
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
				store.EXPECT().DeleteThreshold(uint(threshold.ID)).Times(1).Return(int64(1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteThreshold(gomock.Any()).
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

			url := fmt.Sprintf("/payment/threshold/%d", tc.PID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			assert.NoError(t, err)
			request.Header.Add("x-user-id", "1")
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

// func TestFindAllThresholdAPI(t *testing.T) {
// 	threshold := []models.PaymentThreshold{
// 		{
// 			UserID:         1,
// 			ThresholdLimit: 2000,
// 			Email:          "test@gmail.com",
// 			Active:         true,
// 		},
// 		{
// 			UserID:         1,
// 			ThresholdLimit: 2000,
// 			Email:          "test@gmail.com",
// 			Active:         true,
// 		},
// 	}
// 	threshold[0].ID = 1
// 	threshold[1].ID = 2
// 	type Query struct {
// 		pageID   int
// 		pageSize int
// 	}
// 	testCases := []struct {
// 		name          string
// 		query         Query
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			query: Query{
// 				pageID:   1,
// 				pageSize: 5,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					FindAllThreshold().
// 					Times(1).
// 					Return(threshold, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			query: Query{
// 				pageID:   1,
// 				pageSize: 5,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().
// 					FindAllThreshold().
// 					Times(1).
// 					Return([]models.PaymentThreshold{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			url := "/payment/threshold"
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			request.Header.Add("x-user-id", "1")
// 			require.NoError(t, err)

// 			// Add query parameters to request URL
// 			q := request.URL.Query()
// 			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
// 			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
// 			request.URL.RawQuery = q.Encode()

// 			server.Router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }
