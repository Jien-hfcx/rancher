package cis

const (
	NumberOfRetriesForConfigMapCreate = 3
	NumberOfRetriesForClusterUpdate   = 3
	NumberOfRetriesForClusterGet      = 10
	NumberOfRetriesForPodRemoval      = 50
	RetryIntervalInMilliseconds       = 100
	ConfigFileName                    = "config.json"
	CurrentBenchmarkKey               = "current"

	creatorIDAnno = "field.cattle.io/creatorId"
)
