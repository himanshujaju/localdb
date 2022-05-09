package localdb

import (
  "fmt"
  "errors"
)

type Database struct {
  data map[string]string
}

// Creates a new database.
func CreateDB() (db *Database) {
  return &Database{
    data: make(map[string]string),
  }
}

// Flushes all contents of the database and empties it.
func (db *Database) Clear() {
  db.data = make(map[string]string)
}

// Sets the value of `key` to `value`. Overwrites existing value if the key already exists.
func (db *Database) Set(key string, value string) {
  db.data[key] = value
}

// Returns the stored value for `key` if found, returns an error otherwise.
func (db *Database) Get(key string) (value string, err error) {
  value, present := db.data[key]
  if !present {
    return "", errors.New(fmt.Sprintf("Key %v not found in database.", key))
  }

  return value, nil
}

// Returns all the keys in our database.
func (db *Database) GetKeys() (keys []string) {
  keys = make([]string, 0, len(db.data))
  for k := range db.data {
    keys = append(keys, k)
  }
  return
}

// Removes `key` from our database.
// No OP if the key does not already exist.
func (db *Database) Erase(key string) {
  delete(db.data, key)
}
