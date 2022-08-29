package schemapatcher

import (
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kyaml "sigs.k8s.io/yaml"
)

var RequiredFeatureSet = os.Getenv("OPENSHIFT_REQUIRED_FEATURESET")

// mayHandleFile returns true if this manifest should progress past the file collection stage.
// Currently, the only check is the feature-set annotation.
func mayHandleFile(filename string, rawContent []byte) bool {
	manifest := &unstructured.Unstructured{}
	if err := kyaml.Unmarshal(rawContent, &manifest); err != nil {
		return true
	}

	manifestFeatureSet := manifest.GetAnnotations()["release.openshift.io/feature-set"]
	if len(RequiredFeatureSet) == 0 {
		if len(manifestFeatureSet) == 0 {
			return true
		}
		return false
	}

	return manifestFeatureSet == RequiredFeatureSet
}
