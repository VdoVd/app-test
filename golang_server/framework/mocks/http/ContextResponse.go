// Code generated by mockery. DO NOT EDIT.

package http

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"
)

// ContextResponse is an autogenerated mock type for the ContextResponse type
type ContextResponse struct {
	mock.Mock
}

type ContextResponse_Expecter struct {
	mock *mock.Mock
}

func (_m *ContextResponse) EXPECT() *ContextResponse_Expecter {
	return &ContextResponse_Expecter{mock: &_m.Mock}
}

// Cookie provides a mock function with given fields: cookie
func (_m *ContextResponse) Cookie(cookie http.Cookie) http.ContextResponse {
	ret := _m.Called(cookie)

	if len(ret) == 0 {
		panic("no return value specified for Cookie")
	}

	var r0 http.ContextResponse
	if rf, ok := ret.Get(0).(func(http.Cookie) http.ContextResponse); ok {
		r0 = rf(cookie)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ContextResponse)
		}
	}

	return r0
}

// ContextResponse_Cookie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cookie'
type ContextResponse_Cookie_Call struct {
	*mock.Call
}

// Cookie is a helper method to define mock.On call
//   - cookie http.Cookie
func (_e *ContextResponse_Expecter) Cookie(cookie interface{}) *ContextResponse_Cookie_Call {
	return &ContextResponse_Cookie_Call{Call: _e.mock.On("Cookie", cookie)}
}

func (_c *ContextResponse_Cookie_Call) Run(run func(cookie http.Cookie)) *ContextResponse_Cookie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Cookie))
	})
	return _c
}

func (_c *ContextResponse_Cookie_Call) Return(_a0 http.ContextResponse) *ContextResponse_Cookie_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Cookie_Call) RunAndReturn(run func(http.Cookie) http.ContextResponse) *ContextResponse_Cookie_Call {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields: code, contentType, data
func (_m *ContextResponse) Data(code int, contentType string, data []byte) http.Response {
	ret := _m.Called(code, contentType, data)

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(int, string, []byte) http.Response); ok {
		r0 = rf(code, contentType, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type ContextResponse_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
//   - code int
//   - contentType string
//   - data []byte
func (_e *ContextResponse_Expecter) Data(code interface{}, contentType interface{}, data interface{}) *ContextResponse_Data_Call {
	return &ContextResponse_Data_Call{Call: _e.mock.On("Data", code, contentType, data)}
}

func (_c *ContextResponse_Data_Call) Run(run func(code int, contentType string, data []byte)) *ContextResponse_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *ContextResponse_Data_Call) Return(_a0 http.Response) *ContextResponse_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Data_Call) RunAndReturn(run func(int, string, []byte) http.Response) *ContextResponse_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Download provides a mock function with given fields: filepath, filename
func (_m *ContextResponse) Download(filepath string, filename string) http.Response {
	ret := _m.Called(filepath, filename)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string, string) http.Response); ok {
		r0 = rf(filepath, filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type ContextResponse_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - filepath string
//   - filename string
func (_e *ContextResponse_Expecter) Download(filepath interface{}, filename interface{}) *ContextResponse_Download_Call {
	return &ContextResponse_Download_Call{Call: _e.mock.On("Download", filepath, filename)}
}

func (_c *ContextResponse_Download_Call) Run(run func(filepath string, filename string)) *ContextResponse_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ContextResponse_Download_Call) Return(_a0 http.Response) *ContextResponse_Download_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Download_Call) RunAndReturn(run func(string, string) http.Response) *ContextResponse_Download_Call {
	_c.Call.Return(run)
	return _c
}

// File provides a mock function with given fields: filepath
func (_m *ContextResponse) File(filepath string) http.Response {
	ret := _m.Called(filepath)

	if len(ret) == 0 {
		panic("no return value specified for File")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = rf(filepath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_File_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'File'
type ContextResponse_File_Call struct {
	*mock.Call
}

// File is a helper method to define mock.On call
//   - filepath string
func (_e *ContextResponse_Expecter) File(filepath interface{}) *ContextResponse_File_Call {
	return &ContextResponse_File_Call{Call: _e.mock.On("File", filepath)}
}

func (_c *ContextResponse_File_Call) Run(run func(filepath string)) *ContextResponse_File_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ContextResponse_File_Call) Return(_a0 http.Response) *ContextResponse_File_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_File_Call) RunAndReturn(run func(string) http.Response) *ContextResponse_File_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields:
func (_m *ContextResponse) Flush() {
	_m.Called()
}

// ContextResponse_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type ContextResponse_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
func (_e *ContextResponse_Expecter) Flush() *ContextResponse_Flush_Call {
	return &ContextResponse_Flush_Call{Call: _e.mock.On("Flush")}
}

func (_c *ContextResponse_Flush_Call) Run(run func()) *ContextResponse_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContextResponse_Flush_Call) Return() *ContextResponse_Flush_Call {
	_c.Call.Return()
	return _c
}

