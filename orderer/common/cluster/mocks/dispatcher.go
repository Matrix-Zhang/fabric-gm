// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import orderer "github.com/Matrix-Zhang/fabric-gm/protos/orderer"

// Dispatcher is an autogenerated mock type for the Dispatcher type
type Dispatcher struct {
	mock.Mock
}

// DispatchConsensus provides a mock function with given fields: ctx, request
func (_m *Dispatcher) DispatchConsensus(ctx context.Context, request *orderer.ConsensusRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *orderer.ConsensusRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DispatchSubmit provides a mock function with given fields: ctx, request
func (_m *Dispatcher) DispatchSubmit(ctx context.Context, request *orderer.SubmitRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *orderer.SubmitRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
