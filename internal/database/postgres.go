package database

import (
	"fmt"
	"gotask/internal/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	DB *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDB
)

func NewPostgresDB(cfg *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			cfg.Db.Host,
			cfg.Db.User,
			cfg.Db.Password,
			cfg.Db.Dbname,
			cfg.Db.Port,
			cfg.Db.Sslmode,
			cfg.Db.Timezone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		dbInstance = &postgresDB{DB: db}
	})

	return dbInstance
}

func (p *postgresDB) GetDB() *gorm.DB {
	return dbInstance.DB
}

func (p *postgresDB) CloseDB() error {
	db, err := p.DB.DB()
	if err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
