package localdb

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "os"
)

type Database struct {
  path string
  data map[string]string
}

// Creates a new database.
func CreateDB(path string) (db *Database) {
  contents, err := ioutil.ReadFile(path)
  data := make(map[string]string)
  if err == nil {
    _ = json.Unmarshal(contents, &data)
  }

  return &Database{
    path: path,
    data: data,
  }
}

// Flushes all contents of the database and empties it.
func (db *Database) Clear() (err error) {
  db.data = make(map[string]string)
  return db.persist()
}

// Sets the value of `key` to `value`. Overwrites existing value if the key already exists.
func (db *Database) Set(key string, value string) (err error) {
  db.data[key] = value
  return db.persist()
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
func (db *Database) Erase(key string) (err error){
  _, present := db.data[key]
  if !present {
    return nil
  }

  delete(db.data, key)
  return db.persist()
}

func (db *Database) persist() (err error) {
  file, err := os.Create(db.path)
  if err != nil {
    return err
  }

  defer file.Close()
  jsonString, err := json.Marshal(db.data)
  if err != nil {
    return err
  }

  _, err = file.Write(jsonString)
  if err != nil {
    return err
  }

  err = file.Sync()
  return err
}
