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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KueueOperatorSpec defines the desired state of KueueOperator
type KueueOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// JobSet provides configurations for jobset installation
	JobSet *JobSetSpec `json:"jobSet,omitempty"`
	// LeaderWorkSet provides configurations for LeaderWorkerSet installation
	LeaderWorkerSet *LeaderWorkerSet `json:"leaderWorkerSet,omitempty"`
	// Kueue
	Kueue *Kueue `json:"kueue,omitempty"`
}

// KueueOperatorStatus defines the observed state of KueueOperator
type KueueOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

type JobSetSpec struct {
	JobSetImage string `json:"jobSetImage"`
	Proxy       string `json:"proxy"`
}

type LeaderWorkerSet struct {
	Image string `json:"image"`
}

type Kueue struct {
	Image string `json:"image"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KueueOperator is the Schema for the kueueoperators API
type KueueOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KueueOperatorSpec   `json:"spec,omitempty"`
	Status KueueOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KueueOperatorList contains a list of KueueOperator
type KueueOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KueueOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KueueOperator{}, &KueueOperatorList{})
}
