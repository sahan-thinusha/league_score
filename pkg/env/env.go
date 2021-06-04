package env

import (
	"github.com/dgrijalva/jwt-go"
	gorm "github.com/jinzhu/gorm"
	"os"
)

var RDB *gorm.DB
const REST_PORT = "REST_PORT"
var RestPort = "8080"

const E3_DIALET = "E3_DIALET"

var E3DIALET = "mysql"

const E3_DB = "E3_DB"

var E3db = "league_score"

const E3_HOST = "E3_HOST"

var E3host = "127.0.0.1"

const E3_PORT = "E3_PORT"

var E3port = "3306"

const E3_USER = "E3_USER"

var E3user = "root"

const E3_PWD = "E3_PWD"

var E3pwd = "sahan"

const E3_URL = "E3_URL"

var E3url = "127.0.0.1"

func LoadEnvs() {

	if val := os.Getenv(E3_DIALET); val != "" {
		E3DIALET = val
	}
	if val := os.Getenv(E3_DB); val != "" {
		E3db = val
	}
	if val := os.Getenv(E3_HOST); val != "" {
		E3host = val
	}
	if val := os.Getenv(E3_PORT); val != "" {
		E3port = val
	}
	if val := os.Getenv(E3_USER); val != "" {
		E3user = val
	}
	if val := os.Getenv(E3_PWD); val != "" {
		E3pwd = val
	}
	if val := os.Getenv(E3_URL); val != "" {
		E3url = val
	}
}


const (
	FARMER = "ROLE_FARMER"
	ADMIN = "ROLE_ADMIN"
)

type JwtCustomClaims struct {
	Sub  string
	Auth string
	jwt.StandardClaims
}