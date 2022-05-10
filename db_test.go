package localdb

import (
  "errors"
  "os"
  "reflect"
  "testing"
)

func TestCreateEmptyDB(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  assertEqual(test, 0, len(db.GetKeys())) 
}

func TestSetsNewKey(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  db.Set("key", "value")

  assertEqual(test, []string{"key"}, db.GetKeys())
  assertEqual(test, "value", db.data["key"])
}

func TestOverwritesExistingKey(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  db.Set("key", "value1")
  assertEqual(test, "value1", db.data["key"])

  db.Set("key", "value2")
  assertEqual(test, []string{"key"}, db.GetKeys())
  assertEqual(test, "value2", db.data["key"])
}

func TestGetsExistingKey(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  db.Set("key", "value")

  val, err := db.Get("key")
  assertEqual(test, "value", val)
  assertEqual(test, nil, err)
}

func TestGetsNonExistingKey(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")

  val, err := db.Get("key")
  assertEqual(test, "", val)
  assertEqual(test, errors.New("Key key not found in database."), err)
}

func TestDropDB(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  db.Set("key1", "value1")
  db.Set("key2", "value2")

  db.Clear()
  assertEqual(test, 0, len(db.GetKeys()))
}

func TestEraseExistingKey(test *testing.T) {
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
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
  defer tearDown("test.txt")
  db := CreateDB("test.txt")
  db.Set("key1", "value1")

  db.Erase("key2")

  assertEqual(test, 1, len(db.GetKeys()))

  val, err := db.Get("key1")
  assertEqual(test, "value1", val)
  assertEqual(test, nil, err)
}

func TestLoadsPersistedData(test *testing.T) {
  defer tearDown("test.txt")
  {
    db := CreateDB("test.txt")
    db.Set("key1", "value1")
    db.Set("key2", "value2")
    db.Set("key3", "value3")
  }

  db := CreateDB("test.txt")
  val, _ := db.Get("key1")
  assertEqual(test, "value1", val)

  val, _ = db.Get("key2")
  assertEqual(test, "value2", val)

  val, _ = db.Get("key3")
  assertEqual(test, "value3", val)

  _, err := db.Get("key4")
  assertEqual(test, errors.New("Key key4 not found in database."), err)
}

func TestCreatesEmptyDBIfIncorrectPersistedData(test *testing.T) {
  defer tearDown("test.txt")
  file, err := os.Create("test.txt")
  if err != nil {
    test.Fatal("Could not create test file!")
  }
  file.WriteString("test string")

  db := CreateDB("test.txt")
  assertEqual(test, 0, len(db.GetKeys()))
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %s (type %v)\n Actual: %s (type %v)", expected, reflect.TypeOf(expected),
      actual, reflect.TypeOf(actual))
	}
}

func tearDown(path string) {
  os.Remove(path)
}
