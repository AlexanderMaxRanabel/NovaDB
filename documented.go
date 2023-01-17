package main

import (
    "encoding/json"
    "fmt"
    "sync"
)

// Database is a struct that stores data in key-value format
// It uses a map to store data and a sync.RWMutex for concurrent access
type Database struct {
    data map[string]string
    sync.RWMutex
}

// Set method adds a new key-value pair to the database
// It takes in a key and a value as input and stores it in the data map
// It is thread-safe
func (db *Database) Set(key, value string) error {
    db.Lock()
    defer db.Unlock()
    db.data[key] = value
    return nil
}

// Get method retrieves the value for a given key from the database
// It takes in a key as input and returns the corresponding value and an error
// If the key is not found it returns an error
// It is thread-safe
func (db *Database) Get(key string) (string, error) {
    db.RLock()
    defer db.RUnlock()
    value, ok := db.data[key]
    if !ok {
        return "", fmt.Errorf("Key not found")
    }
    return value, nil
}

// Delete method deletes a key-value pair from the database
// It takes in a key as input and removes the corresponding key-value pair from the data map
// If the key is not found it returns an error
// It is thread-safe
func (db *Database) Delete(key string) error {
    db.Lock()
    defer db.Unlock()
    _, ok := db.data[key]
    if !ok {
        return fmt.Errorf("Key not found")
    }
    delete(db.data, key)
    return nil
}

// List method returns a slice of all keys present in the database
// It is thread-safe
func (db *Database) List() ([]string, error) {
    db.RLock()
    defer db.RUnlock()
    keys := make([]string, 0, len(db.data))
    for k := range db.data {
        keys = append(keys, k)
    }
    return keys, nil
}

// Dump method returns the whole database in json format
// It is thread-safe
func (db *Database) Dump() (string, error) {
    db.RLock()
    defer db.RUnlock()
    jsonData, err := json.Marshal(db.data)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}
