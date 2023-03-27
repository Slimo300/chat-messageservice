// Code generated by mockery v2.20.0. DO NOT EDIT.

package storage

import mock "github.com/stretchr/testify/mock"

// MockStorage is an autogenerated mock type for the StorageLayer type
type MockStorage struct {
	mock.Mock
}

// DeleteFile provides a mock function with given fields: key
func (_m *MockStorage) DeleteFile(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFolder provides a mock function with given fields: folder
func (_m *MockStorage) DeleteFolder(folder string) error {
	ret := _m.Called(folder)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPresignedPutRequest provides a mock function with given fields: key
func (_m *MockStorage) GetPresignedPutRequest(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStorage(t mockConstructorTestingTNewMockStorage) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}