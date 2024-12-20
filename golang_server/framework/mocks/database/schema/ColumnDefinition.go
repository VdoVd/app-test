// Code generated by mockery. DO NOT EDIT.

package schema

import (
	schema "github.com/goravel/framework/contracts/database/schema"
	mock "github.com/stretchr/testify/mock"
)

// ColumnDefinition is an autogenerated mock type for the ColumnDefinition type
type ColumnDefinition struct {
	mock.Mock
}

type ColumnDefinition_Expecter struct {
	mock *mock.Mock
}

func (_m *ColumnDefinition) EXPECT() *ColumnDefinition_Expecter {
	return &ColumnDefinition_Expecter{mock: &_m.Mock}
}

// AutoIncrement provides a mock function with given fields:
func (_m *ColumnDefinition) AutoIncrement() schema.ColumnDefinition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AutoIncrement")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func() schema.ColumnDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_AutoIncrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutoIncrement'
type ColumnDefinition_AutoIncrement_Call struct {
	*mock.Call
}

// AutoIncrement is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) AutoIncrement() *ColumnDefinition_AutoIncrement_Call {
	return &ColumnDefinition_AutoIncrement_Call{Call: _e.mock.On("AutoIncrement")}
}

func (_c *ColumnDefinition_AutoIncrement_Call) Run(run func()) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_AutoIncrement_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_AutoIncrement_Call) RunAndReturn(run func() schema.ColumnDefinition) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Return(run)
	return _c
}

// Comment provides a mock function with given fields: comment
func (_m *ColumnDefinition) Comment(comment string) schema.ColumnDefinition {
	ret := _m.Called(comment)

	if len(ret) == 0 {
		panic("no return value specified for Comment")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func(string) schema.ColumnDefinition); ok {
		r0 = rf(comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Comment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Comment'
type ColumnDefinition_Comment_Call struct {
	*mock.Call
}

// Comment is a helper method to define mock.On call
//   - comment string
func (_e *ColumnDefinition_Expecter) Comment(comment interface{}) *ColumnDefinition_Comment_Call {
	return &ColumnDefinition_Comment_Call{Call: _e.mock.On("Comment", comment)}
}

func (_c *ColumnDefinition_Comment_Call) Run(run func(comment string)) *ColumnDefinition_Comment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ColumnDefinition_Comment_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_Comment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Comment_Call) RunAndReturn(run func(string) schema.ColumnDefinition) *ColumnDefinition_Comment_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllowed provides a mock function with given fields:
func (_m *ColumnDefinition) GetAllowed() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllowed")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ColumnDefinition_GetAllowed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllowed'
type ColumnDefinition_GetAllowed_Call struct {
	*mock.Call
}

// GetAllowed is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetAllowed() *ColumnDefinition_GetAllowed_Call {
	return &ColumnDefinition_GetAllowed_Call{Call: _e.mock.On("GetAllowed")}
}

func (_c *ColumnDefinition_GetAllowed_Call) Run(run func()) *ColumnDefinition_GetAllowed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetAllowed_Call) Return(_a0 []string) *ColumnDefinition_GetAllowed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetAllowed_Call) RunAndReturn(run func() []string) *ColumnDefinition_GetAllowed_Call {
	_c.Call.Return(run)
	return _c
}

// GetAutoIncrement provides a mock function with given fields:
func (_m *ColumnDefinition) GetAutoIncrement() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAutoIncrement")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_GetAutoIncrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAutoIncrement'
type ColumnDefinition_GetAutoIncrement_Call struct {
	*mock.Call
}

// GetAutoIncrement is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetAutoIncrement() *ColumnDefinition_GetAutoIncrement_Call {
	return &ColumnDefinition_GetAutoIncrement_Call{Call: _e.mock.On("GetAutoIncrement")}
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) Run(run func()) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) Return(_a0 bool) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) RunAndReturn(run func() bool) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Return(run)
	return _c
}

// GetComment provides a mock function with given fields:
func (_m *ColumnDefinition) GetComment() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetComment")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnDefinition_GetComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComment'
type ColumnDefinition_GetComment_Call struct {
	*mock.Call
}

// GetComment is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetComment() *ColumnDefinition_GetComment_Call {
	return &ColumnDefinition_GetComment_Call{Call: _e.mock.On("GetComment")}
}

func (_c *ColumnDefinition_GetComment_Call) Run(run func()) *ColumnDefinition_GetComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetComment_Call) Return(comment string) *ColumnDefinition_GetComment_Call {
	_c.Call.Return(comment)
	return _c
}

func (_c *ColumnDefinition_GetComment_Call) RunAndReturn(run func() string) *ColumnDefinition_GetComment_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefault provides a mock function with given fields:
func (_m *ColumnDefinition) GetDefault() any {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDefault")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func() any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

// ColumnDefinition_GetDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefault'
type ColumnDefinition_GetDefault_Call struct {
	*mock.Call
}

// GetDefault is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetDefault() *ColumnDefinition_GetDefault_Call {
	return &ColumnDefinition_GetDefault_Call{Call: _e.mock.On("GetDefault")}
}

