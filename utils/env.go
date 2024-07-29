package utils

import (
	"github.com/spf13/viper"
)

// DBConfig holds database configuration values
type DBConfig struct {
	Host string `mapstructure:"DB_HOST"`
	Port string `mapstructure:"DB_PORT"`
	User string `mapstructure:"DB_USER"`
	Pass string `mapstructure:"DB_PASSWORD"`
	Name string `mapstructure:"DB_NAME"`
}

type JWTConfig struct {
	Secret        string `mapstructure:"JWT_SECRET"`
	SecretRefresh string `mapstructure:"JWT_SECRET_REFRESH"`
}

// Env has environment stored
type Env struct {
	ServerPort  string    `mapstructure:"SERVER_PORT"`
	Environment string    `mapstructure:"ENV"`
	DB          DBConfig  `mapstructure:",squash"`
	JWT         JWTConfig `mapstructure:",squash"`
}

// NewEnv creates a new environment
func NewEnv(log Logger) Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Infof("%+v \n", env)
	return env
}
