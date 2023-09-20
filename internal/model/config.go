package model

type Config struct {
	LogJSON bool `env:"LOG_JSON"`

	Psql
	App
	Redis
}

type App struct {
	Version     string `env:"APP_VERSION"`
	Environment string `env:"APP_ENVIRONMENT"`

	ServerIP   string `env:"SERVER_IP"`
	ServerPort string `env:"SERVER_PORT"`
}

type Psql struct {
	PsqlHost     string `env:"PSQL_HOST"`
	PsqlPort     string `env:"PSQL_PORT"`
	PsqlUser     string `env:"PSQL_USER"`
	PsqlPassword string `env:"PSQL_PASSWORD"`
	PsqlDatabase string `env:"PSQL_DATABASE"`
	PsqlSSL      string `env:"PSQL_SSL"`
}

type Redis struct {
	Prefix   string `env:"REDIS_PREFIX"`
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DATABASE"`
}
