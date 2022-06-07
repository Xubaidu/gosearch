# prd_storage

## 2022.6.6 prd

### requirement

- We need to encapsulate leveldb as a class, in which way we can more conveniently maintain multiple leveldb instances.
- We need to call function from an object, instead of from a package. i.e, replace `storage.Get()` with `leveldbStorage.Get()`

For details, we can check [storage layer of go-found](https://github.com/newpanjing/gofound/blob/main/searcher/storage/leveldb_storage.go) for reference.

### todo

Compulsory TODOs are as followed.

- [ ] methods: `put`, `get`, `delete` and `total`(calculate total entries).
- [ ] members: `db` object, `filePath`(identify each db)
- [ ] methods: `newLeveldb`. We do not have a constructor in go, so we need to encapsulate a function to new an object.

Other functions are optional.




