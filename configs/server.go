package configs

type Server struct {
	Port    string `mapstructure:"port"`
	GinMode string `mapstructure:"gin_mode"`
}
