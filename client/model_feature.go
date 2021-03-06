/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aha

import (
	"time"
)

type Feature struct {
	Id           string    `json:"id,omitempty"`
	ReferenceNum string    `json:"reference_num,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	// Start date in YYYY-MM-DD format.
	StartDate string `json:"start_date,omitempty"`
	// Due date in YYYY-MM-DD format.
	DueDate  string   `json:"due_date,omitempty"`
	Url      string   `json:"url,omitempty"`
	Resource string   `json:"resource,omitempty"`
	Release  Release  `json:"release,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}
