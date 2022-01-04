// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.com/clivern/mint/core/model"
	"gitlab.com/clivern/mint/core/module"
	"gitlab.com/clivern/mint/pkg"

	"github.com/franela/goblin"
	"github.com/labstack/echo/v4"
)

// TestIntegrationURLInfoEndpoint
func TestIntegrationURLInfoEndpoint(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#TestURLInfoEndpoint", func() {
		g.It("validate not found url", func() {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/urlinfo/1/example.com:443/valid-item/item?q=hello", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/urlinfo/1/*")

			err := GetURLInfo(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})

		g.It("validate found url", func() {
			database.CreateURL(&model.URL{
				Hostname: "example.com",
				Port:     443,
				URI:      "/articles/golang-getting-started?ref=10293883&hj=35",
			})

			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/urlinfo/1/example.com:443/articles/golang-getting-started?ref=10293883&hj=35", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/urlinfo/1/*")

			err := GetURLInfo(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
		})

		g.It("validate invalid request", func() {
			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/urlinfo/1/example.com:kl/valid-item/item?q=hello", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/urlinfo/1/*")

			err := GetURLInfo(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
		})
	})
}

// TestIntegrationURLEndpoint
func TestIntegrationURLEndpoint(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#TestCreateURL", func() {
		g.It("validate invalid request (empty hostname)", func() {
			url := &model.URL{
				Hostname: "",
				Port:     443,
				URI:      "/articles/golang-getting-started?ref=10293883&hj=35",
			}

			e := echo.New()
			out, err := url.ConvertToJSON()

			g.Assert(err).Equal(nil)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(out))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/url")

			err = CreateURL(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
		})

		g.It("validate invalid request (empty port)", func() {
			url := &model.URL{
				Hostname: "example.com",
				Port:     0,
				URI:      "/articles/golang-getting-started?ref=10293883&hj=35",
			}

			e := echo.New()
			out, err := url.ConvertToJSON()

			g.Assert(err).Equal(nil)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(out))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/url")

			err = CreateURL(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
		})

		g.It("validate valid request", func() {
			url := &model.URL{
				Hostname: "example.com",
				Port:     443,
				URI:      "/articles/golang-getting-started?ref=10293883&hj=35",
			}

			e := echo.New()
			out, err := url.ConvertToJSON()

			g.Assert(err).Equal(nil)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(out))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/url")

			err = CreateURL(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusCreated)
		})
	})
}
