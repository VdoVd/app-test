// Code generated by mockery. DO NOT EDIT.

package access

import (
	context "context"

	access "github.com/goravel/framework/contracts/auth/access"

	mock "github.com/stretchr/testify/mock"
)

// Gate is an autogenerated mock type for the Gate type
type Gate struct {
	mock.Mock
}

type Gate_Expecter struct {
	mock *mock.Mock
}

func (_m *Gate) EXPECT() *Gate_Expecter {
	return &Gate_Expecter{mock: &_m.Mock}
}

// After provides a mock function with given fields: callback
func (_m *Gate) After(callback func(context.Context, string, map[string]any, access.Response) access.Response) {
	_m.Called(callback)
}

// Gate_After_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'After'
type Gate_After_Call struct {
	*mock.Call
}

// After is a helper method to define mock.On call
//   - callback func(context.Context , string , map[string]any , access.Response) access.Response
func (_e *Gate_Expecter) After(callback interface{}) *Gate_After_Call {
	return &Gate_After_Call{Call: _e.mock.On("After", callback)}
}

func (_c *Gate_After_Call) Run(run func(callback func(context.Context, string, map[string]any, access.Response) access.Response)) *Gate_After_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(context.Context, string, map[string]any, access.Response) access.Response))
	})
	return _c
}

func (_c *Gate_After_Call) Return() *Gate_After_Call {
	_c.Call.Return()
	return _c
}

func (_c *Gate_After_Call) RunAndReturn(run func(func(context.Context, string, map[string]any, access.Response) access.Response)) *Gate_After_Call {
	_c.Call.Return(run)
	return _c
}

// Allows provides a mock function with given fields: ability, arguments
func (_m *Gate) Allows(ability string, arguments map[string]any) bool {
	ret := _m.Called(ability, arguments)

	if len(ret) == 0 {
		panic("no return value specified for Allows")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, map[string]any) bool); ok {
		r0 = rf(ability, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Gate_Allows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Allows'
type Gate_Allows_Call struct {
	*mock.Call
}

// Allows is a helper method to define mock.On call
//   - ability string
//   - arguments map[string]any
func (_e *Gate_Expecter) Allows(ability interface{}, arguments interface{}) *Gate_Allows_Call {
	return &Gate_Allows_Call{Call: _e.mock.On("Allows", ability, arguments)}
}

func (_c *Gate_Allows_Call) Run(run func(ability string, arguments map[string]any)) *Gate_Allows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(map[string]any))
	})
	return _c
}

func (_c *Gate_Allows_Call) Return(_a0 bool) *Gate_Allows_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_Allows_Call) RunAndReturn(run func(string, map[string]any) bool) *Gate_Allows_Call {
	_c.Call.Return(run)
	return _c
}

// Any provides a mock function with given fields: abilities, arguments
func (_m *Gate) Any(abilities []string, arguments map[string]any) bool {
	ret := _m.Called(abilities, arguments)

	if len(ret) == 0 {
		panic("no return value specified for Any")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, map[string]any) bool); ok {
		r0 = rf(abilities, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Gate_Any_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Any'
type Gate_Any_Call struct {
	*mock.Call
}

// Any is a helper method to define mock.On call
//   - abilities []string
//   - arguments map[string]any
func (_e *Gate_Expecter) Any(abilities interface{}, arguments interface{}) *Gate_Any_Call {
	return &Gate_Any_Call{Call: _e.mock.On("Any", abilities, arguments)}
}

func (_c *Gate_Any_Call) Run(run func(abilities []string, arguments map[string]any)) *Gate_Any_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].(map[string]any))
	})
	return _c
}

func (_c *Gate_Any_Call) Return(_a0 bool) *Gate_Any_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_Any_Call) RunAndReturn(run func([]string, map[string]any) bool) *Gate_Any_Call {
	_c.Call.Return(run)
	return _c
}

// Before provides a mock function with given fields: callback
func (_m *Gate) Before(callback func(context.Context, string, map[string]any) access.Response) {
	_m.Called(callback)
}

// Gate_Before_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Before'
type Gate_Before_Call struct {
	*mock.Call
}

// Before is a helper method to define mock.On call
//   - callback func(context.Context , string , map[string]any) access.Response
func (_e *Gate_Expecter) Before(callback interface{}) *Gate_Before_Call {
	return &Gate_Before_Call{Call: _e.mock.On("Before", callback)}
}

func (_c *Gate_Before_Call) Run(run func(callback func(context.Context, string, map[string]any) access.Response)) *Gate_Before_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(context.Context, string, map[string]any) access.Response))
	})
	return _c
}

func (_c *Gate_Before_Call) Return() *Gate_Before_Call {
	_c.Call.Return()
	return _c
}

func (_c *Gate_Before_Call) RunAndReturn(run func(func(context.Context, string, map[string]any) access.Response)) *Gate_Before_Call {
	_c.Call.Return(run)
	return _c
}

// Define provides a mock function with given fields: ability, callback
func (_m *Gate) Define(ability string, callback func(context.Context, map[string]any) access.Response) {
	_m.Called(ability, callback)
}

// Gate_Define_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Define'
type Gate_Define_Call struct {
	*mock.Call
}

