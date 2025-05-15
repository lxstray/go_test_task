package database

import (
	"context"
	"fmt"
	"gotask/internal/config"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresDB struct {
	DB *pgxpool.Pool
}

var (
	once       sync.Once
	dbInstance *postgresDB
)

func NewPostgresDB(cfg *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s pool_max_conns=%d",
			cfg.Db.Host,
			cfg.Db.User,
			cfg.Db.Password,
			cfg.Db.Dbname,
			cfg.Db.Port,
			cfg.Db.Sslmode,
			cfg.Db.Timezone,
			cfg.Db.Max_conn,
		)

		pool, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			panic(fmt.Errorf("failed to create pgxpool: %w", err))
		}

		dbInstance = &postgresDB{DB: pool}
	})

	return dbInstance
}

func (p *postgresDB) GetDB() *pgxpool.Pool {
	return dbInstance.DB
}

func (p *postgresDB) CloseDB() {
	p.DB.Close()
}
