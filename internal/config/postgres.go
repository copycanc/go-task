package config

type PostgresCfg struct {
	User string
	Pass string
	DB   string
	Host string
}

// TODO:Прокинуть конфиг через переменные среды, учесть при старте контейнера и запуска приложения
func NewPostgresSQLCfg() *PostgresCfg {
	return &PostgresCfg{
		User: "app",
		Pass: "app",
		DB:   "app",
		Host: "localhost",
	}
}
