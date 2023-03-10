// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cart "projects/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartService is an autogenerated mock type for the CartService type
type CartService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, productId, qty
func (_m *CartService) Add(token interface{}, productId uint, qty int) (cart.Core, error) {
	ret := _m.Called(token, productId, qty)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, int) cart.Core); ok {
		r0 = rf(token, productId, qty)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, int) error); ok {
		r1 = rf(token, productId, qty)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProductCart provides a mock function with given fields: token, productId
func (_m *CartService) DeleteProductCart(token interface{}, productId uint) error {
	ret := _m.Called(token, productId)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMyCart provides a mock function with given fields: token
func (_m *CartService) GetMyCart(token interface{}) ([]cart.Core, error) {
	ret := _m.Called(token)

	var r0 []cart.Core
	if rf, ok := ret.Get(0).(func(interface{}) []cart.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductCart provides a mock function with given fields: token, productId, qty
func (_m *CartService) UpdateProductCart(token interface{}, productId uint, qty int) error {
	ret := _m.Called(token, productId, qty)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint, int) error); ok {
		r0 = rf(token, productId, qty)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartService creates a new instance of CartService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartService(t mockConstructorTestingTNewCartService) *CartService {
	mock := &CartService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
