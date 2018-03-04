// Code generated by mockery v1.0.0
package mocks

import entities "github.com/vmware/dispatch/pkg/event-manager/drivers/entities"
import mock "github.com/stretchr/testify/mock"

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *Backend) Delete(_a0 *entities.Driver) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Driver) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deploy provides a mock function with given fields: _a0
func (_m *Backend) Deploy(_a0 *entities.Driver) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Driver) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *Backend) Update(_a0 *entities.Driver) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Driver) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}