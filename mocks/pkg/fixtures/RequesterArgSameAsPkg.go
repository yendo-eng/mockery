// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RequesterArgSameAsPkg is an autogenerated mock type for the RequesterArgSameAsPkg type
type RequesterArgSameAsPkg struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0
func (_m *RequesterArgSameAsPkg) Get(_a0 string) {
	_m.Called(_a0)
}

// NewRequesterArgSameAsPkg creates a new instance of RequesterArgSameAsPkg. It also registers a cleanup function to assert the mocks expectations.
func NewRequesterArgSameAsPkg(t testing.TB) *RequesterArgSameAsPkg {
	mock := &RequesterArgSameAsPkg{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
