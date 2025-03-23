package cache

import (
	"context"
	"fmt"
	"syscall/js"
)

type Cache interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	FuzzySearch(ctx context.Context, query string) (string, error)
	Close(ctx context.Context) error
}

type App struct {
	cache Cache
}

func NewApp(cache Cache) *App {
	return &App{cache: cache}
}

type CacheWrapperAPI struct {
	app *App
}

func NewCacheHandler(app *App) *CacheWrapperAPI {
	return &CacheWrapperAPI{app: app}
}

func (r *CacheWrapperAPI) Set(args []js.Value) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid arguments")
	}
	key := args[0].String()
	value := args[1].String()
	err := r.app.cache.Set(context.Background(), key, value)
	return nil, err
}

func (r *CacheWrapperAPI) Get(args []js.Value) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("invalid arguments")
	}
	key := args[0].String()
	fmt.Println(key)
	value, err := r.app.cache.Get(context.Background(), key)
	return value, err
}

func (r *CacheWrapperAPI) Delete(args []js.Value) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("invalid arguments")
	}
	key := args[0].String()
	err := r.app.cache.Delete(context.Background(), key)
	return nil, err
}

func (r *CacheWrapperAPI) FuzzySearch(args []js.Value) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("invalid arguments")
	}
	queryStr := args[0].String()
	results, err := r.app.cache.FuzzySearch(context.Background(), queryStr)
	return results, err
}
