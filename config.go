// config/config.go

package config

import "github.com/spf13/viper"

// AppConfig holds the application's configuration settings.
type AppConfig struct {
	AppName  string
	AppPort  int
	Database DatabaseConfig
	// Add other configuration settings here
}

// DatabaseConfig holds database-related configuration settings.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// LoadConfig loads application configuration from a configuration file.
func LoadConfig() (AppConfig, error) {
	var config AppConfig
	viper.SetConfigName("config") // Name of your configuration file (e.g., config.yaml)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