func (_c *ColumnDefinition_GetDefault_Call) Run(run func()) *ColumnDefinition_GetDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetDefault_Call) Return(_a0 any) *ColumnDefinition_GetDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetDefault_Call) RunAndReturn(run func() any) *ColumnDefinition_GetDefault_Call {
	_c.Call.Return(run)
	return _c
}

// GetLength provides a mock function with given fields:
func (_m *ColumnDefinition) GetLength() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLength")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ColumnDefinition_GetLength_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLength'
type ColumnDefinition_GetLength_Call struct {
	*mock.Call
}

// GetLength is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetLength() *ColumnDefinition_GetLength_Call {
	return &ColumnDefinition_GetLength_Call{Call: _e.mock.On("GetLength")}
}

func (_c *ColumnDefinition_GetLength_Call) Run(run func()) *ColumnDefinition_GetLength_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetLength_Call) Return(_a0 int) *ColumnDefinition_GetLength_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetLength_Call) RunAndReturn(run func() int) *ColumnDefinition_GetLength_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *ColumnDefinition) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnDefinition_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type ColumnDefinition_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetName() *ColumnDefinition_GetName_Call {
	return &ColumnDefinition_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *ColumnDefinition_GetName_Call) Run(run func()) *ColumnDefinition_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetName_Call) Return(_a0 string) *ColumnDefinition_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetName_Call) RunAndReturn(run func() string) *ColumnDefinition_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetNullable provides a mock function with given fields:
func (_m *ColumnDefinition) GetNullable() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNullable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_GetNullable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNullable'
type ColumnDefinition_GetNullable_Call struct {
	*mock.Call
}

// GetNullable is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetNullable() *ColumnDefinition_GetNullable_Call {
	return &ColumnDefinition_GetNullable_Call{Call: _e.mock.On("GetNullable")}
}

func (_c *ColumnDefinition_GetNullable_Call) Run(run func()) *ColumnDefinition_GetNullable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetNullable_Call) Return(_a0 bool) *ColumnDefinition_GetNullable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetNullable_Call) RunAndReturn(run func() bool) *ColumnDefinition_GetNullable_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlaces provides a mock function with given fields:
func (_m *ColumnDefinition) GetPlaces() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPlaces")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ColumnDefinition_GetPlaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlaces'
type ColumnDefinition_GetPlaces_Call struct {
	*mock.Call
}

// GetPlaces is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetPlaces() *ColumnDefinition_GetPlaces_Call {
	return &ColumnDefinition_GetPlaces_Call{Call: _e.mock.On("GetPlaces")}
}

func (_c *ColumnDefinition_GetPlaces_Call) Run(run func()) *ColumnDefinition_GetPlaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetPlaces_Call) Return(_a0 int) *ColumnDefinition_GetPlaces_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetPlaces_Call) RunAndReturn(run func() int) *ColumnDefinition_GetPlaces_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrecision provides a mock function with given fields:
func (_m *ColumnDefinition) GetPrecision() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPrecision")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ColumnDefinition_GetPrecision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrecision'
type ColumnDefinition_GetPrecision_Call struct {
	*mock.Call
}

// GetPrecision is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetPrecision() *ColumnDefinition_GetPrecision_Call {
	return &ColumnDefinition_GetPrecision_Call{Call: _e.mock.On("GetPrecision")}
}

func (_c *ColumnDefinition_GetPrecision_Call) Run(run func()) *ColumnDefinition_GetPrecision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetPrecision_Call) Return(_a0 int) *ColumnDefinition_GetPrecision_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetPrecision_Call) RunAndReturn(run func() int) *ColumnDefinition_GetPrecision_Call {
	_c.Call.Return(run)
	return _c
}

// GetTotal provides a mock function with given fields:
func (_m *ColumnDefinition) GetTotal() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotal")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ColumnDefinition_GetTotal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTotal'
type ColumnDefinition_GetTotal_Call struct {
	*mock.Call
}

// GetTotal is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetTotal() *ColumnDefinition_GetTotal_Call {
	return &ColumnDefinition_GetTotal_Call{Call: _e.mock.On("GetTotal")}
}

func (_c *ColumnDefinition_GetTotal_Call) Run(run func()) *ColumnDefinition_GetTotal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetTotal_Call) Return(_a0 int) *ColumnDefinition_GetTotal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetTotal_Call) RunAndReturn(run func() int) *ColumnDefinition_GetTotal_Call {
	_c.Call.Return(run)
	return _c
}

// GetType provides a mock function with given fields:
func (_m *ColumnDefinition) GetType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnDefinition_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type ColumnDefinition_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetType() *ColumnDefinition_GetType_Call {
	return &ColumnDefinition_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *ColumnDefinition_GetType_Call) Run(run func()) *ColumnDefinition_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetType_Call) Return(_a0 string) *ColumnDefinition_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetType_Call) RunAndReturn(run func() string) *ColumnDefinition_GetType_Call {
	_c.Call.Return(run)
	return _c
}

