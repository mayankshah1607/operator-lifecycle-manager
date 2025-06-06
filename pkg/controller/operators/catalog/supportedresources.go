package catalog

const (
	PrometheusRuleKind        = "PrometheusRule"
	ServiceMonitorKind        = "ServiceMonitor"
	PodDisruptionBudgetKind   = "PodDisruptionBudget"
	PriorityClassKind         = "PriorityClass"
	VerticalPodAutoscalerKind = "VerticalPodAutoscaler"
	ConsoleYAMLSampleKind     = "ConsoleYAMLSample"
	ConsoleQuickStartKind     = "ConsoleQuickStart"
	ConsoleCLIDownloadKind    = "ConsoleCLIDownload"
	ConsoleLinkKind           = "ConsoleLink"
	ConsolePlugin             = "ConsolePlugin"
)

var supportedKinds = map[string]struct{}{
	PrometheusRuleKind:        {},
	ServiceMonitorKind:        {},
	PodDisruptionBudgetKind:   {},
	PriorityClassKind:         {},
	VerticalPodAutoscalerKind: {},
	ConsoleYAMLSampleKind:     {},
	ConsoleQuickStartKind:     {},
	ConsoleCLIDownloadKind:    {},
	ConsoleLinkKind:           {},
	ConsolePlugin:             {},
}

// isSupported returns true if OLM supports this type of CustomResource.
func isSupported(kind string) bool {
	_, ok := supportedKinds[kind]
	return ok
}
