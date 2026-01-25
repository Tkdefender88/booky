package repo

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/glebarez/sqlite"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

type Datastore struct {
	db *sql.DB
}

func (d *Datastore) DB() *sql.DB { return d.db }

func NewDB() (*Datastore, error) {

	dbPath, err := getDBPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get database path: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open db file %w", err)
	}

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Datastore{db: db}, nil
}

func (d *Datastore) Close() error {
	return d.db.Close()
}

func getDBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	dbPath := filepath.Join(homeDir, ".local", "share", "booky", "bookmarks.db")

	_, err = os.Stat(dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(dbPath)
			dbDir := filepath.Dir(dbPath)
			if err := os.MkdirAll(dbDir, 0755); err != nil {
				return "", fmt.Errorf("failed to create db directory: %w", err)
			}
			_, err = os.Create(dbPath)
			if err != nil {
				return "", fmt.Errorf("failed to create db file: %w", err)
			}
			return dbPath, nil
		}
		return "", fmt.Errorf("failed to check if db file exists: %w", err)
	}

	return dbPath, nil
}
