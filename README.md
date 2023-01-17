# NovaDB
A Feature-Less NoSQL Database Written in Go

Features

    Simple key-value store data model
    Concurrency support using sync.RWMutex
    Error handling using Go's error type
    Set(key, value string) error method to add key-value pairs to the database
    Get(key string) (string, error) method to retrieve values by key
    Delete(key string) error method to delete key-value pairs from the database
    List() ([]string, error) method to list all keys in the database
    Dump() (string, error) method to return the whole database in json format

Usage

package main

import (
    "fmt"
)

func main() {
    db := &Database{data: make(map[string]string)}
    db.Set("name", "John Smith")
    db.Set("age", "30")
    name, _ := db.Get("name")
    age, _ := db.Get("age")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    keys, _ := db.List()
    fmt.Println("Keys:", keys)
    jsonData, _ := db.Dump()
    fmt.Println("Database:", jsonData)
    db.Delete("age")
    keys, _ = db.List()
    fmt.Println("Keys after delete:", keys)
}
