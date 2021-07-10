package main

import (
	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`INSERT INTO hdata VALUES (
			'f6014c6c-ddc5-11eb-ba80-0242ac130004',
			'9d5d5c7afe1a9efbb5fcf5f903e15808',
			'congrats'
		)`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec(`
			DELETE FROM hdata
			WHERE guid = 'f6014c6c-ddc5-11eb-ba80-0242ac130004'
			AND hash = '9d5d5c7afe1a9efbb5fcf5f903e15808'
		`)
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210710094202_hdata_test_seeding", up, down, opts)
}
