// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	constraints "github.com/vektra/mockery/v2/pkg/fixtures/constraints"

	test "github.com/vektra/mockery/v2/pkg/fixtures/redefined_type_b"
)

// ReplaceGeneric is an autogenerated mock type for the ReplaceGeneric type
type ReplaceGeneric[TConstraint constraints.String, TKeep any] struct {
	mock.Mock
}

type ReplaceGeneric_Expecter[TConstraint constraints.String, TKeep any] struct {
	mock *mock.Mock
}

func (_m *ReplaceGeneric[TConstraint, TKeep]) EXPECT() *ReplaceGeneric_Expecter[TConstraint, TKeep] {
	return &ReplaceGeneric_Expecter[TConstraint, TKeep]{mock: &_m.Mock}
}

// A provides a mock function with given fields: t1
func (_m *ReplaceGeneric[TConstraint, TKeep]) A(t1 test.B) TKeep {
	ret := _m.Called(t1)

	if len(ret) == 0 {
		panic("no return value specified for A")
	}

	var r0 TKeep
	if rf, ok := ret.Get(0).(func(test.B) TKeep); ok {
		r0 = rf(t1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TKeep)
		}
	}

	return r0
}

// ReplaceGeneric_A_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'A'
type ReplaceGeneric_A_Call[TConstraint constraints.String, TKeep any] struct {
	*mock.Call
}

// A is a helper method to define mock.On call
//   - t1 test.B
func (_e *ReplaceGeneric_Expecter[TConstraint, TKeep]) A(t1 interface{}) *ReplaceGeneric_A_Call[TConstraint, TKeep] {
	return &ReplaceGeneric_A_Call[TConstraint, TKeep]{Call: _e.mock.On("A", t1)}
}

func (_c *ReplaceGeneric_A_Call[TConstraint, TKeep]) Run(run func(t1 test.B)) *ReplaceGeneric_A_Call[TConstraint, TKeep] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(test.B))
	})
	return _c
}

func (_c *ReplaceGeneric_A_Call[TConstraint, TKeep]) Return(_a0 TKeep) *ReplaceGeneric_A_Call[TConstraint, TKeep] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplaceGeneric_A_Call[TConstraint, TKeep]) RunAndReturn(run func(test.B) TKeep) *ReplaceGeneric_A_Call[TConstraint, TKeep] {
	_c.Call.Return(run)
	return _c
}

// B provides a mock function with given fields:
func (_m *ReplaceGeneric[TConstraint, TKeep]) B() test.B {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for B")
	}

	var r0 test.B
	if rf, ok := ret.Get(0).(func() test.B); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(test.B)
		}
	}

	return r0
}

// ReplaceGeneric_B_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'B'
type ReplaceGeneric_B_Call[TConstraint constraints.String, TKeep any] struct {
	*mock.Call
}

// B is a helper method to define mock.On call
func (_e *ReplaceGeneric_Expecter[TConstraint, TKeep]) B() *ReplaceGeneric_B_Call[TConstraint, TKeep] {
	return &ReplaceGeneric_B_Call[TConstraint, TKeep]{Call: _e.mock.On("B")}
}

func (_c *ReplaceGeneric_B_Call[TConstraint, TKeep]) Run(run func()) *ReplaceGeneric_B_Call[TConstraint, TKeep] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReplaceGeneric_B_Call[TConstraint, TKeep]) Return(_a0 test.B) *ReplaceGeneric_B_Call[TConstraint, TKeep] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplaceGeneric_B_Call[TConstraint, TKeep]) RunAndReturn(run func() test.B) *ReplaceGeneric_B_Call[TConstraint, TKeep] {
	_c.Call.Return(run)
	return _c
}

// C provides a mock function with given fields:
func (_m *ReplaceGeneric[TConstraint, TKeep]) C() TConstraint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for C")
	}

	var r0 TConstraint
	if rf, ok := ret.Get(0).(func() TConstraint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TConstraint)
		}
	}

	return r0
}

// ReplaceGeneric_C_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'C'
type ReplaceGeneric_C_Call[TConstraint constraints.String, TKeep any] struct {
	*mock.Call
}

// C is a helper method to define mock.On call
func (_e *ReplaceGeneric_Expecter[TConstraint, TKeep]) C() *ReplaceGeneric_C_Call[TConstraint, TKeep] {
	return &ReplaceGeneric_C_Call[TConstraint, TKeep]{Call: _e.mock.On("C")}
}

func (_c *ReplaceGeneric_C_Call[TConstraint, TKeep]) Run(run func()) *ReplaceGeneric_C_Call[TConstraint, TKeep] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReplaceGeneric_C_Call[TConstraint, TKeep]) Return(_a0 TConstraint) *ReplaceGeneric_C_Call[TConstraint, TKeep] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplaceGeneric_C_Call[TConstraint, TKeep]) RunAndReturn(run func() TConstraint) *ReplaceGeneric_C_Call[TConstraint, TKeep] {
	_c.Call.Return(run)
	return _c
}

// NewReplaceGeneric creates a new instance of ReplaceGeneric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReplaceGeneric[TConstraint constraints.String, TKeep any](t interface {
	mock.TestingT
	Cleanup(func())
}) *ReplaceGeneric[TConstraint, TKeep] {
	mock := &ReplaceGeneric[TConstraint, TKeep]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
