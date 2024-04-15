package compat

import "github.com/gamanlab/go-sqlite"

func init() {
	sqlite.RegisterAsSQLITE3()
}
