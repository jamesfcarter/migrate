# fs

This driver is catered for those that want to source migrations from
a filesystem that supports the [io/fs](https://golang.org/pkg/io/fs/)
FS interface.

A notable example of such a filesystem is the
[embed](https://golang.org/pkg/embed/) package, which allows
migrations to be embedded in a go binary.

Example usage might be to include an `embed.go` file in
the `db/migrations` directory like this:
```go
package migrations

import "embed"

// FS embeds the migrations in this directory into the go binary
//go:embed *.sql
var FS embed.FS
```

Then to perform the migrations code like this could be used:
```go
package main

import (
     "errors"
     "log"

     "github.com/golang-migrate/migrate/v4"
     "github.com/golang-migrate/migrate/v4/source/fs"
     "github.com/my-fantasic-package/db/migrations"

     _ "github.com/golang-migrate/migrate/v4/database/postgres"
     _ "github.com/lib/pq"
)

func main() {
     source, err := fs.WithInstance(migrations.FS, "")
     if err != nil {
             log.Fatalln(err)
     }
     m, err := migrate.NewWithSourceInstance("fs", source, "postgres://postgres@localhost/postgres?sslmode=disable")
     if err != nil {
             log.Fatalln(err)
     }
     if err := m.Up(); errors.Is(err, migrate.ErrNoChange) {
             log.Println(err)
     } else if err != nil {
             log.Fatalln(err)
     }
}
```
