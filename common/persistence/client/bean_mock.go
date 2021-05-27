// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: bean.go

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	"github.com/uber/cadence/common/persistence/visibility"

	gomock "github.com/golang/mock/gomock"

	persistence "github.com/uber/cadence/common/persistence"
)

// MockBean is a mock of Bean interface
type MockBean struct {
	ctrl     *gomock.Controller
	recorder *MockBeanMockRecorder
}

// MockBeanMockRecorder is the mock recorder for MockBean
type MockBeanMockRecorder struct {
	mock *MockBean
}

// NewMockBean creates a new mock instance
func NewMockBean(ctrl *gomock.Controller) *MockBean {
	mock := &MockBean{ctrl: ctrl}
	mock.recorder = &MockBeanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBean) EXPECT() *MockBeanMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockBean) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockBeanMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBean)(nil).Close))
}

// GetMetadataManager mocks base method
func (m *MockBean) GetMetadataManager() persistence.DomainManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataManager")
	ret0, _ := ret[0].(persistence.DomainManager)
	return ret0
}

// GetMetadataManager indicates an expected call of GetMetadataManager
func (mr *MockBeanMockRecorder) GetMetadataManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataManager", reflect.TypeOf((*MockBean)(nil).GetMetadataManager))
}

// SetMetadataManager mocks base method
func (m *MockBean) SetMetadataManager(arg0 persistence.DomainManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMetadataManager", arg0)
}

// SetMetadataManager indicates an expected call of SetMetadataManager
func (mr *MockBeanMockRecorder) SetMetadataManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadataManager", reflect.TypeOf((*MockBean)(nil).SetMetadataManager), arg0)
}

// GetTaskManager mocks base method
func (m *MockBean) GetTaskManager() persistence.TaskManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskManager")
	ret0, _ := ret[0].(persistence.TaskManager)
	return ret0
}

// GetTaskManager indicates an expected call of GetTaskManager
func (mr *MockBeanMockRecorder) GetTaskManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskManager", reflect.TypeOf((*MockBean)(nil).GetTaskManager))
}

// SetTaskManager mocks base method
func (m *MockBean) SetTaskManager(arg0 persistence.TaskManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTaskManager", arg0)
}

// SetTaskManager indicates an expected call of SetTaskManager
func (mr *MockBeanMockRecorder) SetTaskManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTaskManager", reflect.TypeOf((*MockBean)(nil).SetTaskManager), arg0)
}

// GetVisibilityManager mocks base method
func (m *MockBean) GetVisibilityManager() visibility.VisibilityManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVisibilityManager")
	ret0, _ := ret[0].(visibility.VisibilityManager)
	return ret0
}

// GetVisibilityManager indicates an expected call of GetVisibilityManager
func (mr *MockBeanMockRecorder) GetVisibilityManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVisibilityManager", reflect.TypeOf((*MockBean)(nil).GetVisibilityManager))
}

// SetVisibilityManager mocks base method
func (m *MockBean) SetVisibilityManager(arg0 visibility.VisibilityManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVisibilityManager", arg0)
}

// SetVisibilityManager indicates an expected call of SetVisibilityManager
func (mr *MockBeanMockRecorder) SetVisibilityManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVisibilityManager", reflect.TypeOf((*MockBean)(nil).SetVisibilityManager), arg0)
}

// GetDomainReplicationQueueManager mocks base method
func (m *MockBean) GetDomainReplicationQueueManager() persistence.QueueManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainReplicationQueueManager")
	ret0, _ := ret[0].(persistence.QueueManager)
	return ret0
}

// GetDomainReplicationQueueManager indicates an expected call of GetDomainReplicationQueueManager
func (mr *MockBeanMockRecorder) GetDomainReplicationQueueManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainReplicationQueueManager", reflect.TypeOf((*MockBean)(nil).GetDomainReplicationQueueManager))
}

// SetDomainReplicationQueueManager mocks base method
func (m *MockBean) SetDomainReplicationQueueManager(arg0 persistence.QueueManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDomainReplicationQueueManager", arg0)
}

// SetDomainReplicationQueueManager indicates an expected call of SetDomainReplicationQueueManager
func (mr *MockBeanMockRecorder) SetDomainReplicationQueueManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDomainReplicationQueueManager", reflect.TypeOf((*MockBean)(nil).SetDomainReplicationQueueManager), arg0)
}

// GetShardManager mocks base method
func (m *MockBean) GetShardManager() persistence.ShardManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardManager")
	ret0, _ := ret[0].(persistence.ShardManager)
	return ret0
}

// GetShardManager indicates an expected call of GetShardManager
func (mr *MockBeanMockRecorder) GetShardManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardManager", reflect.TypeOf((*MockBean)(nil).GetShardManager))
}

// SetShardManager mocks base method
func (m *MockBean) SetShardManager(arg0 persistence.ShardManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShardManager", arg0)
}

// SetShardManager indicates an expected call of SetShardManager
func (mr *MockBeanMockRecorder) SetShardManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShardManager", reflect.TypeOf((*MockBean)(nil).SetShardManager), arg0)
}

// GetHistoryManager mocks base method
func (m *MockBean) GetHistoryManager() persistence.HistoryManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryManager")
	ret0, _ := ret[0].(persistence.HistoryManager)
	return ret0
}

// GetHistoryManager indicates an expected call of GetHistoryManager
func (mr *MockBeanMockRecorder) GetHistoryManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryManager", reflect.TypeOf((*MockBean)(nil).GetHistoryManager))
}

// SetHistoryManager mocks base method
func (m *MockBean) SetHistoryManager(arg0 persistence.HistoryManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHistoryManager", arg0)
}

// SetHistoryManager indicates an expected call of SetHistoryManager
func (mr *MockBeanMockRecorder) SetHistoryManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHistoryManager", reflect.TypeOf((*MockBean)(nil).SetHistoryManager), arg0)
}

// GetExecutionManager mocks base method
func (m *MockBean) GetExecutionManager(arg0 int) (persistence.ExecutionManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecutionManager", arg0)
	ret0, _ := ret[0].(persistence.ExecutionManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecutionManager indicates an expected call of GetExecutionManager
func (mr *MockBeanMockRecorder) GetExecutionManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionManager", reflect.TypeOf((*MockBean)(nil).GetExecutionManager), arg0)
}

// SetExecutionManager mocks base method
func (m *MockBean) SetExecutionManager(arg0 int, arg1 persistence.ExecutionManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetExecutionManager", arg0, arg1)
}

// SetExecutionManager indicates an expected call of SetExecutionManager
func (mr *MockBeanMockRecorder) SetExecutionManager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExecutionManager", reflect.TypeOf((*MockBean)(nil).SetExecutionManager), arg0, arg1)
}
