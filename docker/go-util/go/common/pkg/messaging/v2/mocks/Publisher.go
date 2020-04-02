// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pubsub "cloud.google.com/go/pubsub"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, msg
func (_m *Publisher) Publish(ctx context.Context, msg *pubsub.Message) *pubsub.PublishResult {
	ret := _m.Called(ctx, msg)

	var r0 *pubsub.PublishResult
	if rf, ok := ret.Get(0).(func(context.Context, *pubsub.Message) *pubsub.PublishResult); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pubsub.PublishResult)
		}
	}

	return r0
}
