// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeMsgClientServiceServer is an autogenerated mock type for the UnsafeMsgClientServiceServer type
type UnsafeMsgClientServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedMsgClientServiceServer provides a mock function with given fields:
func (_m *UnsafeMsgClientServiceServer) mustEmbedUnimplementedMsgClientServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeMsgClientServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeMsgClientServiceServer creates a new instance of UnsafeMsgClientServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeMsgClientServiceServer(t mockConstructorTestingTNewUnsafeMsgClientServiceServer) *UnsafeMsgClientServiceServer {
	mock := &UnsafeMsgClientServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
