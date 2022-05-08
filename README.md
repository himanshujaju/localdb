# localdb
Local key value db for cli tools in golang


## Scratchpad of thoughts
The idea is to provide a library for local persistence that can be used by cli tools. Some requirements:
- Allow the developer where to persist the data, choosing b/w git model (separate store per "folder") or bashrc model (single store per machine)

- DB.create(filePath)
- DB.drop(db)
 
- db.set(key, value) -> Creates a mapping of key:value and fails if key already exists
- db.getKeys() -> Returns a list of keys with optional prefix
- db.get(key) -> Returns value stored against the key
- db.erase(key)
