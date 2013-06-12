package config

import (
	"encoding/json"
	"github.com/stretchr/stew/objects"
	"io/ioutil"
	"sync"
)

// Config represents a configuration object driven by a JSON file.
type Config struct {
	data objects.Map

	dataInitOnce sync.Once
}

// Load loads and parses a JSON configuration file.
func (c *Config) Load(filename string) error {

	bytes, readErr := ioutil.ReadFile(filename)

	if readErr != nil {
		return readErr
	}

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

	var newData objects.Map
	unmarshalErr := json.Unmarshal(data, &newData)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	// merge the data
	c.data = c.Data().Merge(newData)

	// no errors
	return nil
}

// Data gets the raw objects.Map data object for this configuration object.
func (c *Config) Data() objects.Map {

	c.dataInitOnce.Do(func() {
		c.data = make(objects.Map)
	})

	return c.data
}

// Gets a value from the config.  Key paths are supported.
func (c *Config) Get(keypath string) interface{} {
	return c.Data().Get(keypath)
}
