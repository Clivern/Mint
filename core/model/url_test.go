// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitURL
func TestUnitURL(t *testing.T) {
	g := goblin.Goblin(t)

	url1 := &URL{
		ID:       1,
		Hostname: "example.com",
		Port:     443,
		URI:      "/golang-getting-started?ref=23",
	}

	url2 := &URL{
		ID:       2,
		Hostname: "example.com",
		Port:     80,
		URI:      "/ruby-getting-started?ref=23",
	}

	url3 := &URL{}

	g.Describe("#TestLoadFromJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := url1.ConvertToJSON()

			g.Assert(out).Equal(`{"id":1,"hostname":"example.com","port":443,"uri":"/golang-getting-started?ref=23","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}`)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#TestConvertToJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := url2.ConvertToJSON()

			g.Assert(err).Equal(nil)
			url3.LoadFromJSON([]byte(out))

			g.Assert(url3.ID).Equal(2)
			g.Assert(url3.Hostname).Equal("example.com")
			g.Assert(url3.Port).Equal(80)
			g.Assert(url3.URI).Equal("/ruby-getting-started?ref=23")
		})
	})
}
