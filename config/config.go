package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		DB        `yaml:"db"`
		SecretKey string `env:"SECRET_KEY" env-default:"LgszAIV5dBwmy93buOvFDSvmXaam1rzjNAHQ9HJNizE"`
	}

	DB struct {
		Host     string `env:"DB_HOST" env-default:"localhost"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
		Name     string `env:"DB_NAME" env-default:"todo"`
		Port     string `env:"DB_PORT" env-default:"5432"`
	}
)

var Cfg *Config

func New() error {
	Cfg = &Config{}

	if err := cleanenv.ReadEnv(Cfg); err != nil {
		return err
	}

	return nil
}
