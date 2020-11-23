# Golang Example Project Layout

## Available make target

- build, build-all
- test
- clean
- rebuild (clean and build)
- docker-build-app (rebuild with go:1.15 docker)

---

### TODO - cmd/db

- [x] add cmd/db/newCmd for `db new --db-filename example.db`
- [ ] add cmd/db/migrateCmd for `db migrate`
- [ ] add cmd/db/backupCmd for `db backup --backup-dir ./backup`
- [ ] add cmd/db/restoreCmd for `db restore --from-file backup.db`
- [ ] add cmd/db/dumpCmd for `db dump`

### TODO - cmd/db/new.go

- [ ] read db filepath from configuration.
- [ ] check existing sqlite3 db file.
- [ ] create new sqlite3 db if not exists.

### TODO - cmd/db/backup.go

- [ ] read flag --backup-dir
- [ ] generate new backup filename.
- [ ] create target backup db file.
- [ ] do online backup to target db.
