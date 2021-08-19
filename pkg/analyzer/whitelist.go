// Package analyzer Debian CVE Tracker Analyzer
// Copyright 2019 debcvescan authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package analyzer

import (
	"encoding/json"
	"os"
)

// WhitelistEntry represents an entry in the whitelist
type WhitelistEntry struct {
	CVE           string `json:"cve"`
	Justification string `json:"justification"`
}

// object struct for the whitelist
type WhiteList struct {
	Whitelisted []WhitelistEntry `json:"whitelisted"`
}

// NewWhiteList creates a new instance of the WhiteList class
func NewWhitelist() WhiteList {
	result := WhiteList{}
	file, err := os.Open("./debcvescan.whitelist")
	if err == nil {
		err = json.NewDecoder(file).Decode(&result.Whitelisted)
		if err != nil {
			println("Warning: 'debcvescan.whitelist' is not a valid json array")
		}

	}

	return result
}

// Add adds a new CVE whitelist entry together with justification to 'debcvescan.whitelist' file and saves it
func (s *WhiteList) Add(cve string, justification string) {
	s.Whitelisted = append(s.Whitelisted, WhitelistEntry{CVE: cve, Justification: justification})
	s.save()
}

// Remove emoves the given CVE entry from the 'debcvescan.whitelist' file and saves ii
func (s *WhiteList) Remove(cve string) {
	idx := s.findIndex(cve)
	if idx >= 0 {
		s.removeIndex(idx)
		s.save()
	}

}

// IsWhitelisted checks if th given CVE is whitelisted or not
func (s *WhiteList) IsWhitelisted(cve string) bool {
	return s.findIndex(cve) >= 0
}

// findIndex helper function to find the array index for the given CVE
func (s *WhiteList) findIndex(cve string) int {
	for i := 0; i < len(s.Whitelisted); i++ {
		if s.Whitelisted[i].CVE == cve {
			return i
		}
	}

	return -1
}

// removeIndex removes the entry on the given index position
func (s *WhiteList) removeIndex(idx int) {
	newArray := make([]WhitelistEntry, (len(s.Whitelisted) - 1))
	k := 0
	for i := 0; i < len(s.Whitelisted); i++ {
		if i != idx {
			newArray[k] = s.Whitelisted[i]
			k++
		}
	}
	s.Whitelisted = newArray

}

// save persists the whitelist to the './debcvescan.whitelist' file
func (s *WhiteList) save() {
	json, err := json.MarshalIndent(s.Whitelisted, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./debcvescan.whitelist", json, 0600)
	if err != nil {
		panic(err)
	}
}
