/*
 * Aha.io API
 *
 * Articles that matter on social publishing platform
 *
 * OpenAPI spec version: 1.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package aha

import (
	"github.com/grokify/gotilla/time/timeutil"
)

type ReleaseUpdate struct {

	// Release name.
	Name string `json:"name,omitempty"`

	// Start date in YYYY-MM-DD format.
	StartDate timeutil.RFC3339YMDTime `json:"start_date,omitempty"`

	// Release date in YYYY-MM-DD format.
	ReleaseDate timeutil.RFC3339YMDTime `json:"release_date,omitempty"`

	// Date Development started in format YYYY-MM-DD
	DevelopmentStartedOn timeutil.RFC3339YMDTime `json:"development_started_on,omitempty"`

	// The external release date for this feature in format YYYY-MM-DD
	ExternalReleaseDate timeutil.RFC3339YMDTime `json:"external_release_date,omitempty"`

	ParkingLot bool `json:"parking_lot,omitempty"`
}