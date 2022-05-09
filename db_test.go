package localdb

import (
  "errors"
  "reflect"
  "testing"
)

func TestCreateEmptyDB(test *testing.T) {
  db := CreateDB()
  assertEqual(test, 0, len(db.GetKeys())) 
}

func TestSetsNewKey(test *testing.T) {
  db := CreateDB()
  db.Set("key", "value")

  assertEqual(test, []string{"key"}, db.GetKeys())
  assertEqual(test, "value", db.data["key"])
}

func TestOverwritesExistingKey(test *testing.T) {
  db := CreateDB()
  db.Set("key", "value1")
  assertEqual(test, "value1", db.data["key"])

  db.Set("key", "value2")
  assertEqual(test, []string{"key"}, db.GetKeys())
  assertEqual(test, "value2", db.data["key"])
}

func TestGetsExistingKey(test *testing.T) {
  db := CreateDB()
  db.Set("key", "value")

  val, err := db.Get("key")
  assertEqual(test, "value", val)
  assertEqual(test, nil, err)
}

func TestGetsNonExistingKey(test *testing.T) {
  db := CreateDB()

  val, err := db.Get("key")
  assertEqual(test, "", val)
  assertEqual(test, errors.New("Key key not found in database."), err)
}

func TestDropDB(test *testing.T) {
  db := CreateDB()
  db.Set("key1", "value1")
  db.Set("key2", "value2")

  db.Clear()
  assertEqual(test, 0, len(db.GetKeys()))
}

func TestEraseExistingKey(test *testing.T) {
  db := CreateDB()
  db.Set("key1", "value1")
  db.Set("key2", "value2")

  db.Erase("key1")

  assertEqual(test, 1, len(db.GetKeys()))
  
  val, err := db.Get("key1")
  assertEqual(test, "", val)
  assertEqual(test, errors.New("Key key1 not found in database."), err)

  val, err = db.Get("key2")
  assertEqual(test, "value2", val)
  assertEqual(test, nil, err)
}

func TestEraseNonExistentKey(test *testing.T) {
  db := CreateDB()
  db.Set("key1", "value1")

  db.Erase("key2")

  assertEqual(test, 1, len(db.GetKeys()))

  val, err := db.Get("key1")
  assertEqual(test, "value1", val)
  assertEqual(test, nil, err)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %s (type %v)\n Actual: %s (type %v)", expected, reflect.TypeOf(expected),
      actual, reflect.TypeOf(actual))
	}
}
