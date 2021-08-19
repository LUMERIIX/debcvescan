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
	"testing"
)

func TestWhitelistNew(t *testing.T) {
	NewWhitelist()
}

func TestWhitelistAddRemove(t *testing.T) {
	whitelist := NewWhitelist()
	whitelist.Add("CVE-1245", "no risk")
	whitelist.Remove("CVE-1245")
}

func TestWhitelistRemovNotExists(t *testing.T) {
	whitelist := NewWhitelist()
	whitelist.Add("CVE-1245", "no risk")
	whitelist.Remove("CVE-15")
}

func TestIsWhitelistedExists(t *testing.T) {
	whitelist := NewWhitelist()
	whitelist.Add("CVE-1245", "no risk")
	if !whitelist.IsWhitelisted("CVE-1245") {
		t.Fail()
	}
}

func TestIsWhitelistedNotExists(t *testing.T) {
	whitelist := NewWhitelist()
	whitelist.Add("CVE-1245", "no risk")
	if whitelist.IsWhitelisted("CVE-0000") {
		t.Fail()
	}
}