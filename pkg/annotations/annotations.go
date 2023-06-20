package annotations

import (
	vcapi "github.com/ibm/varnish-operator/api/v1alpha1"
	"strings"
)

func AddSpecificCRCRDAnnotations(originalAnnotations map[string]string, instance *vcapi.VarnishCluster) map[string]string {
	rawAnnotations := strings.Split(instance.Annotations["cr-crd-annotations"], ",")
	var parsedAnnotations map[string]string
	for _, annotation := range rawAnnotations {
		annotationValue := strings.Split(annotation, "=")
		if len(annotation) == 0 || len(annotationValue[0]) == 0 || len(annotationValue[1]) == 0 {
			continue
		}
		parsedAnnotations[annotationValue[0]] = annotationValue[1]
	}
	m := make(map[string]string, (len(instance.Annotations))+len(parsedAnnotations))
	for k, v := range originalAnnotations {
		m[k] = v
	}
	for k, v := range parsedAnnotations {
		m[k] = v
	}
	return m
}
