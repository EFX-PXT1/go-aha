/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aha

type FeaturesResponse struct {
	Features   []FeatureMeta `json:"features,omitempty"`
	Pagination Pagination    `json:"pagination,omitempty"`
}
