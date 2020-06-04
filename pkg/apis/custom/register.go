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

package custom

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	// GroupName is the Kubernetes resource group name for CustomTask API types.
	GroupName = "custom.dev"

	// CustomTaskLabelKey is used as the label identifier for a CustomTask CRD objects.
	TaskLabelKey = "/customTask"

	// CustomTaskRunLabelKey is used as the label identifier for a CustomTaskRun CRD objects.
	TaskRunLabelKey = "/customTskRun"
)

var (
	// CustomTaskResource represents a CustomTask
	CustomTaskResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "customtasks",
	}
	// CustomTaskRunResource represents a CustomTaskRun
	CustomTaskRunResource = schema.GroupResource{
		Group:    GroupName,
		Resource: "customtaskruns",
	}
)
