package config

var (
	PORT                string
	ANALYTICS_GRPC_PORT string

	USER_GRPC_SERVER_ADDR string

	LOG_LEVEL int16

	JWT_SECRET_KEY       string
	PLANET_CLIENT_ADDR   string
	PLANET_CLIENT_DOMAIN string

	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
)
