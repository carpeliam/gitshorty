/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// A group of calculated values for this Objective.
type ObjectiveStats struct {
	// The average cycle time (in seconds) of completed stories in this Objective.
	AverageCycleTime int64 `json:"average_cycle_time,omitempty"`
	// The average lead time (in seconds) of completed stories in this Objective.
	AverageLeadTime int64 `json:"average_lead_time,omitempty"`
	// The number of related documents to this Objective.
	NumRelatedDocuments int64 `json:"num_related_documents"`
}
