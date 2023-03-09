package config

import (
	"BootCampT1/logger"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var v *viper.Viper

var c *Config

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()

	env := v.Get("ENV").(string)

	v.SetConfigName(env)
	v.SetConfigType("json")
	v.AddConfigPath("./config/")

	err := v.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	err = v.Unmarshal(&c)
	if err != nil {
		return &Config{}, err
	}

	validate := validator.New()
	if err = validate.Struct(c); err != nil {
		logger.Error.Println("Missing required fields in the struct", err)
		panic("Missing required fields in the struct")
	}

	logger.Info.Println("Succesfully loading config files")
	return c, nil
}

func GetConfig() *Config {
	if c == nil {
		return &Config{}
	}
	return c
}

func GetViper() *viper.Viper {
	return v
}
