package configs

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	PassWord string `mapstructure:"pass_word"`
	Db       int    `mapstructure:"db"`
}
