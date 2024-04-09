// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Copyright 2017 The Kubernetes Authors.

// package v1beta1 is alpha objects from meta that will be introduced.
package v1beta1

import (
	slim_metav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
)

// Table is a tabular representation of a set of API resources. The server transforms the
// object into a set of preferred columns for quickly reviewing the objects.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +protobuf=false
type Table = slim_metav1.Table

// TableColumnDefinition contains information about a column returned in the Table.
// +protobuf=false
type TableColumnDefinition = slim_metav1.TableColumnDefinition

// TableRow is an individual row in a table.
// +protobuf=false
type TableRow = slim_metav1.TableRow

// TableRowCondition allows a row to be marked with additional information.
// +protobuf=false
type TableRowCondition = slim_metav1.TableRowCondition

type RowConditionType = slim_metav1.RowConditionType

type ConditionStatus = slim_metav1.ConditionStatus

type IncludeObjectPolicy = slim_metav1.IncludeObjectPolicy

// TableOptions are used when a Table is requested by the caller.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TableOptions = slim_metav1.TableOptions

// PartialObjectMetadata is a generic representation of any object with ObjectMeta. It allows clients
// to get access to a particular ObjectMeta schema without knowing the details of the version.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PartialObjectMetadata = slim_metav1.PartialObjectMetadata

// IMPORTANT: PartialObjectMetadataList has different protobuf field ids in v1beta1 than
// v1 because ListMeta was accidentally omitted prior to 1.15. Therefore this type must
// remain independent of v1.PartialObjectMetadataList to preserve mappings.

// PartialObjectMetadataList contains a list of objects containing only their metadata.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PartialObjectMetadataList struct {
	slim_metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	slim_metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,2,opt,name=metadata"`

	// items contains each of the included items.
	Items []slim_metav1.PartialObjectMetadata `json:"items" protobuf:"bytes,1,rep,name=items"`
}

const (
	RowCompleted = slim_metav1.RowCompleted

	ConditionTrue    = slim_metav1.ConditionTrue
	ConditionFalse   = slim_metav1.ConditionFalse
	ConditionUnknown = slim_metav1.ConditionUnknown

	IncludeNone     = slim_metav1.IncludeNone
	IncludeMetadata = slim_metav1.IncludeMetadata
	IncludeObject   = slim_metav1.IncludeObject
)
