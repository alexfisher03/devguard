// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	accesscontrol "github.com/l3montree-dev/devguard/internal/accesscontrol"
	mock "github.com/stretchr/testify/mock"
)

// RBACProvider is an autogenerated mock type for the RBACProvider type
type RBACProvider struct {
	mock.Mock
}

type RBACProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *RBACProvider) EXPECT() *RBACProvider_Expecter {
	return &RBACProvider_Expecter{mock: &_m.Mock}
}

// DomainsOfUser provides a mock function with given fields: user
func (_m *RBACProvider) DomainsOfUser(user string) ([]string, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for DomainsOfUser")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RBACProvider_DomainsOfUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DomainsOfUser'
type RBACProvider_DomainsOfUser_Call struct {
	*mock.Call
}

// DomainsOfUser is a helper method to define mock.On call
//   - user string
func (_e *RBACProvider_Expecter) DomainsOfUser(user interface{}) *RBACProvider_DomainsOfUser_Call {
	return &RBACProvider_DomainsOfUser_Call{Call: _e.mock.On("DomainsOfUser", user)}
}

func (_c *RBACProvider_DomainsOfUser_Call) Run(run func(user string)) *RBACProvider_DomainsOfUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RBACProvider_DomainsOfUser_Call) Return(_a0 []string, _a1 error) *RBACProvider_DomainsOfUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RBACProvider_DomainsOfUser_Call) RunAndReturn(run func(string) ([]string, error)) *RBACProvider_DomainsOfUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetDomainRBAC provides a mock function with given fields: domain
func (_m *RBACProvider) GetDomainRBAC(domain string) accesscontrol.AccessControl {
	ret := _m.Called(domain)

	if len(ret) == 0 {
		panic("no return value specified for GetDomainRBAC")
	}

	var r0 accesscontrol.AccessControl
	if rf, ok := ret.Get(0).(func(string) accesscontrol.AccessControl); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(accesscontrol.AccessControl)
		}
	}

	return r0
}

// RBACProvider_GetDomainRBAC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDomainRBAC'
type RBACProvider_GetDomainRBAC_Call struct {
	*mock.Call
}

// GetDomainRBAC is a helper method to define mock.On call
//   - domain string
func (_e *RBACProvider_Expecter) GetDomainRBAC(domain interface{}) *RBACProvider_GetDomainRBAC_Call {
	return &RBACProvider_GetDomainRBAC_Call{Call: _e.mock.On("GetDomainRBAC", domain)}
}

func (_c *RBACProvider_GetDomainRBAC_Call) Run(run func(domain string)) *RBACProvider_GetDomainRBAC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RBACProvider_GetDomainRBAC_Call) Return(_a0 accesscontrol.AccessControl) *RBACProvider_GetDomainRBAC_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RBACProvider_GetDomainRBAC_Call) RunAndReturn(run func(string) accesscontrol.AccessControl) *RBACProvider_GetDomainRBAC_Call {
	_c.Call.Return(run)
	return _c
}

// NewRBACProvider creates a new instance of RBACProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRBACProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *RBACProvider {
	mock := &RBACProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
