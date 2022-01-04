// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gitlab.com/clivern/mint/core/model"
	"gitlab.com/clivern/mint/core/util"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// GetURLInfo controller
func GetURLInfo(c echo.Context, gc *GlobalContext) error {
	err := gc.GetDatabase().Ping()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Unable to connect to the database`)

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errorMessage": "Internal Server Error!",
		})
	}

	path := strings.TrimPrefix(c.Request().URL.Path, "/urlinfo/1/")
	pathItems := strings.SplitN(path, "/", 2)
	hostInfo := strings.Split(pathItems[0], ":")

	hostname := hostInfo[0]
	port, err := strconv.Atoi(hostInfo[1])
	uri := fmt.Sprintf("/%s?%s", pathItems[1], c.Request().URL.RawQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Invalid port value %s: %s", hostInfo[1], err.Error()),
		}).Debug(`Invalid Request`)

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": "Error! HTTP port is invalid",
		})
	}

	log.WithFields(log.Fields{
		"hostname": hostname,
		"port":     port,
		"uri":      uri,
	}).Info(`Search for a URL`)

	url := gc.GetDatabase().GetURL(hostname, port, uri)

	if url.ID > 0 {
		log.WithFields(log.Fields{
			"hostname": hostname,
			"port":     port,
			"uri":      uri,
		}).Info(`Found the URL in the database`)

		return c.NoContent(http.StatusOK)
	}

	log.WithFields(log.Fields{
		"hostname": hostname,
		"port":     port,
		"uri":      uri,
	}).Info(`URL is missing`)

	return c.NoContent(http.StatusNotFound)
}

// CreateURL controller
func CreateURL(c echo.Context, gc *GlobalContext) error {
	err := gc.GetDatabase().Ping()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Unable to connect to the database`)

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"errorMessage": "Internal Server Error!",
		})
	}

	data, _ := ioutil.ReadAll(c.Request().Body)

	url := &model.URL{}
	url.LoadFromJSON(data)

	// Validate Hostname
	if url.Hostname == "" || !util.IsHostnameRFC1123(url.Hostname) {
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Invalid Hostname: %s", url.Hostname),
		}).Debug(`Invalid Request`)

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": "Error! Hostname is required",
		})
	}

	// Validate HTTP Port
	if url.Port == 0 {
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Invalid Port: %d", url.Port),
		}).Debug(`Invalid Request`)

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"errorMessage": "Error! HTTP port is required",
		})
	}

	// Create a new item
	url = gc.GetDatabase().CreateURL(url)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":        url.ID,
		"hostname":  url.Hostname,
		"port":      url.Port,
		"uri":       url.URI,
		"createdAt": url.CreatedAt,
		"updatedAt": url.UpdatedAt,
	})
}
