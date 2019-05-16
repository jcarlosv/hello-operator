// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.Carlos":       schema_pkg_apis_carlos_v1alpha1_Carlos(ref),
		"github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosSpec":   schema_pkg_apis_carlos_v1alpha1_CarlosSpec(ref),
		"github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosStatus": schema_pkg_apis_carlos_v1alpha1_CarlosStatus(ref),
	}
}

func schema_pkg_apis_carlos_v1alpha1_Carlos(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Carlos is the Schema for the carlos API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosSpec", "github.com/jcarlosv/hello-operator/pkg/apis/carlos/v1alpha1.CarlosStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_carlos_v1alpha1_CarlosSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CarlosSpec defines the desired state of Carlos",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_carlos_v1alpha1_CarlosStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CarlosStatus defines the observed state of Carlos",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
