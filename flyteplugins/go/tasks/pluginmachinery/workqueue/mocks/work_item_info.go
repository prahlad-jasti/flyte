// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	workqueue "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/workqueue"
	mock "github.com/stretchr/testify/mock"
)

// WorkItemInfo is an autogenerated mock type for the WorkItemInfo type
type WorkItemInfo struct {
	mock.Mock
}

type WorkItemInfo_Error struct {
	*mock.Call
}

func (_m WorkItemInfo_Error) Return(_a0 error) *WorkItemInfo_Error {
	return &WorkItemInfo_Error{Call: _m.Call.Return(_a0)}
}

func (_m *WorkItemInfo) OnError() *WorkItemInfo_Error {
	c := _m.On("Error")
	return &WorkItemInfo_Error{Call: c}
}

func (_m *WorkItemInfo) OnErrorMatch(matchers ...interface{}) *WorkItemInfo_Error {
	c := _m.On("Error", matchers...)
	return &WorkItemInfo_Error{Call: c}
}

// Error provides a mock function with given fields:
func (_m *WorkItemInfo) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type WorkItemInfo_ID struct {
	*mock.Call
}

func (_m WorkItemInfo_ID) Return(_a0 string) *WorkItemInfo_ID {
	return &WorkItemInfo_ID{Call: _m.Call.Return(_a0)}
}

func (_m *WorkItemInfo) OnID() *WorkItemInfo_ID {
	c := _m.On("ID")
	return &WorkItemInfo_ID{Call: c}
}

func (_m *WorkItemInfo) OnIDMatch(matchers ...interface{}) *WorkItemInfo_ID {
	c := _m.On("ID", matchers...)
	return &WorkItemInfo_ID{Call: c}
}

// ID provides a mock function with given fields:
func (_m *WorkItemInfo) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type WorkItemInfo_Item struct {
	*mock.Call
}

func (_m WorkItemInfo_Item) Return(_a0 workqueue.WorkItem) *WorkItemInfo_Item {
	return &WorkItemInfo_Item{Call: _m.Call.Return(_a0)}
}

func (_m *WorkItemInfo) OnItem() *WorkItemInfo_Item {
	c := _m.On("Item")
	return &WorkItemInfo_Item{Call: c}
}

func (_m *WorkItemInfo) OnItemMatch(matchers ...interface{}) *WorkItemInfo_Item {
	c := _m.On("Item", matchers...)
	return &WorkItemInfo_Item{Call: c}
}

// Item provides a mock function with given fields:
func (_m *WorkItemInfo) Item() workqueue.WorkItem {
	ret := _m.Called()

	var r0 workqueue.WorkItem
	if rf, ok := ret.Get(0).(func() workqueue.WorkItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(workqueue.WorkItem)
		}
	}

	return r0
}

type WorkItemInfo_Status struct {
	*mock.Call
}

func (_m WorkItemInfo_Status) Return(_a0 workqueue.WorkStatus) *WorkItemInfo_Status {
	return &WorkItemInfo_Status{Call: _m.Call.Return(_a0)}
}

func (_m *WorkItemInfo) OnStatus() *WorkItemInfo_Status {
	c := _m.On("Status")
	return &WorkItemInfo_Status{Call: c}
}

func (_m *WorkItemInfo) OnStatusMatch(matchers ...interface{}) *WorkItemInfo_Status {
	c := _m.On("Status", matchers...)
	return &WorkItemInfo_Status{Call: c}
}

// Status provides a mock function with given fields:
func (_m *WorkItemInfo) Status() workqueue.WorkStatus {
	ret := _m.Called()

	var r0 workqueue.WorkStatus
	if rf, ok := ret.Get(0).(func() workqueue.WorkStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(workqueue.WorkStatus)
	}

	return r0
}
