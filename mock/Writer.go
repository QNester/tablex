// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	table "github.com/jedib0t/go-pretty/v6/table"
	mock "github.com/stretchr/testify/mock"
)

// WriterMock is an autogenerated mock type for the Writer type
type WriterMock struct {
	mock.Mock
}

// AppendFooter provides a mock function with given fields: row, configs
func (_m *WriterMock) AppendFooter(row table.Row, configs ...table.RowConfig) {
	_va := make([]interface{}, len(configs))
	for _i := range configs {
		_va[_i] = configs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, row)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// AppendHeader provides a mock function with given fields: row, configs
func (_m *WriterMock) AppendHeader(row table.Row, configs ...table.RowConfig) {
	_va := make([]interface{}, len(configs))
	for _i := range configs {
		_va[_i] = configs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, row)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// AppendRow provides a mock function with given fields: row, configs
func (_m *WriterMock) AppendRow(row table.Row, configs ...table.RowConfig) {
	_va := make([]interface{}, len(configs))
	for _i := range configs {
		_va[_i] = configs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, row)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// AppendRows provides a mock function with given fields: rows, configs
func (_m *WriterMock) AppendRows(rows []table.Row, configs ...table.RowConfig) {
	_va := make([]interface{}, len(configs))
	for _i := range configs {
		_va[_i] = configs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, rows)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Render provides a mock function with given fields:
func (_m *WriterMock) Render() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RenderCSV provides a mock function with given fields:
func (_m *WriterMock) RenderCSV() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RenderHTML provides a mock function with given fields:
func (_m *WriterMock) RenderHTML() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RenderMarkdown provides a mock function with given fields:
func (_m *WriterMock) RenderMarkdown() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewWriterMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriterMock creates a new instance of WriterMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriterMock(t mockConstructorTestingTNewWriterMock) *WriterMock {
	mock := &WriterMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}