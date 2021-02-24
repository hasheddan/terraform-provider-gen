/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:object:root=true

// Test is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Test struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestSpec   `json:"spec"`
	Status TestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Test contains a list of TestList
type TestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Test `json:"items"`
}

// A TestSpec defines the desired state of a Test
type TestSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TestParameters `json:"forProvider"`
}

// A TestParameters defines the desired state of a Test
type TestParameters struct {
	outerName outer
}

type outer struct {
	middleOneName middleOne
	middleTwoName middleTwo
	middleTwoName middleTwo0
}

type middleOne struct {
	duplicatorName duplicator
}

type duplicator struct {
	aString string `json:"a_string"`
}

type middleTwo struct {
	duplicatorName duplicator0
}

type duplicator0 struct {
	aString string `json:"a_string"`
}

type middleTwo0 struct {
	duplicatorName duplicator1
}

type duplicator1 struct {
	bString string `json:"b_string"`
}

// A TestStatus defines the observed state of a Test
type TestStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TestObservation `json:"atProvider"`
}

// A TestObservation records the observed state of a Test
type TestObservation struct{}
