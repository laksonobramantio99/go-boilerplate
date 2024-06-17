package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Env      string
	Port     int
	Database DatabaseConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

// global config variable
var Config Configuration

func InitConfig(env string) (err error) {
	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	Config.Env = env
	return nil
}
