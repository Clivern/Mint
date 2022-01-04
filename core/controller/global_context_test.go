// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"testing"

	"gitlab.com/clivern/mint/core/module"
	"gitlab.com/clivern/mint/pkg"

	"github.com/franela/goblin"
)

// TestUnitGlobalContext
func TestUnitGlobalContext(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#TestGetDatabase", func() {
		g.It("It should satisfy all provided test cases", func() {
			gc := &GlobalContext{Database: database}
			g.Assert(gc.GetDatabase().AutoConnect()).Equal(nil)
			g.Assert(gc.GetDatabase().Ping()).Equal(nil)
		})
	})
}
