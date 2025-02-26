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

// Search Search results containing SDNs, alternate names and/or addreses
type Search struct {
	SDNs              []OfacSdn           `json:"SDNs,omitempty"`
	AltNames          []OfacAlt           `json:"altNames,omitempty"`
	Addresses         []OfacEntityAddress `json:"addresses,omitempty"`
	SectoralSanctions []Ssi               `json:"sectoralSanctions,omitempty"`
	DeniedPersons     []Dpl               `json:"deniedPersons,omitempty"`
	BisEntities       []BisEntities       `json:"bisEntities,omitempty"`
	RefreshedAt       time.Time           `json:"refreshedAt,omitempty"`
}
