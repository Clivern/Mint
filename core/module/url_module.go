// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"gitlab.com/clivern/mint/core/model"
)

// CreateURL creates a new entity
func (db *Database) CreateURL(url *model.URL) *model.URL {
	db.Connection.Create(url)

	return url
}

// GetURL gets an entity by hostname, port and uri
func (db *Database) GetURL(hostname string, port int, uri string) model.URL {
	url := model.URL{}

	db.Connection.Where(
		"hostname = ? AND port = ? AND uri = ?",
		hostname,
		port,
		uri,
	).First(&url)

	return url
}
