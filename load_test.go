package config

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {

	config, loadErr := Load("./test_files/test_config.json")

	if assert.Nil(t, loadErr) {
		assert.Equal(t, "Mat", config.Data()["name"])
	}

}
