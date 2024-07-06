package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Env      string         `mapstructure:"env"`
	Port     int            `mapstructure:"port"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
}

type DatabaseConfig struct {
	Master DBConfig `mapstructure:"master"`
	Slave  DBConfig `mapstructure:"slave"`
}

type DBConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
}

// global config variable
var Config *Configuration

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
