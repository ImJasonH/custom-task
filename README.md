# Change import path

* grep and replace ImJasonH/custom-task with your GitHub user/repo
* grep and replace custom/v1alpha1 with yourgroup/v1alpha1

# Change the apiVersion and kind

* update doc.go: `// +groupName=custom.dev`

Any time you update types.go, run `./hack/update-codegen.sh`.
