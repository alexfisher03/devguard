// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GitLister is an autogenerated mock type for the gitLister type
type GitLister struct {
	mock.Mock
}

type GitLister_Expecter struct {
	mock *mock.Mock
}

func (_m *GitLister) EXPECT() *GitLister_Expecter {
	return &GitLister_Expecter{mock: &_m.Mock}
}

// GetBranchName provides a mock function with given fields: path
func (_m *GitLister) GetBranchName(path string) (string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for GetBranchName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GitLister_GetBranchName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBranchName'
type GitLister_GetBranchName_Call struct {
	*mock.Call
}

// GetBranchName is a helper method to define mock.On call
//   - path string
func (_e *GitLister_Expecter) GetBranchName(path interface{}) *GitLister_GetBranchName_Call {
	return &GitLister_GetBranchName_Call{Call: _e.mock.On("GetBranchName", path)}
}

func (_c *GitLister_GetBranchName_Call) Run(run func(path string)) *GitLister_GetBranchName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GitLister_GetBranchName_Call) Return(_a0 string, _a1 error) *GitLister_GetBranchName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GitLister_GetBranchName_Call) RunAndReturn(run func(string) (string, error)) *GitLister_GetBranchName_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultBranchName provides a mock function with given fields: path
func (_m *GitLister) GetDefaultBranchName(path string) (string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultBranchName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GitLister_GetDefaultBranchName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultBranchName'
type GitLister_GetDefaultBranchName_Call struct {
	*mock.Call
}

// GetDefaultBranchName is a helper method to define mock.On call
//   - path string
func (_e *GitLister_Expecter) GetDefaultBranchName(path interface{}) *GitLister_GetDefaultBranchName_Call {
	return &GitLister_GetDefaultBranchName_Call{Call: _e.mock.On("GetDefaultBranchName", path)}
}

func (_c *GitLister_GetDefaultBranchName_Call) Run(run func(path string)) *GitLister_GetDefaultBranchName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GitLister_GetDefaultBranchName_Call) Return(_a0 string, _a1 error) *GitLister_GetDefaultBranchName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GitLister_GetDefaultBranchName_Call) RunAndReturn(run func(string) (string, error)) *GitLister_GetDefaultBranchName_Call {
	_c.Call.Return(run)
	return _c
}

// GetTags provides a mock function with given fields: path
func (_m *GitLister) GetTags(path string) ([]string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GitLister_GetTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTags'
type GitLister_GetTags_Call struct {
	*mock.Call
}

// GetTags is a helper method to define mock.On call
//   - path string
func (_e *GitLister_Expecter) GetTags(path interface{}) *GitLister_GetTags_Call {
	return &GitLister_GetTags_Call{Call: _e.mock.On("GetTags", path)}
}

func (_c *GitLister_GetTags_Call) Run(run func(path string)) *GitLister_GetTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GitLister_GetTags_Call) Return(_a0 []string, _a1 error) *GitLister_GetTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GitLister_GetTags_Call) RunAndReturn(run func(string) ([]string, error)) *GitLister_GetTags_Call {
	_c.Call.Return(run)
	return _c
}

// GitCommitCount provides a mock function with given fields: path, tag
func (_m *GitLister) GitCommitCount(path string, tag *string) (int, error) {
	ret := _m.Called(path, tag)

	if len(ret) == 0 {
		panic("no return value specified for GitCommitCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string) (int, error)); ok {
		return rf(path, tag)
	}
	if rf, ok := ret.Get(0).(func(string, *string) int); ok {
		r0 = rf(path, tag)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, *string) error); ok {
		r1 = rf(path, tag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GitLister_GitCommitCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GitCommitCount'
type GitLister_GitCommitCount_Call struct {
	*mock.Call
}

// GitCommitCount is a helper method to define mock.On call
//   - path string
//   - tag *string
func (_e *GitLister_Expecter) GitCommitCount(path interface{}, tag interface{}) *GitLister_GitCommitCount_Call {
	return &GitLister_GitCommitCount_Call{Call: _e.mock.On("GitCommitCount", path, tag)}
}

func (_c *GitLister_GitCommitCount_Call) Run(run func(path string, tag *string)) *GitLister_GitCommitCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*string))
	})
	return _c
}

func (_c *GitLister_GitCommitCount_Call) Return(_a0 int, _a1 error) *GitLister_GitCommitCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GitLister_GitCommitCount_Call) RunAndReturn(run func(string, *string) (int, error)) *GitLister_GitCommitCount_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAllPathsAsSafe provides a mock function with no fields
func (_m *GitLister) MarkAllPathsAsSafe() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarkAllPathsAsSafe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GitLister_MarkAllPathsAsSafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAllPathsAsSafe'
type GitLister_MarkAllPathsAsSafe_Call struct {
	*mock.Call
}

// MarkAllPathsAsSafe is a helper method to define mock.On call
func (_e *GitLister_Expecter) MarkAllPathsAsSafe() *GitLister_MarkAllPathsAsSafe_Call {
	return &GitLister_MarkAllPathsAsSafe_Call{Call: _e.mock.On("MarkAllPathsAsSafe")}
}

func (_c *GitLister_MarkAllPathsAsSafe_Call) Run(run func()) *GitLister_MarkAllPathsAsSafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GitLister_MarkAllPathsAsSafe_Call) Return(_a0 error) *GitLister_MarkAllPathsAsSafe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GitLister_MarkAllPathsAsSafe_Call) RunAndReturn(run func() error) *GitLister_MarkAllPathsAsSafe_Call {
	_c.Call.Return(run)
	return _c
}

// NewGitLister creates a new instance of GitLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGitLister(t interface {
	mock.TestingT
	Cleanup(func())
}) *GitLister {
	mock := &GitLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
