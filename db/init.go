package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	C "github.com/RohanDoshi21/messaging-platform/config"
	"github.com/go-pg/pg/v10"

	_ "github.com/lib/pq"
)

var PostgresConn *sql.DB

// Returns postgres connection URL
func GetPostgresURL() string {
	dbHost := C.Conf.PG_HOST
	dbPort := C.Conf.PG_PORT
	dbUser := C.Conf.PG_USER
	dbPass := C.Conf.PG_PASSWORD
	dbName := C.Conf.PG_DATABASE

	if C.Conf.PG_SSL_MODE == "disable" {
		return fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)
	} else {
		return fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s sslrootcert=%s",
			dbHost, dbPort, dbUser, dbPass, dbName, C.Conf.PG_SSL_MODE, C.Conf.PG_ROOT_CERT)
	}
}

// Configure postgres pooling logic
func ConfigurePGConn() {
	pgMaxOpenConns := C.Conf.PG_MAX_OPEN

	PostgresConn.SetMaxOpenConns(pgMaxOpenConns)

	pgMaxIdleConns := C.Conf.PG_MAX_IDLE

	PostgresConn.SetMaxIdleConns(pgMaxIdleConns)

	pgMaxIdleTime := C.Conf.PG_MAX_TIME

	PostgresConn.SetConnMaxIdleTime(pgMaxIdleTime)
}

func GetPGOptions() *pg.Options {
	dbHost := C.Conf.PG_HOST
	dbPort := C.Conf.PG_PORT
	dbUser := C.Conf.PG_USER
	dbPass := C.Conf.PG_PASSWORD
	dbName := C.Conf.PG_DATABASE

	return &pg.Options{
		Addr:        fmt.Sprintf("%v:%v", dbHost, dbPort),
		User:        dbUser,
		Password:    dbPass,
		Database:    dbName,
		PoolSize:    C.Conf.PG_MAX_OPEN,
		IdleTimeout: time.Duration(C.Conf.PG_MAX_IDLE),
	}
}

// Initializes/configures pg and TSDB connection pools
func Init() error {
	db, err := sql.Open("postgres", GetPostgresURL())

	if err != nil {
		return err
	}

	PostgresConn = db

	ConfigurePGConn()

	return nil
}

// Creates and returns a new transaction
func PGTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := PostgresConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Clean up all open connection pools
func Close() {
	err := PostgresConn.Close()

	if err != nil {
		log.Fatalln("Error while trying to close the postgres DB connection!", err)
	}
}


