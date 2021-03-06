/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aha

type FeatureUpdate struct {
	// Name of the feature
	Name string `json:"name,omitempty"`
	// Description of the feature and it can include HTML formatting.
	Description string `json:"description,omitempty"`
	// Email address of user that created the feature.
	CreatedBy string `json:"created_by,omitempty"`
	// Email address of user that is assigned the feature.
	AssignedToUser string `json:"assigned_to_user,omitempty"`
	// Tags can be automatically assigned to the new feature. If more than one tag is used then tags should be separated by commas
	Tags string `json:"tags,omitempty"`
	// Set the original estimated effort in a text format, you can use d, h, min (or 'p' for points) to indicate the units to use.
	OriginalEstimateText string `json:"original_estimate_text,omitempty"`
	//  Set the remaining estimated effort in a text format, you can use d, h, min (or 'p' for points) to indicate the units to use.
	RemainingEstimateText string `json:"remaining_estimate_text,omitempty"`
	// Date that work will start on the feature in format YYYY-MM-DD.
	StartDate string `json:"start_date,omitempty"`
	// Date that work is due to be completed on the feature in format YYYY-MM-DD.
	DueDate string `json:"due_date,omitempty"`
	// Name or id of release phase which the feature belongs to.
	ReleasePhase string `json:"release_phase,omitempty"`
	// Name or id of initiative which the feature belongs to.
	Initiative string `json:"initiative,omitempty"`
	// Name or id of master feature which the feature belongs to.
	MasterFeature string `json:"master_feature,omitempty"`
}
