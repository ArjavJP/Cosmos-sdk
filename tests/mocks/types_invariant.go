// Code generated by MockGen. DO NOT EDIT.
// Source: types/invariant.go

// Package mocks is a generated GoMock package.
package mocks

import (
	types "github.com/ArjavJP/Cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInvariantRegistry is a mock of InvariantRegistry interface
type MockInvariantRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockInvariantRegistryMockRecorder
}

// MockInvariantRegistryMockRecorder is the mock recorder for MockInvariantRegistry
type MockInvariantRegistryMockRecorder struct {
	mock *MockInvariantRegistry
}

// NewMockInvariantRegistry creates a new mock instance
func NewMockInvariantRegistry(ctrl *gomock.Controller) *MockInvariantRegistry {
	mock := &MockInvariantRegistry{ctrl: ctrl}
	mock.recorder = &MockInvariantRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvariantRegistry) EXPECT() *MockInvariantRegistryMockRecorder {
	return m.recorder
}

// RegisterRoute mocks base method
func (m *MockInvariantRegistry) RegisterRoute(moduleName, route string, invar types.Invariant) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRoute", moduleName, route, invar)
}

// RegisterRoute indicates an expected call of RegisterRoute
func (mr *MockInvariantRegistryMockRecorder) RegisterRoute(moduleName, route, invar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRoute", reflect.TypeOf((*MockInvariantRegistry)(nil).RegisterRoute), moduleName, route, invar)
}
