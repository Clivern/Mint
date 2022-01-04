// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"gitlab.com/clivern/mint/core/module"
)

// GlobalContext type
type GlobalContext struct {
	Database *module.Database
}

// GetDatabase gets a database connection
func (c *GlobalContext) GetDatabase() *module.Database {
	return c.Database
}