// Define is a helper method to define mock.On call
//   - ability string
//   - callback func(context.Context , map[string]any) access.Response
func (_e *Gate_Expecter) Define(ability interface{}, callback interface{}) *Gate_Define_Call {
	return &Gate_Define_Call{Call: _e.mock.On("Define", ability, callback)}
}

func (_c *Gate_Define_Call) Run(run func(ability string, callback func(context.Context, map[string]any) access.Response)) *Gate_Define_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func(context.Context, map[string]any) access.Response))
	})
	return _c
}

func (_c *Gate_Define_Call) Return() *Gate_Define_Call {
	_c.Call.Return()
	return _c
}

func (_c *Gate_Define_Call) RunAndReturn(run func(string, func(context.Context, map[string]any) access.Response)) *Gate_Define_Call {
	_c.Call.Return(run)
	return _c
}

// Denies provides a mock function with given fields: ability, arguments
func (_m *Gate) Denies(ability string, arguments map[string]any) bool {
	ret := _m.Called(ability, arguments)

	if len(ret) == 0 {
		panic("no return value specified for Denies")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, map[string]any) bool); ok {
		r0 = rf(ability, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Gate_Denies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Denies'
type Gate_Denies_Call struct {
	*mock.Call
}

// Denies is a helper method to define mock.On call
//   - ability string
//   - arguments map[string]any
func (_e *Gate_Expecter) Denies(ability interface{}, arguments interface{}) *Gate_Denies_Call {
	return &Gate_Denies_Call{Call: _e.mock.On("Denies", ability, arguments)}
}

func (_c *Gate_Denies_Call) Run(run func(ability string, arguments map[string]any)) *Gate_Denies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(map[string]any))
	})
	return _c
}

func (_c *Gate_Denies_Call) Return(_a0 bool) *Gate_Denies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_Denies_Call) RunAndReturn(run func(string, map[string]any) bool) *Gate_Denies_Call {
	_c.Call.Return(run)
	return _c
}

// Inspect provides a mock function with given fields: ability, arguments
func (_m *Gate) Inspect(ability string, arguments map[string]any) access.Response {
	ret := _m.Called(ability, arguments)

	if len(ret) == 0 {
		panic("no return value specified for Inspect")
	}

	var r0 access.Response
	if rf, ok := ret.Get(0).(func(string, map[string]any) access.Response); ok {
		r0 = rf(ability, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Response)
		}
	}

	return r0
}

// Gate_Inspect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Inspect'
type Gate_Inspect_Call struct {
	*mock.Call
}

// Inspect is a helper method to define mock.On call
//   - ability string
//   - arguments map[string]any
func (_e *Gate_Expecter) Inspect(ability interface{}, arguments interface{}) *Gate_Inspect_Call {
	return &Gate_Inspect_Call{Call: _e.mock.On("Inspect", ability, arguments)}
}

func (_c *Gate_Inspect_Call) Run(run func(ability string, arguments map[string]any)) *Gate_Inspect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(map[string]any))
	})
	return _c
}

func (_c *Gate_Inspect_Call) Return(_a0 access.Response) *Gate_Inspect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_Inspect_Call) RunAndReturn(run func(string, map[string]any) access.Response) *Gate_Inspect_Call {
	_c.Call.Return(run)
	return _c
}

// None provides a mock function with given fields: abilities, arguments
func (_m *Gate) None(abilities []string, arguments map[string]any) bool {
	ret := _m.Called(abilities, arguments)

	if len(ret) == 0 {
		panic("no return value specified for None")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, map[string]any) bool); ok {
		r0 = rf(abilities, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Gate_None_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'None'
type Gate_None_Call struct {
	*mock.Call
}

// None is a helper method to define mock.On call
//   - abilities []string
//   - arguments map[string]any
func (_e *Gate_Expecter) None(abilities interface{}, arguments interface{}) *Gate_None_Call {
	return &Gate_None_Call{Call: _e.mock.On("None", abilities, arguments)}
}

func (_c *Gate_None_Call) Run(run func(abilities []string, arguments map[string]any)) *Gate_None_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].(map[string]any))
	})
	return _c
}

func (_c *Gate_None_Call) Return(_a0 bool) *Gate_None_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_None_Call) RunAndReturn(run func([]string, map[string]any) bool) *Gate_None_Call {
	_c.Call.Return(run)
	return _c
}

// WithContext provides a mock function with given fields: ctx
func (_m *Gate) WithContext(ctx context.Context) access.Gate {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WithContext")
	}

	var r0 access.Gate
	if rf, ok := ret.Get(0).(func(context.Context) access.Gate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Gate)
		}
	}

	return r0
}

// Gate_WithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithContext'
type Gate_WithContext_Call struct {
	*mock.Call
}

// WithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Gate_Expecter) WithContext(ctx interface{}) *Gate_WithContext_Call {
	return &Gate_WithContext_Call{Call: _e.mock.On("WithContext", ctx)}
}

func (_c *Gate_WithContext_Call) Run(run func(ctx context.Context)) *Gate_WithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Gate_WithContext_Call) Return(_a0 access.Gate) *Gate_WithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Gate_WithContext_Call) RunAndReturn(run func(context.Context) access.Gate) *Gate_WithContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewGate creates a new instance of Gate. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGate(t interface {
	mock.TestingT
	Cleanup(func())
}) *Gate {
	mock := &Gate{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}