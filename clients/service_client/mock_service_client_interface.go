// Code generated by MockGen. DO NOT EDIT.
// Source: nacos-go/clients/service_client (interfaces: IServiceClient)

// Package service_client is a generated GoMock package.
package service_client

import (
	gomock "github.com/golang/mock/gomock"
	vo "nacos-go/vo"
	reflect "reflect"
)

// MockIServiceClient is a mock of IServiceClient interface
type MockIServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceClientMockRecorder
}

// MockIServiceClientMockRecorder is the mock recorder for MockIServiceClient
type MockIServiceClientMockRecorder struct {
	mock *MockIServiceClient
}

// NewMockIServiceClient creates a new mock instance
func NewMockIServiceClient(ctrl *gomock.Controller) *MockIServiceClient {
	mock := &MockIServiceClient{ctrl: ctrl}
	mock.recorder = &MockIServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIServiceClient) EXPECT() *MockIServiceClientMockRecorder {
	return m.recorder
}

// GetService mocks base method
func (m *MockIServiceClient) GetService(arg0 vo.GetServiceParam) (vo.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(vo.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockIServiceClientMockRecorder) GetService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockIServiceClient)(nil).GetService), arg0)
}

// GetServiceDetail mocks base method
func (m *MockIServiceClient) GetServiceDetail(arg0 vo.GetServiceDetailParam) (vo.ServiceDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceDetail", arg0)
	ret0, _ := ret[0].(vo.ServiceDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceDetail indicates an expected call of GetServiceDetail
func (mr *MockIServiceClientMockRecorder) GetServiceDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceDetail", reflect.TypeOf((*MockIServiceClient)(nil).GetServiceDetail), arg0)
}

// GetServiceInstance mocks base method
func (m *MockIServiceClient) GetServiceInstance(arg0 vo.GetServiceInstanceParam) (vo.ServiceInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceInstance", arg0)
	ret0, _ := ret[0].(vo.ServiceInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceInstance indicates an expected call of GetServiceInstance
func (mr *MockIServiceClientMockRecorder) GetServiceInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceInstance", reflect.TypeOf((*MockIServiceClient)(nil).GetServiceInstance), arg0)
}

// LogoutServiceInstance mocks base method
func (m *MockIServiceClient) LogoutServiceInstance(arg0 vo.LogoutServiceInstanceParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogoutServiceInstance", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogoutServiceInstance indicates an expected call of LogoutServiceInstance
func (mr *MockIServiceClientMockRecorder) LogoutServiceInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutServiceInstance", reflect.TypeOf((*MockIServiceClient)(nil).LogoutServiceInstance), arg0)
}

// ModifyServiceInstance mocks base method
func (m *MockIServiceClient) ModifyServiceInstance(arg0 vo.ModifyServiceInstanceParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyServiceInstance", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyServiceInstance indicates an expected call of ModifyServiceInstance
func (mr *MockIServiceClientMockRecorder) ModifyServiceInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyServiceInstance", reflect.TypeOf((*MockIServiceClient)(nil).ModifyServiceInstance), arg0)
}

// RegisterServiceInstance mocks base method
func (m *MockIServiceClient) RegisterServiceInstance(arg0 vo.RegisterServiceInstanceParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterServiceInstance", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterServiceInstance indicates an expected call of RegisterServiceInstance
func (mr *MockIServiceClientMockRecorder) RegisterServiceInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServiceInstance", reflect.TypeOf((*MockIServiceClient)(nil).RegisterServiceInstance), arg0)
}

// StartBeatTask mocks base method
func (m *MockIServiceClient) StartBeatTask(arg0 vo.BeatTaskParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartBeatTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartBeatTask indicates an expected call of StartBeatTask
func (mr *MockIServiceClientMockRecorder) StartBeatTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartBeatTask", reflect.TypeOf((*MockIServiceClient)(nil).StartBeatTask), arg0)
}

// StopBeatTask mocks base method
func (m *MockIServiceClient) StopBeatTask() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopBeatTask")
}

// StopBeatTask indicates an expected call of StopBeatTask
func (mr *MockIServiceClientMockRecorder) StopBeatTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopBeatTask", reflect.TypeOf((*MockIServiceClient)(nil).StopBeatTask))
}
