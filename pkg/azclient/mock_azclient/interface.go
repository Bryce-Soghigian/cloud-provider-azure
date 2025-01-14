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

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cloud-provider-azure/pkg/azclient (interfaces: ClientFactory)

// Package mock_azclient is a generated GoMock package.
package mock_azclient

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	availabilitysetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	deploymentclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	diskclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	interfaceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	loadbalancerclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	managedclusterclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	privateendpointclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	privatelinkserviceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	privatezoneclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	publicipaddressclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	publicipprefixclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	routetableclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	securitygroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	snapshotclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	subnetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	virtualmachineclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	virtualmachinescalesetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	virtualmachinescalesetvmclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
)

// MockClientFactory is a mock of ClientFactory interface.
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory.
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance.
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// GetavailabilitysetclientInterface mocks base method.
func (m *MockClientFactory) GetavailabilitysetclientInterface(arg0 string) (availabilitysetclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetavailabilitysetclientInterface", arg0)
	ret0, _ := ret[0].(availabilitysetclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetavailabilitysetclientInterface indicates an expected call of GetavailabilitysetclientInterface.
func (mr *MockClientFactoryMockRecorder) GetavailabilitysetclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetavailabilitysetclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetavailabilitysetclientInterface), arg0)
}

// GetdeploymentclientInterface mocks base method.
func (m *MockClientFactory) GetdeploymentclientInterface(arg0 string) (deploymentclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetdeploymentclientInterface", arg0)
	ret0, _ := ret[0].(deploymentclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetdeploymentclientInterface indicates an expected call of GetdeploymentclientInterface.
func (mr *MockClientFactoryMockRecorder) GetdeploymentclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetdeploymentclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetdeploymentclientInterface), arg0)
}

// GetdiskclientInterface mocks base method.
func (m *MockClientFactory) GetdiskclientInterface(arg0 string) (diskclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetdiskclientInterface", arg0)
	ret0, _ := ret[0].(diskclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetdiskclientInterface indicates an expected call of GetdiskclientInterface.
func (mr *MockClientFactoryMockRecorder) GetdiskclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetdiskclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetdiskclientInterface), arg0)
}

// GetinterfaceclientInterface mocks base method.
func (m *MockClientFactory) GetinterfaceclientInterface(arg0 string) (interfaceclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetinterfaceclientInterface", arg0)
	ret0, _ := ret[0].(interfaceclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetinterfaceclientInterface indicates an expected call of GetinterfaceclientInterface.
func (mr *MockClientFactoryMockRecorder) GetinterfaceclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetinterfaceclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetinterfaceclientInterface), arg0)
}

// GetloadbalancerclientInterface mocks base method.
func (m *MockClientFactory) GetloadbalancerclientInterface(arg0 string) (loadbalancerclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetloadbalancerclientInterface", arg0)
	ret0, _ := ret[0].(loadbalancerclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetloadbalancerclientInterface indicates an expected call of GetloadbalancerclientInterface.
func (mr *MockClientFactoryMockRecorder) GetloadbalancerclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetloadbalancerclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetloadbalancerclientInterface), arg0)
}

// GetmanagedclusterclientInterface mocks base method.
func (m *MockClientFactory) GetmanagedclusterclientInterface(arg0 string) (managedclusterclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetmanagedclusterclientInterface", arg0)
	ret0, _ := ret[0].(managedclusterclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetmanagedclusterclientInterface indicates an expected call of GetmanagedclusterclientInterface.
func (mr *MockClientFactoryMockRecorder) GetmanagedclusterclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetmanagedclusterclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetmanagedclusterclientInterface), arg0)
}

// GetprivateendpointclientInterface mocks base method.
func (m *MockClientFactory) GetprivateendpointclientInterface(arg0 string) (privateendpointclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetprivateendpointclientInterface", arg0)
	ret0, _ := ret[0].(privateendpointclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetprivateendpointclientInterface indicates an expected call of GetprivateendpointclientInterface.
func (mr *MockClientFactoryMockRecorder) GetprivateendpointclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetprivateendpointclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetprivateendpointclientInterface), arg0)
}

// GetprivatelinkserviceclientInterface mocks base method.
func (m *MockClientFactory) GetprivatelinkserviceclientInterface(arg0 string) (privatelinkserviceclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetprivatelinkserviceclientInterface", arg0)
	ret0, _ := ret[0].(privatelinkserviceclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetprivatelinkserviceclientInterface indicates an expected call of GetprivatelinkserviceclientInterface.
func (mr *MockClientFactoryMockRecorder) GetprivatelinkserviceclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetprivatelinkserviceclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetprivatelinkserviceclientInterface), arg0)
}

// GetprivatezoneclientInterface mocks base method.
func (m *MockClientFactory) GetprivatezoneclientInterface(arg0 string) (privatezoneclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetprivatezoneclientInterface", arg0)
	ret0, _ := ret[0].(privatezoneclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetprivatezoneclientInterface indicates an expected call of GetprivatezoneclientInterface.
func (mr *MockClientFactoryMockRecorder) GetprivatezoneclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetprivatezoneclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetprivatezoneclientInterface), arg0)
}

