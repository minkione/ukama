// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gen "github.com/ukama/ukama/systems/data-plan/package/pb"
)

// PackagesServiceServer is an autogenerated mock type for the PackagesServiceServer type
type PackagesServiceServer struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *PackagesServiceServer) Add(_a0 context.Context, _a1 *gen.AddPackageRequest) (*gen.AddPackageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.AddPackageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.AddPackageRequest) *gen.AddPackageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.AddPackageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.AddPackageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *PackagesServiceServer) Delete(_a0 context.Context, _a1 *gen.DeletePackageRequest) (*gen.DeletePackageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.DeletePackageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeletePackageRequest) *gen.DeletePackageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.DeletePackageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.DeletePackageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *PackagesServiceServer) Get(_a0 context.Context, _a1 *gen.GetPackageRequest) (*gen.GetPackageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetPackageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetPackageRequest) *gen.GetPackageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetPackageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetPackageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByOrg provides a mock function with given fields: _a0, _a1
func (_m *PackagesServiceServer) GetByOrg(_a0 context.Context, _a1 *gen.GetByOrgPackageRequest) (*gen.GetByOrgPackageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetByOrgPackageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetByOrgPackageRequest) *gen.GetByOrgPackageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetByOrgPackageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetByOrgPackageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *PackagesServiceServer) Update(_a0 context.Context, _a1 *gen.UpdatePackageRequest) (*gen.UpdatePackageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpdatePackageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpdatePackageRequest) *gen.UpdatePackageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpdatePackageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpdatePackageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedPackagesServiceServer provides a mock function with given fields:
func (_m *PackagesServiceServer) mustEmbedUnimplementedPackagesServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewPackagesServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPackagesServiceServer creates a new instance of PackagesServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPackagesServiceServer(t mockConstructorTestingTNewPackagesServiceServer) *PackagesServiceServer {
	mock := &PackagesServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
