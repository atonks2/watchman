// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ofac

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Read will consume the file at path and attempt to parse it was a CSV OFAC file.
//
// For more details on the raw OFAC files see https://docs.moov.io/watchman/file-structure/
func Read(path string) (*Results, error) {
	switch filepath.Base(path) {
	case "add.csv":
		return csvAddressFile(path)
	case "alt.csv":
		return csvAlternateIdentityFile(path)
	case "sdn.csv":
		return csvSDNFile(path)
	case "sdn_comments.csv":
		return csvSDNCommentsFile(path)
	}
	return nil, fmt.Errorf("unknown file %s", path)
}

type Results struct {
	// Addresses returns an array of OFAC Specially Designated National Addresses
	Addresses []*Address `json:"address"`

	// AlternateIdentities returns an array of OFAC Specially Designated National Alternate Identity
	AlternateIdentities []*AlternateIdentity `json:"alternateIdentity"`

	// SDNs returns an array of OFAC Specially Designated Nationals
	SDNs []*SDN `json:"sdn"`

	// SDNComments returns an array of OFAC Specially Designated National Comments
	SDNComments []*SDNComments `json:"sdnComments"`
}

func csvAddressFile(path string) (*Results, error) {
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []*Address

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(Adam): add max line count break here also
			break
		}
		if len(record) != 6 {
			continue
		}

		record = replaceNull(record)
		out = append(out, &Address{
			EntityID:                    record[0],
			AddressID:                   record[1],
			Address:                     record[2],
			CityStateProvincePostalCode: record[3],
			Country:                     record[4],
			AddressRemarks:              record[5],
		})
	}
	return &Results{Addresses: out}, nil
}

func csvAlternateIdentityFile(path string) (*Results, error) {
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []*AlternateIdentity

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(adam)
			break
		}
		if len(record) != 5 {
			continue
		}
		record = replaceNull(record)
		out = append(out, &AlternateIdentity{
			EntityID:         record[0],
			AlternateID:      record[1],
			AlternateType:    record[2],
			AlternateName:    record[3],
			AlternateRemarks: record[4],
		})
	}
	return &Results{AlternateIdentities: out}, nil
}

func csvSDNFile(path string) (*Results, error) {
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []*SDN

	// Read File into a Variable
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil && err == csv.ErrFieldCount {
			continue
		}
		if err == io.EOF { // TODO(Adam): add max line count break here also
			break
		}
		if len(record) != 12 {
			continue
		}
		record = replaceNull(record)
		out = append(out, &SDN{
			EntityID:               record[0],
			SDNName:                record[1],
			SDNType:                record[2],
			Programs:               splitPrograms(record[3]),
			Title:                  record[4],
			CallSign:               record[5],
			VesselType:             record[6],
			Tonnage:                record[7],
			GrossRegisteredTonnage: record[8],
			VesselFlag:             record[9],
			VesselOwner:            record[10],
			Remarks:                record[11],
		})
	}
	return &Results{SDNs: out}, nil
}

func csvSDNCommentsFile(path string) (*Results, error) {
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	// Loop through lines & turn into object
	var out []*SDNComments
	for _, csvLine := range lines {
		csvLine := replaceNull(csvLine)
		out = append(out, &SDNComments{
			EntityID:        csvLine[0],
			RemarksExtended: csvLine[1],
		})
	}
	return &Results{SDNComments: out}, nil
}

// replaceNull replaces a CSV field that contain -0- with "".  Null values for all four formats consist of "-0-"
// (ASCII characters 45, 48, 45).
func replaceNull(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.TrimSpace(strings.Replace(s[i], "-0-", "", -1))
	}
	return s
}

// Some entries in the SDN have a malformed programs list.
// Fields containing lists are supposed to be semicolon delimited, but the programs list doesn't always follow this convention.
// Ex: "SDGT] [IFSR" => "SDGT; IFSR", "SDNTK] [FTO] [SDGT" => "SDNTK; FTO; SDGT"
var prgmReplacer = strings.NewReplacer("] [", "; ", "]", "", "[", "")

func cleanPrgmsList(s string) string {
	return strings.TrimSpace(prgmReplacer.Replace(s))
}

func splitPrograms(in string) []string {
	norm := cleanPrgmsList(in)
	return strings.Split(norm, "; ")
}
