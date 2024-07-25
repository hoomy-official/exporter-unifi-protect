// Code generated by mockery. DO NOT EDIT.

package rest

import (
	context "context"
	http "net/http"

	do "github.com/hoomy-official/go-shared/pkg/net/do"

	mock "github.com/stretchr/testify/mock"
)

// MockRequester is an autogenerated mock type for the Requester type
type MockRequester struct {
	mock.Mock
}

type MockRequester_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequester) EXPECT() *MockRequester_Expecter {
	return &MockRequester_Expecter{mock: &_m.Mock}
}

// CONNECT provides a mock function with given fields: ctx, options
func (_m *MockRequester) CONNECT(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CONNECT")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_CONNECT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CONNECT'
type MockRequester_CONNECT_Call struct {
	*mock.Call
}

// CONNECT is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) CONNECT(ctx interface{}, options ...interface{}) *MockRequester_CONNECT_Call {
	return &MockRequester_CONNECT_Call{Call: _e.mock.On("CONNECT",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_CONNECT_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_CONNECT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_CONNECT_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_CONNECT_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_CONNECT_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_CONNECT_Call {
	_c.Call.Return(run)
	return _c
}

// DELETE provides a mock function with given fields: ctx, options
func (_m *MockRequester) DELETE(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DELETE")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_DELETE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DELETE'
type MockRequester_DELETE_Call struct {
	*mock.Call
}

// DELETE is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) DELETE(ctx interface{}, options ...interface{}) *MockRequester_DELETE_Call {
	return &MockRequester_DELETE_Call{Call: _e.mock.On("DELETE",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_DELETE_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_DELETE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_DELETE_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_DELETE_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_DELETE_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_DELETE_Call {
	_c.Call.Return(run)
	return _c
}

// Do provides a mock function with given fields: ctx, options
func (_m *MockRequester) Do(ctx context.Context, options ...do.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) error); ok {
		r0 = rf(ctx, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequester_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type MockRequester_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) Do(ctx interface{}, options ...interface{}) *MockRequester_Do_Call {
	return &MockRequester_Do_Call{Call: _e.mock.On("Do",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_Do_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_Do_Call) Return(_a0 error) *MockRequester_Do_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequester_Do_Call) RunAndReturn(run func(context.Context, ...do.Option) error) *MockRequester_Do_Call {
	_c.Call.Return(run)
	return _c
}

// GET provides a mock function with given fields: ctx, options
func (_m *MockRequester) GET(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GET")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_GET_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GET'
type MockRequester_GET_Call struct {
	*mock.Call
}

// GET is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) GET(ctx interface{}, options ...interface{}) *MockRequester_GET_Call {
	return &MockRequester_GET_Call{Call: _e.mock.On("GET",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_GET_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_GET_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_GET_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_GET_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_GET_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_GET_Call {
	_c.Call.Return(run)
	return _c
}

// HEAD provides a mock function with given fields: ctx, options
func (_m *MockRequester) HEAD(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for HEAD")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_HEAD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HEAD'
type MockRequester_HEAD_Call struct {
	*mock.Call
}

// HEAD is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) HEAD(ctx interface{}, options ...interface{}) *MockRequester_HEAD_Call {
	return &MockRequester_HEAD_Call{Call: _e.mock.On("HEAD",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_HEAD_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_HEAD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_HEAD_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_HEAD_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_HEAD_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_HEAD_Call {
	_c.Call.Return(run)
	return _c
}

// OPTIONS provides a mock function with given fields: ctx, options
func (_m *MockRequester) OPTIONS(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for OPTIONS")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_OPTIONS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OPTIONS'
type MockRequester_OPTIONS_Call struct {
	*mock.Call
}

// OPTIONS is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) OPTIONS(ctx interface{}, options ...interface{}) *MockRequester_OPTIONS_Call {
	return &MockRequester_OPTIONS_Call{Call: _e.mock.On("OPTIONS",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_OPTIONS_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_OPTIONS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_OPTIONS_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_OPTIONS_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_OPTIONS_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_OPTIONS_Call {
	_c.Call.Return(run)
	return _c
}

// PATCH provides a mock function with given fields: ctx, options
func (_m *MockRequester) PATCH(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PATCH")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_PATCH_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PATCH'
type MockRequester_PATCH_Call struct {
	*mock.Call
}

// PATCH is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) PATCH(ctx interface{}, options ...interface{}) *MockRequester_PATCH_Call {
	return &MockRequester_PATCH_Call{Call: _e.mock.On("PATCH",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_PATCH_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_PATCH_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_PATCH_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_PATCH_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_PATCH_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_PATCH_Call {
	_c.Call.Return(run)
	return _c
}

// POST provides a mock function with given fields: ctx, options
func (_m *MockRequester) POST(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for POST")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_POST_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'POST'
type MockRequester_POST_Call struct {
	*mock.Call
}

// POST is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) POST(ctx interface{}, options ...interface{}) *MockRequester_POST_Call {
	return &MockRequester_POST_Call{Call: _e.mock.On("POST",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_POST_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_POST_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_POST_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_POST_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_POST_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_POST_Call {
	_c.Call.Return(run)
	return _c
}

// PUT provides a mock function with given fields: ctx, options
func (_m *MockRequester) PUT(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PUT")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_PUT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PUT'
type MockRequester_PUT_Call struct {
	*mock.Call
}

// PUT is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) PUT(ctx interface{}, options ...interface{}) *MockRequester_PUT_Call {
	return &MockRequester_PUT_Call{Call: _e.mock.On("PUT",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_PUT_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_PUT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_PUT_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_PUT_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_PUT_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_PUT_Call {
	_c.Call.Return(run)
	return _c
}

// TRACE provides a mock function with given fields: ctx, options
func (_m *MockRequester) TRACE(ctx context.Context, options ...do.Option) (*http.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TRACE")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) (*http.Response, error)); ok {
		return rf(ctx, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...do.Option) *http.Response); ok {
		r0 = rf(ctx, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...do.Option) error); ok {
		r1 = rf(ctx, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRequester_TRACE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TRACE'
type MockRequester_TRACE_Call struct {
	*mock.Call
}

// TRACE is a helper method to define mock.On call
//   - ctx context.Context
//   - options ...do.Option
func (_e *MockRequester_Expecter) TRACE(ctx interface{}, options ...interface{}) *MockRequester_TRACE_Call {
	return &MockRequester_TRACE_Call{Call: _e.mock.On("TRACE",
		append([]interface{}{ctx}, options...)...)}
}

func (_c *MockRequester_TRACE_Call) Run(run func(ctx context.Context, options ...do.Option)) *MockRequester_TRACE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]do.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(do.Option)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRequester_TRACE_Call) Return(_a0 *http.Response, _a1 error) *MockRequester_TRACE_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRequester_TRACE_Call) RunAndReturn(run func(context.Context, ...do.Option) (*http.Response, error)) *MockRequester_TRACE_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequester creates a new instance of MockRequester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequester(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequester {
	mock := &MockRequester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
