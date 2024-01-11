//go:build integration
package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "order_test"
	port     = "5435"
	dsn      = "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var testDb *sqlx.DB
var testRepo RepositoryInterface

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("could not connect docker", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		if resource != nil {
			_ = pool.Purge(resource)
		}
		log.Fatal("could not start resource", err)
	}

	if err := pool.Retry(func() error {
		testDb, err = sqlx.Open("postgres", fmt.Sprintf(dsn, host, port, dbName, user, password))
		if err != nil {
			log.Println("error:", err)
			return err
		}

		return testDb.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatal("could not connect to database:", err)
	}

	err = createTable()
	if err != nil {
		log.Fatal("error creating tables", err)
	}

	testRepo = &repository{db: testDb}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatal("could not purge the resource:", err)
	}

	os.Exit(code)
}

func createTable() error {
	tableSQL, err := os.ReadFile("./../../db/migrations/000001_init.up.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDb.Exec(string(tableSQL))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
