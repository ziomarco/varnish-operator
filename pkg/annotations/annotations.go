package annotations

import (
	vcapi "github.com/ibm/varnish-operator/api/v1alpha1"
)

func MergeAnnotations(annotations map[string]string, instance *vcapi.VarnishCluster) map[string]string {
	m := make(map[string]string, len(instance.Annotations))
	for k, v := range instance.Annotations {
		m[k] = v
	}
	return m
}
