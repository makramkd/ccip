// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"

	merclib "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib"
	mock "github.com/stretchr/testify/mock"
)

// MercuryClient is an autogenerated mock type for the MercuryClient type
type MercuryClient struct {
	mock.Mock
}

// BatchFetchPrices provides a mock function with given fields: ctx, feedIDs
func (_m *MercuryClient) BatchFetchPrices(ctx context.Context, feedIDs [][32]byte) ([]*merclib.ReportWithContext, error) {
	ret := _m.Called(ctx, feedIDs)

	var r0 []*merclib.ReportWithContext
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][32]byte) ([]*merclib.ReportWithContext, error)); ok {
		return rf(ctx, feedIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][32]byte) []*merclib.ReportWithContext); ok {
		r0 = rf(ctx, feedIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*merclib.ReportWithContext)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][32]byte) error); ok {
		r1 = rf(ctx, feedIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMercuryClient creates a new instance of MercuryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMercuryClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MercuryClient {
	mock := &MercuryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
