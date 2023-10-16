package configs

import "github.com/spf13/viper"

type AppConfig struct {
	APP_HOST          string `mapstructure:"APP_HOST"`
	APP_GRPC_HOST     string `mapstructure:"APP_GRPC_HOST"`
	APP_KAFKA_HOST    string `mapstructure:"APP_KAFKA_HOST"`
	DATABASE_NAME     string `mapstructure:"DATABASE_NAME"`
	DATABASE_ADDRESS  string `mapstructure:"DATABASE_ADDRESS"`
	DATABASE_USERNAME string `mapstructure:"DATABASE_USERNAME"`
	DATABASE_PASSWORD string `mapstructure:"DATABASE_PASSWORD"`
	KAFKA_TOPIC       string `mapstructure:"KAFKA_TOPIC"`
	KAFKA_GROUP       string `mapstructure:"KAFKA_GROUP"`
	REDIS_HOST        string `mapstructure:"REDIS_HOST"`
	REDIS_PASSWORD    string `mapstructure:"REDIS_PASSWORD"`
	SECRET_KEY        string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) (config AppConfig, errorReadConfig error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	errorReadConfig = viper.ReadInConfig()
	if errorReadConfig != nil {
		return
	}
	errorReadConfig = viper.Unmarshal(&config)
	return
}
