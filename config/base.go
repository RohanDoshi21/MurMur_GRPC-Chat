package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type confVars struct {
	missing   []string //name of the mandatory environment variable that are missing
	malformed []string //errors describing malformed environment varibale values
}

const (
	default_pg_port = "5432"
)

type Config struct {
	PG_DATABASE  string
	PG_USER      string
	PG_PASSWORD  string
	PG_HOST      string
	PG_PORT      string
	PG_SSL_MODE  string
	PG_ROOT_CERT string
	PG_MAX_OPEN  int
	PG_MAX_IDLE  int
	PG_MAX_TIME  time.Duration

	CASDOOR_ENDPOINT      string
	CASDOOR_CLIENT_ID     string
	CASDOOR_CLIENT_SECRET string
	CASDOOR_CERTIFICATE   string
	CASDOOR_ORG_NAME      string
	CASDOOR_APP_NAME      string

	APP_PORT  int
	GRPC_PORT int
}

var Conf = &Config{}

func Init() (*Config, error) {
	vars := &confVars{}

	Conf.PG_DATABASE = vars.mandatory("PG_DATABASE")
	Conf.PG_USER = vars.mandatory("PG_USER")
	Conf.PG_PASSWORD = vars.mandatory("PG_PASSWORD")
	Conf.PG_HOST = vars.mandatory("PG_HOST")
	Conf.PG_PORT = vars.optional("PG_PORT", default_pg_port)
	Conf.PG_SSL_MODE = vars.optional("PG_SSL_MODE", "disable")
	Conf.PG_ROOT_CERT = vars.optional("PG_ROOT_CERT", "")
	Conf.PG_MAX_OPEN = vars.mandatoryInt("PG_MAX_OPEN")
	Conf.PG_MAX_IDLE = vars.mandatoryInt("PG_MAX_IDLE")
	Conf.PG_MAX_TIME = vars.mandatoryDuration("PG_MAX_TIME")

	Conf.CASDOOR_CERTIFICATE = vars.mandatory("CASDOOR_CERTIFICATE")
	Conf.CASDOOR_CLIENT_ID = vars.mandatory("CASDOOR_CLIENT_ID")
	Conf.CASDOOR_CLIENT_SECRET = vars.mandatory("CASDOOR_CLIENT_SECRET")
	Conf.CASDOOR_ENDPOINT = vars.mandatory("CASDOOR_ENDPOINT")
	Conf.CASDOOR_ORG_NAME = vars.mandatory("CASDOOR_ORG_NAME")
	Conf.CASDOOR_APP_NAME = vars.mandatory("CASDOOR_APP_NAME")

	Conf.APP_PORT = vars.mandatoryInt("APP_PORT")
	Conf.GRPC_PORT = vars.mandatoryInt("GRPC_PORT")

	if len(vars.missing) > 0 {
		return nil, fmt.Errorf("missing environment variables: %v", vars.missing)
	}
	if len(vars.malformed) > 0 {
		return nil, fmt.Errorf("malformed environment variables: %v", vars.malformed)
	}
	return Conf, nil

}

func (vars *confVars) mandatory(key string) string {
	val := os.Getenv(key)

	if val == "" {
		vars.missing = append(vars.missing, key)
		return ""
	}

	return val
}

func (vars *confVars) mandatoryInt(key string) int {
	valStr := vars.mandatory(key)

	val, err := strconv.Atoi(valStr)
	if err != nil {
		vars.malformed = append(vars.malformed, fmt.Sprintf("mandatory %s (value=%q) is not a boolean", key, valStr))
		return 0
	}

	return val

}

func (vars *confVars) optional(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func (vars *confVars) mandatoryDuration(key string) time.Duration {
	valStr := vars.mandatory(key)

	duration, err := time.ParseDuration(valStr)
	if err != nil {
		vars.malformed = append(vars.malformed, fmt.Sprintf("mandatory %s (value=%q) is not a Duration", key, valStr))
		return 0
	}

	return duration
}
