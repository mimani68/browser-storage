package storage

import (
	"sync"
)

// InMemoryCache implementation - likely causing issues in WASM due to bleve's dependencies

type InMemoryCache struct {
	data map[string]string
	// index bleve.Index
	mu sync.RWMutex
}

// func NewInMemoryCache() (*InMemoryCache, error) {
// mapping := bleve.NewIndexMapping()
// index, err := bleve.NewMemOnly(mapping)
// if err != nil {
// return nil, err
// }
// return &InMemoryCache{
// data: make(map[string]string),
// index: index,
// }, nil
// }

// func (c *InMemoryCache) Set(ctx context.Context, key string, value string) error {
// c.mu.Lock()
// defer c.mu.Unlock()
// c.data[key] = value
// err := c.index.Index(key, map[string]interface{}{"value": value})
// if err != nil {
// return err
// }
// return nil
// }

// func (c *InMemoryCache) Get(ctx context.Context, key string) (string, error) {
// c.mu.RLock()
// defer c.mu.RUnlock()
// value, ok := c.data[key]
// if !ok {
// return "", fmt.Errorf("key not found")
// }
// return value, nil
// }

// func (c *InMemoryCache) Delete(ctx context.Context, key string) error {
// c.mu.Lock()
// defer c.mu.Unlock()
// delete(c.data, key)
// err := c.index.Delete(key)
// if err != nil {
// return err
// }
// return nil
// }

// func (c *InMemoryCache) FuzzySearch(ctx context.Context, queryStr string) (string, error) {
// c.mu.RLock()
// defer c.mu.RUnlock()
// q := bleve.NewFuzzyQuery(queryStr)
// search := bleve.NewSearchRequest(q)
// searchResults, err := c.index.Search(search)
// if err != nil {
// return nil, err
// }
// var resultsstring
// for _, hit := range searchResults.Hits {
// results = append(results, hit.ID)
// }
// return results, nil
// }

// func (c *InMemoryCache) Close(ctx context.Context) error {
// return c.index.Close()
// }
