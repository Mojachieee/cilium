// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Copyright 2017 The Kubernetes Authors.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	slim_metav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
)

// GroupName is the group name for this API.
const GroupName = "meta.k8s.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// AddMetaToScheme registers base meta types into schemas.
func AddMetaToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Table{},
		&TableOptions{},
		&PartialObjectMetadata{},
		&PartialObjectMetadataList{},
	)

	return nil
}

// RegisterConversions adds conversion functions to the given scheme.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*PartialObjectMetadataList)(nil), (*slim_metav1.PartialObjectMetadataList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PartialObjectMetadataList_To_v1_PartialObjectMetadataList(a.(*PartialObjectMetadataList), b.(*slim_metav1.PartialObjectMetadataList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*slim_metav1.PartialObjectMetadataList)(nil), (*PartialObjectMetadataList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PartialObjectMetadataList_To_v1beta1_PartialObjectMetadataList(a.(*slim_metav1.PartialObjectMetadataList), b.(*PartialObjectMetadataList), scope)
	}); err != nil {
		return err
	}
	return nil
}
