package markers

import "sigs.k8s.io/controller-tools/pkg/markers"

const OpenShiftFeatureSetMarkerName = "openshift:enable:featureSet"

type OpenShiftFeatureSet struct {
	FeatureSet string `marker:""`
}

func init() {
	FieldOnlyMarkers = append(FieldOnlyMarkers,
		must(markers.MakeDefinition(OpenShiftFeatureSetMarkerName, markers.DescribesField, OpenShiftFeatureSet{})).
			WithHelp(markers.SimpleHelp("OpenShift", "specifies the FeatureSet that is required to generate this field.")),
	)
}
