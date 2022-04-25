package config

import (
	"github.com/spf13/viper"
)

// keys for environment configuration
const (
	UserOperationsHost = "user.operations.host"
)

func init() {
	// env var for db
	_ = viper.BindEnv(UserOperationsHost, "USER_OPERATIONS_HOST")

	viper.SetDefault(UserOperationsHost, "localhost:8080")
}
