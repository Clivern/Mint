// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitDSN
func TestUnitDSN(t *testing.T) {
	g := goblin.Goblin(t)

	dsn1 := &DSN{
		Driver:   "mysql",
		Username: "root",
		Password: "root",
		Hostname: "localhost",
		Port:     3306,
		Name:     "mint.db",
	}

	dsn2 := &DSN{
		Driver:   "sqlite3",
		Username: "root",
		Password: "root",
		Hostname: "localhost",
		Port:     3306,
		Name:     "mint.db",
	}

	dsn3 := &DSN{}

	g.Describe("#TestToString", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(dsn1.ToString()).Equal("root:root@tcp(localhost:3306)/mint.db?charset=utf8&parseTime=True")
			g.Assert(dsn2.ToString()).Equal("mint.db")
		})
	})

	g.Describe("#TestLoadFromJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := dsn1.ConvertToJSON()

			g.Assert(out).Equal(`{"driver":"mysql","username":"root","password":"root","hostname":"localhost","port":3306,"name":"mint.db"}`)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#TestConvertToJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, _ := dsn1.ConvertToJSON()
			dsn3.LoadFromJSON([]byte(out))

			g.Assert(dsn3.Driver).Equal("mysql")
			g.Assert(dsn3.Username).Equal("root")
			g.Assert(dsn3.Password).Equal("root")
			g.Assert(dsn3.Hostname).Equal("localhost")
			g.Assert(dsn3.Name).Equal("mint.db")
		})
	})
}
