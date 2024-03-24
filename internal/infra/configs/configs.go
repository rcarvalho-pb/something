package configs

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	Port         uint16 `mapstructure:"PORT"`
	DBConnString string `mapstructure:"DB_CONN_STRING"`
	SecretKey    string `mapstructure:"JWT_SECRET_KEY"`
	HashKey      string `mapstructure:"HASH_KEY"`
	BlockKey     string `mapstructure:"BLOCK_KEY"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading env file: ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("error parsing config values: ", err)
	}

	return
}