func (_c *ContextResponse_Flush_Call) RunAndReturn(run func()) *ContextResponse_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields: key, value
func (_m *ContextResponse) Header(key string, value string) http.ContextResponse {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 http.ContextResponse
	if rf, ok := ret.Get(0).(func(string, string) http.ContextResponse); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ContextResponse)
		}
	}

	return r0
}

// ContextResponse_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ContextResponse_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *ContextResponse_Expecter) Header(key interface{}, value interface{}) *ContextResponse_Header_Call {
	return &ContextResponse_Header_Call{Call: _e.mock.On("Header", key, value)}
}

func (_c *ContextResponse_Header_Call) Run(run func(key string, value string)) *ContextResponse_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ContextResponse_Header_Call) Return(_a0 http.ContextResponse) *ContextResponse_Header_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Header_Call) RunAndReturn(run func(string, string) http.ContextResponse) *ContextResponse_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Json provides a mock function with given fields: code, obj
func (_m *ContextResponse) Json(code int, obj any) http.Response {
	ret := _m.Called(code, obj)

	if len(ret) == 0 {
		panic("no return value specified for Json")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(int, any) http.Response); ok {
		r0 = rf(code, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_Json_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Json'
type ContextResponse_Json_Call struct {
	*mock.Call
}

// Json is a helper method to define mock.On call
//   - code int
//   - obj any
func (_e *ContextResponse_Expecter) Json(code interface{}, obj interface{}) *ContextResponse_Json_Call {
	return &ContextResponse_Json_Call{Call: _e.mock.On("Json", code, obj)}
}

func (_c *ContextResponse_Json_Call) Run(run func(code int, obj any)) *ContextResponse_Json_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(any))
	})
	return _c
}

func (_c *ContextResponse_Json_Call) Return(_a0 http.Response) *ContextResponse_Json_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Json_Call) RunAndReturn(run func(int, any) http.Response) *ContextResponse_Json_Call {
	_c.Call.Return(run)
	return _c
}

// NoContent provides a mock function with given fields: code
func (_m *ContextResponse) NoContent(code ...int) http.Response {
	_va := make([]interface{}, len(code))
	for _i := range code {
		_va[_i] = code[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for NoContent")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(...int) http.Response); ok {
		r0 = rf(code...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_NoContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoContent'
type ContextResponse_NoContent_Call struct {
	*mock.Call
}

// NoContent is a helper method to define mock.On call
//   - code ...int
func (_e *ContextResponse_Expecter) NoContent(code ...interface{}) *ContextResponse_NoContent_Call {
	return &ContextResponse_NoContent_Call{Call: _e.mock.On("NoContent",
		append([]interface{}{}, code...)...)}
}

func (_c *ContextResponse_NoContent_Call) Run(run func(code ...int)) *ContextResponse_NoContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ContextResponse_NoContent_Call) Return(_a0 http.Response) *ContextResponse_NoContent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_NoContent_Call) RunAndReturn(run func(...int) http.Response) *ContextResponse_NoContent_Call {
	_c.Call.Return(run)
	return _c
}

// Origin provides a mock function with given fields:
func (_m *ContextResponse) Origin() http.ResponseOrigin {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Origin")
	}

	var r0 http.ResponseOrigin
	if rf, ok := ret.Get(0).(func() http.ResponseOrigin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseOrigin)
		}
	}

	return r0
}

// ContextResponse_Origin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Origin'
type ContextResponse_Origin_Call struct {
	*mock.Call
}

// Origin is a helper method to define mock.On call
func (_e *ContextResponse_Expecter) Origin() *ContextResponse_Origin_Call {
	return &ContextResponse_Origin_Call{Call: _e.mock.On("Origin")}
}

