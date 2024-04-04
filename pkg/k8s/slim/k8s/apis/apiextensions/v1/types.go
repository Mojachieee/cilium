/*
Copyright 2019 The Kubernetes Authors.

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
	metav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
)

// CustomResourceDefinitionConditionType is a valid value for CustomResourceDefinitionCondition.Type
type CustomResourceDefinitionConditionType string

const (
	// Established means that the resource has become active. A resource is established when all names are
	// accepted without a conflict for the first time. A resource stays established until deleted, even during
	// a later NamesAccepted due to changed names. Note that not all names can be changed.
	Established CustomResourceDefinitionConditionType = "Established"
	// NamesAccepted means the names chosen for this CustomResourceDefinition do not conflict with others in
	// the group and are therefore accepted.
	NamesAccepted CustomResourceDefinitionConditionType = "NamesAccepted"
	// NonStructuralSchema means that one or more OpenAPI schema is not structural.
	//
	// A schema is structural if it specifies types for all values, with the only exceptions of those with
	// - x-kubernetes-int-or-string: true — for fields which can be integer or string
	// - x-kubernetes-preserve-unknown-fields: true — for raw, unspecified JSON values
	// and there is no type, additionalProperties, default, nullable or x-kubernetes-* vendor extenions
	// specified under allOf, anyOf, oneOf or not.
	//
	// Non-structural schemas will not be allowed anymore in v1 API groups. Moreover, new features will not be
	// available for non-structural CRDs:
	// - pruning
	// - defaulting
	// - read-only
	// - OpenAPI publishing
	// - webhook conversion
	NonStructuralSchema CustomResourceDefinitionConditionType = "NonStructuralSchema"
	// Terminating means that the CustomResourceDefinition has been deleted and is cleaning up.
	Terminating CustomResourceDefinitionConditionType = "Terminating"
	// KubernetesAPIApprovalPolicyConformant indicates that an API in *.k8s.io or *.kubernetes.io is or is not approved.  For CRDs
	// outside those groups, this condition will not be set.  For CRDs inside those groups, the condition will
	// be true if .metadata.annotations["api-approved.kubernetes.io"] is set to a URL, otherwise it will be false.
	// See https://github.com/kubernetes/enhancements/pull/1111 for more details.
	KubernetesAPIApprovalPolicyConformant CustomResourceDefinitionConditionType = "KubernetesAPIApprovalPolicyConformant"
)

// CustomResourceCleanupFinalizer is the name of the finalizer which will delete instances of
// a CustomResourceDefinition
const CustomResourceCleanupFinalizer = "customresourcecleanup.apiextensions.k8s.io"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format
// <.spec.name>.<.spec.group>.
type CustomResourceDefinition struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomResourceDefinitionList is a list of CustomResourceDefinition objects.
type CustomResourceDefinitionList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items list individual CustomResourceDefinition objects
	Items []CustomResourceDefinition `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// CustomResourceSubresources defines the status and scale subresources for CustomResources.
type CustomResourceSubresources struct {
	// status indicates the custom resource should serve a `/status` subresource.
	// When enabled:
	// 1. requests to the custom resource primary endpoint ignore changes to the `status` stanza of the object.
	// 2. requests to the custom resource `/status` subresource ignore changes to anything other than the `status` stanza of the object.
	// +optional
	Status *CustomResourceSubresourceStatus `json:"status,omitempty" protobuf:"bytes,1,opt,name=status"`
	// scale indicates the custom resource should serve a `/scale` subresource that returns an `autoscaling/v1` Scale object.
	// +optional
	Scale *CustomResourceSubresourceScale `json:"scale,omitempty" protobuf:"bytes,2,opt,name=scale"`
}

// CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources.
// Status is represented by the `.status` JSON path inside of a CustomResource. When set,
// * exposes a /status subresource for the custom resource
// * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the status stanza
// * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza
type CustomResourceSubresourceStatus struct{}

// CustomResourceSubresourceScale defines how to serve the scale subresource for CustomResources.
type CustomResourceSubresourceScale struct {
	// specReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale `spec.replicas`.
	// Only JSON paths without the array notation are allowed.
	// Must be a JSON Path under `.spec`.
	// If there is no value under the given path in the custom resource, the `/scale` subresource will return an error on GET.
	SpecReplicasPath string `json:"specReplicasPath" protobuf:"bytes,1,name=specReplicasPath"`
	// statusReplicasPath defines the JSON path inside of a custom resource that corresponds to Scale `status.replicas`.
	// Only JSON paths without the array notation are allowed.
	// Must be a JSON Path under `.status`.
	// If there is no value under the given path in the custom resource, the `status.replicas` value in the `/scale` subresource
	// will default to 0.
	StatusReplicasPath string `json:"statusReplicasPath" protobuf:"bytes,2,opt,name=statusReplicasPath"`
	// labelSelectorPath defines the JSON path inside of a custom resource that corresponds to Scale `status.selector`.
	// Only JSON paths without the array notation are allowed.
	// Must be a JSON Path under `.status` or `.spec`.
	// Must be set to work with HorizontalPodAutoscaler.
	// The field pointed by this JSON path must be a string field (not a complex selector struct)
	// which contains a serialized label selector in string form.
	// More info: https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions#scale-subresource
	// If there is no value under the given path in the custom resource, the `status.selector` value in the `/scale`
	// subresource will default to the empty string.
	// +optional
	LabelSelectorPath *string `json:"labelSelectorPath,omitempty" protobuf:"bytes,3,opt,name=labelSelectorPath"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConversionReview describes a conversion request/response.
type ConversionReview struct {
	metav1.TypeMeta `json:",inline"`
	// request describes the attributes for the conversion request.
	// +optional
	Request *ConversionRequest `json:"request,omitempty" protobuf:"bytes,1,opt,name=request"`
}

// ConversionRequest describes the conversion request parameters.
type ConversionRequest struct {
	// desiredAPIVersion is the version to convert given objects to. e.g. "myapi.example.com/v1"
	DesiredAPIVersion string `json:"desiredAPIVersion" protobuf:"bytes,2,name=desiredAPIVersion"`
}
