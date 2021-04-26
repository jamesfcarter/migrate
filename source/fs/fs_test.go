package fs_test

import (
	"testing"

	"github.com/golang-migrate/migrate/v4/source/fs"
	"github.com/golang-migrate/migrate/v4/source/fs/testdata"
	st "github.com/golang-migrate/migrate/v4/source/testing"
)

func TestFS(t *testing.T) {
	d, err := fs.WithInstance(testdata.FS, "")
	if err != nil {
		t.Fatal(err)
	}
	st.Test(t, d)
}

func TestOpen(t *testing.T) {
	_, err := (&fs.FS{}).Open("")
	if err != fs.ErrNotImplemented {
		t.Errorf("unexpected error: %v", err)
	}
}
