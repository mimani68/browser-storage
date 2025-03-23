# Browser Storage Engine Wrapper

A lightweight Go library providing unified access to browser storage mechanisms through WebAssembly (WASM), with automatic fallback strategies and fuzzy search capabilities.

## Features

- **Multi-Storage Support**
  - localStorage (default)
  - IndexedDB
  - Redis* (remote connection placeholder)
  - Easy extension point for custom storage engines

- **WASM Optimized**  
  Full compatibility with browser execution environment

- **Advanced Operations**
  - Fuzzy key search
  - TTL support
  - Cross-storage caching strategies

- **Zero-Dependency Core**  
  Pure Go implementation (except IndexedDB syscall/js binding)

## Installation

```bash
# Build WASM target
GOARCH=wasm GOOS=js go build -o storage.wasm main.go

# Include in your web project
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
