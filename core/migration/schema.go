// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package migration

import (
	"gitlab.com/clivern/mint/core/util"

	"github.com/jinzhu/gorm"
)

// URL struct
type URL struct {
	gorm.Model

	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	URI      string `json:"uri"`
}

// LoadFromJSON update object from json
func (a *URL) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *URL) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}
