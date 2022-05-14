# localdb

## Why?
Key-value db backed by local file storage system. `localdb` was created mainly for cli tools in golang to store/retrieve user settings/preferences locally. This can be used by tools such as `git` (different settings per folder) as well as tools like `bash` (single `bashrc` file per user).

## API Docs

- `CreateDB(filepath)`: Takes in a relative `filepath` as input. If the file already exists, then it tries to load the contents to an in memory db, otherwise it creates the file.
- `Set(key, value)`: Saves `value` against `key` in our database. Overwrites the value if `key` already exists in the db.
- `GetKeys()`: Returns all the keys in our database.
- `Clear()`: Flushes out all data. Note - this also flushes out persisted data so the operation is permanent and not reversible.
- `Erase(key)`: Erases the `key` and the corresponding value from the db.

## Example

Example to store the author's name for a git tool.

```
...
import (
  ...
  "github.com/himanshujaju/localdb"
  ...
  )

// Creates a database that stores/retrieves data from .settings file.
database := CreateDB(".git")
// Stores the author's name in db.
database.Set("author", "himanshu")
// Outputs the author's name
val, _ := database.Get("author")
fmt.Println("Author:", val)
...
```

