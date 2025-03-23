package storage

import (
	"context"
	"fmt"
	"syscall/js"
)

type LocalStorageCache struct {
	window  js.Value
	storage js.Value
}

func NewLocalStorageCache() (*LocalStorageCache, error) {
	window := js.Global()
	storage := window.Get("localStorage")
	if storage.IsUndefined() || storage.IsNull() {
		return nil, fmt.Errorf("localStorage is not available")
	}
	return &LocalStorageCache{
		window:  window,
		storage: storage,
	}, nil
}

func (c *LocalStorageCache) Set(ctx context.Context, key string, value string) error {
	c.storage.Call("setItem", key, value)
	return nil
}

func (c *LocalStorageCache) Get(ctx context.Context, key string) (string, error) {
	val := c.storage.Call("getItem", key)
	if val.IsNull() {
		return "", fmt.Errorf("key not found")
	}
	return val.String(), nil
}

func (c *LocalStorageCache) Delete(ctx context.Context, key string) error {
	c.storage.Call("removeItem", key)
	return nil
}

func (c *LocalStorageCache) FuzzySearch(ctx context.Context, query string) (string, error) {
	// LocalStorage does not natively support fuzzy search.
	// This is a simplified example that checks if the query is a substring of any key.
	// For more advanced fuzzy search, consider using a client-side JavaScript library.
	keys := c.getKeys()
	for _, k := range keys {
		val, _ := c.Get(ctx, k)
		if contains(val, query) {
			return val, nil
		}
	}

	return "", fmt.Errorf("no match found")
}

func (c *LocalStorageCache) Close(ctx context.Context) error {
	return nil
}

func (c *LocalStorageCache) getKeys() []string {
	length := c.storage.Get("length").Int()
	keys := make([]string, length)
	for i := 0; i < length; i++ {
		keys[i] = c.storage.Call("key", i).String()
	}
	return keys
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
