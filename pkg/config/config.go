package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	Db_Port      string `mapstructure:"DB_PORT"`
	Db_Host      string `mapstructure:"DB_HOST"`
	Db_User      string `mapstructure:"DB_USER"`
	Db_Password  string `mapstructure:"DB_PASSWORD"`
	Db_Name      string `mapstructure:"DB_NAME"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
