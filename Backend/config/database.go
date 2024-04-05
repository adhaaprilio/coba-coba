package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Config() *pgxpool.Config {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		GetString("DB_USERNAME"),
		GetString("DB_PASSWORD"),
		GetString("DB_HOST"),
		GetString("DB_PORT"),
		GetString("DB_NAME"),
	)
	// const defaultMaxConns = int32(4)
	// const defaultMinConns = int32(0)
	// const defaultMaxConnLifetime = time.Hour
	// const defaultMaxConnIdletime = time.Minute * 30
	// const defaultHealthCheckPeriod = time.Minute
	// const defaultConnectTimeout = time.Second * 5

	dbconfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	// dbconfig.MaxConns = defaultMaxConns
	// dbconfig.MinConns = defaultMinConns
	// dbconfig.MaxConnLifetime = defaultMaxConnLifetime
	// dbconfig.MaxConnIdleTime = defaultMaxConnIdletime
	// dbconfig.HealthCheckPeriod = defaultHealthCheckPeriod
	// dbconfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
	dbconfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		// log.Println("Before aqcuiring the connection pool to the database")
		return true
	}
	dbconfig.AfterRelease = func(c *pgx.Conn) bool {
		// log.Println("After releasing the connection pool to the database")
		return true
	}
	dbconfig.BeforeClose = func(c *pgx.Conn) {
		// log.Println("Closed the connection pool to the database")
	}

	return dbconfig
}

func ConnectDatabase() *pgxpool.Pool {
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	GetString("DB_USERNAME"),
	// 	GetString("DB_PASSWORD"),
	// 	GetString("DB_HOST"),
	// 	GetString("DB_PORT"),
	// 	GetString("DB_NAME"),
	// )
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config())
	log.Println("Connected to the database!!")
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	// connection, err := connPool.Acquire(context.Background())
	// if err != nil {
	// 	log.Fatal("Error while acquiring connection from the database pool!!")
	// }
	// defer connection.Release()
	// err = connection.Ping(context.Background())
	// if err != nil {
	// 	log.Fatal("Could not ping database")
	// }

	return connPool
}
