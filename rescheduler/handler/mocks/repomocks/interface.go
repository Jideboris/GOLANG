// Code generated by mockery v2.38.0. DO NOT EDIT.

package repomocks

import mock "github.com/stretchr/testify/mock"

// ComputeHandler is an autogenerated mock type for the ComputeHandler type
type ComputeHandler struct {
	mock.Mock
}

// GenerateTable provides a mock function with given fields:
func (_m *ComputeHandler) GenerateTable() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateTable")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateTableData provides a mock function with given fields:
func (_m *ComputeHandler) GenerateTableData() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateTableData")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewComputeHandler creates a new instance of ComputeHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComputeHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *ComputeHandler {
	mock := &ComputeHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
