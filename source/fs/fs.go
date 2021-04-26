// Package fs container a driver that can read migrations from any
// filesystem that implements the io/fs FS interface, notably
// embed.FS.
package fs

import (
	"errors"
	"io/fs"
	"net/http"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

var (
	// ErrNotImplemented is retured from Open
	ErrNotImplemented = errors.New("not implemented")
)

// FS is a source driver that uses any underlying filesystem that
// supports the io/fs FS interface.
type FS struct {
	httpfs.PartialDriver
}

// Open implements the source.Driver interface for VFS.
// This function always returns ErrNotImplemented, use the
// WithInstance function instead.
func (f *FS) Open(url string) (source.Driver, error) {
	return nil, ErrNotImplemented
}

// WithInstance creates a new driver from the underlying fs
// Filesystem. The path arguments specifies the location of
// the migrations within the filesystem, defaulting to "/".
func WithInstance(fs fs.FS, path string) (source.Driver, error) {
	if path == "" {
		path = "/"
	}

	source := &FS{}
	if err := source.Init(http.FS(fs), path); err != nil {
		return nil, err
	}
	return source, nil
}
