package infrastructure

type App struct {
	Port int64 `mapstructure:"port" default:"7100"`
}

type Config struct {
	App App `mapstructure:"app"`
}
