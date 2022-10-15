package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthService    string `mapstructure:"AUTH_SERVICE"`
	ProductService string `mapstructure:"PRODUCT_SERVICE"`
	OrderService   string `mapstructure:"ORDER_SERVICE"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
