// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     CRDTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package crd

import (
	_ "embed"

	"github.com/grafana/grafana/pkg/kinds/serviceaccount"
	"github.com/grafana/grafana/pkg/kindsys/k8ssys"
)

// The CRD YAML representation of the ServiceAccount kind.
//
//go:embed serviceaccount.crd.yml
var CRDYaml []byte

// ServiceAccount is the Go CRD representation of a single ServiceAccount object.
// It implements [runtime.Object], and is used in k8s scheme construction.
type ServiceAccount struct {
	k8ssys.Base[serviceaccount.ServiceAccount]
}

// ServiceAccountList is the Go CRD representation of a list ServiceAccount objects.
// It implements [runtime.Object], and is used in k8s scheme construction.
type ServiceAccountList struct {
	k8ssys.ListBase[serviceaccount.ServiceAccount]
}