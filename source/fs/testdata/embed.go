package testdata

import "embed"

// FS embeds the migrations in this directory into the go binary
//go:embed *.sql
var FS embed.FS
