package services

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log/slog"
	"os"
	"wails_tables/internal/database"
	"wails_tables/internal/models"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Migrator struct{}

var MigratorService = application.NewService(&Migrator{})

var db = database.GetInstance()

func CalculateFileSha256Checksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func dropDatabase() error {
	_ = database.Shutdown()
	if _, err := os.Stat(database.Path); err == nil {
		err := os.Remove(database.Path)
		if err != nil {

			return err
		}
	}
	return nil
}

func createDatabase(entities ...any) error {
	// Close current connections and create new database
	err := dropDatabase()
	if err != nil {
		return err
	}

	db = database.GetInstance()

	err = db.AutoMigrate(entities...)
	if err != nil {
		return err
	}

	InsertDefaultData()
	return nil
}

func Migrate(entities ...any) error {
	hashBeforeMigration, err := CalculateFileSha256Checksum(database.Path)

	slog.Info("Calculating hash...")

	if err != nil {
		return err
	}

	slog.Info("Apply migrations")
	err = db.AutoMigrate(entities...)

	if err != nil {
		slog.Info(fmt.Sprintf("Error occurred while migrations: %s, Recreate database...", err))
		if err = createDatabase(entities...); err != nil {
			slog.Error(fmt.Sprintf("Error occurred again: %s. Panic!", err))
			return err
		}
	} else {
		slog.Info("Calculating hash after migrations...")
		var hashAfterMigration string
		hashAfterMigration, err = CalculateFileSha256Checksum(database.Path)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to calc hash: %s", err))
			return err
		}

		if hashAfterMigration != hashBeforeMigration {
			slog.Info("Hashes before and after migrations are different. Recreate database...")
			err = createDatabase(entities...)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to create new database: %s", err))
				return err
			}
		}
	}
	slog.Info("Migrations proceeded")
	return nil
}

func (migrator *Migrator) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	return Migrate(models.Entities...)
}

func (migrator *Migrator) OnShutdown() {
	err := database.Shutdown()
	if err != nil {
		panic(err)
	}
}
