package app

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/san-sp/GoCallAI/backend/internal/config"
	"github.com/san-sp/GoCallAI/backend/internal/storage/postgres"
)

type App struct {
	DB                 *pgxpool.Pool
	Redis              *redis.Client
	SettingsRepository *postgres.SettingsRepository
	CallsRepository    *postgres.CallsRepository
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database config: %w", err)
	}
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL pool: %w", err)
	}
	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
	})
	if err := cache.Ping(ctx).Err(); err != nil {
		_ = cache.Close()
		db.Close()
		return nil, fmt.Errorf("ping Redis: %w", err)
	}

	if err := postgres.RunMigrations(ctx, db); err != nil {
		_ = cache.Close()
		db.Close()
		return nil, fmt.Errorf("run database migrations: %w", err)
	}

	return &App{
		DB:                 db,
		Redis:              cache,
		SettingsRepository: postgres.NewSettingsRepository(db),
		CallsRepository:    postgres.NewCallsRepository(db),
	}, nil
}

func (a *App) Close() {
	if a == nil {
		return
	}
	if a.Redis != nil {
		_ = a.Redis.Close()
	}
	if a.DB != nil {
		a.DB.Close()
	}
}

func (a *App) CheckDependencies(ctx context.Context) error {
	if err := a.DB.Ping(ctx); err != nil {
		return fmt.Errorf("PostgreSQL unavailable: %w", err)
	}
	if err := a.Redis.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("Redis unavailable: %w", err)
	}
	return nil
}
