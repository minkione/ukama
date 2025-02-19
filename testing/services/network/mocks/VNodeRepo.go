// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import db "github.com/ukama/ukama/testing/services/network/internal/db"
import mock "github.com/stretchr/testify/mock"

// VNodeRepo is an autogenerated mock type for the VNodeRepo type
type VNodeRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: nodeId
func (_m *VNodeRepo) Delete(nodeId string) error {
	ret := _m.Called(nodeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(nodeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInfo provides a mock function with given fields: nodeId
func (_m *VNodeRepo) GetInfo(nodeId string) (*db.VNode, error) {
	ret := _m.Called(nodeId)

	var r0 *db.VNode
	if rf, ok := ret.Get(0).(func(string) *db.VNode); ok {
		r0 = rf(nodeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.VNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nodeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: nodeId, status
func (_m *VNodeRepo) Insert(nodeId string, status string) error {
	ret := _m.Called(nodeId, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(nodeId, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *VNodeRepo) List() (*[]db.VNode, error) {
	ret := _m.Called()

	var r0 *[]db.VNode
	if rf, ok := ret.Get(0).(func() *[]db.VNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]db.VNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PowerOff provides a mock function with given fields: nodeId
func (_m *VNodeRepo) PowerOff(nodeId string) error {
	ret := _m.Called(nodeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(nodeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PowerOn provides a mock function with given fields: nodeId
func (_m *VNodeRepo) PowerOn(nodeId string) error {
	ret := _m.Called(nodeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(nodeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: nodeId, status
func (_m *VNodeRepo) Update(nodeId string, status string) error {
	ret := _m.Called(nodeId, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(nodeId, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
