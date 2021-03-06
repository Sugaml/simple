// Code generated by MockGen. DO NOT EDIT.
// Source: internal/models/store.go

// Package mockdb is a generated GoMock package.
package mockdb

import (
	models "01cloud-payment/internal/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateDeduction mocks base method.
func (m *MockStore) CreateDeduction(data models.Deduction) (models.Deduction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeduction", data)
	ret0, _ := ret[0].(models.Deduction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeduction indicates an expected call of CreateDeduction.
func (mr *MockStoreMockRecorder) CreateDeduction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeduction", reflect.TypeOf((*MockStore)(nil).CreateDeduction), data)
}

// CreateGateway mocks base method.
func (m *MockStore) CreateGateway(data models.Gateway) (models.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGateway", data)
	ret0, _ := ret[0].(models.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGateway indicates an expected call of CreateGateway.
func (mr *MockStoreMockRecorder) CreateGateway(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGateway", reflect.TypeOf((*MockStore)(nil).CreateGateway), data)
}

// CreateInvoice mocks base method.
func (m *MockStore) CreateInvoice(data models.Invoice) (models.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", data)
	ret0, _ := ret[0].(models.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoice indicates an expected call of CreateInvoice.
func (mr *MockStoreMockRecorder) CreateInvoice(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockStore)(nil).CreateInvoice), data)
}

// CreateInvoiceItems mocks base method.
func (m *MockStore) CreateInvoiceItems(data models.InvoiceItems) (models.InvoiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoiceItems", data)
	ret0, _ := ret[0].(models.InvoiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoiceItems indicates an expected call of CreateInvoiceItems.
func (mr *MockStoreMockRecorder) CreateInvoiceItems(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoiceItems", reflect.TypeOf((*MockStore)(nil).CreateInvoiceItems), data)
}

// CreatePaymentHistory mocks base method.
func (m *MockStore) CreatePaymentHistory(data models.PaymentHistory) (models.PaymentHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePaymentHistory", data)
	ret0, _ := ret[0].(models.PaymentHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePaymentHistory indicates an expected call of CreatePaymentHistory.
func (mr *MockStoreMockRecorder) CreatePaymentHistory(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePaymentHistory", reflect.TypeOf((*MockStore)(nil).CreatePaymentHistory), data)
}

// CreatePaymentSetting mocks base method.
func (m *MockStore) CreatePaymentSetting(data models.PaymentSetting) (models.PaymentSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePaymentSetting", data)
	ret0, _ := ret[0].(models.PaymentSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePaymentSetting indicates an expected call of CreatePaymentSetting.
func (mr *MockStoreMockRecorder) CreatePaymentSetting(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePaymentSetting", reflect.TypeOf((*MockStore)(nil).CreatePaymentSetting), data)
}

// CreatePromocode mocks base method.
func (m *MockStore) CreatePromocode(data models.PromoCode) (models.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePromocode", data)
	ret0, _ := ret[0].(models.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePromocode indicates an expected call of CreatePromocode.
func (mr *MockStoreMockRecorder) CreatePromocode(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePromocode", reflect.TypeOf((*MockStore)(nil).CreatePromocode), data)
}

// CreateThreshold mocks base method.
func (m *MockStore) CreateThreshold(data models.PaymentThreshold) (models.PaymentThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateThreshold", data)
	ret0, _ := ret[0].(models.PaymentThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateThreshold indicates an expected call of CreateThreshold.
func (mr *MockStoreMockRecorder) CreateThreshold(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateThreshold", reflect.TypeOf((*MockStore)(nil).CreateThreshold), data)
}

// CreateTransaction mocks base method.
func (m *MockStore) CreateTransaction(data models.Transaction) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", data)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockStoreMockRecorder) CreateTransaction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockStore)(nil).CreateTransaction), data)
}

// DeleteDeduction mocks base method.
func (m *MockStore) DeleteDeduction(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeduction", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDeduction indicates an expected call of DeleteDeduction.
func (mr *MockStoreMockRecorder) DeleteDeduction(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeduction", reflect.TypeOf((*MockStore)(nil).DeleteDeduction), pid)
}

// DeleteGateway mocks base method.
func (m *MockStore) DeleteGateway(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGateway", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGateway indicates an expected call of DeleteGateway.
func (mr *MockStoreMockRecorder) DeleteGateway(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGateway", reflect.TypeOf((*MockStore)(nil).DeleteGateway), pid)
}

// DeleteInvoice mocks base method.
func (m *MockStore) DeleteInvoice(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvoice", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInvoice indicates an expected call of DeleteInvoice.
func (mr *MockStoreMockRecorder) DeleteInvoice(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvoice", reflect.TypeOf((*MockStore)(nil).DeleteInvoice), pid)
}

// DeleteInvoiceItems mocks base method.
func (m *MockStore) DeleteInvoiceItems(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvoiceItems", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInvoiceItems indicates an expected call of DeleteInvoiceItems.
func (mr *MockStoreMockRecorder) DeleteInvoiceItems(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvoiceItems", reflect.TypeOf((*MockStore)(nil).DeleteInvoiceItems), pid)
}

// DeletePaymentSetting mocks base method.
func (m *MockStore) DeletePaymentSetting(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePaymentSetting", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePaymentSetting indicates an expected call of DeletePaymentSetting.
func (mr *MockStoreMockRecorder) DeletePaymentSetting(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePaymentSetting", reflect.TypeOf((*MockStore)(nil).DeletePaymentSetting), pid)
}

// DeletePromocode mocks base method.
func (m *MockStore) DeletePromocode(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePromocode", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePromocode indicates an expected call of DeletePromocode.
func (mr *MockStoreMockRecorder) DeletePromocode(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePromocode", reflect.TypeOf((*MockStore)(nil).DeletePromocode), pid)
}

// DeleteThreshold mocks base method.
func (m *MockStore) DeleteThreshold(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteThreshold", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteThreshold indicates an expected call of DeleteThreshold.
func (mr *MockStoreMockRecorder) DeleteThreshold(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThreshold", reflect.TypeOf((*MockStore)(nil).DeleteThreshold), pid)
}

// DeleteTransaction mocks base method.
func (m *MockStore) DeleteTransaction(pid uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", pid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockStoreMockRecorder) DeleteTransaction(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockStore)(nil).DeleteTransaction), pid)
}

// FindAllByUser mocks base method.
func (m *MockStore) FindAllByUser(userID uint, startDate, endDate time.Time) ([]models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByUser", userID, startDate, endDate)
	ret0, _ := ret[0].([]models.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByUser indicates an expected call of FindAllByUser.
func (mr *MockStoreMockRecorder) FindAllByUser(userID, startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByUser", reflect.TypeOf((*MockStore)(nil).FindAllByUser), userID, startDate, endDate)
}

// FindAllDeduction mocks base method.
func (m *MockStore) FindAllDeduction() ([]models.Deduction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllDeduction")
	ret0, _ := ret[0].([]models.Deduction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllDeduction indicates an expected call of FindAllDeduction.
func (mr *MockStoreMockRecorder) FindAllDeduction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllDeduction", reflect.TypeOf((*MockStore)(nil).FindAllDeduction))
}

// FindAllGateway mocks base method.
func (m *MockStore) FindAllGateway() ([]models.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllGateway")
	ret0, _ := ret[0].([]models.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllGateway indicates an expected call of FindAllGateway.
func (mr *MockStoreMockRecorder) FindAllGateway() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllGateway", reflect.TypeOf((*MockStore)(nil).FindAllGateway))
}

// FindAllInvoiceItems mocks base method.
func (m *MockStore) FindAllInvoiceItems() ([]models.InvoiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllInvoiceItems")
	ret0, _ := ret[0].([]models.InvoiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllInvoiceItems indicates an expected call of FindAllInvoiceItems.
func (mr *MockStoreMockRecorder) FindAllInvoiceItems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllInvoiceItems", reflect.TypeOf((*MockStore)(nil).FindAllInvoiceItems))
}

// FindAllPaymentSetting mocks base method.
func (m *MockStore) FindAllPaymentSetting() ([]models.PaymentSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllPaymentSetting")
	ret0, _ := ret[0].([]models.PaymentSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllPaymentSetting indicates an expected call of FindAllPaymentSetting.
func (mr *MockStoreMockRecorder) FindAllPaymentSetting() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllPaymentSetting", reflect.TypeOf((*MockStore)(nil).FindAllPaymentSetting))
}

// FindAllPromocode mocks base method.
func (m *MockStore) FindAllPromocode() ([]models.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllPromocode")
	ret0, _ := ret[0].([]models.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllPromocode indicates an expected call of FindAllPromocode.
func (mr *MockStoreMockRecorder) FindAllPromocode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllPromocode", reflect.TypeOf((*MockStore)(nil).FindAllPromocode))
}

// FindAllThreshold mocks base method.
func (m *MockStore) FindAllThreshold() ([]models.PaymentThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllThreshold")
	ret0, _ := ret[0].([]models.PaymentThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllThreshold indicates an expected call of FindAllThreshold.
func (mr *MockStoreMockRecorder) FindAllThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllThreshold", reflect.TypeOf((*MockStore)(nil).FindAllThreshold))
}

// FindAllTransaction mocks base method.
func (m *MockStore) FindAllTransaction() ([]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTransaction")
	ret0, _ := ret[0].([]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTransaction indicates an expected call of FindAllTransaction.
func (mr *MockStoreMockRecorder) FindAllTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTransaction", reflect.TypeOf((*MockStore)(nil).FindAllTransaction))
}

// FindByCountryDeduction mocks base method.
func (m *MockStore) FindByCountryDeduction(country string) (models.Deduction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCountryDeduction", country)
	ret0, _ := ret[0].(models.Deduction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCountryDeduction indicates an expected call of FindByCountryDeduction.
func (mr *MockStoreMockRecorder) FindByCountryDeduction(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCountryDeduction", reflect.TypeOf((*MockStore)(nil).FindByCountryDeduction), country)
}

// FindByIdDeduction mocks base method.
func (m *MockStore) FindByIdDeduction(uid uint) (models.Deduction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdDeduction", uid)
	ret0, _ := ret[0].(models.Deduction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdDeduction indicates an expected call of FindByIdDeduction.
func (mr *MockStoreMockRecorder) FindByIdDeduction(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdDeduction", reflect.TypeOf((*MockStore)(nil).FindByIdDeduction), uid)
}

// FindByIdGateway mocks base method.
func (m *MockStore) FindByIdGateway(uid uint) (models.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdGateway", uid)
	ret0, _ := ret[0].(models.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdGateway indicates an expected call of FindByIdGateway.
func (mr *MockStoreMockRecorder) FindByIdGateway(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdGateway", reflect.TypeOf((*MockStore)(nil).FindByIdGateway), uid)
}

// FindByIdInvoice mocks base method.
func (m *MockStore) FindByIdInvoice(uid uint) (models.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdInvoice", uid)
	ret0, _ := ret[0].(models.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdInvoice indicates an expected call of FindByIdInvoice.
func (mr *MockStoreMockRecorder) FindByIdInvoice(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdInvoice", reflect.TypeOf((*MockStore)(nil).FindByIdInvoice), uid)
}

// FindByIdInvoiceItems mocks base method.
func (m *MockStore) FindByIdInvoiceItems(uid uint) (models.InvoiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdInvoiceItems", uid)
	ret0, _ := ret[0].(models.InvoiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdInvoiceItems indicates an expected call of FindByIdInvoiceItems.
func (mr *MockStoreMockRecorder) FindByIdInvoiceItems(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdInvoiceItems", reflect.TypeOf((*MockStore)(nil).FindByIdInvoiceItems), uid)
}

// FindByIdPaymentSetting mocks base method.
func (m *MockStore) FindByIdPaymentSetting(uid uint) (models.PaymentSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdPaymentSetting", uid)
	ret0, _ := ret[0].(models.PaymentSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdPaymentSetting indicates an expected call of FindByIdPaymentSetting.
func (mr *MockStoreMockRecorder) FindByIdPaymentSetting(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdPaymentSetting", reflect.TypeOf((*MockStore)(nil).FindByIdPaymentSetting), uid)
}

// FindByIdPromocode mocks base method.
func (m *MockStore) FindByIdPromocode(uid uint) (models.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdPromocode", uid)
	ret0, _ := ret[0].(models.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdPromocode indicates an expected call of FindByIdPromocode.
func (mr *MockStoreMockRecorder) FindByIdPromocode(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdPromocode", reflect.TypeOf((*MockStore)(nil).FindByIdPromocode), uid)
}

// FindByIdThreshold mocks base method.
func (m *MockStore) FindByIdThreshold(uid uint) (models.PaymentThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdThreshold", uid)
	ret0, _ := ret[0].(models.PaymentThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdThreshold indicates an expected call of FindByIdThreshold.
func (mr *MockStoreMockRecorder) FindByIdThreshold(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdThreshold", reflect.TypeOf((*MockStore)(nil).FindByIdThreshold), uid)
}

// FindByIdTransaction mocks base method.
func (m *MockStore) FindByIdTransaction(uid uint) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIdTransaction", uid)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIdTransaction indicates an expected call of FindByIdTransaction.
func (mr *MockStoreMockRecorder) FindByIdTransaction(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIdTransaction", reflect.TypeOf((*MockStore)(nil).FindByIdTransaction), uid)
}

// FindByPromoCode mocks base method.
func (m *MockStore) FindByPromoCode(promocode string) (models.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPromoCode", promocode)
	ret0, _ := ret[0].(models.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPromoCode indicates an expected call of FindByPromoCode.
func (mr *MockStoreMockRecorder) FindByPromoCode(promocode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPromoCode", reflect.TypeOf((*MockStore)(nil).FindByPromoCode), promocode)
}

// FindByUserIDPaymentHistory mocks base method.
func (m *MockStore) FindByUserIDPaymentHistory(uid uint) ([]models.PaymentHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIDPaymentHistory", uid)
	ret0, _ := ret[0].([]models.PaymentHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIDPaymentHistory indicates an expected call of FindByUserIDPaymentHistory.
func (mr *MockStoreMockRecorder) FindByUserIDPaymentHistory(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIDPaymentHistory", reflect.TypeOf((*MockStore)(nil).FindByUserIDPaymentHistory), uid)
}

// FindByUserIDPaymentSetting mocks base method.
func (m *MockStore) FindByUserIDPaymentSetting(uid uint) (models.PaymentSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIDPaymentSetting", uid)
	ret0, _ := ret[0].(models.PaymentSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIDPaymentSetting indicates an expected call of FindByUserIDPaymentSetting.
func (mr *MockStoreMockRecorder) FindByUserIDPaymentSetting(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIDPaymentSetting", reflect.TypeOf((*MockStore)(nil).FindByUserIDPaymentSetting), uid)
}

// FindByUserIDThreshold mocks base method.
func (m *MockStore) FindByUserIDThreshold(uid uint) (models.PaymentThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserIDThreshold", uid)
	ret0, _ := ret[0].(models.PaymentThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserIDThreshold indicates an expected call of FindByUserIDThreshold.
func (mr *MockStoreMockRecorder) FindByUserIDThreshold(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserIDThreshold", reflect.TypeOf((*MockStore)(nil).FindByUserIDThreshold), uid)
}

// FindSubscription mocks base method.
func (m *MockStore) FindSubscription(pid uint) (models.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSubscription", pid)
	ret0, _ := ret[0].(models.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSubscription indicates an expected call of FindSubscription.
func (mr *MockStoreMockRecorder) FindSubscription(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSubscription", reflect.TypeOf((*MockStore)(nil).FindSubscription), pid)
}

// MigrateDB mocks base method.
func (m *MockStore) MigrateDB() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MigrateDB")
}

// MigrateDB indicates an expected call of MigrateDB.
func (mr *MockStoreMockRecorder) MigrateDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateDB", reflect.TypeOf((*MockStore)(nil).MigrateDB))
}

// UpdateDeduction mocks base method.
func (m *MockStore) UpdateDeduction(data models.Deduction) (models.Deduction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeduction", data)
	ret0, _ := ret[0].(models.Deduction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDeduction indicates an expected call of UpdateDeduction.
func (mr *MockStoreMockRecorder) UpdateDeduction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeduction", reflect.TypeOf((*MockStore)(nil).UpdateDeduction), data)
}

// UpdateGateway mocks base method.
func (m *MockStore) UpdateGateway(data models.Gateway) (models.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGateway", data)
	ret0, _ := ret[0].(models.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGateway indicates an expected call of UpdateGateway.
func (mr *MockStoreMockRecorder) UpdateGateway(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGateway", reflect.TypeOf((*MockStore)(nil).UpdateGateway), data)
}

// UpdateInvoice mocks base method.
func (m *MockStore) UpdateInvoice(data models.Invoice) (models.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvoice", data)
	ret0, _ := ret[0].(models.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInvoice indicates an expected call of UpdateInvoice.
func (mr *MockStoreMockRecorder) UpdateInvoice(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvoice", reflect.TypeOf((*MockStore)(nil).UpdateInvoice), data)
}

// UpdateInvoiceItems mocks base method.
func (m *MockStore) UpdateInvoiceItems(data models.InvoiceItems) (models.InvoiceItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvoiceItems", data)
	ret0, _ := ret[0].(models.InvoiceItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInvoiceItems indicates an expected call of UpdateInvoiceItems.
func (mr *MockStoreMockRecorder) UpdateInvoiceItems(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvoiceItems", reflect.TypeOf((*MockStore)(nil).UpdateInvoiceItems), data)
}

// UpdatePaymentSetting mocks base method.
func (m *MockStore) UpdatePaymentSetting(data models.PaymentSetting) (models.PaymentSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentSetting", data)
	ret0, _ := ret[0].(models.PaymentSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentSetting indicates an expected call of UpdatePaymentSetting.
func (mr *MockStoreMockRecorder) UpdatePaymentSetting(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentSetting", reflect.TypeOf((*MockStore)(nil).UpdatePaymentSetting), data)
}

// UpdatePromocode mocks base method.
func (m *MockStore) UpdatePromocode(data models.PromoCode) (models.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePromocode", data)
	ret0, _ := ret[0].(models.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePromocode indicates an expected call of UpdatePromocode.
func (mr *MockStoreMockRecorder) UpdatePromocode(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePromocode", reflect.TypeOf((*MockStore)(nil).UpdatePromocode), data)
}

// UpdateThreshold mocks base method.
func (m *MockStore) UpdateThreshold(data models.PaymentThreshold) (models.PaymentThreshold, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThreshold", data)
	ret0, _ := ret[0].(models.PaymentThreshold)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateThreshold indicates an expected call of UpdateThreshold.
func (mr *MockStoreMockRecorder) UpdateThreshold(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThreshold", reflect.TypeOf((*MockStore)(nil).UpdateThreshold), data)
}

// UpdateTransaction mocks base method.
func (m *MockStore) UpdateTransaction(data models.Transaction) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", data)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockStoreMockRecorder) UpdateTransaction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockStore)(nil).UpdateTransaction), data)
}
