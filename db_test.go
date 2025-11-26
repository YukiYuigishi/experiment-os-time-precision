package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// testDB keeps a connection that integration tests can reuse.
var testDB *gorm.DB

func TestMain(m *testing.M) {
	exitCode := func() int {
		ctx := context.Background()
		dbName := "users"
		dbUser := "user"
		dbPassword := "password"

		container, err := postgres.Run(ctx,
			"postgres:14-alpine",
			postgres.WithDatabase(dbName),
			postgres.WithUsername(dbUser),
			postgres.WithPassword(dbPassword),
			postgres.BasicWaitStrategies(),
		)
		if err != nil {
			log.Fatalf("failed to start postgres container: %v", err)
		}
		defer func() {
			if terminateErr := container.Terminate(ctx); terminateErr != nil {
				log.Printf("failed to terminate postgres container: %v", terminateErr)
			}
		}()

		dsn, err := container.ConnectionString(ctx, "sslmode=disable")
		if err != nil {
			log.Fatalf("failed to get postgres connection string: %v", err)
		}

		testDB, err = gorm.Open(gormpostgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open gorm connection: %v", err)
		}

		if err := testDB.AutoMigrate(&TimePrecisionExperience{}); err != nil {
			log.Fatalf("failed to automigrate schema: %v", err)
		}

		return m.Run()
	}()

	os.Exit(exitCode)
}
