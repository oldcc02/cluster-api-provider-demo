/*
Copyright 2022.

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
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DemoClusterTemplateSpec defines the desired state of DemoClusterTemplate
type DemoClusterTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Template DemoClusterTemplateResource `json:"template"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DemoClusterTemplate is the Schema for the democlustertemplates API
type DemoClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DemoClusterTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// DemoClusterTemplateList contains a list of DemoClusterTemplate
type DemoClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoClusterTemplate{}, &DemoClusterTemplateList{})
}

type DemoClusterTemplateResource struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata,omitempty"`
	Spec       DemoClusterSpec      `json:"spec"`
}
