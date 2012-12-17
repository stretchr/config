package config

// Load loads a configuration file and returns a Config object.
func Load(filename string) (*Config, error) {
	config := new(Config)
	loadErr := config.Load(filename)
	return config, loadErr
}
