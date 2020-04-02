// Code generated by mockery v1.0.0. DO NOT EDIT.

package templates

import (
	context "context"

	datastore "cloud.google.com/go/datastore"
	datastoreclient "github.com/ForgeCloud/saas/tree/master/services/go/common/pkg/datastoreclient"

	key "github.com/ForgeCloud/saas/tree/master/services/go/common/pkg/models/key"

	mock "github.com/stretchr/testify/mock"

	template "github.com/ForgeCloud/saas/tree/master/services/go/common/pkg/models/template"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, encKey
func (_m *MockStore) Delete(ctx context.Context, encKey string) error {
	ret := _m.Called(ctx, encKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, encKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTemplateByType provides a mock function with given fields: ctx, tmplType
func (_m *MockStore) FindTemplateByType(ctx context.Context, tmplType string) (*template.DatastoreEntity, error) {
	ret := _m.Called(ctx, tmplType)

	var r0 *template.DatastoreEntity
	if rf, ok := ret.Get(0).(func(context.Context, string) *template.DatastoreEntity); ok {
		r0 = rf(ctx, tmplType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*template.DatastoreEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tmplType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, encKey, dst
func (_m *MockStore) Get(ctx context.Context, encKey string, dst key.Keyer) error {
	ret := _m.Called(ctx, encKey, dst)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, key.Keyer) error); ok {
		r0 = rf(ctx, encKey, dst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: ctx, qf, dst
func (_m *MockStore) Query(ctx context.Context, qf datastoreclient.QueryFilter, dst interface{}) error {
	ret := _m.Called(ctx, qf, dst)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, datastoreclient.QueryFilter, interface{}) error); ok {
		r0 = rf(ctx, qf, dst)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, _a1, src
func (_m *MockStore) Save(ctx context.Context, _a1 *datastore.Key, src key.Keyer) error {
	ret := _m.Called(ctx, _a1, src)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datastore.Key, key.Keyer) error); ok {
		r0 = rf(ctx, _a1, src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
