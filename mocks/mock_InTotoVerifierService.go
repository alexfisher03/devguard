// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InTotoVerifierService is an autogenerated mock type for the InTotoVerifierService type
type InTotoVerifierService struct {
	mock.Mock
}

type InTotoVerifierService_Expecter struct {
	mock *mock.Mock
}

func (_m *InTotoVerifierService) EXPECT() *InTotoVerifierService_Expecter {
	return &InTotoVerifierService_Expecter{mock: &_m.Mock}
}

// VerifySupplyChain provides a mock function with given fields: supplyChainID
func (_m *InTotoVerifierService) VerifySupplyChain(supplyChainID string) (bool, error) {
	ret := _m.Called(supplyChainID)

	if len(ret) == 0 {
		panic("no return value specified for VerifySupplyChain")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(supplyChainID)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(supplyChainID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(supplyChainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InTotoVerifierService_VerifySupplyChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifySupplyChain'
type InTotoVerifierService_VerifySupplyChain_Call struct {
	*mock.Call
}

// VerifySupplyChain is a helper method to define mock.On call
//   - supplyChainID string
func (_e *InTotoVerifierService_Expecter) VerifySupplyChain(supplyChainID interface{}) *InTotoVerifierService_VerifySupplyChain_Call {
	return &InTotoVerifierService_VerifySupplyChain_Call{Call: _e.mock.On("VerifySupplyChain", supplyChainID)}
}

func (_c *InTotoVerifierService_VerifySupplyChain_Call) Run(run func(supplyChainID string)) *InTotoVerifierService_VerifySupplyChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChain_Call) Return(_a0 bool, _a1 error) *InTotoVerifierService_VerifySupplyChain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChain_Call) RunAndReturn(run func(string) (bool, error)) *InTotoVerifierService_VerifySupplyChain_Call {
	_c.Call.Return(run)
	return _c
}

// VerifySupplyChainByDigestOnly provides a mock function with given fields: digest
func (_m *InTotoVerifierService) VerifySupplyChainByDigestOnly(digest string) (bool, error) {
	ret := _m.Called(digest)

	if len(ret) == 0 {
		panic("no return value specified for VerifySupplyChainByDigestOnly")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(digest)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(digest)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InTotoVerifierService_VerifySupplyChainByDigestOnly_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifySupplyChainByDigestOnly'
type InTotoVerifierService_VerifySupplyChainByDigestOnly_Call struct {
	*mock.Call
}

// VerifySupplyChainByDigestOnly is a helper method to define mock.On call
//   - digest string
func (_e *InTotoVerifierService_Expecter) VerifySupplyChainByDigestOnly(digest interface{}) *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call {
	return &InTotoVerifierService_VerifySupplyChainByDigestOnly_Call{Call: _e.mock.On("VerifySupplyChainByDigestOnly", digest)}
}

func (_c *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call) Run(run func(digest string)) *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call) Return(_a0 bool, _a1 error) *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call) RunAndReturn(run func(string) (bool, error)) *InTotoVerifierService_VerifySupplyChainByDigestOnly_Call {
	_c.Call.Return(run)
	return _c
}

// VerifySupplyChainWithOutputDigest provides a mock function with given fields: supplyChainID, digest
func (_m *InTotoVerifierService) VerifySupplyChainWithOutputDigest(supplyChainID string, digest string) (bool, error) {
	ret := _m.Called(supplyChainID, digest)

	if len(ret) == 0 {
		panic("no return value specified for VerifySupplyChainWithOutputDigest")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(supplyChainID, digest)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(supplyChainID, digest)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(supplyChainID, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifySupplyChainWithOutputDigest'
type InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call struct {
	*mock.Call
}

// VerifySupplyChainWithOutputDigest is a helper method to define mock.On call
//   - supplyChainID string
//   - digest string
func (_e *InTotoVerifierService_Expecter) VerifySupplyChainWithOutputDigest(supplyChainID interface{}, digest interface{}) *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call {
	return &InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call{Call: _e.mock.On("VerifySupplyChainWithOutputDigest", supplyChainID, digest)}
}

func (_c *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call) Run(run func(supplyChainID string, digest string)) *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call) Return(_a0 bool, _a1 error) *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call) RunAndReturn(run func(string, string) (bool, error)) *InTotoVerifierService_VerifySupplyChainWithOutputDigest_Call {
	_c.Call.Return(run)
	return _c
}

// NewInTotoVerifierService creates a new instance of InTotoVerifierService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInTotoVerifierService(t interface {
	mock.TestingT
	Cleanup(func())
}) *InTotoVerifierService {
	mock := &InTotoVerifierService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
