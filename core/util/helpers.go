// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"
	"regexp"
)

const (
	// hostnameRegexStringRFC1123 accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	hostnameRegexStringRFC1123 = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`
)

var (
	// hostnameRegexRFC1123 regex
	hostnameRegexRFC1123 = regexp.MustCompile(hostnameRegexStringRFC1123)
)

// LoadFromJSON update object from json
func LoadFromJSON(item interface{}, data []byte) error {
	err := json.Unmarshal(data, &item)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func ConvertToJSON(item interface{}) (string, error) {
	data, err := json.Marshal(&item)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// IsHostnameRFC1123 validates a hostname
func IsHostnameRFC1123(value string) bool {
	return hostnameRegexRFC1123.MatchString(value)
}
