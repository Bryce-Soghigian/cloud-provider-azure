/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
)

var (
	testVmssFlexID1 = "subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/virtualMachineScaleSets/vmssflex1"
	testNodeName1   = "vmssflex1000001"
	testNode1       = &v1.Node{
		ObjectMeta: meta.ObjectMeta{
			Name: testNodeName1,
		},
	}

	testVmssFlexID2 = "subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/virtualMachineScaleSets/vmssflex2"
	testNodeName2   = "vmssflex2000001"
	testNode2       = &v1.Node{
		ObjectMeta: meta.ObjectMeta{
			Name: testNodeName2,
		},
	}
)

func TestGetNodeVMSetNameVmssFlex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		description       string
		expectedVMSetName string
		expectedErr       error
	}{
		{
			description:       "GetNodeVMSetName should return the correct VMSetName of the node",
			expectedVMSetName: "vmssflex1",
			expectedErr:       nil,
		},
	}

	for _, tc := range testCases {
		fs, err := NewTestFlexScaleSet(ctrl)
		assert.NoError(t, err, "unexpected error when creating test FlexScaleSet")
		fs.vmssFlexVMNameToVmssID.Store(testNodeName1, testVmssFlexID1)
		fs.vmssFlexVMNameToVmssID.Store(testNodeName2, testVmssFlexID2)

		vmSetName, err := fs.GetNodeVMSetName(testNode1)
		assert.Equal(t, tc.expectedVMSetName, vmSetName, tc.description)
		assert.Equal(t, tc.expectedErr, err, tc.description)
	}

}

func TestGetAgentPoolVMSetNamesVmssFlex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		description                 string
		nodes                       []*v1.Node
		expectedAgentPoolVMSetNames *[]string
		expectedErr                 error
	}{
		{
			description:                 "GetNodeVMSetName should return the correct VMSetName of the node",
			nodes:                       []*v1.Node{testNode1, testNode2},
			expectedAgentPoolVMSetNames: &[]string{"vmssflex1", "vmssflex2"},
			expectedErr:                 nil,
		},
	}

	for _, tc := range testCases {
		fs, err := NewTestFlexScaleSet(ctrl)
		assert.NoError(t, err, "unexpected error when creating test FlexScaleSet")
		fs.vmssFlexVMNameToVmssID.Store(testNodeName1, testVmssFlexID1)
		fs.vmssFlexVMNameToVmssID.Store(testNodeName2, testVmssFlexID2)

		agentPoolVMSetNames, err := fs.GetAgentPoolVMSetNames(tc.nodes)
		assert.Equal(t, tc.expectedAgentPoolVMSetNames, agentPoolVMSetNames, tc.description)
		assert.Equal(t, tc.expectedErr, err, tc.description)
	}
}

func TestGetVMSetNamesVmssFlex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		description        string
		service            *v1.Service
		nodes              []*v1.Node
		useSingleSLB       bool
		expectedVMSetNames *[]string
		expectedErr        error
	}{
		{
			description:        "GetVMSetNames should return the primary vm set name if the service has no mode annotation",
			service:            &v1.Service{},
			expectedVMSetNames: &[]string{"vmss"},
		},
		{
			description: "GetVMSetNames should return the primary vm set name when using the single SLB",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{consts.ServiceAnnotationLoadBalancerMode: consts.ServiceAnnotationLoadBalancerAutoModeValue}},
			},
			useSingleSLB:       true,
			expectedVMSetNames: &[]string{"vmss"},
		},
		{
			description: "GetVMSetNames should return all scale sets if the service has auto mode annotation",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{consts.ServiceAnnotationLoadBalancerMode: consts.ServiceAnnotationLoadBalancerAutoModeValue}},
			},
			nodes:              []*v1.Node{testNode1, testNode2},
			expectedVMSetNames: &[]string{"vmssflex1", "vmssflex2"},
		},
		{
			description: "GetVMSetNames should report the error if there's no such vmss",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{consts.ServiceAnnotationLoadBalancerMode: "vmssflex3"}},
			},
			nodes:       []*v1.Node{testNode1, testNode2},
			expectedErr: fmt.Errorf("scale set (vmssflex3) - not found"),
		},
		{
			description: "GetVMSetNames should return the correct vmss names",
			service: &v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{consts.ServiceAnnotationLoadBalancerMode: "vmssflex1"}},
			},
			nodes:              []*v1.Node{testNode1, testNode2},
			expectedVMSetNames: &[]string{"vmssflex1"},
		},
	}

	for _, tc := range testCases {
		fs, err := NewTestFlexScaleSet(ctrl)
		assert.NoError(t, err, "unexpected error when creating test FlexScaleSet")
		fs.vmssFlexVMNameToVmssID.Store(testNodeName1, testVmssFlexID1)
		fs.vmssFlexVMNameToVmssID.Store(testNodeName2, testVmssFlexID2)

		if tc.useSingleSLB {
			fs.EnableMultipleStandardLoadBalancers = false
			fs.LoadBalancerSku = consts.LoadBalancerSkuStandard
		}

		vmSetNames, err := fs.GetVMSetNames(tc.service, tc.nodes)
		assert.Equal(t, tc.expectedVMSetNames, vmSetNames, tc.description)
		assert.Equal(t, tc.expectedErr, err, tc.description)
	}
}

func TestGetInstanceIDByNodeNameVmssFlex(t *testing.T) {

}

func TestGetInstanceTypeByNodeNameVmssFlex(t *testing.T) {

}

func TestGetZoneByNodeNameVmssFlex(t *testing.T) {

}

func TestGetProvisioningStateByNodeNameVmssFlex(t *testing.T) {

}

func TestGetPowerStatusByNodeNameVmssFlex(t *testing.T) {

}

func TestGetPrimaryInterfaceVmssFlex(t *testing.T) {

}

func TestGetIPByNodeNameVmssFlex(t *testing.T) {

}

func TestGetPrivateIPsByNodeNameVmssFlex(t *testing.T) {

}

func TestGetNodeNameByIPConfigurationIDVmssFlex(t *testing.T) {

}

func TestGetNodeCIDRMasksByProviderIDVmssFlex(t *testing.T) {

}

func TestEnsureHostInPoolVmssFlex(t *testing.T) {

}

func TestEnsureHostsInPoolVmssFlex(t *testing.T) {

}

func TestEnsureBackendPoolDeletedFromVMSetsVmssFlex(t *testing.T) {

}

func TestEnsureBackendPoolDeletedVmssFlex(t *testing.T) {

}