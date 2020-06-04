#!/usr/bin/env bash

# Copyright 2019 The Tekton Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

source $(git rev-parse --show-toplevel)/vendor/github.com/tektoncd/plumbing/scripts/library.sh

PREFIX=${GOBIN:-${GOPATH}/bin}

OLDGOFLAGS="${GOFLAGS:-}"
GOFLAGS="-mod=vendor"
# This generates deepcopy,client,informer and lister for the custom/v1alpha1 package
bash ${REPO_ROOT_DIR}/hack/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/ImJasonH/custom-task/pkg/client github.com/ImJasonH/custom-task/pkg/apis \
  "custom:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

# Knative Injection
# This generates the knative inject packages for the custom/v1alpha1 package
bash ${REPO_ROOT_DIR}/hack/generate-knative.sh "injection" \
  github.com/ImJasonH/custom-task/pkg/client github.com/ImJasonH/custom-task/pkg/apis \
  "custom:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt
GOFLAGS="${OLDGOFLAGS}"

# Make sure our dependencies are up-to-date
${REPO_ROOT_DIR}/hack/update-deps.sh
