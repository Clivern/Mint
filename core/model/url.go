// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"time"

	"gitlab.com/clivern/mint/core/util"
)

// URL struct
type URL struct {
	ID        int       `json:"id"`
	Hostname  string    `json:"hostname"`
	Port      int       `json:"port"`
	URI       string    `json:"uri"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// LoadFromJSON update object from json
func (a *URL) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *URL) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}
