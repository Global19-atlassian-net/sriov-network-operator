package common

const (
	// GeneratedByControllerVersionAnnotationKey is used to tag the machineconfigs generated by the controller with the version of the controller.
	GeneratedByControllerVersionAnnotationKey = "machineconfiguration.openshift.io/generated-by-controller-version"

	// ControllerConfigName is the name of the ControllerConfig object that controllers use
	ControllerConfigName = "machine-config-controller"

	// KernelTypeDefault denominates the default kernel type
	KernelTypeDefault = "default"

	// KernelTypeRealtime denominates the realtime kernel type
	KernelTypeRealtime = "realtime"

	// MasterLabel defines the label associated with master node. The master taint uses the same label as taint's key
	MasterLabel = "node-role.kubernetes.io/master"
)