package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v81"
)

var (
	PORT            string
	API_LISTEN_HOST string

	POSTGRESQL_CONN_STRING_MASTER string
	POSTGRESQL_CONN_STRING_SLAVE  string
	POSTGRESQL_MAX_IDLE_CONNS     int
	POSTGRESQL_MAX_OPEN_CONNS     int
	PRODUCT_SERVICE_ADDR          string
	CART_SERVICE_ADDR             string
	ORDER_SERVICE_ADDR            string
	JWT_SECRET                    string
	STRIPE_API_KEY                string
)

func InitEnv() {
	// loads environment variables
	envPath := "/app/secrets/.env"
	if os.Getenv("ENV") == "dev" {
		envPath = "./secrets/testing.env"
	}
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Error loading env file")
	}
	// jwt secret
	JWT_SECRET = getEnv("JWT_SECRET", "secret")

	// rest api
	PORT = getEnv("API_PORT", "8080")
	API_LISTEN_HOST = getEnv("API_LISTEN_HOST", "0.0.0.0")

	// postgress
	POSTGRESQL_CONN_STRING_MASTER = getEnv("POSTGRESQL_CONN_STRING_MASTER", "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	POSTGRESQL_CONN_STRING_SLAVE = getEnv("POSTGRESQL_CONN_STRING_SLAVE", "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	maxOpenConns, err := strconv.Atoi(getEnv("POSTGRESQL_MAX_OPEN_CONNS", "10"))
	if err != nil {
		panic("Invalid value for POSTGRESQL_MAX_OPEN_CONNS")
	}
	POSTGRESQL_MAX_OPEN_CONNS = maxOpenConns
	maxIdleConns, err := strconv.Atoi(getEnv("POSTGRESQL_MAX_IDLE_CONNS", "5"))
	if err != nil {
		panic("Invalid value for POSTGRESQL_MAX_IDLE_CONNS")
	}
	POSTGRESQL_MAX_IDLE_CONNS = maxIdleConns

	// grpc
	PRODUCT_SERVICE_ADDR = getEnv("PRODUCT_SERVICE_ADDR", "product.default.svc.cluster.local:50051")
	CART_SERVICE_ADDR = getEnv("CART_SERVICE_ADDR", "cart.default.svc.cluster.local:50051")
	ORDER_SERVICE_ADDR = getEnv("ORDER_SERVICE_ADDR", "order.default.svc.cluster.local:50051")

	// stripe
	STRIPE_API_KEY = getEnv("STRIPE_API_KEY", "some-api-key")
	stripe.Key = getEnv("STRIPE_SECRET_KEY", "some-secret-key")
}

// get env with default if the value is empty
// getEnv("ENV_VAR", "default")
func getEnv(s ...string) string {

	if len(s) <= 0 {

		// only one arg, don't provide defaults
		return ""

	} else if val := os.Getenv(s[0]); len(s) >= 2 && val != "" {

		// two args and the env var provides empty value
		return val

	} else {

		return s[1]

	}

}
