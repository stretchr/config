package config

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestConfig_Load(t *testing.T) {

	c := new(Config)
	loadErr := c.Load("./test_files/test_config.json")

	if assert.Nil(t, loadErr) {
		assert.Equal(t, "Mat", c.Data()["name"])
	}

}

func TestConfig_Parse(t *testing.T) {

	c := new(Config)
	err := c.Parse([]byte(`{"name":"Mat","age":29,"location":["London, UK", "Boulder, CO"],"nested":{"value":1}}`))

	if assert.Nil(t, err) {
		assert.Equal(t, "Mat", c.Data()["name"])
	}

}

func TestConfig_Parse_Inherited(t *testing.T) {

	c := new(Config)
	err1 := c.Parse([]byte(`{"name":"Mathew","age":29,"location":["London, UK", "Boulder, CO"],"nested":{"value":1}}`))
	err2 := c.Parse([]byte(`{"name":"Mat"}`))

	if assert.Nil(t, err1) && assert.Nil(t, err2) {
		assert.Equal(t, "Mat", c.Data()["name"])
		assert.Equal(t, 29, c.Data()["age"])
	}

}

func TestConfig_Data(t *testing.T) {

	c := new(Config)
	assert.NotNil(t, c.Data())

}
