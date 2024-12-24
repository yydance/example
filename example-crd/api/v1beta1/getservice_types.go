/*
Copyright 2024.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GetServiceSpec defines the desired state of GetService.
type GetServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Regex     bool   `json:"regex,omitempty"`
	MatchStr  string `json:"matchStr,omitempty"`
	GetAll    bool   `json:"getAll,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Action    string `json:"action,omitempty"`
}

// GetServiceStatus defines the observed state of GetService.
type GetServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Completed bool   `json:"completed,omitempty"`
	Status    string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GetService is the Schema for the getservices API.
type GetService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GetServiceSpec   `json:"spec,omitempty"`
	Status GetServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GetServiceList contains a list of GetService.
type GetServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GetService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GetService{}, &GetServiceList{})
}