func (_c *ContextResponse_Origin_Call) Run(run func()) *ContextResponse_Origin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContextResponse_Origin_Call) Return(_a0 http.ResponseOrigin) *ContextResponse_Origin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Origin_Call) RunAndReturn(run func() http.ResponseOrigin) *ContextResponse_Origin_Call {
	_c.Call.Return(run)
	return _c
}

// Redirect provides a mock function with given fields: code, location
func (_m *ContextResponse) Redirect(code int, location string) http.Response {
	ret := _m.Called(code, location)

	if len(ret) == 0 {
		panic("no return value specified for Redirect")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(int, string) http.Response); ok {
		r0 = rf(code, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_Redirect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Redirect'
type ContextResponse_Redirect_Call struct {
	*mock.Call
}

// Redirect is a helper method to define mock.On call
//   - code int
//   - location string
func (_e *ContextResponse_Expecter) Redirect(code interface{}, location interface{}) *ContextResponse_Redirect_Call {
	return &ContextResponse_Redirect_Call{Call: _e.mock.On("Redirect", code, location)}
}

func (_c *ContextResponse_Redirect_Call) Run(run func(code int, location string)) *ContextResponse_Redirect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string))
	})
	return _c
}

func (_c *ContextResponse_Redirect_Call) Return(_a0 http.Response) *ContextResponse_Redirect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Redirect_Call) RunAndReturn(run func(int, string) http.Response) *ContextResponse_Redirect_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: code
func (_m *ContextResponse) Status(code int) http.ResponseStatus {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 http.ResponseStatus
	if rf, ok := ret.Get(0).(func(int) http.ResponseStatus); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseStatus)
		}
	}

	return r0
}

// ContextResponse_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type ContextResponse_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - code int
func (_e *ContextResponse_Expecter) Status(code interface{}) *ContextResponse_Status_Call {
	return &ContextResponse_Status_Call{Call: _e.mock.On("Status", code)}
}

func (_c *ContextResponse_Status_Call) Run(run func(code int)) *ContextResponse_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ContextResponse_Status_Call) Return(_a0 http.ResponseStatus) *ContextResponse_Status_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Status_Call) RunAndReturn(run func(int) http.ResponseStatus) *ContextResponse_Status_Call {
	_c.Call.Return(run)
	return _c
}

// Stream provides a mock function with given fields: code, step
func (_m *ContextResponse) Stream(code int, step func(http.StreamWriter) error) http.Response {
	ret := _m.Called(code, step)

	if len(ret) == 0 {
		panic("no return value specified for Stream")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(int, func(http.StreamWriter) error) http.Response); ok {
		r0 = rf(code, step)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_Stream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stream'
type ContextResponse_Stream_Call struct {
	*mock.Call
}

// Stream is a helper method to define mock.On call
//   - code int
//   - step func(http.StreamWriter) error
func (_e *ContextResponse_Expecter) Stream(code interface{}, step interface{}) *ContextResponse_Stream_Call {
	return &ContextResponse_Stream_Call{Call: _e.mock.On("Stream", code, step)}
}

func (_c *ContextResponse_Stream_Call) Run(run func(code int, step func(http.StreamWriter) error)) *ContextResponse_Stream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(func(http.StreamWriter) error))
	})
	return _c
}

func (_c *ContextResponse_Stream_Call) Return(_a0 http.Response) *ContextResponse_Stream_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Stream_Call) RunAndReturn(run func(int, func(http.StreamWriter) error) http.Response) *ContextResponse_Stream_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields: code, format, values
func (_m *ContextResponse) String(code int, format string, values ...any) http.Response {
	var _ca []interface{}
	_ca = append(_ca, code, format)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(int, string, ...any) http.Response); ok {
		r0 = rf(code, format, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ContextResponse_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type ContextResponse_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
//   - code int
//   - format string
//   - values ...any
func (_e *ContextResponse_Expecter) String(code interface{}, format interface{}, values ...interface{}) *ContextResponse_String_Call {
	return &ContextResponse_String_Call{Call: _e.mock.On("String",
		append([]interface{}{code, format}, values...)...)}
}

func (_c *ContextResponse_String_Call) Run(run func(code int, format string, values ...any)) *ContextResponse_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]any, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		run(args[0].(int), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *ContextResponse_String_Call) Return(_a0 http.Response) *ContextResponse_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_String_Call) RunAndReturn(run func(int, string, ...any) http.Response) *ContextResponse_String_Call {
	_c.Call.Return(run)
	return _c
}

