/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aha

type Pagination struct {
	TotalRecords int64 `json:"total_records,omitempty"`
	TotalPages   int64 `json:"total_pages,omitempty"`
	CurrentPage  int64 `json:"current_page,omitempty"`
}