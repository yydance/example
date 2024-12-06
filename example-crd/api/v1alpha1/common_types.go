package v1alpha1

import (
	"k8s.io/apimachinery/pkg/types"
)

var OutputService []string

type OutputServiceSpec struct {
	types.NamespacedName `json:"namespacedName"`
	PortName             string `json:"portName,omitempty"`
	TargetPort           int32  `json:"targetPort,omitempty"`
	Port                 int32  `json:"port,omitempty"`
}
