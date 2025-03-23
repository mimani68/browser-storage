package storage

import (
	"context"
	"fmt"
)

type IndexedDBCache struct{}

func NewIndexedDBCache() (*IndexedDBCache, error) {
	return &IndexedDBCache{}, nil
}

func (c *IndexedDBCache) Set(ctx context.Context, key string, value string) error {
	return fmt.Errorf("IndexedDBCache not implemented")
}

func (c *IndexedDBCache) Get(ctx context.Context, key string) (string, error) {
	return "", fmt.Errorf("IndexedDBCache not implemented")
}

func (c *IndexedDBCache) Delete(ctx context.Context, key string) error {
	return fmt.Errorf("IndexedDBCache not implemented")
}

func (c *IndexedDBCache) FuzzySearch(ctx context.Context, query string) (string, error) {
	return "", fmt.Errorf("IndexedDBCache not implemented")
}

func (c *IndexedDBCache) Close(ctx context.Context) error {
	return nil
}
