package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	GRPCAddress   string `mapstructure:"GRPC_ADDRESS"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	MigrationURL  string `mapstructure:"MIGRATION_URL"`
	GinMode       string `mapstructure:"GIN_MODE"`
	StripeKey     string `mapstructure:"STRIPE_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