// GetpublicipaddressclientInterface mocks base method.
func (m *MockClientFactory) GetpublicipaddressclientInterface(arg0 string) (publicipaddressclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetpublicipaddressclientInterface", arg0)
	ret0, _ := ret[0].(publicipaddressclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetpublicipaddressclientInterface indicates an expected call of GetpublicipaddressclientInterface.
func (mr *MockClientFactoryMockRecorder) GetpublicipaddressclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetpublicipaddressclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetpublicipaddressclientInterface), arg0)
}

// GetpublicipprefixclientInterface mocks base method.
func (m *MockClientFactory) GetpublicipprefixclientInterface(arg0 string) (publicipprefixclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetpublicipprefixclientInterface", arg0)
	ret0, _ := ret[0].(publicipprefixclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetpublicipprefixclientInterface indicates an expected call of GetpublicipprefixclientInterface.
func (mr *MockClientFactoryMockRecorder) GetpublicipprefixclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetpublicipprefixclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetpublicipprefixclientInterface), arg0)
}

// GetroutetableclientInterface mocks base method.
func (m *MockClientFactory) GetroutetableclientInterface(arg0 string) (routetableclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetroutetableclientInterface", arg0)
	ret0, _ := ret[0].(routetableclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetroutetableclientInterface indicates an expected call of GetroutetableclientInterface.
func (mr *MockClientFactoryMockRecorder) GetroutetableclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetroutetableclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetroutetableclientInterface), arg0)
}

// GetsecuritygroupclientInterface mocks base method.
func (m *MockClientFactory) GetsecuritygroupclientInterface(arg0 string) (securitygroupclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetsecuritygroupclientInterface", arg0)
	ret0, _ := ret[0].(securitygroupclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetsecuritygroupclientInterface indicates an expected call of GetsecuritygroupclientInterface.
func (mr *MockClientFactoryMockRecorder) GetsecuritygroupclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetsecuritygroupclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetsecuritygroupclientInterface), arg0)
}

// GetsnapshotclientInterface mocks base method.
func (m *MockClientFactory) GetsnapshotclientInterface(arg0 string) (snapshotclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetsnapshotclientInterface", arg0)
	ret0, _ := ret[0].(snapshotclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetsnapshotclientInterface indicates an expected call of GetsnapshotclientInterface.
func (mr *MockClientFactoryMockRecorder) GetsnapshotclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetsnapshotclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetsnapshotclientInterface), arg0)
}

// GetsubnetclientInterface mocks base method.
func (m *MockClientFactory) GetsubnetclientInterface(arg0 string) (subnetclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetsubnetclientInterface", arg0)
	ret0, _ := ret[0].(subnetclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetsubnetclientInterface indicates an expected call of GetsubnetclientInterface.
func (mr *MockClientFactoryMockRecorder) GetsubnetclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetsubnetclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetsubnetclientInterface), arg0)
}

// GetvirtualmachineclientInterface mocks base method.
func (m *MockClientFactory) GetvirtualmachineclientInterface(arg0 string) (virtualmachineclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetvirtualmachineclientInterface", arg0)
	ret0, _ := ret[0].(virtualmachineclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetvirtualmachineclientInterface indicates an expected call of GetvirtualmachineclientInterface.
func (mr *MockClientFactoryMockRecorder) GetvirtualmachineclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetvirtualmachineclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetvirtualmachineclientInterface), arg0)
}

// GetvirtualmachinescalesetclientInterface mocks base method.
func (m *MockClientFactory) GetvirtualmachinescalesetclientInterface(arg0 string) (virtualmachinescalesetclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetvirtualmachinescalesetclientInterface", arg0)
	ret0, _ := ret[0].(virtualmachinescalesetclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetvirtualmachinescalesetclientInterface indicates an expected call of GetvirtualmachinescalesetclientInterface.
func (mr *MockClientFactoryMockRecorder) GetvirtualmachinescalesetclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetvirtualmachinescalesetclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetvirtualmachinescalesetclientInterface), arg0)
}

// GetvirtualmachinescalesetvmclientInterface mocks base method.
func (m *MockClientFactory) GetvirtualmachinescalesetvmclientInterface(arg0 string) (virtualmachinescalesetvmclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetvirtualmachinescalesetvmclientInterface", arg0)
	ret0, _ := ret[0].(virtualmachinescalesetvmclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetvirtualmachinescalesetvmclientInterface indicates an expected call of GetvirtualmachinescalesetvmclientInterface.
func (mr *MockClientFactoryMockRecorder) GetvirtualmachinescalesetvmclientInterface(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetvirtualmachinescalesetvmclientInterface", reflect.TypeOf((*MockClientFactory)(nil).GetvirtualmachinescalesetvmclientInterface), arg0)
}
