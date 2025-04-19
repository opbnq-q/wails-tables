package database

import (
	"context"
	"database/sql"
	"github.com/ncruces/go-sqlite3"
	"github.com/ncruces/go-sqlite3/driver"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"
	"wails_tables/internal/dal"

	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/ext/unicode"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

const Path = "database.db"

func initialize() error {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err = gorm.Open(gormlite.Open("file:"+Path+"?_fk=1"), &gorm.Config{
		Logger:               newLogger,
		FullSaveAssociations: false,
	})
	if err != nil {
		return err
	}
	RegisterUnicodeExtension(db)
	if res := db.Exec(`PRAGMA foreign_keys = ON`); res.Error != nil {
		return res.Error
	}

	if err := limitConnectionPool(); err != nil {
		return err
	}

	dal.SetDefault(db)

	log.Println("Initialized database")
	return nil
}

func limitConnectionPool() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func Shutdown() error {
	db, err := db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	once = sync.Once{}
	return nil
}

func GetInstance() *gorm.DB {
	once.Do(func() {
		err := initialize()
		if err != nil {
			panic(err)
		}
	})
	return db
}

func RegisterUnicodeExtension(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	conn, err := sqlDB.Conn(ctx)
	if err != nil {
		panic(err)
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(conn)

	err = conn.Raw(func(driverConn any) error {
		c := driverConn.(driver.Conn)
		sqliteConn := c.Raw()

		if err := unicode.Register(sqliteConn); err != nil {
			return err
		}

		if err := sqliteConn.Exec(`SELECT icu_load_collation('ru-RU', 'russian')`); err != nil {
			return err
		}

		if err := sqliteConn.Exec(`SELECT icu_load_collation('en-US', 'english')`); err != nil {
			return err
		}

		stmt, _, err := sqliteConn.Prepare(`SELECT 'ы' LIKE 'Ы'`)
		if err != nil {
			return err
		}
		defer func(stmt *sqlite3.Stmt) {
			err := stmt.Close()
			if err != nil {
				slog.Error(err.Error())
			}
		}(stmt)

		if stmt.Step() {
			slog.Info("ICU test result", "value", stmt.ColumnBool(0))
		}
		return stmt.Err()
	})

	if err != nil {
		panic(err)
	}
}
