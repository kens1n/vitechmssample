package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE hdata
			(
				guid uuid NOT NULL,
				hash VARCHAR(32) UNIQUE NOT NULL,
				data VARCHAR(32) UNIQUE NOT NULL
			);
			
			CREATE UNIQUE INDEX hdata_unique_idx ON hdata (guid, hash);
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			DROP INDEX hdata_unique_idx;
			DROP TABLE hdata;
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210710090940_create_hdata_table", up, down, opts)
}
