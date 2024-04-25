/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	v1alpha1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1"
	v1beta1 "k8s.io/sample-apiserver/pkg/apis/wardle/v1beta1"
	wardlev1alpha1 "k8s.io/sample-apiserver/pkg/generated/applyconfiguration/wardle/v1alpha1"
	wardlev1beta1 "k8s.io/sample-apiserver/pkg/generated/applyconfiguration/wardle/v1beta1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=wardle.example.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Fischer"):
		return &wardlev1alpha1.FischerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Flunder"):
		return &wardlev1alpha1.FlunderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FlunderSpec"):
		return &wardlev1alpha1.FlunderSpecApplyConfiguration{}

		// Group=wardle.example.com, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Flunder"):
		return &wardlev1beta1.FlunderApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlunderSpec"):
		return &wardlev1beta1.FlunderSpecApplyConfiguration{}

	}
	return nil
}
