/*
Copyright 2019 The Tekton Authors

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
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type CustomTask struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata"`

	// Spec holds the desired state of the CustomTask from the client
	// +optional
	Spec CustomTaskSpec `json:"spec"`
}

type CustomTaskSpec struct {
	// TODO: Fill in fields of a CustomTask's desired functionality.
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type CustomTaskRun struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata"`

	// +optional
	Spec CustomTaskRunSpec `json:"spec"`

	// +optional
	Status CustomTaskRunStatus `json:"status,omitempty"`
}

type CustomTaskRunSpec struct {
	// TODO: Fill in fields of a CustomTask's desired functionality.
}

type CustomTaskRunStatus struct {
	duckv1beta1.Status `json:",inline"`

	// StartTime is the time this run actually started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CompletionTime is the time this run completed.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
}
