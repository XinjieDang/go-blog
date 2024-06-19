package configs

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	PassWord string `mapstructure:"pass_word"`
	DataBase string `mapstructure:"data_base"`
}
