package config

const GRPCAddr = ":8183"
const HTTPAddr = ":8180"

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

func GetConfig() *Config {

	return &Config{
		DB: DBConfig{
			User:     "todoadmin",
			Password: "tododo",
			Host:     "postgres",
			Port:     5432,
			Dbname:   "tododb",
			Sslmode:  "",
		},
	}
}
