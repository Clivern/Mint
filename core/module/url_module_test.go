// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"testing"

	"gitlab.com/clivern/mint/core/model"
	"gitlab.com/clivern/mint/pkg"

	"github.com/franela/goblin"
)

// TestUnitURLCRUD
func TestUnitURLCRUD(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#TestCreateURL", func() {
		g.It("It should satisfy test cases", func() {
			url := database.CreateURL(&model.URL{
				Hostname: "example.com",
				Port:     443,
				URI:      "/tutorials/getting-started-with-ruby?ref=233&fd=34",
			})

			g.Assert(url.ID == 0).Equal(false)
		})
	})

	g.Describe("#TestGetURL", func() {
		g.It("It should satisfy test cases", func() {
			url := database.GetURL(
				"example.com",
				443,
				"/tutorials/getting-started-with-ruby?ref=233&fd=34",
			)

			g.Assert(url.ID == 0).Equal(false)
			g.Assert(url.Hostname).Equal("example.com")
			g.Assert(url.Port).Equal(443)
			g.Assert(url.URI).Equal("/tutorials/getting-started-with-ruby?ref=233&fd=34")
		})
	})
}
