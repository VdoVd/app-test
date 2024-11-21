// Code generated by mockery. DO NOT EDIT.

package database

import (
	database "github.com/goravel/framework/contracts/database"
	mock "github.com/stretchr/testify/mock"
)

// ConfigBuilder is an autogenerated mock type for the ConfigBuilder type
type ConfigBuilder struct {
	mock.Mock
}

type ConfigBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *ConfigBuilder) EXPECT() *ConfigBuilder_Expecter {
	return &ConfigBuilder_Expecter{mock: &_m.Mock}
}

// Reads provides a mock function with given fields:
func (_m *ConfigBuilder) Reads() []database.FullConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Reads")
	}

	var r0 []database.FullConfig
	if rf, ok := ret.Get(0).(func() []database.FullConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.FullConfig)
		}
	}

	return r0
}

// ConfigBuilder_Reads_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reads'
type ConfigBuilder_Reads_Call struct {
	*mock.Call
}

// Reads is a helper method to define mock.On call
func (_e *ConfigBuilder_Expecter) Reads() *ConfigBuilder_Reads_Call {
	return &ConfigBuilder_Reads_Call{Call: _e.mock.On("Reads")}
}

func (_c *ConfigBuilder_Reads_Call) Run(run func()) *ConfigBuilder_Reads_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigBuilder_Reads_Call) Return(_a0 []database.FullConfig) *ConfigBuilder_Reads_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigBuilder_Reads_Call) RunAndReturn(run func() []database.FullConfig) *ConfigBuilder_Reads_Call {
	_c.Call.Return(run)
	return _c
}

// Writes provides a mock function with given fields:
func (_m *ConfigBuilder) Writes() []database.FullConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Writes")
	}

	var r0 []database.FullConfig
	if rf, ok := ret.Get(0).(func() []database.FullConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.FullConfig)
		}
	}

	return r0
}

// ConfigBuilder_Writes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Writes'
type ConfigBuilder_Writes_Call struct {
	*mock.Call
}

// Writes is a helper method to define mock.On call
func (_e *ConfigBuilder_Expecter) Writes() *ConfigBuilder_Writes_Call {
	return &ConfigBuilder_Writes_Call{Call: _e.mock.On("Writes")}
}

func (_c *ConfigBuilder_Writes_Call) Run(run func()) *ConfigBuilder_Writes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConfigBuilder_Writes_Call) Return(_a0 []database.FullConfig) *ConfigBuilder_Writes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConfigBuilder_Writes_Call) RunAndReturn(run func() []database.FullConfig) *ConfigBuilder_Writes_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfigBuilder creates a new instance of ConfigBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfigBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfigBuilder {
	mock := &ConfigBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
