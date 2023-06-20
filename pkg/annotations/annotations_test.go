package annotations

import (
	vcapi "github.com/ibm/varnish-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

const (
	annotationVarnishClusterName      = "varnish-cluster-name"
	annotationVarnishClusterNamespace = "varnish-cluster-namespace"
)

func TestParseCRCRDAnnotations(t *testing.T) {
	rawAnnotation := "argocd.argoproj.io/compare-options=IgnoreExtraneous,"
	parsed := parseCRCRDAnnotations(rawAnnotation)
	if parsed["argocd.argoproj.io/compare-options"] != "IgnoreExtraneous" {
		t.Error("failed to parse annotations!")
	}
}

func TestAddSpecificCRCRDAnnotations(t *testing.T) {
	originalAnnotations := map[string]string{
		annotationVarnishClusterNamespace: "mockNamespace",
		annotationVarnishClusterName:      "mockName",
	}
	mockInstance := vcapi.VarnishCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mock",
			Annotations: map[string]string{
				"cr-crd-annotations": "argocd.argoproj.io/compare-options=IgnoreExtraneous,",
			},
		},
	}
	value := AddSpecificCRCRDAnnotations(originalAnnotations, &mockInstance)
	addedAnnotations := parseCRCRDAnnotations(mockInstance.Annotations["cr-crd-annotations"])

	for _, addedAnnotation := range mockInstance.Annotations {
		if value[addedAnnotation] != addedAnnotations[addedAnnotation] {
			t.Error("merge annotations failed")
		}
	}
}
