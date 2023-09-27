package config

import "github.com/spf13/viper"

type Config struct {
	UserServiceRestPort string `mapstructure:"USER_SERVICE_REST_PORT"`

	UserServiceGrpcHost string `mapstructure:"USER_SERVICE_GRPC_HOST"`
	UserServiceGrpcPort string `mapstructure:"USER_SERVICE_GRPC_PORT"`

	DBHost     string `mapstructure:"USER_DB_HOST"`
	DBPort     string `mapstructure:"USER_DB_PORT"`
	DBUser     string `mapstructure:"USER_DB_USER"`
	DBPassword string `mapstructure:"USER_DB_PASSWORD"`
	DBName     string `mapstructure:"USER_DB_NAME"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
}

var envs = []string{
	"USER_DB_HOST", "USER_DB_PORT", "USER_DB_USER", "USER_DB_PASSWORD", "USER_DB_NAME",
	"USER_SERVICE_REST_PORT",
	"USER_SERVICE_GRPC_PORT", "USER_SERVICE_GRPC_HOST",
	"REDIS_HOST", "REDIS_PORT",
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {

		if err = viper.BindEnv(env); err != nil {
			return Config{}, err
		}
	}

	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
