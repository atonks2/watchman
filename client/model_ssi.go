/*
 * Watchman API
 *
 * Moov Watchman is an HTTP API and Go library to download, parse and offer search functions over numerous trade sanction lists from the United States, European Union governments, agencies, and non profits for complying with regional laws. Also included is a web UI and async webhook notification service to initiate processes on remote systems.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Ssi Treasury Department Sectoral Sanctions Identifications List (SSI)
type Ssi struct {
	// The ID assigned to an entity by the Treasury Department
	EntityID string `json:"entityID,omitempty"`
	// Entity type (e.g. individual, vessel, aircraft, etc)
	Type string `json:"type,omitempty"`
	// Sanction programs for which the entity is flagged
	Programs []string `json:"programs,omitempty"`
	// The name of the entity
	Name string `json:"name,omitempty"`
	// Addresses associated with the entity
	Addresses []string `json:"addresses,omitempty"`
	// Additional details regarding the entity
	Remarks []string `json:"remarks,omitempty"`
	// Known aliases associated with the entity
	AlternateNames []string `json:"alternateNames,omitempty"`
	// IDs on file for the entity
	Ids []string `json:"ids,omitempty"`
	// The link to the official SSI list
	SourceListURL string `json:"sourceListURL,omitempty"`
	// The link for information regarding the source
	SourceInfoURL string `json:"sourceInfoURL,omitempty"`
}
