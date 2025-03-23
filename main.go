package main

import (
	"client-cache/internal/libs/cache"
	"client-cache/internal/libs/storage"
	"fmt"
	"syscall/js"
)

func registerCallbacks(api *cache.CacheWrapperAPI) {
	js.Global().Set("Set", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := api.Set(args)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return result
	}))

	js.Global().Set("Get", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := api.Get(args)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return result
	}))

	js.Global().Set("Delete", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := api.Delete(args)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return result
	}))

	js.Global().Set("FuzzySearch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result, err := api.FuzzySearch(args)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		return result
	}))
}

func main() {
	// cacheType := js.Global().Get("cacheType").String()
	cacheType := "local-storage"

	var cacheEngine cache.Cache
	var err error

	switch cacheType {
	case "local-storage":
		cacheEngine, err = storage.NewLocalStorageCache()
	case "indexeddb":
		cacheEngine, err = storage.NewIndexedDBCache()
	case "redis":
		// TODO: Connect to remote redis
		cacheEngine, err = storage.NewLocalStorageCache()
	case "bleve":
		// FIXME:
		// Commenting out bleve as it has dependencies not compatible with WASM
		// cache, err = NewInMemoryCache()
		fmt.Println("Warning: 'bleve' cache is not suitable for WASM. Using Local Storage instead.")
		cacheEngine, err = storage.NewLocalStorageCache()
	default:
		cacheEngine, err = storage.NewLocalStorageCache()
	}

	if err != nil {
		fmt.Println("Error creating cache:", err)
		return
	}

	app := cache.NewApp(cacheEngine)
	api := cache.NewCacheHandler(app)
	registerCallbacks(api)

	select {}
}
