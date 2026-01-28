package postgres

import "go-graphql/internal/infra/config"

func LoadConfig(prefix string) Config {
	return Config{
		Host:     config.GetEnv(prefix+"_HOST", "localhost"),
		Port:     config.GetEnvAsInt(prefix+"_PORT", 5432),
		User:     config.GetEnv(prefix+"_USER", "postgres"),
		Password: config.GetEnv(prefix+"_PASSWORD", "rahasia"),
		DBName:   config.GetEnv(prefix+"_NAME", "postgres"),
		SSLMode:  config.GetEnv(prefix+"_SSLMODE", "disable"),
	}
}