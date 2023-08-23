package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/redhatinsights/platform-changelog-go/internal/config"
	l "github.com/redhatinsights/platform-changelog-go/internal/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnectorImpl struct {
	db *gorm.DB
}

func NewDBConnector(cfg *config.Config) *DBConnectorImpl {
	dsn, err := buildPostgresDSN(cfg)
	if err != nil {
		l.Log.Fatal("Error building postgres DSN: ", err)
		return nil
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Info("DB initialization complete")

	return &DBConnectorImpl{db: db}
}

func SetDBConnector(db *gorm.DB) *DBConnectorImpl {
	return &DBConnectorImpl{db: db}
}

func OpenPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn, err := buildPostgresDSN(cfg)
	if err != nil {
		l.Log.Fatal("Error building postgres DSN: ", err)
		return nil, err
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		l.Log.Fatal("Error opening DB: ", err)
		return nil, err
	}

	l.Log.Info("DB initialization complete")

	return db, nil
}

func buildPostgresDSN(cfg *config.Config) (string, error) {
	var (
		user     = cfg.DatabaseConfig.DBUser
		password = cfg.DatabaseConfig.DBPassword
		dbname   = cfg.DatabaseConfig.DBName
		host     = cfg.DatabaseConfig.DBHost
		port     = cfg.DatabaseConfig.DBPort
	)

	sslConfigString, err := buildPostgresSslConfigString(cfg)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s %s", user, password, dbname, host, port, sslConfigString), nil
}

func buildPostgresSslConfigString(cfg *config.Config) (string, error) {
	if cfg.DatabaseConfig.DBSSLMode == "disable" {
		return "sslmode=disable", nil
	} else if cfg.DatabaseConfig.DBSSLMode == "verify-full" {
		return "sslmode=verify-full sslrootcert=" + cfg.DatabaseConfig.RDSCa, nil
	} else {
		return "", errors.New("Invalid SSL configuration for database connection: " + cfg.DatabaseConfig.DBSSLMode)
	}
}

func (conn *DBConnectorImpl) Exec(sql string) error {
	return conn.db.Exec(sql).Error
}