// IsSetComment provides a mock function with given fields:
func (_m *ColumnDefinition) IsSetComment() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSetComment")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_IsSetComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSetComment'
type ColumnDefinition_IsSetComment_Call struct {
	*mock.Call
}

// IsSetComment is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) IsSetComment() *ColumnDefinition_IsSetComment_Call {
	return &ColumnDefinition_IsSetComment_Call{Call: _e.mock.On("IsSetComment")}
}

func (_c *ColumnDefinition_IsSetComment_Call) Run(run func()) *ColumnDefinition_IsSetComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_IsSetComment_Call) Return(_a0 bool) *ColumnDefinition_IsSetComment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_IsSetComment_Call) RunAndReturn(run func() bool) *ColumnDefinition_IsSetComment_Call {
	_c.Call.Return(run)
	return _c
}

// Nullable provides a mock function with given fields:
func (_m *ColumnDefinition) Nullable() schema.ColumnDefinition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Nullable")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func() schema.ColumnDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Nullable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Nullable'
type ColumnDefinition_Nullable_Call struct {
	*mock.Call
}

// Nullable is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) Nullable() *ColumnDefinition_Nullable_Call {
	return &ColumnDefinition_Nullable_Call{Call: _e.mock.On("Nullable")}
}

func (_c *ColumnDefinition_Nullable_Call) Run(run func()) *ColumnDefinition_Nullable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_Nullable_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_Nullable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Nullable_Call) RunAndReturn(run func() schema.ColumnDefinition) *ColumnDefinition_Nullable_Call {
	_c.Call.Return(run)
	return _c
}

// Places provides a mock function with given fields: places
func (_m *ColumnDefinition) Places(places int) schema.ColumnDefinition {
	ret := _m.Called(places)

	if len(ret) == 0 {
		panic("no return value specified for Places")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func(int) schema.ColumnDefinition); ok {
		r0 = rf(places)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Places_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Places'
type ColumnDefinition_Places_Call struct {
	*mock.Call
}

// Places is a helper method to define mock.On call
//   - places int
func (_e *ColumnDefinition_Expecter) Places(places interface{}) *ColumnDefinition_Places_Call {
	return &ColumnDefinition_Places_Call{Call: _e.mock.On("Places", places)}
}

func (_c *ColumnDefinition_Places_Call) Run(run func(places int)) *ColumnDefinition_Places_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ColumnDefinition_Places_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_Places_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Places_Call) RunAndReturn(run func(int) schema.ColumnDefinition) *ColumnDefinition_Places_Call {
	_c.Call.Return(run)
	return _c
}

// Total provides a mock function with given fields: total
func (_m *ColumnDefinition) Total(total int) schema.ColumnDefinition {
	ret := _m.Called(total)

	if len(ret) == 0 {
		panic("no return value specified for Total")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func(int) schema.ColumnDefinition); ok {
		r0 = rf(total)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Total_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Total'
type ColumnDefinition_Total_Call struct {
	*mock.Call
}

// Total is a helper method to define mock.On call
//   - total int
func (_e *ColumnDefinition_Expecter) Total(total interface{}) *ColumnDefinition_Total_Call {
	return &ColumnDefinition_Total_Call{Call: _e.mock.On("Total", total)}
}

func (_c *ColumnDefinition_Total_Call) Run(run func(total int)) *ColumnDefinition_Total_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ColumnDefinition_Total_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_Total_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Total_Call) RunAndReturn(run func(int) schema.ColumnDefinition) *ColumnDefinition_Total_Call {
	_c.Call.Return(run)
	return _c
}

// Unsigned provides a mock function with given fields:
func (_m *ColumnDefinition) Unsigned() schema.ColumnDefinition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Unsigned")
	}

	var r0 schema.ColumnDefinition
	if rf, ok := ret.Get(0).(func() schema.ColumnDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Unsigned_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unsigned'
type ColumnDefinition_Unsigned_Call struct {
	*mock.Call
}

// Unsigned is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) Unsigned() *ColumnDefinition_Unsigned_Call {
	return &ColumnDefinition_Unsigned_Call{Call: _e.mock.On("Unsigned")}
}

func (_c *ColumnDefinition_Unsigned_Call) Run(run func()) *ColumnDefinition_Unsigned_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_Unsigned_Call) Return(_a0 schema.ColumnDefinition) *ColumnDefinition_Unsigned_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Unsigned_Call) RunAndReturn(run func() schema.ColumnDefinition) *ColumnDefinition_Unsigned_Call {
	_c.Call.Return(run)
	return _c
}

// NewColumnDefinition creates a new instance of ColumnDefinition. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewColumnDefinition(t interface {
	mock.TestingT
	Cleanup(func())
}) *ColumnDefinition {
	mock := &ColumnDefinition{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
