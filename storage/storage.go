package storage

import (
	"context"
	"time"
)

type IStorage interface {
	CloseDB()

	Redis() IRedisStorage
}

type IRedisStorage interface {
	SetX(context.Context, string, interface{}, time.Duration) error
	Get(context.Context, string) interface{}
	Del(context.Context, string) error
}
