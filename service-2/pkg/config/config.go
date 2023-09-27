package config

import "github.com/spf13/viper"

type Config struct {
	Service2RestPort string `mapstructure:"SERVICE2_REST_PORT"`

	UserServiceGrpcHost string `mapstructure:"USER_SERVICE_GRPC_HOST"`
	UserServiceGrpcPort string `mapstructure:"USER_SERVICE_GRPC_PORT"`
}

var envs = []string{
	"SERVICE2_REST_PORT",
	"USER_SERVICE_GRPC_HOST", "USER_SERVICE_GRPC_PORT",
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
