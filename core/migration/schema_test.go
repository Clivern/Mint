// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package migration

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitSchema
func TestUnitSchema(t *testing.T) {
	g := goblin.Goblin(t)

	lookup := &URL{
		Hostname: "example.com",
		Port:     443,
		URI:      "bad?q=scam",
	}

	g.Describe("#TestConvertToJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := lookup.ConvertToJSON()

			g.Assert(out != "").Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
