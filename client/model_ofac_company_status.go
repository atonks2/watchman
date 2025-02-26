/*
 * Watchman API
 *
 * Moov Watchman is an HTTP API and Go library to download, parse and offer search functions over numerous trade sanction lists from the United States, European Union governments, agencies, and non profits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// OfacCompanyStatus Status properties of an OFAC Company
type OfacCompanyStatus struct {
	// User ID provided when updating status
	UserID string `json:"userID,omitempty"`
	// Optional note from updating status
	Note string `json:"note,omitempty"`
	// Manually applied status for OFAC Company
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
