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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePromocodeSettingAPI(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"title":      "test",
				"code":       111,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				promocode := models.PromoCode{
					Title:     "test",
					Code:      111,
					IsPercent: true,
					Discount:  11,
					Limit:     3,
					Count:     1,
					Active:    true,
				}
				store.EXPECT().
					CreatePromocode(gomock.Any()).
					Times(1).
					Return(promocode, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"title":      "",
				"code":       111,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePromocode(gomock.Any()).
					Times(1).
					Return(models.PromoCode{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UnprocessableEnitity",
			body: gin.H{
				"title":      67,
				"code":       111,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePromocode(gomock.Any()).
					Times(0).
					Return(models.PromoCode{}, nil)
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

			url := "/payment/promocode"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Add("x-user-id", "1")
			assert.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func TestUpdatePromocodeAPI(t *testing.T) {
	promocode := models.PromoCode{
		Title:     "test",
		Code:      111,
		IsPercent: true,
		Discount:  11,
		Limit:     3,
		Count:     1,
		Active:    true,
	}
	promocode.ID = 1
	testCases := []struct {
		name          string
		PID           uint
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			PID:  promocode.ID,
			body: gin.H{
				"title":      "test",
				"code":       111,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPromocode(gomock.Eq(uint(promocode.ID))).Times(1).Return(promocode, nil)
				store.EXPECT().UpdatePromocode(promocode).Times(1).Return(promocode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "UnprocessedEnitity",
			PID:  promocode.ID,
			body: gin.H{
				"title":      "",
				"code":       "234",
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPromocode(gomock.Eq(uint(promocode.ID))).Times(1).Return(promocode, nil)
				store.EXPECT().UpdatePromocode(gomock.Any()).Times(0).Return(models.PromoCode{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name: "FindInternalError",
			PID:  promocode.ID,
			body: gin.H{
				"title":      "test",
				"code":       1234,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPromocode(gomock.Any()).Times(1).Return(models.PromoCode{}, sql.ErrConnDone)
				store.EXPECT().UpdatePromocode(gomock.Any()).Times(0).Return(promocode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			PID:  promocode.ID,
			body: gin.H{
				"title":      "test",
				"code":       1234,
				"is_percent": true,
				"discount":   11,
				"limit":      3,
				"count":      1,
				"active":     true,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().FindByIdPromocode(gomock.Eq(uint(promocode.ID))).Times(1).Return(promocode, nil)
				store.EXPECT().UpdatePromocode(gomock.Any()).Times(1).Return(models.PromoCode{}, sql.ErrConnDone)
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
			url := fmt.Sprintf("/payment/promocode/%d", tc.PID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			request.Header.Add("x-user-id", "1")
			require.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}

}

func TestGetByIDPromocodeAPI(t *testing.T) {
	promocode := models.PromoCode{
		Title:      "test",
		Code:       111,
		IsPercent:  true,
		ExpiryDate: time.Now(),
		Discount:   11,
		Limit:      3,
		Count:      1,
		Active:     true,
	}
	promocode.ID = 1
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
					FindByIdPromocode(gomock.Eq(promocode.ID)).
					Times(1).
					Return(promocode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
				//requireBodyMatchPromocode(t, recorder.Body, promocode)
			},
		},
		{
			name: "InternalError",
			PID:  promocode.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					FindByIdPromocode(gomock.Eq(promocode.ID)).
					Times(1).
					Return(models.PromoCode{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/payment/promocode/%d", tc.PID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeletePromocodeAPI(t *testing.T) {
	promocode := models.PromoCode{
		Title:      "test",
		Code:       111,
		IsPercent:  true,
		ExpiryDate: time.Now(),
		Discount:   11,
		Limit:      3,
		Count:      1,
		Active:     true,
	}
	promocode.ID = 1
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
				store.EXPECT().DeletePromocode(uint(promocode.ID)).Times(1).Return(int64(1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, recorder.Code)
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeletePromocode(gomock.Any()).
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

			url := fmt.Sprintf("/payment/promocode/%d", tc.PID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			assert.NoError(t, err)
			server.Router.ServeHTTP(recorder, request)
			//check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestFindAllPromocodeAPI(t *testing.T) {
	promocode := []models.PromoCode{
		{
			Title:     "test",
			Code:      111,
			IsPercent: true,
			Discount:  11,
			Limit:     3,
			Count:     1,
			Active:    true,
		},
		{
			Title:     "test",
			Code:      111,
			IsPercent: true,
			Discount:  11,
			Limit:     3,
			Count:     1,
			Active:    true,
		},
	}
	promocode[0].ID = 1
	promocode[1].ID = 2
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
					FindAllPromocode().
					Times(1).
					Return(promocode, nil)
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
					FindAllPromocode().
					Times(1).
					Return([]models.PromoCode{}, sql.ErrConnDone)
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

			url := "/payment/promocode"
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
}
