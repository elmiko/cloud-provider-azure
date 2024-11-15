// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: fileshareclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_fileshareclient -source fileshareclient/interface.go -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_fileshareclient is a generated GoMock package.
package mock_fileshareclient

import (
	context "context"
	reflect "reflect"

	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockInterface) Create(ctx context.Context, resourceGroupName, resourceName, parentResourceName string, resource armstorage.FileShare) (*armstorage.FileShare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, resourceGroupName, resourceName, parentResourceName, resource)
	ret0, _ := ret[0].(*armstorage.FileShare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInterfaceMockRecorder) Create(ctx, resourceGroupName, resourceName, parentResourceName, resource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInterface)(nil).Create), ctx, resourceGroupName, resourceName, parentResourceName, resource)
}

// Delete mocks base method.
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, parentResourceName, resourceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, parentResourceName, resourceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, parentResourceName, resourceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, parentResourceName, resourceName)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, parentResourceName, resourceName string) (*armstorage.FileShare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, parentResourceName, resourceName)
	ret0, _ := ret[0].(*armstorage.FileShare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, parentResourceName, resourceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, parentResourceName, resourceName)
}

// Update mocks base method.
func (m *MockInterface) Update(ctx context.Context, resourceGroupName, resourceName, parentResourceName string, resource armstorage.FileShare) (*armstorage.FileShare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, resourceGroupName, resourceName, parentResourceName, resource)
	ret0, _ := ret[0].(*armstorage.FileShare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockInterfaceMockRecorder) Update(ctx, resourceGroupName, resourceName, parentResourceName, resource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), ctx, resourceGroupName, resourceName, parentResourceName, resource)
}
