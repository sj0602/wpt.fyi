// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/shared (interfaces: MetadataFetcher)

// Package sharedtest is a generated GoMock package.
package sharedtest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetadataFetcher is a mock of MetadataFetcher interface.
type MockMetadataFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataFetcherMockRecorder
}

// MockMetadataFetcherMockRecorder is the mock recorder for MockMetadataFetcher.
type MockMetadataFetcherMockRecorder struct {
	mock *MockMetadataFetcher
}

// NewMockMetadataFetcher creates a new mock instance.
func NewMockMetadataFetcher(ctrl *gomock.Controller) *MockMetadataFetcher {
	mock := &MockMetadataFetcher{ctrl: ctrl}
	mock.recorder = &MockMetadataFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataFetcher) EXPECT() *MockMetadataFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockMetadataFetcher) Fetch() (*string, map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(map[string][]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Fetch indicates an expected call of Fetch.
func (mr *MockMetadataFetcherMockRecorder) Fetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockMetadataFetcher)(nil).Fetch))
}
