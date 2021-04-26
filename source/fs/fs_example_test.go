package fs_test

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/fs"
	"github.com/golang-migrate/migrate/v4/source/fs/testdata"
)

func Example_fs() {
	d, err := fs.WithInstance(testdata.FS, "")
	if err != nil {
		panic("bad migrations found!")
	}
	m, err := migrate.NewWithSourceInstance("godoc-vfs", d, "database://foobar")
	if err != nil {
		panic("error creating the migrations")
	}
	err = m.Up()
	if err != nil {
		panic("up failed")
	}
}
