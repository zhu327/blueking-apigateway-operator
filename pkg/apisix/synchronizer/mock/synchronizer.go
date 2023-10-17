// Code generated by MockGen. DO NOT EDIT.
// Source: synchronizer.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	apisix "github.com/TencentBlueKing/blueking-apigateway-operator/pkg/apisix"
	gomock "github.com/golang/mock/gomock"
)

// MockApisixConfigSynchronizer is a mock of ApisixConfigSynchronizer interface.
type MockApisixConfigSynchronizer struct {
	ctrl     *gomock.Controller
	recorder *MockApisixConfigSynchronizerMockRecorder
}

// MockApisixConfigSynchronizerMockRecorder is the mock recorder for MockApisixConfigSynchronizer.
type MockApisixConfigSynchronizerMockRecorder struct {
	mock *MockApisixConfigSynchronizer
}

// NewMockApisixConfigSynchronizer creates a new mock instance.
func NewMockApisixConfigSynchronizer(ctrl *gomock.Controller) *MockApisixConfigSynchronizer {
	mock := &MockApisixConfigSynchronizer{ctrl: ctrl}
	mock.recorder = &MockApisixConfigSynchronizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApisixConfigSynchronizer) EXPECT() *MockApisixConfigSynchronizerMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockApisixConfigSynchronizer) Flush(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush", ctx)
}

// Flush indicates an expected call of Flush.
func (mr *MockApisixConfigSynchronizerMockRecorder) Flush(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockApisixConfigSynchronizer)(nil).Flush), ctx)
}

// RemoveNotExistStage mocks base method.
func (m *MockApisixConfigSynchronizer) RemoveNotExistStage(ctx context.Context, existStageKeys []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveNotExistStage", ctx, existStageKeys)
}

// RemoveNotExistStage indicates an expected call of RemoveNotExistStage.
func (mr *MockApisixConfigSynchronizerMockRecorder) RemoveNotExistStage(ctx, existStageKeys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNotExistStage", reflect.TypeOf((*MockApisixConfigSynchronizer)(nil).RemoveNotExistStage), ctx, existStageKeys)
}

// Sync mocks base method.
func (m *MockApisixConfigSynchronizer) Sync(ctx context.Context, gatewayName, stageName string, config *apisix.ApisixConfiguration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sync", ctx, gatewayName, stageName, config)
}

// Sync indicates an expected call of Sync.
func (mr *MockApisixConfigSynchronizerMockRecorder) Sync(ctx, gatewayName, stageName, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockApisixConfigSynchronizer)(nil).Sync), ctx, gatewayName, stageName, config)
}
