package config

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/objx"
	"io/ioutil"
	"sync"
)

// Config represents a configuration object driven by a JSON file.
type Config struct {
	data objx.Map

	dataInitOnce sync.Once
}

// Load loads and parses a JSON configuration file.
func (c *Config) Load(filename string) error {

	rawBytes, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		return readErr
	}

	bytes := make([]byte, 0, len(rawBytes))

	skip := false
	for i := 0; i < len(rawBytes); i++ {
		if skip == true {
			if rawBytes[i] == '\n' {
				skip = false
			} else {
				continue
			}
		}
		if rawBytes[i] == '#' {
			skip = true
			continue
		}
		bytes = append(bytes, rawBytes[i])
	}

	fmt.Printf("%#v\n", string(bytes))

	parseErr := c.Parse(bytes)

	if parseErr != nil {
		return parseErr
	}

	// no errors
	return nil
}

// Parse parses the JSON bytes into the data object.  Multiple calls to Parse will
// merge the data fields together.
func (c *Config) Parse(data []byte) error {

	var newData objx.Map
	unmarshalErr := json.Unmarshal(data, &newData)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	// merge the data
	c.data = c.Data().Merge(newData)

	// no errors
	return nil
}

// Data gets the raw objx.Map data object for this configuration object.
func (c *Config) Data() objx.Map {

	c.dataInitOnce.Do(func() {
		c.data = make(objx.Map)
	})

	return c.data
}

// Gets a value from the config.  Key paths are supported.
func (c *Config) Get(keypath string) interface{} {
	return c.Data().Get(keypath).Data()
}