// Success provides a mock function with given fields:
func (_m *ContextResponse) Success() http.ResponseStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Success")
	}

	var r0 http.ResponseStatus
	if rf, ok := ret.Get(0).(func() http.ResponseStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseStatus)
		}
	}

	return r0
}

// ContextResponse_Success_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Success'
type ContextResponse_Success_Call struct {
	*mock.Call
}

// Success is a helper method to define mock.On call
func (_e *ContextResponse_Expecter) Success() *ContextResponse_Success_Call {
	return &ContextResponse_Success_Call{Call: _e.mock.On("Success")}
}

func (_c *ContextResponse_Success_Call) Run(run func()) *ContextResponse_Success_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContextResponse_Success_Call) Return(_a0 http.ResponseStatus) *ContextResponse_Success_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Success_Call) RunAndReturn(run func() http.ResponseStatus) *ContextResponse_Success_Call {
	_c.Call.Return(run)
	return _c
}

// View provides a mock function with given fields:
func (_m *ContextResponse) View() http.ResponseView {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for View")
	}

	var r0 http.ResponseView
	if rf, ok := ret.Get(0).(func() http.ResponseView); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseView)
		}
	}

	return r0
}

// ContextResponse_View_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'View'
type ContextResponse_View_Call struct {
	*mock.Call
}

// View is a helper method to define mock.On call
func (_e *ContextResponse_Expecter) View() *ContextResponse_View_Call {
	return &ContextResponse_View_Call{Call: _e.mock.On("View")}
}

func (_c *ContextResponse_View_Call) Run(run func()) *ContextResponse_View_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContextResponse_View_Call) Return(_a0 http.ResponseView) *ContextResponse_View_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_View_Call) RunAndReturn(run func() http.ResponseView) *ContextResponse_View_Call {
	_c.Call.Return(run)
	return _c
}

// WithoutCookie provides a mock function with given fields: name
func (_m *ContextResponse) WithoutCookie(name string) http.ContextResponse {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for WithoutCookie")
	}

	var r0 http.ContextResponse
	if rf, ok := ret.Get(0).(func(string) http.ContextResponse); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ContextResponse)
		}
	}

	return r0
}

// ContextResponse_WithoutCookie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithoutCookie'
type ContextResponse_WithoutCookie_Call struct {
	*mock.Call
}

// WithoutCookie is a helper method to define mock.On call
//   - name string
func (_e *ContextResponse_Expecter) WithoutCookie(name interface{}) *ContextResponse_WithoutCookie_Call {
	return &ContextResponse_WithoutCookie_Call{Call: _e.mock.On("WithoutCookie", name)}
}

func (_c *ContextResponse_WithoutCookie_Call) Run(run func(name string)) *ContextResponse_WithoutCookie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ContextResponse_WithoutCookie_Call) Return(_a0 http.ContextResponse) *ContextResponse_WithoutCookie_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_WithoutCookie_Call) RunAndReturn(run func(string) http.ContextResponse) *ContextResponse_WithoutCookie_Call {
	_c.Call.Return(run)
	return _c
}

// Writer provides a mock function with given fields:
func (_m *ContextResponse) Writer() nethttp.ResponseWriter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Writer")
	}

	var r0 nethttp.ResponseWriter
	if rf, ok := ret.Get(0).(func() nethttp.ResponseWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.ResponseWriter)
		}
	}

	return r0
}

// ContextResponse_Writer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Writer'
type ContextResponse_Writer_Call struct {
	*mock.Call
}

// Writer is a helper method to define mock.On call
func (_e *ContextResponse_Expecter) Writer() *ContextResponse_Writer_Call {
	return &ContextResponse_Writer_Call{Call: _e.mock.On("Writer")}
}

func (_c *ContextResponse_Writer_Call) Run(run func()) *ContextResponse_Writer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ContextResponse_Writer_Call) Return(_a0 nethttp.ResponseWriter) *ContextResponse_Writer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContextResponse_Writer_Call) RunAndReturn(run func() nethttp.ResponseWriter) *ContextResponse_Writer_Call {
	_c.Call.Return(run)
	return _c
}

// NewContextResponse creates a new instance of ContextResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContextResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContextResponse {
	mock := &ContextResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}