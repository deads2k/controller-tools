package crd

import (
	"fmt"
	"os"

	crdmarkers "sigs.k8s.io/controller-tools/pkg/crd/markers"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var RequiredFeatureSet = os.Getenv("OPENSHIFT_REQUIRED_FEATURESET")

// mayHandleField returns true if the field should be considered by this invocation of the generator.
// Right now, the only sip is based on the featureset marker.
func mayHandleField(field markers.FieldInfo) bool {
	uncastFeatureSet := field.Markers.Get(crdmarkers.OpenShiftFeatureSetMarkerName)
	if uncastFeatureSet == nil {
		return true
	}

	featureSetMarker, ok := uncastFeatureSet.(crdmarkers.OpenShiftFeatureSet)
	if !ok {
		panic(fmt.Sprintf("actually got %t", uncastFeatureSet))
	}
	return featureSetMarker.FeatureSet == RequiredFeatureSet
}
