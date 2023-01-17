package main

import (
    "encoding/json"
    "fmt"
    "sync"
)

// Database struct
type Database struct {
    data map[string]string
    sync.RWMutex
}

// Set method to add key and value to the database
func (db *Database) Set(key, value string) error {
    db.Lock()
    defer db.Unlock()
    db.data[key] = value
    return nil
}

// Get method to retrieve value by key
func (db *Database) Get(key string) (string, error) {
    db.RLock()
    defer db.RUnlock()
    value, ok := db.data[key]
    if !ok {
        return "", fmt.Errorf("Key not found")
    }
    return value, nil
}

// Delete method to delete key and value from the database
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

// List method to list all keys
func (db *Database) List() ([]string, error) {
    db.RLock()
    defer db.RUnlock()
    keys := make([]string, 0, len(db.data))
    for k := range db.data {
        keys = append(keys, k)
    }
    return keys, nil
}

// Dump method to return database in json format
func (db *Database) Dump() (string, error) {
    db.RLock()
    defer db.RUnlock()
    jsonData, err := json.Marshal(db.data)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}

func main() {
    db := &Database{data: make(map[string]string)}
    db.Set("name", "John Smith")
    db.Set("age", "30")
    db.Set("gender", "male")
    name, _ := db.Get("name")
    age, _ := db.Get("age")
    gender, _ := db.Get("gender")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Gender:", gender)
    keys, _ := db.List()
    fmt.Println("Keys:", keys)
    jsonData, _ := db.Dump()
    fmt.Println("Database:", jsonData)
    db.Delete("gender")
    keys, _ = db.List()
    fmt.Println("Keys after delete:", keys)
}
